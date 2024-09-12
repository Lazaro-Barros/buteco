package order

import (
	"errors"

	"github.com/google/uuid"

	// fiquei com duvida nesse import, se essa é a melhor forma.
	item "github.com/Lazaro-Barros/buteco/command/domain/entities/item"
)

type Order struct {
	uuid			string
	order_number	int
	total			int
	payed			bool
	items			[]item.Item
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
	return o.total
}

func (o *Order) Payed() bool {
	return o.payed
}

func (o *Order) OrderNumber() int {
	return o.order_number
}

func (o *Order) Items() []item.Item {
	return o.items
}

func (o *Order) AddItem(item item.Item) error {
	if o.payed {
		return errors.New("cannot add item to a paid order")
	}
	o.items = append(o.items, item)
	o.total += item.Price()
	return nil
}

func (o *Order) Pay() {
	o.payed = true
}
