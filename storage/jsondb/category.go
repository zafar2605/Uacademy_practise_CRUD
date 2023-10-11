package jsondb

import (
	"os"
	"practise_CRUD/config"
	"practise_CRUD/models"
	"practise_CRUD/pkg"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type CategoryRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewCategoryRepo(cfg *config.Config, file *os.File) *CategoryRepo {
	return &CategoryRepo{
		file: file,
		cfg:  cfg,
	}
}

func (c *CategoryRepo) Create(req *models.CreateCategory) (*models.Category, error) {

	data, err := pkg.Read(c.cfg.CategoryPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.Category{
		Id:   uuid.New().String(),
		Name: req.Name,
	}

	data = append(data, resp)
	err = pkg.Write(c.cfg.CategoryPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CategoryRepo) GetByID(req *models.CategoryPrimaryKey) (*models.Category, error) {

	data, err := pkg.Read(c.cfg.CategoryPath)
	if err != nil {
		return nil, err
	}

	for _, categoryObject := range data {
		var category = cast.ToStringMap(categoryObject)

		if cast.ToString(category["id"]) == req.Id {
			return &models.Category{
				Id:   cast.ToString(category["id"]),
				Name: cast.ToString(category["name"]),
			}, nil
		}
	}

	return nil, err
}

func (c *CategoryRepo) GetList(req *models.GetListCategoryRequests) (*models.GetListCategoryResponse, error) {

	data, err := pkg.Read(c.cfg.CategoryPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListCategoryResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var category = cast.ToStringMap(data[i])
			resp.Categories = append(resp.Categories, &models.Category{
				Id:   cast.ToString(category["id"]),
				Name: cast.ToString(category["name"]),
			})
		} else {
			break
		}
	}

	resp.Count = len(data)

	
	return resp, nil
}

func (c *CategoryRepo) Update(req *models.UpdateCategory) (*models.Category, error) {

	data, err := pkg.Read(c.cfg.CategoryPath)
	if err != nil {
		return nil, err
	}

	for index, ctg_obg := range data {
		var category = cast.ToStringMap(ctg_obg)
		if cast.ToString(category["id"]) == req.Id {
			category["name"] = req.Name

			data[index] = category
			break
		}
	}

	err = pkg.Write(c.cfg.CategoryPath, data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *CategoryRepo) Delete(req *models.CategoryPrimaryKey) error {

	data, err := pkg.Read(c.cfg.CategoryPath)
	if err != nil {
		return err
	}

	for ind, ctg_obg := range data {
		var category = cast.ToStringMap(ctg_obg)
		if cast.ToString(category["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = pkg.Write(c.cfg.CategoryPath, data)
	if err != nil {
		return err
	}

	return nil
}
