package books

import (
	"context"
	"database/sql"

	"github.com/ryirwansyah/clean-arch-v1/models"
	"github.com/ryirwansyah/clean-arch-v1/helper"
)

type booksRepo struct {
	Conn *sql.DB
}

func NewBooksRepository(Conn *sql.DB) models.BooksRepository {
	return &booksRepo{Conn}
}


func (m *booksRepo) fetch(ctx context.Context, query string, args ...interface{}) (res models.Books) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return models.Books{}, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			fmt.Println("got some error")
		}
	}

	result = make([]models.Books, 0)
	for rows.Next() {
		r := models.Books{}
		authorID := int64(0)
		err = rows.Scan(
			&r.id,
			&AuthorID,
			&r.title,
			&r.release_date,
			&r.created_at,
			&r.updated_at,
		)

		if err != nil {
			return nil, err
		}

		r.Books = models.Books{
			id: authorID
		}

		result = append(result, r)
	}

	return result, nil

}

func (m *booksRepo) getOne(ctx context.Context, query string, args ...interface{}) (res models.Books) {

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return models.Books{}, err
	}

	row := stmt.QueryRowContext(ctx, args...)
	res := models.Books{}

	err = row.Scan(
		&res.id,
		&res.authors_id,
		&res.release_date,
		&res.created_at,
		&res.updated_at,
	)

	return
}


func (m *booksRepo) GetByID(ctx context.Context, id int64) (models.Books, err error) {
	query := `SELECT * FROM books WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m *NewBooksRepository) Fetch(ctx context.Context, cursor string, id int64) (res []models.Books, nextCursor string, err error) {
	query := `SELECT * FROM books WHERE created_at > ? ORDER BY created_at LIMIT ?`

	decodedCursor, err  := helper.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", models.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = helper.EncodedCursor(res[len(res) - 1].created_at)
	}
}
