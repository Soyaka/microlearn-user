package handlers

import (
	"context"
	"errors"
	"time"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"github.com/Soyaka/microlearn-user/internal/database"
	"github.com/Soyaka/microlearn-user/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImplementUserMethods struct {
	Db    *database.Service
	Cache *database.RedisClient
	proto.UnimplementedUserServiceServer
}

func NewImplementUserMethods() *ImplementUserMethods {
	return &ImplementUserMethods{
		Db:    database.NewDatabase(),
		Cache: database.NewCache(),
	}
}

func (u *ImplementUserMethods) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.Cache.GetUserFromCache(req.Email)
	if err != nil || user == nil {
		user, err = u.Db.GetUserFromDb(ctx, req.Email)
		if err != nil || user == nil {
			return nil, err
		}
		u.Cache.AddUserToCache(user, 5*time.Minute)
	}

	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.Email, user.Name, req.Agent, user.Id)
	if err != nil {
		return nil, err
	}

	session := &proto.Session{
		Id:        uuid.New().String(),
		Email:     user.Email,
		Name:      user.Name,
		Token:     token,
		Agent:     req.Agent,
		ExpiresAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}

	_, err = u.Db.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}
	resp := &proto.Token{
		Id:    user.Id,
		Token: token,
		Email: user.Email,
	}
	return resp, nil
}

func (u *ImplementUserMethods) RegisterUser(ctx context.Context, req *proto.RegisterRequest) (*proto.OK, error) {
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

	res, err := u.Db.AddUserToDb(ctx, user)
	if err == gorm.ErrDuplicatedKey {
		return &proto.OK{Ok: false}, errors.New("user already exists")
	}

	_ = u.Cache.AddUserToCache(user, 5*time.Minute)

	return res, nil
}

func (u *ImplementUserMethods) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.OK, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := u.Db.UpdateUserInDb(ctx, req)

	if err != nil {
		return &proto.OK{Ok: false}, err

	}

	_ = u.Cache.DeleteUserFromCache(req.Id)

	return res, nil
}

func (u *ImplementUserMethods) GetUser(ctx context.Context, req *proto.ID) (*proto.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := u.Db.GetUserByID(ctx, req)

	if err != nil {
		return &proto.User{}, err
	}
	return res, nil
}
func (u *ImplementUserMethods) DeleteUser(ctx context.Context, req *proto.ID) (*proto.OK, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := u.Db.DeleteUserFromDb(ctx, req)

	if err != nil {
		return &proto.OK{Ok: false}, err
	}

	//destroy cache key and session

	_ = u.Cache.DeleteUserFromCache(req.Id)

	return res, nil
}
func (u *ImplementUserMethods) VerifyToken(ctx context.Context, req *proto.Token) (*proto.Claims, error) {
	if req.Token == "" {
		return &proto.Claims{}, nil
	}
	claims, err := utils.ValidateToken(req.Token)

	if err != nil {
		return &proto.Claims{}, err
	}

	protoClaims := &proto.Claims{
		ID:     claims.ID,
		Email:  claims.Email,
		Name:   claims.Name,
		Agent:  claims.Agent,
		Iat:    claims.RegisteredClaims.IssuedAt.Time.Format(time.RFC3339),
		Exp:    claims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339),
		Iss:    claims.RegisteredClaims.Issuer,
		Sub:    claims.RegisteredClaims.Subject,
		Jti:    claims.RegisteredClaims.ID,
		Nbf:    claims.RegisteredClaims.NotBefore.Time.Format(time.RFC3339),
		UserID: claims.UserID,
	}

	return protoClaims, nil

}

func (u *ImplementUserMethods) Logout(ctx context.Context, req *proto.Token) (*proto.OK, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	session, err := u.Db.GetSessionByToken(ctx, req.GetToken())
	if err != nil {
		return &proto.OK{Ok: false}, err
	}

	ok, err := u.Db.DeleteSession(ctx, session)
	if err != nil {
		return &proto.OK{Ok: ok.GetOk()}, err
	}

	return &proto.OK{Ok: true}, nil
}

func (u *ImplementUserMethods) RefreshToken(ctx context.Context, req *proto.Token) (*proto.Token, error) {
	if req.Token == "" {
		return &proto.Token{}, nil
	}
	claims, err := utils.ValidateToken(req.Token)

	if err != nil {
		return &proto.Token{}, err
	}
	token, err := utils.GenerateToken(claims.Email, claims.Name, claims.Agent, claims.UserID)

	if err != nil {
		return &proto.Token{}, err
	}
	resp := &proto.Token{
		Id:    req.Id,
		Token: token,
		Email: req.Email,
	}

	return resp, nil
}
func (u *ImplementUserMethods) CreateSession(ctx context.Context, req *proto.Session) (*proto.OK, error) {
	return &proto.OK{Ok: true}, nil
}

func (u *ImplementUserMethods) GetSessionByToken(ctx context.Context, req *proto.Token) (*proto.Session, error) {
	return &proto.Session{}, nil
}

// OTP Methods

func (u *ImplementUserMethods) CreateOtp(ctx context.Context, req *proto.CreateOtpRequest) (*proto.OK, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	otp := &proto.Otp{
		Id:        uuid.New().String(),
		Email:     req.GetEmail(),
		Otp:       utils.GenerateOtp(),
		ExpiresAt: time.Now().Add(10 * time.Minute).Format(time.RFC3339),
	}

	res, err := u.Db.CreateOtp(ctx, otp)
	if err != nil {
		return &proto.OK{Ok: false}, err
	}

	// TODO: send otp via email

	return res, nil
}

func (u *ImplementUserMethods) VerifyOtp(ctx context.Context, req *proto.VerifyOtpRequest) (*proto.OK, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return u.Db.VerifyOtp(ctx, req)
}

func (u *ImplementUserMethods) GetOtpById(ctx context.Context, req *proto.ID) (*proto.Otp, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return u.Db.GetOtpById(ctx, req.Id)

}

func (u *ImplementUserMethods) CleanupExpiredOtps(ctx context.Context, req *proto.OK) (*proto.OK, error) {

	return u.Db.CleanupExpiredOtps(ctx)
}
