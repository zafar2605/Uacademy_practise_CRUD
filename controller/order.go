package controller

import (
	"log"
	"practise_CRUD/models"
)

func (c *Con) CreateOrder(req *models.CreateOrder) (*models.Order, error) {

	log.Println("CreateOrder req ----> ", req)
	resp, err := c.storage.Order().Create(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (c *Con) GetByIDOrder(req *models.OrderPrimaryKey) (*models.Order, error) {

	log.Println("GetByIDOrder req ----> ", req)
	resp, err := c.storage.Order().GetByID(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) GetListOrder(req *models.GetListOrderRequests) (*models.GetListOrderResponse, error) {

	log.Println("GetListOrder req ----> ", req)
	resp, err := c.storage.Order().GetList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) UpdateOrder(req *models.UpdateOrder) (*models.Order, error) {

	log.Println("UpdateOrder req ----> ", req)
	resp, err := c.storage.Order().Update(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) DeleteOrder(req *models.OrderPrimaryKey) error {

	log.Println("GetByIDOrder req ----> ", req)
	err := c.storage.Order().Delete(req)
	if err != nil {
		return err
	}
	return nil
}
