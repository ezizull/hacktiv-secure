package controllers

import (
	"net/http"
	"secure/challenge-2/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get All Book
func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, domain.Books)
}

// Get Book
func GetBook(ctx *gin.Context) {
	// Convert id to int
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Find book with matching ID
	for _, book := range domain.Books {
		if book.ID == bookID {
			ctx.JSON(http.StatusOK, book)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// Add Book
func AddBook(ctx *gin.Context) {
	var (
		data domain.Book
		err  error
	)

	// Get json data
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	// Validate book title
	if data.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title are required fields"})
		return
	}

	// Validate book author
	if data.Author == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author are required fields"})
		return
	}

	// Add Book to Books
	data.ID = len(domain.Books) + 1
	domain.Books = append(domain.Books, data)

	ctx.JSON(http.StatusCreated, data)
}

// Update Book
func UpdateBook(ctx *gin.Context) {
	var (
		data struct {
			Title  *string `json:"title"`
			Author *string `json:"author"`
			Descr  *string `json:"description"`
		}

		found bool
		err   error
	)

	// Convert id to int
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Get json data
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	// Update book
	for i, book := range domain.Books {
		if book.ID == bookID {
			// update data ID
			domain.Books[i].ID = bookID

			// if update title
			if data.Title != nil {
				domain.Books[i].Title = *data.Title
			}

			// if update author
			if data.Author != nil {
				domain.Books[i].Author = *data.Author
			}

			// if update description
			if data.Descr != nil {
				domain.Books[i].Descr = *data.Descr
			}

			found = true
			break
		}
	}

	// Handle update
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// Delete Book
func DeleteBook(ctx *gin.Context) {
	// Convert id to int
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Delete book
	for i, book := range domain.Books {
		if book.ID == bookID {
			domain.Books = append(domain.Books[:i], domain.Books[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "delete"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
