package user

import (
	"context"
	"log"

	"authentication-service/entity"
	"authentication-service/infrastructure/repository"

	svcErr "authentication-service/usecase/error"

	"golang.org/x/crypto/bcrypt"
)

const (
	MINCOST     int = 4  // the minimum allowable cost as passed into GenerateFromPassword
	MAXCOST     int = 31 // the maximum allowable cost as passed into GenerateFromPassword
	DEFAULTCOST int = 10 // the cost that will actually be set if a cost below mincost is passed into GenerateFromPassword
)

// service interface
type Service struct {
	userConfigRepository *repository.UserRepository
}

// new service
func NewService(userConfigRepository *repository.UserRepository) *Service {
	return &Service{
		userConfigRepository: userConfigRepository,
	}
}

// create an user
// bool - whether a user create or not
// error - error object, if error occured
func (service *Service) CreateUser(ctx context.Context, userConfig *entity.User) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(userConfig.Password), DEFAULTCOST)
	if err != nil {
		log.Fatalf("error hashing password. %v", err)
		return false, &svcErr.ServiceError{
			Err:  err,
			Code: svcErr.PROCESSING_ERROR,
		}
	}
	userConfig.Password = string(hash)
	err = service.userConfigRepository.AddUserDetails(ctx, userConfig)
	if err != nil {
		log.Fatalf("error creating user details. %v", err)
		return false, &svcErr.ServiceError{
			Err:  err,
			Code: svcErr.PROCESSING_ERROR,
		}
	}
	return true, nil
}
