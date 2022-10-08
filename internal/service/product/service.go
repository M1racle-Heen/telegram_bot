package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	return &allProducts[idx], nil
}

func (s *Service) Set(elem string) (*Product, error) {
	allProducts = append(allProducts, Product{elem})
	return &allProducts[len(allProducts)-1], nil
}
