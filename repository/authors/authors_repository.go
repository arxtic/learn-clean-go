package authors

import (
	"context"
	"database/sql"

	"github.com/ryirwansyah/clean-arch-v1/models"
)

type authorRepo struct {
	Conn *sql.DB
}

func NewAuthorRepository(Conn *sql.DB) models.AuthorRepository {
	return &authorRepo{Conn}
}

func (m *authorRepo) getOne(ctx Context, query string, args ...interface{}) (res models.Authors, err error) {
	/*
		define statement
	*/
	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		return models.Authors{}, err
	}

	row := stmt.QueryRowContext(ctx, args...)
	res := models.Authors{}

	err = row.Scan(
		&res.id,
		&res.full_name,
		&res.twitter,
		&res.created_at,
		&res.updated_at,
	)

	return
}

func (m *authorRepo) GetByID(ctx context.Context, id int64) (models.Authors , err error) {
	query := `SELECT id, full_name, twitter, created_at, updated_at FROM authors WHERE id=?`
	return m.getOne(ctx, query, id)
}