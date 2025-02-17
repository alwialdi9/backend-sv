package controllers

import (
	"fmt"

	"github.com/alwialdi9/backend-sv/config"
	"github.com/alwialdi9/backend-sv/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type ArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

// Inisialisasi Validator
var validate = validator.New()

// Custom Pesan Error
var customMessages = map[string]string{
	"Title.required":    "Judul wajib diisi",
	"Title.min":         "Judul minimal harus 20 karakter",
	"Content.required":  "Konten wajib diisi",
	"Content.min":       "Konten minimal harus 200 karakter",
	"Category.required": "Kategori wajib diisi",
	"Category.min":      "Kategori minimal harus 3 karakter",
	"Status.required":   "Status wajib diisi",
	"Status.oneof":      "Status harus diantara publish, draft, thrash",
}

func ValidateStruct(s interface{}) map[string]string {
	errors := make(map[string]string)
	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			key := fmt.Sprintf("%s.%s", field, tag)

			if msg, exists := customMessages[key]; exists {
				errors[field] = msg
			} else {
				errors[field] = fmt.Sprintf("Field %s tidak valid", field)
			}
		}
	}
	return errors
}

// @Summary Add article.
// @Description Add article and save to database.
// @Tags Article
// @Accept json
// @Produce json
// @Param Body body ArticleRequest true "the body to create a Article"
// @Success 200 object map[string]interface{}
// @Router /article [POST]
func CreateArticle(c *fiber.Ctx) error {

	var req ArticleRequest
	err := c.BodyParser(&req)
	if err != nil {
		log.WithError(err).Error("Realtime sync handler: parse bad request")
		return fiber.ErrBadRequest
	}

	// Validasi request
	errors := ValidateStruct(req)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	result := config.DB.Table("posts").Create(&req)

	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to create article")
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Failed to create article",
		})
	}

	response := map[string]interface{}{
		"status": "success",
		"data":   req,
	}
	return c.JSON(response)
}

// @Summary Get article with limit and offset.
// @Description Get article by Limit and offset.
// @Tags Article
// @Accept json
// @Produce json
// @Param limit path integer true "article limit"
// @Param offset path integer true "article offset"
// @Success 200 object map[string]interface{}
// @Router /article/:limit/:offset [GET]
func GetArticle(c *fiber.Ctx) error {
	limit, _ := c.ParamsInt("limit", 10)
	offset, _ := c.ParamsInt("offset", 0)

	var data []models.Posts

	result := config.DB.Select("id", "title", "content", "category", "status").Limit(limit).Offset(offset).Find(&data)

	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to Get article")
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Failed to Get article",
		})
	}

	response := map[string]interface{}{
		"status": "success",
		"limit":  limit,
		"offset": offset,
		"data":   data,
	}
	return c.JSON(response)
}

// @Summary Get article By ID.
// @Description Get article by Id.
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "article id"
// @Success 200 object map[string]interface{}
// @Router /article/:id [GET]
func GetArticleById(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.Posts

	result := config.DB.Select("id", "title", "content", "category", "status").Where("id = ?", id).First(&data)

	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to Get article")
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Failed to Get article",
		})
	}

	response := map[string]interface{}{
		"status": "success",
		"data":   data,
	}
	return c.JSON(response)
}

// @Summary Edit article By ID.
// @Description Edit article by Id.
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "article id"
// @Param Body body ArticleRequest true "the body to update a Article"
// @Success 200 object map[string]interface{}
// @Router /article/:id [POST]
func EditArticle(c *fiber.Ctx) error {
	var req ArticleRequest
	err := c.BodyParser(&req)
	if err != nil {
		log.WithError(err).Error("parse bad request")
		return fiber.ErrBadRequest
	}

	// Validasi request
	errors := ValidateStruct(req)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	id, err := c.ParamsInt("id", 0)

	if err != nil {
		log.WithError(err).Error("Failed to parse article id")
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Failed to parse article id",
		})
	}

	if id == 0 {
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Id not found",
		})
	}

	var data models.Posts

	result := config.DB.Where("id = ?", id).First(&data)

	if result.Error != nil {
		log.WithError(result.Error).Error(fmt.Sprintf("Failed to Get article with id %d", id))
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": fmt.Sprintf("Failed to Get article with id %d", id),
		})
	}

	data.Title = req.Title
	data.Content = req.Content
	data.Category = req.Category
	data.Status = req.Status
	data.ID = uint(id)

	config.DB.Save(&data)

	response := map[string]interface{}{
		"status": "success",
	}
	return c.JSON(response)
}

// @Summary Delete article By ID.
// @Description Delete article by Id.
// @Tags Article
// @Accept json
// @Produce json
// @Param id path integer true "article id"
// @Success 200 object map[string]interface{}
// @Router /article/:id [DELETE]
func DeleteArticle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)

	if err != nil {
		log.WithError(err).Error("Failed to parse article id")
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Failed to parse article id",
		})
	}

	if id == 0 {
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": "Id not found",
		})
	}

	var data models.Posts

	result := config.DB.Delete(&data, id)

	if result.Error != nil {
		log.WithError(result.Error).Error(fmt.Sprintf("Failed to Delete article with id %d", id))
		return c.JSON(map[string]interface{}{
			"status":  "failed",
			"message": fmt.Sprintf("Failed to Delete article with id %d", id),
		})
	}

	response := map[string]interface{}{
		"status": "success",
	}
	return c.JSON(response)
}
