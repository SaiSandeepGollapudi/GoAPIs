package dataservice

import (
	"/api/model"
	"database/sql"
)

func CreateBook(db *sql.DB, book model.Book) error {
	query := `INSERT INTO library(title, author, year) VALUES (?, ?, ?)`
	_, err := db.Exec(query, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}
