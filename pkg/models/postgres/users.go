package postgres

import (
	"context"

	"github.com/TommasoAmici/golang-postgres-github-action/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// UserModel implements postgres bindings for models.UserModel
type UserModel struct {
	Pool *pgxpool.Pool
}

// Get fetches details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	s := &models.User{}

	stmt := "SELECT id, username FROM users WHERE id = $1"
	row := m.Pool.QueryRow(context.Background(), stmt, id)
	err := row.Scan(&s.ID, &s.Username)
	if err != nil {
		return nil, err
	}

	return s, nil
}
