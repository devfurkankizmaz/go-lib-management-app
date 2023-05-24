package service

import "github.com/devfurkankizmaz/go-lib-management-app/models"

type bookService struct {
	bookRepository models.BookRepository
}

func NewBookService(repo models.BookRepository) models.BookService {
	return &bookService{bookRepository: repo}
}

func (bs *bookService) Create(book *models.Book) error {
	err := bs.bookRepository.Create(book)
	if err != nil {
		return err
	}
	return nil
}

func (bs *bookService) FetchAllByUserID(id string, limit int, page int) ([]models.Book, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
	offset := (page - 1) * page
	books, err := bs.bookRepository.FetchAllByUserID(id, limit, offset)
	if err != nil {
		return books, err
	}
	return books, nil
}

func (bs *bookService) FetchByID(id string) (models.Book, error) {
	book, err := bs.bookRepository.FetchByID(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (bs *bookService) UpdateByID(book *models.Book, id string) error {
	err := bs.bookRepository.UpdateByID(book, id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *bookService) DeleteByID(id string) error {
	err := bs.bookRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
