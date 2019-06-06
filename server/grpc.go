package server

import (
	"go_user_microservice/middleware"
	pb "go_user_microservice/usecase/user"

	"github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
)

//StartGRPCServer : start grpc server
func StartGRPCServer(impl pb.UserServiceHandler) (err error) {

	// Create a new service. Optionally include some options here.
	srv := grpc.NewService(
		micro.Name("user.service"),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
		// Middleware for error logging
		micro.WrapHandler(middleware.Logger),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterUserServiceMicroHandler(srv.Server(), impl)

	return srv.Run()

}
