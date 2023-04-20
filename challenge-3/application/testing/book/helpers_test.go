package book

import (
	bookDomain "secure/challenge-3/domain/book"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupDatabase(its *IntTestSuite, db *gorm.DB) {
	its.T().Log("setting up database")

	tx := db.Exec(createDatabase)
	if tx.Error != nil {
		its.FailNowf("unable to create database", tx.Error.Error())
	}

	var err error
	db, err = gorm.Open(postgres.Open(connectionTest), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		its.FailNowf("unable to connect to database", err.Error())
	}

	tablesMigrate := []interface{}{
		&bookDomain.Book{},
	}

	err = db.AutoMigrate(tablesMigrate...)
	if err != nil {
		its.FailNowf("unable to create table", err.Error())
	}
}

func seedTestTable(its *IntTestSuite, db *gorm.DB) {
	its.T().Log("seeding test table")

	for i := 1; i <= 2; i++ {
		idStr := strconv.Itoa(i)
		book := bookDomain.Book{ID: uint(i), Title: "Tittle " + idStr, UserID: i, Description: "Description " + idStr}

		tx := db.Create(&book)
		if tx.Error != nil {
			its.FailNowf("unable to seed table", tx.Error.Error())
		}
	}
}

func cleanTable(its *IntTestSuite) {
	its.T().Log("cleaning database")

	tx := its.db.Exec(deleteFromTable)
	if tx.Error != nil {
		its.FailNowf("unable to clean table", tx.Error.Error())
	}
}

func tearDownDatabase(its *IntTestSuite) {
	its.T().Log("tearing down database")

	tx := its.db.Exec(dropTable)
	if tx.Error != nil {
		its.FailNowf("unable to drop table", tx.Error.Error())
	}

	tx = its.db.Exec(dropaDatabase)
	if tx.Error != nil {
		its.FailNowf("unable to drop database", tx.Error.Error())
	}

	inDB, err := its.db.DB()
	if err != nil {
		its.FailNowf("unable to close database", err.Error())
	}

	err = inDB.Close()
	if err != nil {
		its.FailNowf("unable to close database", err.Error())
	}
}
