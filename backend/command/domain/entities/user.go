package entities

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	uuid         string
	name         string
	email        string
	hashPassword string
}

type IUserRepository interface {
	Save(user *User) error
	Update(user *User) error
}

func NewUser(
	name string,
	email string,
	password string,
) (obj *User, err error) {
	if err := validatePassword(password); err != nil {
		return nil, err
	}
	// other validations

	obj = &User{
		uuid:         uuid.New().String(),
		name:         name,
		email:        email,
		hashPassword: hashPassword(password),
	}
	return
}

func (u *User) Uuid() string {
	return u.uuid
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) MatchPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.hashPassword), []byte(password))
	return err == nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("a senha deve ter pelo menos 8 caracteres")
	}
	// Add other validations here
	return nil
}

// HashPassword gera um hash de uma senha usando bcrypt
func hashPassword(password string) string {
	// Gera o hash da senha, o segundo parâmetro é o custo do algoritmo (valor comum é 14)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
