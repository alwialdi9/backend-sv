package controllers

import (
	"github.com/alwialdi9/backend-sv/config"
	"github.com/alwialdi9/backend-sv/internal/models"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type CreateArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

// @Summary Add article.
// @Description Add article and save to database.
// @Tags Article
// @Accept json
// @Produce json
// @Param body json CreateArticleRequest true "the body to create a Article"
// @Success 200 object CreateArticleRequest
// @Router /article [POST]
func CreateArticle(c *fiber.Ctx) error {

	var req CreateArticleRequest
	err := c.BodyParser(&req)
	if err != nil {
		log.WithError(err).Error("Realtime sync handler: parse bad request")
		return fiber.ErrBadRequest
	}

	result := config.DB.Create(&req)

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
// @Success 200 object []models.Posts{}
// @Router /article/:limit/:offset [GET]
func GetArticle(c *fiber.Ctx) error {
	limit, _ := c.ParamsInt("limit", 10)
	offset, _ := c.ParamsInt("offset", 0)

	var data []models.Posts

	result := config.DB.Select("title", "content", "category", "status").Limit(limit).Offset(offset).Find(&data)

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
// @Success 200 object models.Posts{}
// @Router /article/:id [GET]
func GetArticleById(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.Posts

	result := config.DB.Select("title", "content", "category", "status").Where("id = ?", id).Find(&data)

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
