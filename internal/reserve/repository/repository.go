package repository

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")

type repository struct {
	pool *pgxpool.Pool
}

var ErrAlreadyExists = errors.New("already exist")

func New(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}
