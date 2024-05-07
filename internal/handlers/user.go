package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "main/api/proto/gen"
	"main/internal/database"
	"main/internal/utils"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UmplimentUserMethods struct {
	Db    *database.Service
	Cache *database.RedisClient
	proto.UnimplementedUserServiceServer
}

func NewUmplimentUserMethods() *UmplimentUserMethods {
	return &UmplimentUserMethods{
		Db:    database.NewDatabase(),
		Cache: database.NewCache(),
	}
}

func (U *UmplimentUserMethods) GetUser(ctx context.Context, req *proto.ID) (*proto.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := U.Db.GetUserByID(ctx, req)

	if err != nil {
		return &proto.User{}, err
	}
	return res, nil
}
func (U *UmplimentUserMethods) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.OK, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	user := &proto.User{
		Id:    uuid.New().String(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}
	pwd, err := utils.HashPassword(req.GetPassword())

	if err != nil {
		return &proto.OK{Ok: false}, err
	}
	user.Password = pwd

	res, err := U.Db.AddUserToDb(ctx, user)
	if err == gorm.ErrDuplicatedKey {
		return &proto.OK{Ok: false}, errors.New("user already exists")
	}

	_ = U.Cache.AddUserToCache(user, 5*time.Minute)

	return res, nil

}
func (U *UmplimentUserMethods) Logout(ctx context.Context, req *proto.Token) (*proto.OK, error) {
	//ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//defer cancel()

	return &proto.OK{Ok: true}, nil
}
func (U *UmplimentUserMethods) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user *proto.User

	user, err := U.Cache.GetUserFromCache(req.Email)

	if err != nil && err != redis.Nil || user == nil {
		fmt.Print("cache error")
		user, err = U.Db.GetUserFromDb(ctx, req.Email)
		if err != nil {
			return &proto.Token{}, err
		}
		_ = U.Cache.AddUserToCache(user, 5*time.Minute)
	}

	if !utils.VerifyPassword(user.Password, req.Password) {
		return &proto.Token{}, err
	}
	token, err := utils.GenerateToken(user.Email, user.Name, user.Id)

	if err != nil {
		return &proto.Token{}, err
	}

	return &proto.Token{Token: token}, nil

}

func (*UmplimentUserMethods) RefreshToken(ctx context.Context, req *proto.Token) (*proto.Token, error) {
	if req.Token == "" {
		return &proto.Token{}, nil
	}
	_, err := utils.ValidateToken(req.Token)

	if err != nil {
		return &proto.Token{}, err
	}
	token, err := utils.RefreshToken(req.Token)

	if err != nil {
		return &proto.Token{}, err
	}
	return &proto.Token{Token: token}, nil
}

func (*UmplimentUserMethods) VerifyToken(ctx context.Context, req *proto.Token) (*proto.OK, error) {

	if req.Token == "" {
		return &proto.OK{}, nil
	}
	_, err := utils.ValidateToken(req.Token)

	if err != nil {
		return &proto.OK{Ok: false}, err
	}

	return &proto.OK{Ok: true}, nil
}
func (U *UmplimentUserMethods) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.OK, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return U.Db.UpdateUserInDb(ctox, req)
}

func (U *UmplimentUserMethods) DeleteUser(ctx context.Context, req *proto.ID) (*proto.OK, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return U.Db.DeleteUserFromDb(ctox, req)
}
