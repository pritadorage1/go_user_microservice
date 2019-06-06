package main

import (
	"context"

	"go_user_microservice/config"
	"go_user_microservice/server"
)

func init() {
	config.Load()
}

func main() {
	conf := config.Gateway()
	ctx := context.Background()
	server.StartRESTGateway(ctx, conf)
}
