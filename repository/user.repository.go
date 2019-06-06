package repository

import (
	pb "go_user_microservice/usecase/user"

	"github.com/jinzhu/gorm"
)

//Repository interface
type Repository interface {
	Create(user *pb.User) error
	GetAll() ([]*pb.User, error)
}

//UserRepository struct
type UserRepository struct {
	Db *gorm.DB
}

//Create func
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetAll func
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
