package repository

import (
	"fmt"

	"github.com/litonshil/crud_go_echo/pkg/database"
	"github.com/litonshil/crud_go_echo/pkg/models"
)

var db = database.GetDB()

func CreateUser(user *models.User) error {
	err := db.Create(&user).Error
	return err
}
func GetAllUsers() ([]models.User, error) {
	var all_users []models.User
	err := db.Find(&all_users).Error
	fmt.Println(all_users)
	return all_users, err

}

func GetAUsers(id int) ([]models.User, error) {
	var user []models.User
	err := db.Where("id = ?", id).Find(&user).Error
	fmt.Println("user", user)
	return user, err
}

func UpdateUser(id int, user *models.User) (*models.User, error) {

	err := db.Model(&user).Where("id = ?", id).Update(&user).Error
	fmt.Println("user", user)
	return user, err
}

func DeleteUser(id int) (error) {
	var user []models.User
	err := db.Where("id = ?", id).Delete(&user).Error
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := db.Where("email = ?", email).Find(&user).Error
	return &user, err
}