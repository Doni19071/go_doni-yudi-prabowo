package service

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type DBUserService struct {
	db *gorm.DB
}

func (ps DBUserService) Add(user model.User) (model.User, error) {
	tx := ps.db.Save(&user)
	err := tx.Error
	return user, err
}

func (ps DBUserService) Get() ([]model.User, error) {
	users := []model.User{}
	tx := ps.db.Find(&users)
	err := tx.Error
	return users, err
}

func NewDBUserService(db *gorm.DB) DBUserService {
	return DBUserService{
		db: db,
	}
}
