package config

import (
	"fmt"
	"os"

	"github.com/alwialdi9/backend-sv/internal/models"
	"github.com/alwialdi9/backend-sv/internal/utils"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables
	var dsn string
	var dialector gorm.Dialector
	var err error

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	dialector = postgres.Open(dsn)

	// Connect to database
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		utils.LogError("Failed to connect to database", err, log.Fields{})
	} else {
		utils.LogInfo("✅ Database connected successfully", log.Fields{})
	}

	DB.AutoMigrate(&models.Posts{})

}
