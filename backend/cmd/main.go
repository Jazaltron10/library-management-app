// cmd/main.go
package main

import (
	"os"

	"github.com/jazaltron10/libraryApp/backend/internal/handler"
	"github.com/jazaltron10/libraryApp/backend/internal/model"
	"github.com/jazaltron10/libraryApp/backend/internal/repository"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Enable CORS
	e.Use(middleware.CORS())

	// Create a new BookRepository and BookHandler
	bookRepo := repository.NewBookRepository()

	// Initialize repository with sample data (for testing purposes)
	bookRepo.AddBook(model.Book{ID: 1, Title: "Sample Book 1", Author: "Author 1"})
	bookRepo.AddBook(model.Book{ID: 2, Title: "Sample Book 2", Author: "Author 2"})

	bookHandler := handler.NewBookHandler(bookRepo)

	// Define routes
	e.GET("/books", bookHandler.GetAllBooksHandler)
	e.GET("/books/:id", bookHandler.GetBookByIDHandler)
	e.POST("/books", bookHandler.AddBookHandler)
	e.PUT("/books/:id", bookHandler.UpdateBookHandler)
	e.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
