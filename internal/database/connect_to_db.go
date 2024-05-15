package database

import (
	"fmt"
	"log"

	//proto "github.com/Soyaka/user/api/proto/gen"kan

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"github.com/Soyaka/microlearn-user/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service struct {
	Db *gorm.DB
}

func NewDatabase() *Service {
	database := utils.GetEnv("POSTGRES_DB", "microlearn")
	password := utils.GetEnv("POSTGRES_PASSWORD", "password")
	username := utils.GetEnv("POSTGRES_USER", "user")
	port := utils.GetEnv("DB_PORT", "5432")
	host := utils.GetEnv("DB_HOST", "localhost")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// err = AutoMigrate(db)
	// if err != nil {
	// 	log.Fatalf("Failed to migrate the database: %v", err)
	// 	return nil
	// }

	log.Println("Database connected and migrated successfully.")
	return &Service{Db: db}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&proto.User{}, &proto.Session{}, &proto.Otp{})
}
