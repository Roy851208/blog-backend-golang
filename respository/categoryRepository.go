package respository

import (
	"blog/common"
	"blog/model"

	"gorm.io/gorm"
)

type CategoryRespository struct {
	DB *gorm.DB
}

// func NewCategoryRepository() CategoryRespository {
// 	return CategoryRepository{}
// }

func (c CategoryRespository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}
	if err := common.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRespository) Update(category model.Category, name string) (*model.Category, error) {
	if err := common.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRespository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := common.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRespository) DeleteById(id int) error {
	if err := common.DB.Delete(model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
