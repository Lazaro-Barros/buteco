package entities_test

import (
	"testing"

	establishmentEntity "github.com/Lazaro-Barros/buteco/command/domain/entities/establishment"
	"github.com/stretchr/testify/assert"
)

func TestEstablishment(t *testing.T) {
	t.Run("Test Create establishment", func(t *testing.T) {
		establishment, err := establishmentEntity.NewEstablishment(
			"Bar do Pote",
			"Rua Araibú 335",
		)
		assert.NoError(t, err)
		assert.NotNil(t, establishment)
		assert.NotEmpty(t, establishment.Uuid())
		assert.Equal(t, "Bar do Pote", establishment.Name())
		assert.Equal(t, "Rua Araibú 335", establishment.Address())
	})
}
