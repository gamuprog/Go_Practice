package repository

import (
	"errors"

	"gorm.io/gorm"

	"Go_Practice/adapter/database/model"
	"Go_Practice/domain/entity"
	"Go_Practice/usecase/interactor"
	"Go_Practice/usecase/port"
)
type UserRepository struct {
	db   *gorm.DB
	ulid port.ULID
}

func NewUserRepository(
	db *gorm.DB,
	ulid port.ULID,
) port.UserRepository {
	return &UserRepository{db: db, ulid: ulid}
}


func (r *UserRepository) FindById(userID string) (entity.User, error) {
	res := model.User{}
	// TODO: SQL書く
	err := r.db.
		Where("user_id = ?", userID).
		First(&res).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, interactor.ErrUserNotFound
	}
	if err != nil {
		return entity.User{}, err
	}
	return res.Entity(), nil
}

