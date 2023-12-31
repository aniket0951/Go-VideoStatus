// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createAdminUser = `-- name: CreateAdminUser :one
insert into users (
    name,
    email,
    contact,
    password,
    user_type,
    is_account_active
) values (
    $1,$2,$3,$4,$5,$6
) returning id, name, email, contact, password, user_type, is_account_active, created_at, updated_at
`

type CreateAdminUserParams struct {
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Contact         string       `json:"contact"`
	Password        string       `json:"password"`
	UserType        string       `json:"user_type"`
	IsAccountActive sql.NullBool `json:"is_account_active"`
}

func (q *Queries) CreateAdminUser(ctx context.Context, arg CreateAdminUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createAdminUser,
		arg.Name,
		arg.Email,
		arg.Contact,
		arg.Password,
		arg.UserType,
		arg.IsAccountActive,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserType,
		&i.IsAccountActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :execresult
delete from users
where ID = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, id)
}

const getUser = `-- name: GetUser :one
select id, name, email, contact, password, user_type, is_account_active, created_at, updated_at from users
where ID = $1 limit 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserType,
		&i.IsAccountActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, name, email, contact, password, user_type, is_account_active, created_at, updated_at from users
where email = $1 limit 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserType,
		&i.IsAccountActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select id, name, email, contact, password, user_type, is_account_active, created_at, updated_at from users
order by ID 
limit $1
offset $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]Users, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Contact,
			&i.Password,
			&i.UserType,
			&i.IsAccountActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserAccountActive = `-- name: UpdateUserAccountActive :one
update users
set is_account_active = $2
where ID = $1
returning id, name, email, contact, password, user_type, is_account_active, created_at, updated_at
`

type UpdateUserAccountActiveParams struct {
	ID              uuid.UUID    `json:"id"`
	IsAccountActive sql.NullBool `json:"is_account_active"`
}

func (q *Queries) UpdateUserAccountActive(ctx context.Context, arg UpdateUserAccountActiveParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, updateUserAccountActive, arg.ID, arg.IsAccountActive)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Contact,
		&i.Password,
		&i.UserType,
		&i.IsAccountActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
