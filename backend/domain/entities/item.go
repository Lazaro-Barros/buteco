package entities

import "github.com/google/uuid"

type Item struct {
	uuid  string
	name  string
	price int
}

func NewItem(
	name string,
	price int,
) (obj *Item, err error) {
	// other validations

	obj = &Item{
		uuid:  uuid.New().String(),
		name:  name,
		price: price,
	}
	return
}

type IItemRepository interface {
	Save(item *Item) error
}

func (obj Item) Uuid() string {
	return obj.uuid
}

func (obj Item) Name() string {
	return obj.name
}
