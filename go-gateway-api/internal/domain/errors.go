package domain

import "errors"

var (
	// ErrAccountNotFound is returned when an account is not found in the database.
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicatedKey is returned when an API key is duplicated.
	ErrDuplicatedAPIKey = errors.New("duplicated API key")
	// ErrInvoiceNotFound is returned when an invoice is not found in the database.
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess is returned when an unauthorized access is attempted.
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
