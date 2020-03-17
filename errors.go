package corona

import (
	"errors"
)

var (
	// ErrorNoCasesFound is returned when no filtered cases are found.
	ErrorNoCasesFound = errors.New("no cases found")
)
