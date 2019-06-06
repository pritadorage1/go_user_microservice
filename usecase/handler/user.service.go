package handler

import (
	context "context"
	fmt "fmt"
	repo "go_user_microservice/repository"
	pb "go_user_microservice/usecase/user"
)

//Service : repo service
type Service struct {
	Repo repo.Repository
}

//Create : create user
func (srv *Service) Create(ctx context.Context, req *pb.User, res *pb.UserCreateResponse) error {
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	fmt.Println("user created", req)
	return nil
}

//GetAll : get all user
func (srv *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.AllUserResponse) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
