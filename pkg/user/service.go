package user

import "strconv"

type UserService struct {
	userRepo UserRepo
}

func New(uRepo UserRepo) *UserService {
	return &UserService{uRepo}
}

func (u *UserService) CreateUser(newUser User) (User, error) {
	err := u.userRepo.CreateUser(newUser)
	return newUser, err
}

func (u *UserService) GetUser(idStr string) (User, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return User{}, err
	}
	return u.userRepo.GetUser(id)
}
