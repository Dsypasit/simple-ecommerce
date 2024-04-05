package user

type UserRepo interface {
	CreateUser(User) error
	GetUser(int) (User, error)
}
