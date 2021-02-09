package customerror

import "errors"

var (
	ErrMongo          = errors.New("Error searching for records")
	ErrRecordNotFound = errors.New("Record not found")
	ErrRangeVal       = errors.New("The field does not supply a correct value range")
)
