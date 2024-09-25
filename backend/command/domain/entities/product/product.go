package entities

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrInvalidName        = fmt.Errorf("name is required")
	ErrInvalidValue       = fmt.Errorf("value must be greater than 0")
	ErrInvalidDescription = fmt.Errorf("description is required")
)

type Product struct {
	uuid        string
	name        string
	description string
	value       int
	visible     bool
}

type ProductRepository interface {
	Save(product *Product) error
	List() (map[string]Product, error)
	Update(uuid string, product *Product) error
	Delete(uuid string) error
	GetByUuid(uuid string) (*Product, error)
}

func NewProduct(
	name string,
	description string,
	value int,
	visible bool,
) (obj *Product, err error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if value <= 0 {
		return nil, ErrInvalidValue
	}
	if description == "" {
		return nil, ErrInvalidDescription
	}
	obj = &Product{
		uuid:        uuid.New().String(),
		name:        name,
		description: description,
		value:       value,
		visible:     visible,
	}
	return
}

func (p *Product) Uuid() string {
	return p.uuid
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Value() int {
	return p.value
}

func (p *Product) Visible() bool {
	return p.visible
}
