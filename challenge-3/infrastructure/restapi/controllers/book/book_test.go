package book

import "testing"

// TestGetAllBooks godoc
// @Tags book
// @Summary Testing Get all Books
// @Description Get all Books on the system
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {array} useCaseBook.PaginationResultBook
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /book [get]
func TestGetAllBooks(t *testing.T) {

}

// TestGetBookByID godoc
// @Tags book
// @Summary Testing Get book by ID
// @Description Get book by ID on the system
// @Accept json
// @Produce json
// @Param book_id path int true "ID of book"
// @Success 200 {object} domainBook.Book
// @Failure 400 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /book/{book_id} [get]
func TestGetBookByID(t *testing.T) {

}
