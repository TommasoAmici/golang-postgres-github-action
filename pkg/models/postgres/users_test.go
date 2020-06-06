package postgres

import (
	"reflect"
	"testing"

	"github.com/TommasoAmici/golang-postgres-github-action/pkg/models"
)

func TestUserModel(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	pool, teardown := newTestDB(t)
	defer teardown()
	// set up database
	u := UserModel{Pool: pool}

	// Set up a suite of table-driven tests and expected results.
	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError bool
	}{
		{
			name:   "get valid guest",
			userID: 1,
			wantUser: &models.User{
				ID:       1,
				Username: "User name",
			},
			wantError: false,
		},
		{
			name:      "get Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize a connection pool to our test database, and defer a
			// call to the teardown function, so it is always run immediately
			// before this sub-test returns.

			// Call the UserModel.Get() method and check that the return value
			// and error match the expected values for the sub-test.
			user, err := u.Get(tt.userID)

			if (err != nil) != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(user, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, user)
			}
		})
	}
}
