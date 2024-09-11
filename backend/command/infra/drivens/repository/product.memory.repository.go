package repository

import ProductEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/product"

type ProductMemoryRepository struct {
	Products map[string]ProductEntity.Product
}

func NewProductMemoryRepository() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		Products: make(map[string]ProductEntity.Product),
	}
}

func (p ProductMemoryRepository) Save(product *ProductEntity.Product) error {
	p.Products[product.Uuid()] = *product
	return nil
}

func (p ProductMemoryRepository) Update(uuid string, product *ProductEntity.Product) error {
	p.Products[uuid] = *product
	return nil
}

func (p ProductMemoryRepository) Delete(uuid string) error {
	delete(p.Products, uuid)
	return nil
}

func (p ProductMemoryRepository) GetByUuid(uuid string) (*ProductEntity.Product, error) {
	product, ok := p.Products[uuid]
	if !ok {
		return nil, nil
	}
	return &product, nil
}
