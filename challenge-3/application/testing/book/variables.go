package book

import "fmt"

// initial
var (
	dbHost     = "localhost"
	dbName     = "hacktiv"
	dbUser     = "root"
	dbPassword = "root"
	dbPort     = "5432"
	dbTimezone = "Asia/Jakarta"

	dbTestHost     = "localhost"
	dbTestName     = "books_test"
	dbTestUser     = "root"
	dbTestPassword = "root"
	dbTestPort     = "5432"
	dbTestTimezone = "Asia/Jakarta"

	dbTable = "books"
)

// connection
var (
	connection     = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbTimezone)
	connectionTest = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbTestHost, dbTestUser, dbTestPassword, dbTestName, dbTestPort, dbTestTimezone)
)

// table
var (
	createDatabase  = fmt.Sprintf("CREATE DATABASE %s", dbTestName)
	deleteFromTable = fmt.Sprintf("DELETE FROM %s", dbTable)
	dropTable       = fmt.Sprintf("DROP TABLE %s", dbTable)
	dropaDatabase   = fmt.Sprintf("DROP DATABASE %s", dbTestName)
)
