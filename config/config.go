package config

type Config struct {
	CategoryPath string
	OrderPath    string
	ProductPath  string
	UserPath     string
}

func Path() Config {

	var cfg Config

	cfg.CategoryPath = "./data/category.json"
	cfg.OrderPath = "./data/order.json"
	cfg.ProductPath = "./data/product.json"
	cfg.UserPath = "./data/user.json"

	return cfg
}
