package main

import (
	"fmt"
	"log"
	"practise_CRUD/config"
	"practise_CRUD/controller"
	"practise_CRUD/models"
	"practise_CRUD/storage/jsondb"
)

func main() {

	var cfg = config.Path()

	jdb, err := jsondb.NewJsonDBConnection(&cfg)
	if err != nil {
		panic(err)
	}

	con := controller.NewController(&cfg, jdb)

	//	CATEGORY

	//	CREATE
	// resp, err := con.CreateCategory(&models.CreateCategory{
	// 	Name: "AAAAAAA",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// GETBYID
	// resp, err := con.GetByIDCategory(&models.CategoryPrimaryKey{
	// 	Id: "4db823aa-1024-4594-bbfa-ce7a6555b477",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// Delete
	// err = con.DeleteCategory(&models.CategoryPrimaryKey{
	// 	Id: "a6959407-a18a-4046-961e-89f3d63f97f2",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//	UPDATE
	// resp, err := con.UpdateCategory(&models.UpdateCategory{
	// 	Id:   "f2e6032d-4263-4399-9ead-23cbc9993500",
	// 	Name: "zzzzzzzzz",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// GETLIST
	// resp, err := con.GetListCategory(&models.GetListCategoryRequests{
	// 	Offset: 0,
	// 	Limit:  10,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// non := resp.Categories
	// for _, category := range non {
	// 	fmt.Println(category)
	// }

	//	ORDER

	// resp, err := con.CreateOrder(&models.CreateOrder{
	// 	UserId:    "userId",
	// 	Sum:       2321,
	// 	SumCount:  234,
	// 	DateTime:  "sdfsf",
	// 	Status:    "status",
	// 	ProductId: "product_id",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// GETBYID
	// resp, err := con.GetByIDOrder(&models.OrderPrimaryKey{
	// 	Id: "c54a4aa3-1b80-449c-90c5-692c19a1aa65",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// Delete
	// err = con.DeleteOrder(&models.OrderPrimaryKey{
	// 	Id: "c54a4aa3-1b80-449c-90c5-692c19a1aa65",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//	UPDATE
	// resp, err := con.UpdateOrder(&models.UpdateOrder{
	// 	Id:        "6b5d8f1d-602a-425f-80ce-7b9bf47da152",
	// 	UserId:    "user_id",
	// 	Sum:       122,
	// 	SumCount:  324,
	// 	DateTime:  "date_time",
	// 	Status:    "status",
	// 	ProductId: "product_id",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// GETLIST
	// resp, err := con.GetListOrder(&models.GetListOrderRequests{
	// 	Offset: 0,
	// 	Limit:  10,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// non := resp.Orders
	// for _, category := range non {
	// 	fmt.Println(category)
	// }

	//	USER

	// 	CREATE
	// resp, err := con.CreateUser(&models.CreateUser{
	// 	FirstName: "neww",
	// 	LastName:  "newwnew",
	// 	Balance:   345,
	// })
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(resp)

	// GETBYID
	// resp, err := con.GetByIDUser(&models.UserPrimaryKey{
	// 	Id: "c54a4aa3-1b80-449c-90c5-692c19a1aa65",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// Delete
	// err = con.DeleteUser(&models.UserPrimaryKey{
	// 	Id: "c54a4aa3-1b80-449c-90c5-692c19a1aa65",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//	UPDATE
	// resp, err := con.UpdateUser(&models.UpdateUser{
	// 	Id:        "6b5d8f1d-602a-425f-80ce-7b9bf47da152",
	// 	FirstName: "sdsdsa",
	// 	LastName: "asdasda",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// GETLIST
	// resp, err := con.GetListUser(&models.GetListUserRequests{
	// 	Offset: 0,
	// 	Limit:  10,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// non := resp.Users
	// for _, user := range non {
	// 	fmt.Println(user)
	// }

	//	PRODUCT

	// CREATE
	// resp, err := con.CreateProduct(&models.CreateProduct{
	// 	Name: "Name",
	// 	Price: 234,
	// 	Discount: 234,
	// 	DiscountType: "sdfs",
	// 	CategoryID: "jfiownfoiw-fsfsdse",
	// })
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(resp)

	// GETBYID
	// resp, err := con.GetByIDProduct(&models.ProductPrimaryKey{
	// 	Id: "c54a4aa3-1b80-449c-90c5-692c19a1aa65",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// Delete
	// err = con.DeleteProduct(&models.ProductPrimaryKey{
	// 	Id: "c54a4aa3-1b80-449c-90c5-692c19a1aa65",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//	UPDATE
	// resp, err := con.UpdateProduct(&models.UpdateProduct{
	// 	Id:           "6b5d8f1d-602a-425f-80ce-7b9bf47da152",
	// 	Name:         "user_id",
	// 	Price:        3423,
	// 	Discount:     122,
	// 	DiscountType: "dsd324",
	// 	CategoryID:   "fsadasdas",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp)

	// GETLIST
	resp, err := con.GetListProduct(&models.GetListProductRequests{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		log.Println(err.Error())
	}
	non := resp.Products
	for _, product := range non {
		fmt.Println(product)
	}

}
