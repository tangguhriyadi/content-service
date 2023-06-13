package grpc

import (
	"context"
	"log"

	pb "github.com/tangguhriyadi/grpc-user/user"
	"google.golang.org/grpc"
)

type Grpc interface {
	GetUserData(id int32) (*pb.ResponseMessage, error)
	GetUserById(client pb.UserServiceClient, id int32) (*pb.ResponseMessage, error)
}

type GrpcImpl struct {
}

func NewGrpcDial() Grpc {
	return &GrpcImpl{}
}

func (g GrpcImpl) GetUserById(client pb.UserServiceClient, id int32) (*pb.ResponseMessage, error) {
	req := pb.RequestMessage{Id: id}
	c := context.Background()
	data, err := client.GetUserById(c, &req)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (g GrpcImpl) GetUserData(id int32) (*pb.ResponseMessage, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dialing")
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// var c *fiber.Ctx
	// ctx := c.Context()

	data, err := g.GetUserById(client, 1)
	if err != nil {
		return nil, err
	}

	return data, nil

}
