package models

import (
	"context"
	"time"
)

type Books struct {
	id           int64     `json:id`
	authors_id   int64     `json:authors_id`
	title        string    `json:title`
	release_date string    `json:release_date`
	created_at   time.Time `json:"created_at"`
	updated_at   time.Time `json:"updated_at"`
}

type BookService interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Books, string, error)
	GetByID(ctx context.Context, id int64) (Books, error)
	Update(ctx context.Context, id int64) (Books, error)
	Create(ctx context.Context, *Books) error
	Delete(ctx context.Context, id int64) error
}

type BooksRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Books, string, error)
	GetByID(ctx context.Context, id int64) (Books, error)
	Update(ctx context.Context, id int64) (Books, error)
	Create(ctx context.Context, b *Books) error
	Delete(ctx context.Context, id int64) error
}
