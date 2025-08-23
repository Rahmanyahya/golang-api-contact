package main

import (
	"golang-api-contact/config"
	"golang-api-contact/controller"
	"golang-api-contact/helpers"
	"golang-api-contact/repositories"
	"golang-api-contact/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error Loading .env file")
	} 

	config.InitDB()

	mainController := controller.NewMainController()
	healthController := controller.NewHealthController()
	contactRepository := repositories.NewContactRepository(config.DB)
	contactService := services.NewContactService(contactRepository)
	contactController := controller.NewContactController(contactService)

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     helper.ParseEnvList("CORS_ALLOWED_ORIGINS"),
		AllowMethods:     helper.ParseEnvList("CORS_ALLOWED_METHODS"),
		AllowHeaders:     helper.ParseEnvList("CORS_ALLOWED_HEADERS"),
		AllowCredentials: helper.GetEnvBool("CORS_ALLOW_CREDENTIALS", false),
		ExposeHeaders:    helper.ParseEnvList("CORS_EXPOSE_HEADERS"),
		MaxAge:           12 * 60 * 60, // 12 hours
	}

	router.Use(cors.New(corsConfig))

	router.GET("/",  mainController.MainController)
	router.GET("/health", healthController.HealthCheck)
	router.GET("/contacts", contactController.GetContacts)
	router.GET("/contacts/:id", contactController.GetContact)
	router.POST("/contacts", contactController.CreateContact)
	router.PUT("/contacts/:id", contactController.UpdateContact)
	router.DELETE("/contacts/:id", contactController.DeleteContact)

	appPort := config.GetEnv("APP_PORT", "8080")

	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Failed running service : %v", err)
	}
}