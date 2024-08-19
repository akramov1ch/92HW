package services

import (
	"errors"
	"92HW/models"
	"92HW/utils"
)

var users = map[string]string{}

func RegisterUser(user models.User) error {
	if _, exists := users[user.Username]; exists {
		return errors.New("user already exists")
	}
	users[user.Username] = user.Password
	return nil
}

func LoginUser(user models.User) (string, error) {
	storedPassword, exists := users[user.Username]
	if !exists || storedPassword != user.Password {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
