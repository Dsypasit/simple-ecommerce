package user

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUserRepo struct {
	user User
}

func (r *MockUserRepo) CreateUser(user User) error {
	r.user = user
	return nil
}
func (r *MockUserRepo) GetUser(id int) (User, error) { return r.user, nil }

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string
		newUser  User
		expected User
	}{
		{
			"create user success",
			User{
				Name:  "test",
				Age:   10,
				Email: "a@gmail.com",
			},
			User{
				Name:  "test",
				Age:   10,
				Email: "a@gmail.com",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uRepo := MockUserRepo{}

			uService := New(&uRepo)

			actual, err := uService.CreateUser(test.newUser)

			assert.Equal(t, test.expected, actual)
			assert.Nil(t, err)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name        string
		newUser     User
		expected    User
		expectedErr error
		id          string
	}{
		{
			name: "get user success",
			newUser: User{
				Name:  "test",
				Age:   10,
				Email: "a@gmail.com",
			},
			expected: User{
				Name:  "test",
				Age:   10,
				Email: "a@gmail.com",
			},
			id:          "10",
			expectedErr: nil,
		},
		{
			name: "failed to get user when id is not integer",
			newUser: User{
				Name:  "test",
				Age:   10,
				Email: "a@gmail.com",
			},
			expected:    User{},
			id:          "str",
			expectedErr: &strconv.NumError{"Atoi", "str", strconv.ErrSyntax},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uRepo := MockUserRepo{}

			uService := New(&uRepo)

			uService.CreateUser(test.newUser)

			actual, err := uService.GetUser(test.id)

			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
