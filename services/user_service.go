package services

import (
	"go-gin-api/models"
)

type UserService struct {
	users []models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]models.User, 0),
	}
}

func (s *UserService) GetUsers() []models.User {
	return s.users
}

func (s *UserService) CreateUser(user models.User) models.User {
	// In a real application, you would save to a database
	user.ID = uint(len(s.users) + 1)
	s.users = append(s.users, user)
	return user
}
