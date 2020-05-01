package models

import (
	"context"
	"time"
)

type Authors struct {
	id           int64     `json:id`
	full_name    string    `json:full_name`
	phone_number string    `json:phone_number`
	twitter      string    `json:twitter`
	created_at   time.Time `json:"created_at"`
	updated_at   time.Time `json:"updated_at"`
}

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Authors, error)
}
