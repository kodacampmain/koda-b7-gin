package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b7-gin/internal/model"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) AddNewUser(ctx context.Context, username, hashedPwd string) (model.User, error) {
	sql := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username, password, created_at"
	args := []any{username, hashedPwd}
	var user model.User
	if err := a.db.QueryRow(ctx, sql, args...).Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *AuthRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	sql := "SELECT id, password FROM users WHERE username = $1"
	args := []any{username}
	var user model.User
	if err := a.db.QueryRow(ctx, sql, args...).Scan(&user.Id, &user.Password); err != nil {
		return model.User{}, err
	}
	return user, nil
}
