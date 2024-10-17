package repository

import (
	errors "bookstore/internal/app"
	"bookstore/internal/models"
	"bookstore/internal/store/pgx"
	"context"
)

func AddAuthor(ctx context.Context, a models.Author) error {
	_, err := pgx.GetDBPool().Exec(ctx, "INSERT INTO authors (name) VALUES ($1)", a.Name)
	if err != nil {
		return errors.NewCustomError("Failed to add author", 500)
	}

	return nil
}

func GetAllAuthor(ctx context.Context) ([]models.Author, error) {
	rows, err := pgx.GetDBPool().Query(ctx, "SELECT id, name FROM authors")
	if err != nil {
		return nil, errors.NewCustomError("Failed to retrieve authors", 500)
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var a models.Author
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, errors.NewCustomError("Error scanning author data", 500)
		}
		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.NewCustomError("Error processing author data", 500)
	}

	return authors, nil
}

func GetAuthorByID(ctx context.Context, id string) (models.Author, error) {
	var author models.Author

	err := pgx.GetDBPool().QueryRow(ctx, "SELECT id, name FROM authors WHERE id = $1", id).
		Scan(&author.ID, &author.Name)

	if err != nil {
		return author, errors.NewCustomError("Failed to retrieve author", 500)
	}

	return author, nil
}

func UpdateAuthor(ctx context.Context, id string, updatedAuthor models.Author) error {
	var exists bool
	err := pgx.GetDBPool().QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM authors WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return errors.NewCustomError("Failed to check if author exists", 500)
	}

	if !exists {
		return errors.NewCustomError("Author not found", 404)
	}

	result, err := pgx.GetDBPool().Exec(ctx, "UPDATE authors SET name=$1 WHERE id=$2", updatedAuthor.Name, id)
	if err != nil {
		return errors.NewCustomError("Failed to update author", 500)
	}

	if result.RowsAffected() == 0 {
		return errors.NewCustomError("Author not found", 404)
	}

	return nil
}

func DeleteAuthor(ctx context.Context, id string) error {
	var exists bool
	err := pgx.GetDBPool().QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM authors WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return errors.NewCustomError("Failed to check if author exists", 500)
	}

	if !exists {
		return errors.NewCustomError("Author not found", 404)
	}

	result, err := pgx.GetDBPool().Exec(ctx, "DELETE FROM authors WHERE id=$1", id)
	if err != nil {
		return errors.NewCustomError("Failed to delete author", 500)
	}

	if result.RowsAffected() == 0 {
		return errors.NewCustomError("Author not found", 404)
	}

	return nil
}
