package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (*Service) ListProduct() []Product {
	return allProducts
}
