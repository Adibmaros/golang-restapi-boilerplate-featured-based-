package user

import "golang-restapi-big-structure/internal/pkg/jwt"

type RegisterUserDTO struct {
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Role     jwt.Role `json:"role"`
	Password string   `json:"password" binding:"required"`
}
