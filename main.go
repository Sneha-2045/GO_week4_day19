package main

import (
	"blogapi/config"
	"blogapi/controllers"
	"blogapi/models"
	"blogapi/repository"
	"blogapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDatabase()

	db.AutoMigrate(&models.Blog{})

	repo := repository.NewBlogRepository(db)

	controller := controllers.NewBlogController(repo)

	r := gin.Default()

	routes.SetupRoutes(r, controller)

	r.Run(":8080")
}