package database

import (
	"context"
	"time"

	proto "main/api/proto/gen"
	"gorm.io/gorm"
)

func (s *Service) AddUserToDb(ctx context.Context, req *proto.User) (*proto.OK, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.Db.Create(&req).WithContext(ctox)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	return &proto.OK{Ok: true}, nil
}

func (S *Service) GetUserByID(ctx context.Context, req *proto.ID) (*proto.User, error) {

	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user *proto.User

	res := S.Db.First(&user, "id = ?", req.Id).WithContext(ctxx)

	if res.Error != nil {
		return user, res.Error
	}

	return user, nil

}

func (s *Service) GetUserFromDb(ctx context.Context, email string) (*proto.User, error) {

	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user proto.User
	res := s.Db.First(&user, "email = ?", email).WithContext(ctox)
	if res.Error != nil || res.Error == gorm.ErrRecordNotFound {
		return &proto.User{}, res.Error
	}
	return &user, nil
}

func (s *Service) UpdateUserInDb(ctx context.Context, req *proto.UpdateUserRequest) (*proto.OK, error) {
	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.Db.Save(&req).WithContext(ctxx)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	return &proto.OK{Ok: true}, nil
}

func (s *Service) DeleteUserFromDb(ctx context.Context, req *proto.ID) (*proto.OK, error) {
	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.Db.Delete(&proto.User{}, req.Id).WithContext(ctxx)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	return &proto.OK{Ok: true}, nil
}
