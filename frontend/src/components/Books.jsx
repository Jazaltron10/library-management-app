// src/components/Books.js
import React, { useState, useEffect } from 'react';
import { MainContainer, Wrapper,Wrap, BooksListContainer, BookListItem } from './BooksList';

const Books = () => {
  const [books, setBooks] = useState([]);
  const [newBook, setNewBook] = useState({ title: '', author: '' });
  const [selectedBook, setSelectedBook] = useState(null);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8080/books');
      const data = await response.json();
      setBooks(data);
    } catch (error) {
      console.error('Error fetching books:', error);
    }
  };

  const handleInputChange = (e) => {
    setNewBook({ ...newBook, [e.target.name]: e.target.value });
  };

  const handleAddBook = async () => {
    try {
      await fetch('http://localhost:8080/books', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newBook),
      });
      setNewBook({ title: '', author: '' });
      fetchData();
    } catch (error) {
      console.error('Error adding book:', error);
    }
  };

  const handleUpdateBook = async () => {
    if (!selectedBook) return;

    try {
      await fetch(`http://localhost:8080/books/${selectedBook.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(selectedBook),
      });
      setSelectedBook(null);
      fetchData();
    } catch (error) {
      console.error('Error updating book:', error);
    }
  };

  const handleDeleteBook = async (id) => {
    try {
      await fetch(`http://localhost:8080/books/${id}`, {
        method: 'DELETE',
      });
      fetchData();
    } catch (error) {
      console.error('Error deleting book:', error);
    }
  };

  return (
    <MainContainer>
    <Wrapper>
    
      <h2>Library Books</h2>

      {/* Add Book Form */}

        <h3>Add Book</h3>
        <form>
          <Wrap>
          <label>Title:</label>
          <input type="text" name="title" value={newBook.title} onChange={handleInputChange} />
          
          <label>Author:</label>
          <input type="text" name="author" value={newBook.author} onChange={handleInputChange} />
          
          <button type="button" onClick={handleAddBook}>Add Book</button>
          </Wrap>
        </form>

      {/* Update Book Form */}
      {selectedBook && (
        <div>
          <h3>Update Book</h3>
          <form>
            <label>Title:</label>
            <input
              type="text"
              name="title"
              value={selectedBook.title}
              onChange={(e) => setSelectedBook({ ...selectedBook, title: e.target.value })}
            />
            <label>Author:</label>
            <input
              type="text"
              name="author"
              value={selectedBook.author}
              onChange={(e) => setSelectedBook({ ...selectedBook, author: e.target.value })}
            />
            <button type="button" onClick={handleUpdateBook}>Update Book</button>
          </form>
        </div>
      )}

      {/* Book List */}
      <BooksListContainer>
              

        <ul>
          {books.map((book) => (
            <BookListItem key={book.id}>
              <span>
                {book.title} by {book.author}
              </span>
              <div>
                <button onClick={() => setSelectedBook(book)}>Edit</button>
                <button onClick={() => handleDeleteBook(book.id)}>Delete</button>
              </div>
            </BookListItem>
          ))}
        </ul>
      
      </BooksListContainer>
      </Wrapper>

    </MainContainer>
  );
};

export default Books;
