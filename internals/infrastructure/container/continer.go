package container

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/tangguhriyadi/content-service/internals/helper"
	"github.com/tangguhriyadi/content-service/internals/infrastructure/db"
)

type Container interface {
	Init() error
}

type ContainerImpl struct {
	db.Postgre
}

func NewContainer() Container {
	return &ContainerImpl{}
}

func (cont ContainerImpl) Init() error {
	var validate = validator.New()
	helper.RegisterValidation(validate)

	db := db.NewPostgreConnection()

	_, err := db.Connect()

	if err != nil {
		log.Println(err)
	}

	return nil
}
