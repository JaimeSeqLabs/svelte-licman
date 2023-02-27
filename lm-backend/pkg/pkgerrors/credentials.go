package pkgerrors

import "errors"

var (
	ErrCredsNotFound = errors.New("credentials not found")
	ErrNoTokensFound = errors.New("user does not have any issued token")
)
