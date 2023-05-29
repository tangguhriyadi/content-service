package container

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/db"
	"gorm.io/gorm"
)

type Container struct {
	Validator *validator.Validate
	Postgre   *gorm.DB
}

func InitContainer() (cont *Container) {
	// setup validation
	validate := validator.New()

	// setup postgre
	postgre := db.NewPostgreConnection()

	_, err := postgre.Connect()

	if err != nil {
		log.Println(err)
	}

	var db *gorm.DB

	return &Container{
		Validator: validate,
		Postgre:   db,
	}

}
