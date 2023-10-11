package jsondb

import (
	"os"
	"practise_CRUD/config"
	"practise_CRUD/models"
	"practise_CRUD/pkg"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type ProductRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewProductRepo(cfg *config.Config, file *os.File) *ProductRepo {
	return &ProductRepo{
		file: file,
		cfg:  cfg,
	}
}

func (p *ProductRepo) Create(req *models.CreateProduct) (*models.Product, error) {

	data, err := pkg.Read(p.cfg.ProductPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.Product{
		Id:   uuid.New().String(),
		Name: req.Name,
		Price: req.Price,
		Discount: req.Discount,
		DiscountType: req.DiscountType,
		CategoryID: req.CategoryID,
	}

	data = append(data, resp)
	err = pkg.Write(p.cfg.ProductPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProductRepo) GetByID(req *models.ProductPrimaryKey) (*models.Product, error) {

	data, err := pkg.Read(p.cfg.ProductPath)
	if err != nil {
		return nil, err
	}

	for _, productObject := range data {
		var product = cast.ToStringMap(productObject)

		if cast.ToString(product["id"]) == req.Id {
			return &models.Product{
				Name:         cast.ToString(product["name"]),
				Price:        cast.ToInt(product["price"]),
				Discount:     cast.ToInt(product["discount"]),
				DiscountType: cast.ToString(product["discount_type"]),
				CategoryID:   cast.ToString(product["category_id"]),
			}, nil
		}
	}

	return nil, err
}

func (p *ProductRepo) GetList(req *models.GetListProductRequests) (*models.GetListProductResponse, error) {

	data, err := pkg.Read(p.cfg.ProductPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListProductResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var product = cast.ToStringMap(data[i])
			resp.Products = append(resp.Products, &models.Product{
				Id:           cast.ToString(product["id"]),
				Name:         cast.ToString(product["name"]),
				Price:        cast.ToInt(product["price"]),
				Discount:     cast.ToInt(product["discount"]),
				DiscountType: cast.ToString(product["discount_type"]),
				CategoryID:   cast.ToString(product["category_id"]),
			})
		} else {
			break
		}
	}

	resp.Count = len(data)

	return resp, nil
}

func (p *ProductRepo) Update(req *models.UpdateProduct) (*models.Product, error) {

	data, err := pkg.Read(p.cfg.ProductPath)
	if err != nil {
		return nil, err
	}

	for index, ctg_obg := range data {
		var product = cast.ToStringMap(ctg_obg)
		if cast.ToString(product["id"]) == req.Id {
			product["name"] = req.Name
			product["price"] = req.Price
			product["discount"] = req.Discount
			product["discount_type"] = req.DiscountType
			product["category_id"] = req.CategoryID

			data[index] = product
			break
		}
	}

	err = pkg.Write(p.cfg.ProductPath, data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (p *ProductRepo) Delete(req *models.ProductPrimaryKey) error {

	data, err := pkg.Read(p.cfg.ProductPath)
	if err != nil {
		return err
	}

	for ind, ctg_obg := range data {
		var product = cast.ToStringMap(ctg_obg)
		if cast.ToString(product["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = pkg.Write(p.cfg.ProductPath, data)
	if err != nil {
		return err
	}

	return nil
}
