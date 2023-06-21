package db

import (
	"fmt"
	"os"
	"time"

	"github.com/tangguhriyadi/content-service/internals/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgre interface {
	Connect() (*gorm.DB, error)
}

type PostgreImpl struct {
}

func NewPostgreConnection() Postgre {
	return &PostgreImpl{}
}

var DB *gorm.DB

func (ps PostgreImpl) Connect() (*gorm.DB, error) {
	// env := config.New()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_HOST"),
		os.Getenv("DB_POSTGRES_PORT"),
		os.Getenv("DB_POSTGRES_DBNAME"),
	)

	//start connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// database migration
	go func() {
		db.AutoMigrate(&entity.Content{})
		db.AutoMigrate(&entity.ContentCategory{})
		db.AutoMigrate(&entity.ContentType{})
		db.AutoMigrate(&entity.ContentLikeHistory{})
		db.AutoMigrate(&entity.ContentComment{})
	}()

	// rules of usage
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	DB = db

	return DB, nil

}
