package entities_test

import (
	"testing"

	itemEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/item"
	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	t.Run("should create a item correctly", func(t *testing.T) {
		item, err := itemEntity.NewItem(
			"Café com leite vegetal",
			15000,
		)
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.NotEmpty(t, item.Uuid())
		assert.Equal(t, "Café com leite vegetal", item.Name())
		assert.Equal(t, 15000, item.Price())
	})
}
