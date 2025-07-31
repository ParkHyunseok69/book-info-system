package handler

import (
	"book_info_system/database"
	"book_info_system/model"
	"database/sql"
	"fmt"
)

var db = database.Connect()

func ViewBook(title string) (model.Book, error) {
	var book model.Book

	row := db.QueryRow("SELECT * FROM bookTable WHERE title = ?", title)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Summary, &book.Genre,
		&book.PublicateYear, &book.Pages, &book.DateAcquired); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("error: %v: no such book", err)
		}
		return book, fmt.Errorf("error: %v", err)
	}
	return book, nil
}

func AddBook(book model.Book) (int64, error) {
	result, err := db.Exec("INSERT INTO bookTable (book_id, title, author, summary, genre, publicate_year, pages, date_acquired) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		book.ID, book.Title, book.Author, book.Summary, book.Genre,
		book.PublicateYear, book.Pages, book.DateAcquired)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}
	return id, nil
}
