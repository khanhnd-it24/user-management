package repositories

import (
	"gorm.io/gorm"
	"log"
	"user-management/core/entities"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Save(user entities.User) (entities.User, error)
	FindAllUsers() (entities.Users, error)
	FindUserByUsername(username string) (entities.User, error)
	UpdatePassword(id uint, password []byte) error
}

func New(db *gorm.DB) UserRepository {
	return userRepository{DB: db}
}

func (u userRepository) Save(user entities.User) (entities.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) FindAllUsers() (users entities.Users, err error) {
	log.Print("[UserRepository]...FindAllUsers")
	err = u.DB.Find(&users).Error
	return users, err
}

func (u userRepository) FindUserByUsername(username string) (user entities.User, err error) {
	log.Print("[UserRepository]...FindUserByUsername")
	err = u.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u userRepository) UpdatePassword(id uint, password []byte) error {
	log.Print("[UserRepository]...UpdatePassword")
	err := u.DB.Model(&entities.User{}).Where("id = ?", id).Update("password", password).Error
	return err
}
