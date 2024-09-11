package entities

import "github.com/google/uuid"

type Establishment struct {
	uuid    string
	name    string
	address string
}

func NewEstablishment(
	name string,
	address string,
) (obj *Establishment, err error) {

	obj = &Establishment{
		uuid:    uuid.New().String(),
		name:    name,
		address: address,
	}
	return
}

type IEstablishmentRepository interface {
	Save(establishment *Establishment) error
}

func (obj Establishment) Uuid() string {
	return obj.uuid
}

func (obj Establishment) Name() string {
	return obj.name
}

func (obj Establishment) Address() string {
	return obj.address
}
