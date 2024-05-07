package database

import (
	"fmt"
	"os"

	//proto "github.com/Soyaka/user/api/proto/gen"kan

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func NewDatabase() *Service {
	database := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	username := os.Getenv("POSTGRES_USER")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("host=%s port=%s  user=%s password=%s dbname=%s  sslmode=disable", host, port, username, password, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&proto.User{})
	return &Service{Db: db}
}
