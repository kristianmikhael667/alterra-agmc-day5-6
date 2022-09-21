package database

import (
	"main/config"
	"main/models"
)

// get all user
func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// get single user
func GetSingleUser(id string) (interface{}, error) {
	var users []models.User

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// delete user
func DeleteUser(id string) (interface{}, error) {

	var users []models.User

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func EditUser(id string, biod *models.User) (interface{}, error) {

	// var users []User
	var user []models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Model(&user).Where("id = ?", id).Updates(models.User{Name: biod.Name, Email: biod.Email, Password: biod.Password}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// create new user
func CreateUser(name string, email string, password string) (interface{}, error) {

	var users = []models.User{{Name: name, Email: email, Password: password}}

	if err := config.DB.Create(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Login User
func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
