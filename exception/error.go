package exception

import "errors"

var ErrInternalServer = errors.New("internal server error")

var ErrIdIsNotFound = errors.New("id is not found")
