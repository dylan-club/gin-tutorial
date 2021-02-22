package repository

import (
	"com.nicklaus/ginpractice/common"
	"com.nicklaus/ginpractice/model"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	//添加
	Create(name string) (*model.Category, error)
	//更新分类名称
	UpdateForName(category *model.Category, name string) (*model.Category, error)
	//根据id删除分类
	DeleteById(id int) error
	//根据id查找分类
	SelectById(id int) (*model.Category, error)
	//根据name查找分类
	SelectByName(name string) (*model.Category, error)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() ICategoryRepository {
	db := common.GetDB()
	_ = db.AutoMigrate(model.Category{})
	return &CategoryRepository{DB: db}
}

func (c *CategoryRepository) Create(name string) (*model.Category, error) {
	var createCategory = new(model.Category)
	createCategory.Name = name
	if err := c.DB.Create(createCategory).Error; err != nil {
		return nil, err
	}
	return createCategory, nil
}

func (c *CategoryRepository) UpdateForName(category *model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (c *CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category = new(model.Category)
	if err := c.DB.First(category, id).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepository) SelectByName(name string) (*model.Category, error) {
	var category = new(model.Category)
	if err := c.DB.Where("name = ?", name).First(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
