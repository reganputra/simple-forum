package membership

import (
	"context"
	"database/sql"
	"errors"
	"simple-forum/internal/model/membership"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*membership.UserModel, error) {
	query := "SELECT id, email, username, created_at, updated_at, created_by, updated_by FROM users WHERE " +
		"email = ? OR username = ?"

	row := r.db.QueryRowContext(ctx, query, email, username)
	var user membership.UserModel

	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy,
		&user.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, user *membership.UserModel) error {
	query := "INSERT INTO users (email, username, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}
