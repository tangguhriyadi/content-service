package db

import (
	"fmt"
	"time"

	"github.com/tangguhriyadi/content-service/internals/config"
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

func (ps PostgreImpl) Connect() (*gorm.DB, error) {
	env := config.New()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.Get("DB_POSTGRE_USER"),
		env.Get("DB_POSTGRE_PASSWORD"),
		env.Get("DB_POSTGRE_HOST"),
		env.Get("DB_POSTGRE_PORT"),
		env.Get("DB_POSTGRE_DBNAME"),
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
	}()

	// rules of usage
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil

}
