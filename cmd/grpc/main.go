package main

import (
	"go_user_microservice/config"
	logger "go_user_microservice/middleware"
	"go_user_microservice/model"
	repo "go_user_microservice/repository"
	"go_user_microservice/server"
	handler "go_user_microservice/usecase/handler"
	pb "go_user_microservice/usecase/user"

	log "github.com/sirupsen/logrus"
)

func init() {
	config.Load()
	logger.Setup()
}

func main() {
	conf := config.Db()
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := model.CreateConnection(conf)
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(&pb.User{})
	repository := &repo.UserRepository{Db: db}
	userService := &handler.Service{Repo: repository}

	// Run the server
	if err := server.StartGRPCServer(userService); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
