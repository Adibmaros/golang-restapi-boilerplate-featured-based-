package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(u *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	if err := r.db.Select("id", "name", "email", "role").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) FindByID(id uint) (*User, error) {
	var user User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(u *User) error {
	return r.db.Create(u).Error
}
