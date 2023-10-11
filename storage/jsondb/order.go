package jsondb

import (
	"os"
	"practise_CRUD/config"
	"practise_CRUD/models"
	"practise_CRUD/pkg"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type OrderRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewOrderRepo(cfg *config.Config, file *os.File) *OrderRepo {
	return &OrderRepo{
		file: file,
		cfg:  cfg,
	}
}

func (o *OrderRepo) Create(req *models.CreateOrder) (*models.Order, error) {

	data, err := pkg.Read(o.cfg.OrderPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.Order{
		Id:        uuid.New().String(),
		UserId:    req.UserId,
		Sum:       req.Sum,
		SumCount:  req.SumCount,
		DateTime:  req.DateTime,
		Status:    req.Status,
		ProductId: req.ProductId,
	}

	data = append(data, resp)
	err = pkg.Write(o.cfg.OrderPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OrderRepo) GetByID(req *models.OrderPrimaryKey) (*models.Order, error) {

	data, err := pkg.Read(o.cfg.OrderPath)
	if err != nil {
		return nil, err
	}

	for _, OrderObject := range data {
		var order = cast.ToStringMap(OrderObject)

		if cast.ToString(order["id"]) == req.Id {
			return &models.Order{
				Id:        cast.ToString(order["id"]),
				UserId:    cast.ToString(order["user_id"]),
				Sum:       cast.ToInt(order["sum"]),
				SumCount:  cast.ToInt(order["sum_count"]),
				DateTime:  cast.ToString(order["date_time"]),
				Status:    cast.ToString(order["status"]),
				ProductId: cast.ToString(order["product_id"]),
			}, nil
		}
	}

	return nil, err
}

func (o *OrderRepo) GetList(req *models.GetListOrderRequests) (*models.GetListOrderResponse, error) {

	data, err := pkg.Read(o.cfg.OrderPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListOrderResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var order = cast.ToStringMap(data[i])
			resp.Orders = append(resp.Orders, &models.Order{
				Id:        cast.ToString(order["id"]),
				UserId:    cast.ToString(order["user_id"]),
				Sum:       cast.ToInt(order["sum"]),
				SumCount:  cast.ToInt(order["sum_count"]),
				DateTime:  cast.ToString(order["date_time"]),
				Status:    cast.ToString(order["status"]),
				ProductId: cast.ToString(order["product_id"]),
			})
		} else {
			break
		}
	}

	resp.Count = len(data)

	return resp, nil
}

func (o *OrderRepo) Update(req *models.UpdateOrder) (*models.Order, error) {

	data, err := pkg.Read(o.cfg.OrderPath)
	if err != nil {
		return nil, err
	}

	for index, ctg_obg := range data {
		var order = cast.ToStringMap(ctg_obg)
		if cast.ToString(order["id"]) == req.Id {
			order["user_id"] = req.UserId
			order["sum"] = req.Sum
			order["sum_count"] = req.SumCount
			order["date_time"] = req.DateTime
			order["status"] = req.Status
			order["product_id"] = req.ProductId

			data[index] = order
			break
		}
	}

	err = pkg.Write(o.cfg.OrderPath, data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (o *OrderRepo) Delete(req *models.OrderPrimaryKey) error {

	data, err := pkg.Read(o.cfg.OrderPath)
	if err != nil {
		return err
	}

	for ind, ctg_obg := range data {
		var order = cast.ToStringMap(ctg_obg)
		if cast.ToString(order["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = pkg.Write(o.cfg.OrderPath, data)
	if err != nil {
		return err
	}

	return nil
}
