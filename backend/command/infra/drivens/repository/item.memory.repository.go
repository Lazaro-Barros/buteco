package repository

import entities "github.com/Lazaro-Barros/buteco/command/domain/entities/item"

type ItemMemoryRepository struct {
	Items map[string]entities.Item
}

func (i ItemMemoryRepository) Save(item *entities.Item) error {
	i.Items[item.Uuid()] = *item
	return nil
}
