package routes

import (
	"blogapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, controller *controllers.BlogController) {

	r.POST("/blogs", controller.CreateBlog)

	r.GET("/blogs", controller.GetBlogs)

	r.GET("/blogs/:id", controller.GetBlog)

	r.PUT("/blogs/:id", controller.UpdateBlog)

	r.DELETE("/blogs/:id", controller.DeleteBlog)
}