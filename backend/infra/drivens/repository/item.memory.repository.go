package repository

import "github.com/Lazaro-Barros/buteco/domain/entities"

type ItemMemoryRepository struct {
	Items map[string]entities.Item
}

func (i ItemMemoryRepository) Save(item *entities.Item) error {
	i.Items[item.Uuid()] = *item
	return nil
}
