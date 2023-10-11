package controller

import (
	"log"
	"practise_CRUD/models"
)

func (c *Con) CreateCategory(req *models.CreateCategory) (*models.Category, error) {

	log.Println("CreateCategory req ----> ", req)
	resp, err := c.storage.Category().Create(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (c *Con) GetByIDCategory(req *models.CategoryPrimaryKey) (*models.Category, error) {

	log.Println("GetByIDCategory req ----> ", req)
	resp, err := c.storage.Category().GetByID(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) GetListCategory(req *models.GetListCategoryRequests) (*models.GetListCategoryResponse, error) {

	log.Println("GetListCategory req ----> ", req)
	resp, err := c.storage.Category().GetList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) UpdateCategory(req *models.UpdateCategory) (*models.Category, error) {

	log.Println("UpdateCategory req ----> ", req)
	resp, err := c.storage.Category().Update(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) DeleteCategory(req *models.CategoryPrimaryKey) error {

	log.Println("GetByIDCategory req ----> ", req)
	err := c.storage.Category().Delete(req)
	if err != nil {
		return err
	}
	return nil
}
