package entities

import "fmt"

type Status string

const (
	Open   Status = "open"
	Closed Status = "closed"
)

type Order struct {
	uuid   string
	itens  []Item
	status Status
}

func (o Order) RemoveItem(uuidItem string) error {
	if o.status == Open {
		return fmt.Errorf("order is closed")
	}
	for i, item := range o.itens {
		if item.Uuid() == uuidItem {
			o.itens = append(o.itens[:i], o.itens[i+1:]...)
			return nil
		}
	}
	return nil
}
