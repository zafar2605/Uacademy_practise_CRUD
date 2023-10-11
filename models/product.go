package models

type Product struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type CreateProduct struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type UpdateProduct struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	CategoryID   string `json:"category_id"`
}

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type GetListProductRequests struct {
	Offset int
	Limit  int
}

type GetListProductResponse struct {
	Count    int
	Products []*Product
}
