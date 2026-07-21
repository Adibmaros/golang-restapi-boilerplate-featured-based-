package user

import (
	"golang-restapi-big-structure/internal/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers() ([]User, error)
	GetUserByID(userID uint) (*User, error)
	RegisterUser(req RegisterUserDTO) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetUserByID(userID uint) (*User, error) {
	user, err := s.repo.FindByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) RegisterUser(req RegisterUserDTO) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	input := &User{
		Name:     req.Name,
		Email:    req.Email,
		Role:     jwt.RoleUser,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(input); err != nil {
		return nil, err
	}

	return input, nil
}
