package book

// type UnitTestSuite struct {
// 	suite.Suite
// 	book     bookDomain.Testing
// 	bookMock *bookDomain.MockTesting
// }

// func TestUnitTestSuite(t *testing.T) {
// 	suite.Run(t, &UnitTestSuite{})
// }

// func (uts *UnitTestSuite) SetupTest() {
// 	bookMock := bookDomain.MockTesting{}
// 	book := NewPriceIncreasebook(&bookMock)

// 	uts.book = book
// 	uts.bookMock = &bookMock
// }

// func (uts *UnitTestSuite) TestCalculate() {
// 	uts.bookMock.On("List", mock.Anything).Return([]*stocks.PriceData{}, nil)

// 	actual, err := uts.book.PriceIncrease()

// 	uts.Equal(0.0, actual)
// 	uts.EqualError(err, "not enough data")
// }

// func (uts *UnitTestSuite) TestCalculate_ErrorFromPriceProvider() {
// 	expectedError := errors.New("oh my god")

// 	uts.bookMock.On("List", mock.Anything).Return([]*stocks.PriceData{}, expectedError)

// 	actual, err := uts.book.PriceIncrease()

// 	uts.Equal(0.0, actual)
// 	uts.Equal(expectedError, err)

// }
