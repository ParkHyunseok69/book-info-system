package handler

import (
	"book_info_system/database"
	"book_info_system/model"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()

func ViewBook(c *gin.Context) {
	title := c.Param("title")
	var book model.Book

	row := db.QueryRow("SELECT * FROM bookTable WHERE title = ?", title)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Summary, &book.Genre,
		&book.PublicateYear, &book.Pages, &book.DateAcquired); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func ListBooks(c *gin.Context) {
	var books []model.Book

	rows, err := db.Query("SELECT * FROM bookTable")
	if err != nil {
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Summary, &book.Genre,
			&book.PublicateYear, &book.Pages, &book.DateAcquired); err != nil {
			log.Println("error:", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var newBook model.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	stmt, err := db.Prepare("INSERT INTO bookTable (book_id, title, author, summary, genre, publicate_year, pages, date_acquired) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(newBook.ID, newBook.Title, newBook.Author, newBook.Summary, newBook.Genre,
		newBook.PublicateYear, newBook.Pages, newBook.DateAcquired)
	if err2 != nil {
		log.Println("error:", err2)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook)
}

func DeleteBook(c *gin.Context) {
	title := c.Param("title")

	result, err := db.Exec("DELETE FROM bookTable WHERE title = ?", title)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}

func UpdateBook(c *gin.Context) {
	var newBook model.Book
	title := c.Param("title")
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	stmt, err := db.Prepare("UPDATE bookTable SET title = ?, author = ?, summary = ?, genre = ?, publicate_year = ?, pages = ?, date_acquired = ? WHERE title = ?")
	if err != nil {
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(newBook.Title, newBook.Author, newBook.Summary, newBook.Genre,
		newBook.PublicateYear, newBook.Pages, newBook.DateAcquired, title)
	if err2 != nil {
		log.Println("error:", err2)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book updated successfully"})
}
