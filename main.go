package main

import (
	"geo-data-backend/handlers"
	"geo-data-backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to the PostgreSQL database
	models.ConnectDB()

	// Use CORS middleware
	r.Use(cors.Default())

	// Routes for authentication
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.GET("/logout", handlers.Logout)
	}

	// Routes for managing files
	files := r.Group("/files")
	{
		files.POST("/upload", handlers.UploadFile)
		files.GET("/:id", handlers.GetFile)
		files.DELETE("/:id", handlers.DeleteFile)
	}

	// Routes for geospatial data
	geo := r.Group("/geospatial")
	{
		geo.POST("/create", handlers.CreateGeoData)
		geo.GET("/:id", handlers.GetGeoData)
		geo.PUT("/:id", handlers.UpdateGeoData)
		geo.DELETE("/:id", handlers.DeleteGeoData)
	}

	r.Run(":8080")
}
