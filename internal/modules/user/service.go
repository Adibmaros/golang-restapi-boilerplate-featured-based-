package user

import (
	"errors"
	"golang-restapi-big-structure/internal/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUsers() ([]User, error)
	GetUserByID(userID uint) (*User, error)
	RegisterUser(req RegisterUserDTO) (*User, error)
	LoginUser(req LoginUserDTO) (string, *User, error)
}

type service struct {
	repo       Repository
	jwtService jwt.Service
}

func NewService(repo Repository, jwtService jwt.Service) *service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
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

	role := req.Role
	if role == "" {
		role = jwt.RoleUser
	}

	input := &User{
		Name:     req.Name,
		Email:    req.Email,
		Role:     role,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(input); err != nil {
		return nil, err
	}

	return input, nil
}

func (s *service) LoginUser(req LoginUserDTO) (string, *User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
