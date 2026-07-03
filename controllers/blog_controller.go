package controllers

import (
	"blogapi/models"
	"blogapi/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Repo *repository.BlogRepository
}

func NewBlogController(repo *repository.BlogRepository) *BlogController {
	return &BlogController{
		Repo: repo,
	}
}
func (bc *BlogController) CreateBlog(c *gin.Context) {

	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bc.Repo.Create(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, blog)
}
func (bc *BlogController) GetBlogs(c *gin.Context) {

	blogs, err := bc.Repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, blogs)
}
func (bc *BlogController) GetBlog(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	blog, err := bc.Repo.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, blog)
}
func (bc *BlogController) UpdateBlog(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bc.Repo.Update(uint(id), &blog); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
	})
}
func (bc *BlogController) DeleteBlog(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	if err := bc.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})
}