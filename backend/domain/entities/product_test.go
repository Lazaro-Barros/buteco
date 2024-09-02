package entities_test

import (
	"testing"

	"github.com/Lazaro-Barros/buteco/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("Test Create product", func(t *testing.T) {
		product, err := entities.NewProduct(
			"Café Moscão",
			"Café especial em grãos torrados 250g",
			3500,
			true,
		)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.NotEmpty(t, product.Uuid())
		assert.Equal(t, "Café Moscão", product.Name())
		assert.Equal(t, "Café especial em grãos torrados 250g", product.Description())
		assert.Equal(t, 3500, product.Value())
		assert.Equal(t, true, product.Visible())
	})
}
