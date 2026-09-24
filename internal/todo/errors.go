package todo

import (
	"errors"
)

var ErrTodoNotFound = errors.New("todo with provided id not found")
var ErrInvalidID = errors.New("invalid or missing id")
var ErrInvalidJson = errors.New("json payload provided is invalid")
