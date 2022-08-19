package repository

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/litonshil/crud_go_echo/pkg/domain"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

type dbs struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// NewUsersRepository will create an object that represent the User.Repository implementations
func NewUsersRepository(dbc *gorm.DB, redis *redis.Client) domain.IUsersRepo {
	return &dbs{
		DB:    dbc,
		Redis: redis,
	}
}

func (db *dbs) CreateUser(user *models.User) error {
	fmt.Println(user)
	_, errExist := db.GetUserByEmail(user.Email)
	if errExist == nil {
		return errors.New("email exist")
	}
	err := db.DB.Create(&user).Error
	return err
}

func (db *dbs) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := db.DB.Where("email = ?", email).Find(&user).Error
	return &user, err
}

func (db *dbs) GetUsers() ([]models.User, error) {
	var all_users []models.User
	err := db.DB.Find(&all_users).Error
	fmt.Println(all_users)
	return all_users, err
}

func (db *dbs) GetUserById(id int) (models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).Find(&user).Error
	return user, err
}

func checkEmptyUserField(user *models.User, old_user models.User) *models.User {
	if user.Name == "" {
		user.Name = old_user.Name
	}
	if user.Address == "" {
		user.Address = old_user.Address
	}
	if user.Email == "" {
		user.Email = old_user.Email
	}
	if user.Type == "" {
		user.Type = old_user.Type
	}
	if user.Password == "" {
		user.Password = old_user.Password
	}
	return user
}

func (db *dbs) UpdateUser(id int, user *models.User, existed_user models.User) (*models.User, error) {

	user.Id = id
	checkedUser := checkEmptyUserField(user, existed_user)

	err := db.DB.Model(&user).Where("id = ?", id).Update(&checkedUser).Error
	return user, err
}

func (db *dbs) DeleteUser(id int) error {
	var user models.User
	err := db.DB.Where("id = ?", id).Delete(&user).Error
	return err
}
