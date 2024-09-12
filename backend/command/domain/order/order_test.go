package order_test

import (
	"testing"

	orderEntity "github.com/Lazaro-Barros/buteco/command/domain/order"
	itemEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/item"
	"github.com/stretchr/testify/assert"
)

func sut(t *testing.T, orderNumber int) *orderEntity.Order {
    order, err := orderEntity.NewOrder(orderNumber)
    assert.NoError(t, err)
    assert.NotNil(t, order)
    return order
}

func TestItem(t *testing.T) {
	t.Run("should create a new Ordem with value 0 by default", func(t *testing.T) {
		order := sut(t, 1)
        assert.Equal(t, 0, order.Total())
        assert.NotEmpty(t, order.Uuid())
	})

	t.Run("should create a new unpaid Order by default", func(t *testing.T) {
		order := sut(t, 1)
		assert.NotEmpty(t, order.Uuid())
		assert.Equal(t, false, order.Payed())
	})

	t.Run("should create a new Order with number correctly", func(t *testing.T) {
		order := sut(t, 1)
		assert.NotEmpty(t, order.Uuid())
		assert.Equal(t, 1, order.OrderNumber())
	})

	t.Run("should create a new Order without items by default", func(t *testing.T) {
		order := sut(t, 1)
		assert.NotEmpty(t, order.Uuid())
		assert.Len(t, order.Items(), 0)
	})

	t.Run("should add a new Item on Order", func(t *testing.T) {
		order := sut(t, 1)
		item, errItem := itemEntity.NewItem(
			"Café Coado",
			15000,
		)
		order.AddItem(
			*item,
		)
		assert.NoError(t, errItem)
		assert.NotEmpty(t, order.Uuid())
		assert.Len(t, order.Items(), 1)
	})

	t.Run("should update total of order when a new Item is added", func(t *testing.T) {
		order := sut(t, 1)
		assert.Equal(t, 0, order.Total())
		item, errItem := itemEntity.NewItem(
			"Café Coado",
			15000,
		)
		order.AddItem(
			*item,
		)
		assert.NoError(t, errItem)
		assert.Equal(t, 15000, order.Total())
	})

	t.Run("should pay a Order", func(t *testing.T) {
		order := sut(t, 1)
		order.Pay()
		assert.Equal(t, true, order.Payed())
	})

	t.Run("shouldnt add a Item when the Order is payed", func(t *testing.T) {
		order := sut(t, 1)
		item, errItem := itemEntity.NewItem(
			"Café Coado",
			15000,
		)
		order.Pay()
		errAddItem := order.AddItem(
			*item,
		)
		assert.NoError(t, errItem)
		assert.Error(t, errAddItem)
	})
}
