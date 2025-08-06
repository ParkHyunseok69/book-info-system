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
	id := c.Param("id")
	var book model.Book

	row := db.QueryRow("SELECT * FROM bookTable WHERE book_id = ?", id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Summary, &book.Genre,
		&book.PublicateYear, &book.Pages, &book.DateAcquired, &book.Status); err != nil {
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
			&book.PublicateYear, &book.Pages, &book.DateAcquired, &book.Status); err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	stmt, err := db.Prepare("INSERT INTO bookTable (title, author, summary, genre, publicate_year, pages, date_acquired, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("error:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(newBook.Title, newBook.Author, newBook.Summary, newBook.Genre,
		newBook.PublicateYear, newBook.Pages, newBook.DateAcquired, newBook.Status)
	if err2 != nil {
		log.Println("error:", err2)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM bookTable WHERE book_id = ?", id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
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
	id := c.Param("id")
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var oldBook model.Book
	err := db.QueryRow(`SELECT title, author, summary, genre, publicate_year, pages, date_acquired, status 
                        FROM bookTable WHERE book_id = ?`, id).
		Scan(&oldBook.Title, &oldBook.Author, &oldBook.Summary, &oldBook.Genre,
			&oldBook.PublicateYear, &oldBook.Pages, &oldBook.DateAcquired, &oldBook.Status)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if newBook.Title != "" {
		oldBook.Title = newBook.Title
	}
	if newBook.Author != "" {
		oldBook.Author = newBook.Author
	}
	if newBook.Summary != "" {
		oldBook.Summary = newBook.Summary
	}
	if newBook.Genre != "" {
		oldBook.Genre = newBook.Genre
	}
	if newBook.PublicateYear != "" {
		oldBook.PublicateYear = newBook.PublicateYear
	}
	if newBook.Pages != 0 {
		oldBook.Pages = newBook.Pages
	}
	if newBook.DateAcquired != "" {
		oldBook.DateAcquired = newBook.DateAcquired
	}
	if newBook.Status != "" {
		oldBook.Status = newBook.Status
	}

	_, err = db.Exec(`UPDATE bookTable 
                      SET title=?, author=?, summary=?, genre=?, publicate_year=?, pages=?, date_acquired=?, status=?
                      WHERE book_id=?`,
		oldBook.Title, oldBook.Author, oldBook.Summary, oldBook.Genre,
		oldBook.PublicateYear, oldBook.Pages, oldBook.DateAcquired, oldBook.Status,
		id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
