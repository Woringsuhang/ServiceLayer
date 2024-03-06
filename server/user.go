package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/Woringsuhang/ServiceLayer/model"
	"github.com/Woringsuhang/mess/user"
)

type Servers struct {
	user.UnimplementedUserServer
}

func (Servers) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {

	mobile := request.Mobile
	password := request.Password
	fmt.Println(mobile)
	if mobile == "" {
		return &user.LoginResponse{}, errors.New("手机号不能为空")
	}

	if password == "" {
		return &user.LoginResponse{}, errors.New("password不能为空")
	}

	mobileUser, err := model.GetMobileUser(mobile)

	if err != nil {
		return nil, err
	}

	data := &user.Users{
		Id:       int64(mobileUser.ID),
		Username: mobileUser.Username,
		Password: mobileUser.Password,
		Mobile:   mobileUser.Mobile,
		Age:      int64(mobileUser.Age),
		Sex:      user.Sex(mobileUser.Sex),
		Address:  mobileUser.Address,
	}
	return &user.LoginResponse{Data: data}, nil
}

func (Servers) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	id := request.Id
	if id == 0 {
		return nil, errors.New("id不能为0")
	}

	get, err := model.Get(id)
	if err != nil {
		return nil, err
	}
	users := &user.Users{
		Id:       int64(get.ID),
		Username: get.Username,
		Password: get.Password,
		Mobile:   get.Mobile,
		Age:      int64(get.Age),
		Sex:      user.Sex(get.Sex),
		Address:  get.Address,
	}
	return &user.GetUserResponse{
		Data: users,
	}, nil
}
func (Servers) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	mobile := request.User.Mobile
	password := request.User.Password

	if mobile == "" {
		return &user.CreateUserResponse{}, errors.New("手机号不能为空")
	}

	if password == "" {
		return &user.CreateUserResponse{}, errors.New("password不能为空")
	}
	var users = model.Users{
		Username: request.User.Username,
		Password: password,
		Mobile:   mobile,
		Sex:      int(request.User.Sex),
		Age:      int(request.User.Age),
		Address:  request.User.Address,
	}

	create, err := model.Create(&users)
	if err == nil {
		return nil, err
	}

	u := &user.Users{
		Username: create.Username,
		Password: create.Password,
		Mobile:   create.Mobile,
		Age:      int64(create.Age),
		Sex:      user.Sex(create.Sex),
		Address:  create.Address,
	}
	return &user.CreateUserResponse{Data: u}, nil
}
func (Servers) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	id := request.Id
	if id == 0 {
		return &user.DeleteUserResponse{}, errors.New("id不能为0")
	}

	err := model.DeleteUser(id)
	if err != nil {
		return &user.DeleteUserResponse{}, err
	}

	return &user.DeleteUserResponse{}, nil
}
func (Servers) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	id := request.Data.Id
	if id == 0 {
		return &user.UpdateUserResponse{}, errors.New("id不能为0")
	}
	//var users = model.Users{
	//	Username: request,
	//	Password: "",
	//	Mobile:   "",
	//	Sex:      0,
	//	Age:      0,
	//	Address:  "",
	//}
	return &user.UpdateUserResponse{}, nil
}
