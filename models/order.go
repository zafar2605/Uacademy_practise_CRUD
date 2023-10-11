package models

type Order struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	SumCount  int    `json:"sum_count"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	ProductId string `json:"product_id"`
}

type CreateOrder struct {
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	SumCount  int    `json:"sum_count"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	ProductId string `json:"product_id"`
}

type UpdateOrder struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	SumCount  int    `json:"sum_count"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	ProductId string `json:"product_id"`
}

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type GetListOrderRequests struct {
	Offset int
	Limit  int
}

type GetListOrderResponse struct {
	Count  int
	Orders []*Order
}
