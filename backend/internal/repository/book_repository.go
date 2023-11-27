// internal/repository/book_repository.go
package repository

import "github.com/jazaltron10/libraryApp/backend/internal/model"

type BookRepository struct {
	books []model.Book
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) GetAllBooks() []model.Book {
	return r.books
}

func (r *BookRepository) GetBookByID(id int) (*model.Book, error) {
	// Implementation for getting a book by ID
	return nil, nil
}

func (r *BookRepository) AddBook(book model.Book) {
	r.books = append(r.books, book)
}

func (r *BookRepository) UpdateBook(id int, updatedBook model.Book) error {
	// Implementation for updating a book
	return nil
}

func (r *BookRepository) DeleteBook(id int) error {
	// Implementation for deleting a book
	return nil
}
