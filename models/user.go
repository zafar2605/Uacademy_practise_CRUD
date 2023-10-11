package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type GetListUserRequests struct {
	Offset int
	Limit  int
}

type GetListUserResponse struct {
	Count int
	Users []*User
}
