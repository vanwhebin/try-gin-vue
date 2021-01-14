package repo


import (
	"gorm.io/gorm"
	"vanwhebin/try-gin-vue/common"
	"vanwhebin/try-gin-vue/model"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository() PostRepository {
	return PostRepository{DB: common.GetDB()}
}

func (c PostRepository) Create(post model.Post) (*model.Post, error) {
	if err := c.DB.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (c PostRepository) Update(category model.Post, name string) (*model.Post, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c PostRepository) SelectByID(id int) (*model.Post, error) {
	var category model.Post
	err := c.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c PostRepository) DeleteByID(id int) error {
	err := c.DB.Delete(model.Post{}, id).Error
	if err != nil {
		return err
	}
	return nil
}