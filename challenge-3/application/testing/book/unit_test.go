package book

import (
	"testing"

	"secure/challenge-3/application/mocks"
	bookUsecase "secure/challenge-3/application/usecases/book"

	bookDomain "secure/challenge-3/domain/book"
	errorDomain "secure/challenge-3/domain/errors"

	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	suite.Suite
	book     bookUsecase.BookTesting
	bookMock *mocks.BookTesting
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	bookMock := mocks.BookTesting{}
	book := bookUsecase.NewTesting(&bookMock)

	uts.book = book
	uts.bookMock = &bookMock
}

func (uts *UnitTestSuite) TestGetAll() {
	uts.bookMock.On("GetAll", 1, 20).Return([]*bookDomain.PaginationResultBook{}, nil)

	actual, err := uts.book.GetAll(1, 20)

	uts.GreaterOrEqual(1, len(*actual.Data))
	uts.EqualError(err, errorDomain.NotFound)
}

// func (uts *UnitTestSuite) TestGetAll_Error() {
// 	expectedError := errors.New(errorDomain.NotFound)

// 	uts.bookMock.On("List", mock.Anything).Return([]*bookDomain.Book{}, expectedError)

// 	actual, err := uts.book.GetAll(0, 0)

// 	uts.Equal(0, len(*actual.Data))
// 	uts.Equal(expectedError, err)

// }
