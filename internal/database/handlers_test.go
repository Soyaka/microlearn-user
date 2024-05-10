package database

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"gorm.io/gorm"
)

func TestService_AddUserToDb(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *proto.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.OK
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Db: tt.fields.Db,
			}
			got, err := s.AddUserToDb(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddUserToDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddUserToDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUserByID(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *proto.ID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			S := &Service{
				Db: tt.fields.Db,
			}
			got, err := S.GetUserByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUserFromDb(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.User
		wantErr bool
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Db: tt.fields.Db,
			}
			got, err := s.GetUserFromDb(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUserFromDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUserFromDb() = %v, want %v", got, tt.want)
			}
		})
	}
}
