package jsondb

import (
	"os"
	"practise_CRUD/config"
	"practise_CRUD/models"
	"practise_CRUD/pkg"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

type UserRepo struct {
	cfg  *config.Config
	file *os.File
}

func NewUserRepo(cfg *config.Config, file *os.File) *UserRepo {
	return &UserRepo{
		file: file,
		cfg:  cfg,
	}
}

func (u *UserRepo) Create(req *models.CreateUser) (*models.User, error) {

	data, err := pkg.Read(u.cfg.UserPath)
	if err != nil {
		return nil, err
	}

	var resp = &models.User{
		Id:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Balance:   req.Balance,
	}

	data = append(data, resp)
	err = pkg.Write(u.cfg.UserPath, data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) GetByID(req *models.UserPrimaryKey) (*models.User, error) {

	data, err := pkg.Read(u.cfg.UserPath)
	if err != nil {
		return nil, err
	}

	for _, UserObject := range data {
		var user_oder = cast.ToStringMap(UserObject)

		if cast.ToString(user_oder["id"]) == req.Id {
			return &models.User{
				Id:        cast.ToString(user_oder["id"]),
				FirstName: cast.ToString(user_oder["firs_name"]),
				LastName:  cast.ToString(user_oder["last_name"]),
				Balance:   cast.ToInt(user_oder["balance"]),
			}, nil
		}
	}

	return nil, err
}

func (u *UserRepo) GetList(req *models.GetListUserRequests) (*models.GetListUserResponse, error) {

	data, err := pkg.Read(u.cfg.UserPath)
	if err != nil {
		return nil, err
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	var resp = &models.GetListUserResponse{}

	for i := req.Offset; i < req.Limit+req.Offset; i++ {
		if len(data) > i {
			var User = cast.ToStringMap(data[i])
			resp.Users = append(resp.Users, &models.User{
				Id:        cast.ToString(User["id"]),
				FirstName: cast.ToString(User["firs_name"]),
				LastName:  cast.ToString(User["last_name"]),
				Balance:   cast.ToInt(User["balance"]),
			})
		} else {
			break
		}
	}

	resp.Count = len(data)

	return resp, nil
}

func (u *UserRepo) Update(req *models.UpdateUser) (*models.User, error) {

	data, err := pkg.Read(u.cfg.UserPath)
	if err != nil {
		return nil, err
	}

	for index, ctg_obg := range data {
		var user = cast.ToStringMap(ctg_obg)
		if cast.ToString(user["id"]) == req.Id {
			user["firs_name"] = req.FirstName
			user["last_name"] = req.LastName
			user["balanca"] = req.Balance

			data[index] = user
			break
		}
	}

	err = pkg.Write(u.cfg.UserPath, data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *UserRepo) Delete(req *models.UserPrimaryKey) error {

	data, err := pkg.Read(u.cfg.UserPath)
	if err != nil {
		return err
	}

	for ind, ctg_obg := range data {
		var User = cast.ToStringMap(ctg_obg)
		if cast.ToString(User["id"]) == req.Id {
			data = append(data[:ind], data[ind+1:]...)
			break
		}
	}

	err = pkg.Write(u.cfg.UserPath, data)
	if err != nil {
		return err
	}

	return nil
}
