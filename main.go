package main

import (
	"book_info_system/handler"
	"book_info_system/model"
	"fmt"
	"log"
)

func main() {
	bookID, err := handler.AddBook(model.Book{
		ID:            0, // If your DB auto-generates ID, you can omit or leave it zero
		Title:         "The Pragmatic Programmer",
		Author:        "Andrew Hunt",
		Summary:       "A classic guide to software craftsmanship.",
		Genre:         "Programming",
		PublicateYear: "1999-10-20",
		Pages:         352,
		DateAcquired:  "2025-07-30",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", bookID)
}
