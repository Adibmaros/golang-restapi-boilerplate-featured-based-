package product

type Service interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(id uint) (*Product, error)
	GetProductsByUserID(userID uint) ([]Product, error)
	CreateProduct(userID uint, req CreateProductDTO) (*Product, error)
	DeleteProduct(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllProducts() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *service) GetProductByID(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetProductsByUserID(userID uint) ([]Product, error) {
	return s.repo.FindByUserID(userID)
}

func (s *service) CreateProduct(userID uint, req CreateProductDTO) (*Product, error) {
	input := &Product{
		Name:   req.Name,
		Price:  req.Price,
		Stock:  req.Stock,
		UserID: userID,
	}

	if err := s.repo.Create(input); err != nil {
		return nil, err
	}

	return input, nil
}

func (s *service) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
