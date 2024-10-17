package repository

import (
	errors "bookstore/internal/app"
	"bookstore/internal/models"
	"bookstore/internal/store/pgx"
	"context"
	"fmt"
)

func AddBook(ctx context.Context, b models.Book) error {
	var isAuthorExists bool

	err := pgx.GetDBPool().QueryRow(ctx, "SELECT EXISTS(SELECT * FROM authors WHERE id = $1)", b.AuthorID).Scan(&isAuthorExists)

	fmt.Println(isAuthorExists)
	if err != nil {
		return errors.NewCustomError("Failed to check author", 500)
	}
	if !isAuthorExists {
		return errors.NewCustomError("Invalid author ID", 400)
	}

	_, err = pgx.GetDBPool().Exec(ctx, "INSERT INTO books (title, author_id, quantity, price, file_path) VALUES ($1, $2, $3, $4, $5)",
		b.Title, b.AuthorID, b.Quantity, b.Price, b.FilePath)

	if err != nil {
		return errors.NewCustomError("Failed to add book", 500)
	}

	return nil
}

func GetAllBooks(ctx context.Context) ([]models.Book, error) {
	rows, err := pgx.GetDBPool().Query(ctx, `
	SELECT b.id, b.title, b.author_id, b.quantity, b.price, b.file_path, a.id, a.name 
	FROM books b 
	LEFT JOIN authors a ON b.author_id = a.id`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		var author models.Author

		if err := rows.Scan(&b.ID, &b.Title, &b.AuthorID, &b.Quantity, &b.Price, &b.FilePath, &author.ID, &author.Name); err != nil {
			return nil, errors.NewCustomError("Error scanning book data", 500)
		}

		b.Author = author
		books = append(books, b)
	}

	return books, nil
}

func GetBookByID(ctx context.Context, id string) (models.Book, error) {
	var book models.Book
	var author models.Author

	err := pgx.GetDBPool().QueryRow(ctx, `
		SELECT b.id, b.title, b.author_id, b.quantity, b.price, b.file_path, a.id, a.name
		FROM books b
		LEFT JOIN authors a ON b.author_id = a.id
		WHERE b.id = $1`, id).
		Scan(&book.ID, &book.Title, &book.AuthorID, &book.Quantity, &book.Price, &book.FilePath, &author.ID, &author.Name)

	if err != nil {
		return book, errors.NewCustomError("Failed to retrieve book details", 500)
	}

	book.Author = author

	return book, nil
}

func UpdateBook(ctx context.Context, id string, updatedBook models.Book) error {

	var isAuthorExists bool
	err := pgx.GetDBPool().QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM authors WHERE id = $1)", updatedBook.AuthorID).Scan(&isAuthorExists)
	if err != nil {
		return errors.NewCustomError("Failed to check author", 500)
	}
	if !isAuthorExists {
		return errors.NewCustomError("Invalid author ID", 400)
	}

	result, err := pgx.GetDBPool().Exec(ctx, "UPDATE books SET title=$1, author_id=$2, quantity=$3, price=$4, file_path=$5 WHERE id=$6",
		updatedBook.Title, updatedBook.AuthorID, updatedBook.Quantity, updatedBook.Price, updatedBook.FilePath, id)

	if err != nil {
		return errors.NewCustomError("Failed to update book", 500)
	}

	if result.RowsAffected() == 0 {
		return errors.NewCustomError("Book not found", 404)
	}

	return nil
}

func DeleteBook(ctx context.Context, id string) error {
	var exists bool
	err := pgx.GetDBPool().QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM books WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return errors.NewCustomError("Failed to check if book exists", 500)
	}

	if !exists {
		return errors.NewCustomError("Book not found", 404)
	}

	result, err := pgx.GetDBPool().Exec(ctx, "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return errors.NewCustomError("Failed to delete book", 500)
	}

	if result.RowsAffected() == 0 {
		return errors.NewCustomError("Book not found", 404)
	}

	return nil
}

func UpdateBookFilePath(ctx context.Context, id, filePath string) error {
	var exists bool
	err := pgx.GetDBPool().QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM books WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return errors.NewCustomError("Failed to check if book exists", 500)
	}

	if !exists {
		return errors.NewCustomError("Book not found", 404)
	}

	result, err := pgx.GetDBPool().Exec(ctx, "UPDATE books SET file_path=$1 WHERE id=$2", filePath, id)
	if err != nil {
		return errors.NewCustomError("Failed to update book file path", 500)
	}

	if result.RowsAffected() == 0 {
		return errors.NewCustomError("Book not found", 404)
	}

	return nil
}
