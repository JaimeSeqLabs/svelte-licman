package pkgerrors

import "errors"

var (
	ErrCredsNotFound = errors.New("credentials not found")
)
