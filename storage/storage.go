package storage

import "practise_CRUD/models"

type StorageI interface {
	Category() CategoryRepoI
	Order() OrderRepoI
	Product() ProductRepoI
	User() UserRepoI
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	GetByID(req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(req *models.GetListCategoryRequests) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (*models.Category, error)
	Delete(req *models.CategoryPrimaryKey) error
}

type OrderRepoI interface {
	Create(req *models.CreateOrder) (*models.Order, error)
	GetByID(req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(req *models.GetListOrderRequests) (*models.GetListOrderResponse, error)
	Update(req *models.UpdateOrder) (*models.Order, error)
	Delete(req *models.OrderPrimaryKey) error
}

type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	GetByID(req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(req *models.GetListProductRequests) (*models.GetListProductResponse, error)
	Update(req *models.UpdateProduct) (*models.Product, error)
	Delete(req *models.ProductPrimaryKey) error
}

type UserRepoI interface {
	Create(req *models.CreateUser) (*models.User, error)
	GetByID(req *models.UserPrimaryKey) (*models.User, error)
	GetList(req *models.GetListUserRequests) (*models.GetListUserResponse, error)
	Update(req *models.UpdateUser) (*models.User, error)
	Delete(req *models.UserPrimaryKey) error
}
