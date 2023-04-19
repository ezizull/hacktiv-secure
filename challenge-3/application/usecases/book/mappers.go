package book

// Package medicine provides the use case for medicine

import (
	domainBook "secure/challenge-3/domain/book"
)

func (n *NewBook) toDomainMapper() *domainBook.Book {
	return &domainBook.Book{
		Title:       n.Title,
		UserID:      n.UserID,
		Description: n.Description,
	}
}
