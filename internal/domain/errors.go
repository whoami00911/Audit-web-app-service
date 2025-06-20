package domain

import "errors"

var (
	ErrNoAction = errors.New("action doesn't exists")
	ErrNoMethod = errors.New("method doesn't exists")
)
