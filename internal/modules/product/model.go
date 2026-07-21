package product

import "golang-restapi-big-structure/internal/modules/user"

type Product struct {
	ID     uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string     `gorm:"not null" json:"name"`
	Price  float64    `gorm:"not null" json:"price"`
	Stock  int        `gorm:"not null;default:0" json:"stock"`
	UserID uint       `gorm:"not null" json:"user_id"`
	User   *user.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}
