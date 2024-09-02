package entities

type IUserRepository interface {
	Save(user *User) error
	Update(user *User) error
}
