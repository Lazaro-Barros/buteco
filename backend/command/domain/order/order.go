package order

import (
	"errors"

	"github.com/google/uuid"

	itemEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/item"
	productEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/product"
)

type Order struct {
	uuid			string
	order_number	int
	total			int
	payed			bool
	items			[]itemEntity.Item
}

func NewOrder(
	order_number int,
) (obj *Order, err error) {
	obj = &Order{
		uuid:			uuid.New().String(),
		order_number:	order_number,
		total: 			0,
		payed: 			false,
	}
	return
}

func (o *Order) Uuid() string {
	return o.uuid
}

func (o *Order) Total() int {
	o.total = 0
	for _, item := range o.items {
		o.total += item.Price()
	}
	return o.total
}

func (o *Order) Payed() bool {
	return o.payed
}

func (o *Order) OrderNumber() int {
	return o.order_number
}

func (o *Order) Items() []itemEntity.Item {
	return o.items
}

func (o *Order) AddItem(product productEntity.Product) error {
	if o.payed {
		return errors.New("cannot add product to a paid order")
	}
	if !product.Visible() {
		return errors.New("cannot add a invisible product")
	}
	item, _ := itemEntity.NewItem(
		product.Name(),
		product.Value(),
	)
	o.items = append(o.items, *item)
	return nil
}

func (o *Order) Pay() {
	o.payed = true
}
