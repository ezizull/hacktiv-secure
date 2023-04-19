package book

import (
	domainBook "secure/challenge-3/domain/book"

	"gorm.io/gorm"
)

func MigrateTestPostgre(inGormDB *gorm.DB) error {
	tablesMigrate := []interface{}{
		&domainBook.TestBook{},
	}

	err := inGormDB.AutoMigrate(tablesMigrate...)
	if err != nil {
		return err
	}
	return nil
}
