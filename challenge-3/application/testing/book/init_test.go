package book

import (
	"fmt"
	"testing"

	bookUsecase "secure/challenge-3/application/usecases/book"
	errorDomain "secure/challenge-3/domain/errors"
	bookRepository "secure/challenge-3/infrastructure/repository/postgres/book"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type IntTestSuite struct {
	suite.Suite
	db       *gorm.DB
	bookCase bookUsecase.Service
}

func TestIntTestSuite(t *testing.T) {
	suite.Run(t, &IntTestSuite{})
}

func (its *IntTestSuite) SetupSuite() {
	inDB, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		its.FailNowf("unable to connect to database", err.Error())
	}

	inDB = setupDatabase(its, inDB)

	bookRepo := bookRepository.Repository{DB: inDB}
	bookCase := bookUsecase.Service{BookRepository: bookRepo}

	its.db = inDB
	its.bookCase = bookCase
}

func (its *IntTestSuite) BeforeTest(suiteName, testName string) {
	if testName == "TestGetAll_Error" {
		return
	}
	seedTestTable(its, its.db)
}

func (its *IntTestSuite) TearDownSuite() {
	tearDownDatabase(its)
}

func (its *IntTestSuite) TearDownTest() {
	cleanTable(its)
}

func (its *IntTestSuite) TestGetByID() {
	actual, err := its.bookCase.GetByID(1)

	fmt.Println("check ", actual)

	its.Nil(err)
	its.Equal(uint(1), actual.ID)

}

func (its *IntTestSuite) TestGetByID_Error() {
	actual, err := its.bookCase.GetByID(0)

	its.EqualError(err, errorDomain.NotFoundMessage)
	its.Equal(uint(0), actual.ID)

}

func (its *IntTestSuite) TestGetAll() {
	actual, err := its.bookCase.GetAll(1, 20)

	its.Nil(err)
	its.GreaterOrEqual(1, len(*actual.Data))

}

func (its *IntTestSuite) TestGetAll_Error() {
	actual, err := its.bookCase.GetAll(1, 1)

	its.Nil(err)
	its.Equal(0, len(*actual.Data))

}
