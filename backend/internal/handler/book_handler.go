// internal/handler/book_handler.go
package handler

import (
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"

	"github.com/jazaltron10/libraryApp/backend/internal/model"
	"github.com/jazaltron10/libraryApp/backend/internal/repository"
)

type BookHandler struct {
	bookRepo *repository.BookRepository
}


func NewBookHandler(bookRepo *repository.BookRepository) *BookHandler {
	return &BookHandler{
		bookRepo: bookRepo,
	}
}


func (h *BookHandler) GetAllBooksHandler(c echo.Context) error {
	books := h.bookRepo.GetAllBooks()
	return c.JSON(http.StatusOK, books)
}


func (h *BookHandler) GetBookByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid book ID")
	}

	book, err := h.bookRepo.GetBookByID(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting book")
	}
	if book == nil {
		return c.String(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) AddBookHandler(c echo.Context) error {
	var book model.Book
	err := c.Bind(&book)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	h.bookRepo.AddBook(book)
	return c.String(http.StatusCreated, "Book added successfully")
}

func (h *BookHandler) UpdateBookHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid book ID")
	}

	var updatedBook model.Book
	err = c.Bind(&updatedBook)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	err = h.bookRepo.UpdateBook(id, updatedBook)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error updating book")
	}

	return c.String(http.StatusOK, "Book updated successfully")
}

func (h *BookHandler) DeleteBookHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid book ID")
	}

	err = h.bookRepo.DeleteBook(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting book")
	}

	return c.String(http.StatusOK, "Book deleted successfully")
}
