package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type UpdateCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type GetListCategoryRequests struct {
	Offset int
	Limit  int
}

type GetListCategoryResponse struct {
	Count      int
	Categories []*Category
}
