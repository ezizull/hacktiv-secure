package book

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type IntTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func TestIntTestSuite(t *testing.T) {
	suite.Run(t, &IntTestSuite{})
}
