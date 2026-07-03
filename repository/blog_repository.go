package repository

import (
	"blogapi/models"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DB *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		DB: db,
	}
}
func (r *BlogRepository) Create(blog *models.Blog) error {

	tx := r.DB.Begin()

	if err := tx.Create(blog).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *BlogRepository) GetAll() ([]models.Blog, error) {

	var blogs []models.Blog

	err := r.DB.Find(&blogs).Error

	return blogs, err
}
func (r *BlogRepository) GetByID(id uint) (*models.Blog, error) {

	var blog models.Blog

	err := r.DB.First(&blog, id).Error

	if err != nil {
		return nil, err
	}

	return &blog, nil
}
func (r *BlogRepository) Update(id uint, updatedBlog *models.Blog) error {

	tx := r.DB.Begin()

	var blog models.Blog

	if err := tx.First(&blog, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	blog.Title = updatedBlog.Title
	blog.Content = updatedBlog.Content
	blog.Author = updatedBlog.Author

	if err := tx.Save(&blog).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *BlogRepository) Delete(id uint) error {

	tx := r.DB.Begin()

	var blog models.Blog

	if err := tx.First(&blog, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&blog).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}