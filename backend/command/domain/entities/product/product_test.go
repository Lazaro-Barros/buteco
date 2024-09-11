package entities_test

import (
	"testing"

	entities "github.com/Lazaro-Barros/buteco/command/domain/entities/product"

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

	t.Run("Test Create product with invalid value", func(t *testing.T) {
		product, err := entities.NewProduct(
			"Café Moscão",
			"Café especial em grãos torrados 250g",
			0,
			true,
		)
		assert.Error(t, err)
		assert.Equal(t, entities.ErrInvalidValue.Error(), err.Error())
		assert.Nil(t, product)
	})

	t.Run("Test Create product with invalid name", func(t *testing.T) {
		product, err := entities.NewProduct(
			"",
			"Café especial em grãos torrados 250g",
			3500,
			true,
		)
		assert.Error(t, err)
		assert.Equal(t, entities.ErrInvalidName.Error(), err.Error())
		assert.Nil(t, product)
	})

	t.Run("Test Create product with invalid description", func(t *testing.T) {
		product, err := entities.NewProduct(
			"Café Moscão",
			"",
			3500,
			true,
		)
		assert.Error(t, err)
		assert.Equal(t, entities.ErrInvalidDescription.Error(), err.Error())
		assert.Nil(t, product)
	})
}
