package controller

import (
	"log"
	"practise_CRUD/models"
)

func (c *Con) CreateUser(req *models.CreateUser) (*models.User, error) {

	log.Println("CreateUser req ----> ", req)
	resp, err := c.storage.User().Create(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (c *Con) GetByIDUser(req *models.UserPrimaryKey) (*models.User, error) {

	log.Println("GetByIDUser req ----> ", req)
	resp, err := c.storage.User().GetByID(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) GetListUser(req *models.GetListUserRequests) (*models.GetListUserResponse, error) {

	log.Println("GetListUser req ----> ", req)
	resp, err := c.storage.User().GetList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) UpdateUser(req *models.UpdateUser) (*models.User, error) {

	log.Println("UpdateUser req ----> ", req)
	resp, err := c.storage.User().Update(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Con) DeleteUser(req *models.UserPrimaryKey) error {

	log.Println("GetByIDUser req ----> ", req)
	err := c.storage.User().Delete(req)
	if err != nil {
		return err
	}
	return nil
}
