package controller

import (
	"log"
	"practise_CRUD/models"
)

func (c *Con) CreateProduct(req *models.CreateProduct) (*models.Product, error) {

	log.Println("CreateProduct req ----> ", req)
	resp, err := c.storage.Product().Create(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (c *Con) GetByIDProduct(req *models.ProductPrimaryKey) (*models.Product, error) {

	log.Println("GetByIDProduct req ----> ", req)
	resp, err := c.storage.Product().GetByID(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) GetListProduct(req *models.GetListProductRequests) (*models.GetListProductResponse, error) {

	log.Println("GetListProduct req ----> ", req)
	resp, err := c.storage.Product().GetList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) UpdateProduct(req *models.UpdateProduct) (*models.Product, error) {

	log.Println("UpdateProduct req ----> ", req)
	resp, err := c.storage.Product().Update(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) DeleteProduct(req *models.ProductPrimaryKey) error {

	log.Println("GetByIDProduct req ----> ", req)
	err := c.storage.Product().Delete(req)
	if err != nil {
		return err
	}
	return nil
}
