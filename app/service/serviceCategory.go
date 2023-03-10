package service

import (
	"github.com/Tabed23/article-category-crud/app/types"
	u "github.com/Tabed23/article-category-crud/app/utils"
)

type CategoryService struct {
}

func (CategoryService) CreatCategory(cat *types.Category) (string, error) {
	c := &types.Category{
		CategoryID: cat.CategoryID,
		Name:       cat.Name,
	}
	u.WriteOnFile(c)

	err := db.Table("categories").Create(&c).Error

	return "ok", err
}

func (CategoryService) CategoryFindById(id string) (*types.Category, error) {
	c := &types.Category{}
	c, err := u.FindByIdFile(id, c)
	if err != nil {
		return nil, err
	}

	err = db.Table("categories").Where("category_id = ?", id).First(c).Error
	return c, err
}

func (CategoryService) DeleteCategory(id string) (bool, error) {
	c := types.Category{}
	ok, err := u.DeleteByIdFromFile(id)
	if err != nil && !ok {
		return false, err
	}

	err = db.Table("categories").Where("category_id = ?", id).First(&c).Error
	if err != nil {
		return false, err
	}
	err = db.Delete(c, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (CategoryService) UpdateCategory(id string, catUpdate types.Category) (*types.Category, error) {
	c := types.Category{}
	err := db.Table("categories").Where("category_id = ?", id).First(&c).Error
	if err != nil {
		return nil, err
	}
	c.Name = catUpdate.Name
	err = db.Save(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (CategoryService) GetCategory() (*[]types.Category, error) {
	cat := make([]types.Category, 0)
	err := db.Table("categories").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	return &cat, nil
}
