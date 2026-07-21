package user

import "golang-restapi-big-structure/internal/pkg/jwt"

type User struct {
	ID       uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string   `gorm:"not null" json:"name"`
	Email    string   `gorm:"not null" json:"email"`
	Role     jwt.Role `gorm:"default:user" json:"role"`
	Password string   `gorm:"not null" json:"-"`
}
