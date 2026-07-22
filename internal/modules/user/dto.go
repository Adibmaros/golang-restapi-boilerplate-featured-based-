package user

import "golang-restapi-big-structure/internal/pkg/jwt"

type RegisterUserDTO struct {
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required,email"`
	Role     jwt.Role `json:"role"`
	Password string   `json:"password" binding:"required"`
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
