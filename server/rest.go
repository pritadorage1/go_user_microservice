package server

import (
	"flag"
	"fmt"
	"net/http"

	"context"

	"go_user_microservice/config"
	usecase "go_user_microservice/usecase/user"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

//StartRESTGateway : start rest server
func StartRESTGateway(ctx context.Context, conf config.GatewayConfig) error {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Println("rest", conf)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	address := fmt.Sprintf("%s:%d", conf.Host, conf.Grpcport)
	endpoint := flag.String("endpoint", address, "usecase address")

	err := usecase.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)

	if err != nil {
		return err
	}

	server := fmt.Sprintf(":%d", conf.Port)

	fmt.Println("Server started at port:", server)
	return http.ListenAndServe(server, mux)
}
