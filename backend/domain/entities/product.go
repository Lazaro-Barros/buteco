package entities

import (
	"github.com/google/uuid"
)

type Product struct {
	uuid         string
	name         string
	description  string
	value				 int
	visible      bool
}

func NewProduct(
	name string,
	description string,
	value int,
	visible bool,
) (obj *Product, err error) {
	
	obj = &Product {
		uuid:         uuid.New().String(),
		name:         name,
		description:  description,
		value:        value,
		visible:      visible,
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
