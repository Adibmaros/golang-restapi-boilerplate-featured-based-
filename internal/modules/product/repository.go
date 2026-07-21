package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(id uint) (*Product, error)
	FindByUserID(userID uint) ([]Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	if err := r.db.Preload("User").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindByID(id uint) (*Product, error) {
	var product Product
	if err := r.db.Preload("User").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) FindByUserID(userID uint) ([]Product, error) {
	var products []Product
	if err := r.db.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *repository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}
