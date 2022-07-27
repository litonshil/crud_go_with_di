package repository

import (
	"fmt"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

func (db *dbs) CreateUser(user *models.User) error {
	err := db.DB.Create(&user).Error
	return err
}

func (db *dbs) GetAllUsers() ([]models.User, error) {
	var all_users []models.User
	err := db.DB.Find(&all_users).Error
	fmt.Println(all_users)
	return all_users, err
}

func (db *dbs) GetAUsers(id int) ([]models.User, error) {
	var user []models.User
	err := db.DB.Where("id = ?", id).Find(&user).Error
	fmt.Println("user", user)
	return user, err
}

func checkEmptyUserField(user *models.User, old_user *models.User) *models.User {
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

func (db *dbs) UpdateUser(id int, user *models.User, old_user *models.User) (*models.User, error) {

	old_err := db.DB.Model(old_user).Where("id = ?", id).Find(&old_user).Error

	if old_err != nil {
		return user,old_err
	}

	user.Id = id

	checkedUser := checkEmptyUserField(user, old_user)

	err := db.DB.Model(&user).Where("id = ?", id).Update(&checkedUser).Error
	fmt.Println("user", user)
	return user, err
}

func (db *dbs) DeleteUser(id int) error {
	var user []models.User
	err := db.DB.Where("id = ?", id).Delete(&user).Error
	return err
}

func (db *dbs) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := db.DB.Where("email = ?", email).Find(&user).Error
	return &user, err
}
