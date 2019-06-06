module go_user_microservice

go 1.12

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

replace github.com/hashicorp/consul => github.com/hashicorp/consul/api v1.1.0

require (
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.0
	github.com/jinzhu/gorm v1.9.8
	github.com/micro/go-grpc v1.0.1
	github.com/micro/go-micro v1.5.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.4.0
	google.golang.org/grpc v1.21.1
)
