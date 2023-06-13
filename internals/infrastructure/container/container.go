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
	// GrpcClient pb.UserServiceClient
}

func InitContainer() (cont *Container) {
	// setup validation
	validate := validator.New()

	// setup postgre
	postgre := db.NewPostgreConnection()

	db, err := postgre.Connect()
	if err != nil {
		log.Println(err)
	}

	// // setup gRPC
	// GRPC := grpc.NewGrpcDial()

	// grpcClient, err := GRPC.Connect()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	return &Container{
		Validator: validate,
		Postgre:   db,
		// GrpcClient: grpcClient,
	}

}
