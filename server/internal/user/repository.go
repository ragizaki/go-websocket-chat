// This file defines the repository layer which communicates directly with the database.
// It is responsible for all database operations.
package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertedId int64
	createUserQuery := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, createUserQuery, user.Username, user.Password, user.Email).Scan(&lastInsertedId)

	if err != nil {
		return nil, err
	}
	user.ID = int64(lastInsertedId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	getUserQuery := `SELECT * FROM users WHERE email=$1`
	err := r.db.QueryRowContext(ctx, getUserQuery, email).Scan(&u.ID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
