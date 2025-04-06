package interactor

import (
	"Go_Practice/domain/entity"
	"Go_Practice/usecase/port"
	"errors"
)

var (
	ErrUserAlreadyExists                 = errors.New("user already exists")
	ErrUserNotFound                      = errors.New("user not found")
	ErrUserNotWorkerNorCompany           = errors.New("user is not worker nor company")
	ErrEmailNotChanged                   = errors.New("email not changed")
	ErrEmailAlreadyUsed                  = errors.New("email already used")
	ErrYouAreNotWGSalesmanNorCompanyUser = errors.New("you are not company user nor WGSalesman")
)

type UserUseCase struct {
	userRepo   port.UserRepository
}

func NewUserUseCase(
	userRepo port.UserRepository,
) IUserUseCase {
	return &UserUseCase{
		userRepo:   userRepo,
	}
}

func (u *UserUseCase) FindById(userID string) (entity.User, error) {
	return u.userRepo.FindById(userID)
}

