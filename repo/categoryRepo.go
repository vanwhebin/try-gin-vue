package repo

import (
	"gorm.io/gorm"
	"vanwhebin/try-gin-vue/common"
	"vanwhebin/try-gin-vue/model"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	var category = model.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) SelectByID(id int) (*model.Category, error) {
	var category model.Category
	err := c.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) DeleteByID(id int) error {
	err := c.DB.Delete(model.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
