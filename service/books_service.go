package service

import (
	"context"
	"time"

	"github.com/ryirwansyah/clean-arch-v1/models"
)

type booksService struct {
	bookRepo       models.BooksRepository
	contextTimeout time.Duration
}

func NewBooksService(b models.BooksRepository, timeout time.Duration) models.BookService {
	return &booksService{
		bookRepo:       b,
		contextTimeout: timeout,
	}
}

func (a *booksService) GetByID(c context.Context, id int64) (res models.Books, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.bookRepo.GetByID(ctx, id)
	if err != nil {
		return models.Books{}, err
	}

	return
}
