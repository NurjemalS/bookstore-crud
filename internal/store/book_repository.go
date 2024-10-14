package bookRepository

import (
	"context"
	"bookstore/internal/store/pgx"
	"bookstore/internal/models"
)

func AddBook(ctx context.Context, b models.Book) error {
	_, err := pgx.GetDBPool().Exec(ctx, "INSERT INTO books (title, author, quantity, price) VALUES ($1, $2, $3, $4)", b.Title, b.Author, b.Quantity, b.Price)
	return err
}

func GetAllBooks(ctx context.Context) ([]models.Book, error) {
	rows, err := pgx.GetDBPool().Query(ctx, "SELECT id, title, author, quantity, price FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity, &b.Price); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func GetBookByID(ctx context.Context, id string) (models.Book, error) {
	var book models.Book
	err := pgx.GetDBPool().QueryRow(ctx, "SELECT id, title, author, quantity, price FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Quantity, &book.Price)
	return book, err
}

func UpdateBook(ctx context.Context, id string, updatedBook models.Book) error {
	_, err := pgx.GetDBPool().Exec(ctx, "UPDATE books SET title=$1, author=$2, quantity=$3, price=$4 WHERE id=$5",
		updatedBook.Title, updatedBook.Author, updatedBook.Quantity, updatedBook.Price, id)
	return err
}

func DeleteBook(ctx context.Context, id string) error {
	_, err := pgx.GetDBPool().Exec(ctx, "DELETE FROM books WHERE id=$1", id)
	return err
}
