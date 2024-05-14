package database

import (
	"context"
	"errors"
	"time"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"github.com/fatih/color"
)

// AddUserToDb adds a new user to the database
func (s *Service) AddUserToDb(ctx context.Context, req *proto.User) (*proto.OK, error) {
	if req.Email == "" || req.Password == "" {
		return &proto.OK{Ok: false}, errors.New("email and password are required")
	}
	if user , err := s.getUserByEmail(ctx, req.Email); err == nil && user != nil {
		return &proto.OK{Ok: false}, errors.New("user already exists")
	}

	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.Db.WithContext(ctox).Create(&req)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	color.Blue("User added")
	return &proto.OK{Ok: true}, nil
}

func (s *Service) GetUserByID(ctx context.Context, req *proto.ID) (*proto.User, error) {
	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user proto.User
	res := s.Db.WithContext(ctxx).First(&user, "id = ?", req.Id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (s *Service) GetUserFromDb(ctx context.Context, email string) (*proto.User, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user proto.User
	res := s.Db.WithContext(ctox).First(&user, "email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (s *Service) UpdateUserInDb(ctx context.Context, req *proto.UpdateUserRequest) (*proto.OK, error) {
	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user proto.User
	if err := s.Db.WithContext(ctxx).First(&user, "id = ?", req.Id).Error; err != nil {
		return &proto.OK{Ok: false}, err
	}

	user.Name = req.Name
	user.Email = req.Email

	res := s.Db.WithContext(ctxx).Save(&user)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	return &proto.OK{Ok: true}, nil
}

func (s *Service) DeleteUserFromDb(ctx context.Context, req *proto.ID) (*proto.OK, error) {
	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res := s.Db.WithContext(ctxx).Delete(&proto.User{}, req.Id)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	return &proto.OK{Ok: true}, nil
}

func (s *Service) CreateSession(ctx context.Context, session *proto.Session) (*proto.OK, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.Db.WithContext(ctox).Create(&session)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	color.Green("Session created")
	return &proto.OK{Ok: true}, nil
}

func (s *Service) GetSessionByToken(ctx context.Context, token string) (*proto.Session, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var session proto.Session
	res := s.Db.WithContext(ctox).First(&session, "token = ?", token)
	if res.Error != nil {
		return nil, res.Error
	}
	return &session, nil
}

func (s *Service) CreateOtp(ctx context.Context, otp *proto.Otp) (*proto.OK, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.Db.WithContext(ctox).Create(&otp)
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	color.Yellow("OTP created")
	return &proto.OK{Ok: true}, nil
}

func (s *Service) GetOtpByEmail(ctx context.Context, email string) (*proto.Otp, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var otp proto.Otp
	res := s.Db.WithContext(ctox).First(&otp, "email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &otp, nil
}

func (s *Service) CleanupExpiredOtps(ctx context.Context) (*proto.OK, error) {
	ctox, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	now := time.Now().Format(time.RFC3339)
	res := s.Db.WithContext(ctox).Where("expires_at <= ?", now).Delete(&proto.Otp{})
	if res.Error != nil {
		return &proto.OK{Ok: false}, res.Error
	}
	color.Red("Expired OTPs cleaned up")
	return &proto.OK{Ok: true}, nil
}

func (s *Service) getUserByEmail(ctx context.Context, email string) (*proto.User, error) {
	ctxx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user proto.User
	res := s.Db.WithContext(ctxx).First(&user, "email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
