package mocks

import (
	"errors"

	"github.com/TommasoAmici/golang-postgres-github-action/pkg/models"
)

var mockUser = &models.User{
	ID:       1,
	Username: "User name",
}

// UserModel mocks the behavior of postgres.UserModel
type UserModel struct{}

// Get fetches details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, errors.New("no records in db")
	}
}
