package corona

import (
	"errors"
)

var (
	// ErrorNoReportsFound is returned when no filtered cases are found.
	ErrorNoReportsFound = errors.New("no cases found")
)
