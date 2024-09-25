package product

import (
	"errors"
	ProductEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/product"
	entities "github.com/Lazaro-Barros/buteco/command/domain/entities/product"
)

// ProductService lida com os casos de uso para produtos.
type ProductService struct {
	repo entities.ProductRepository
}

// NewProductService cria um novo ProductService.
func NewProductService(repo entities.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct é o caso de uso para criar um novo produto.
func (s *ProductService) CreateProduct(
	name string,
	description string,
	value int,
	visible bool,
) (uuid string, err error) {
	product, err := entities.NewProduct(name, description, value, visible)
	if err != nil {
		return "", err
	}

	err = s.repo.Save(product)
	if err != nil {
		return "", err
	}

	return product.Uuid(), nil
}

func (s *ProductService) ListProducts() (Products map[string]ProductEntity.Product, err error) {
	return s.repo.List()
}

// UpdateProduct é o caso de uso para atualizar um produto existente.
func (s *ProductService) UpdateProduct(
	uuid string,
	name string,
	description string,
	value int,
	visible bool,
) error {
	_, err := s.repo.GetByUuid(uuid)
	if err != nil {
		return errors.New("product not found")
	}

	updatedProduct, err := entities.NewProduct(name, description, value, visible)
	if err != nil {
		return err
	}

	err = s.repo.Update(uuid, updatedProduct)
	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct é o caso de uso para excluir um produto.
func (s *ProductService) DeleteProduct(uuid string) error {
	product, err := s.repo.GetByUuid(uuid)
	if err != nil {
		return errors.New("product not found")
	}

	err = s.repo.Delete(product.Uuid())
	if err != nil {
		return err
	}

	return nil
}
