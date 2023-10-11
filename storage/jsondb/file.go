package jsondb

import (
	"os"
	"practise_CRUD/config"
	"practise_CRUD/storage"
)

type StoreJson struct {
	user     storage.UserRepoI
	order    storage.OrderRepoI
	product  storage.ProductRepoI
	category storage.CategoryRepoI
}

func NewJsonDBConnection(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.UserPath)
	if err != nil {
		return nil, err
	}

	orderFile, err := os.Open(cfg.OrderPath)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.ProductPath)
	if err != nil {
		return nil, err
	}

	categoryFile, err := os.Open(cfg.CategoryPath)
	if err != nil {
		return nil, err
	}

	return &StoreJson{
		category: NewCategoryRepo(cfg, categoryFile),
		order:    NewOrderRepo(cfg, orderFile),
		user:     NewUserRepo(cfg, userFile),
		product:  NewProductRepo(cfg, productFile),
	}, nil

}

func (s StoreJson) Category() storage.CategoryRepoI {
	return s.category
}

func (s StoreJson) Order() storage.OrderRepoI {
	return s.order
}

func (s StoreJson) User() storage.UserRepoI {
	return s.user
}

func (s StoreJson) Product() storage.ProductRepoI {
	return s.product
}
