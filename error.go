package ro

import "errors"

var (
	ErrInvalidResultCount = errors.New("invalid result count")
	ErrConfigUndefined    = errors.New("redis config undefined")
)
