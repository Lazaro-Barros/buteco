package entities_test

import (
	"testing"

	"github.com/Lazaro-Barros/buteco/command/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("Test Create user and match password", func(t *testing.T) {
		usr, err := entities.NewUser(
			"John Doe",
			"test@test.com",
			"12345678",
		)
		assert.NoError(t, err)
		assert.NotNil(t, usr)
		assert.NotEmpty(t, usr.Uuid())
		assert.Equal(t, "John Doe", usr.Name())
		assert.Equal(t, "test@test.com", usr.Email())

		assert.True(t, usr.MatchPassword("12345678"))
	})

	t.Run("Test Create user with invalid password", func(t *testing.T) {
		usr, err := entities.NewUser(
			"John Doe",
			"test@test.com",
			"showrtP",
		)
		assert.Error(t, err)
		assert.Equal(t, "a senha deve ter pelo menos 8 caracteres", err.Error())
		assert.Nil(t, usr)
	})

}
