GOPATH:=$(shell printenv GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
PROTOCMD=protoc

grpc:
	# building grpc binary
	$(GOBUILD) -o bin/grpc cmd/grpc/main.go
rest:
	# building rest gateway binary
	$(GOBUILD) -o bin/rest cmd/rest/main.go

proto:
	# # compiling protobuffer code:
	protoc --proto_path=api \
	-I /home/prita/prita/go_user_microservice/api \
	--micro_out=./usecase/user \
	--go_out=plugins=grpc:./usecase/user user.proto

	#compiling rest gateway
	protoc --proto_path=api \
	-I /home/prita/prita/go_user_microservice/api \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=api/user.yaml:./usecase/user user.proto

	# # # compiling open api info
	protoc --proto_path=api \
	-I /home/prita/prita/go_user_microservice/api \
	--swagger_out=logtostderr=true:api user.proto

	sed -i -e "s/RegisterUserServiceHandler/RegisterUserServiceMicroHandler/g" ./usecase/user/user.micro.go

	
