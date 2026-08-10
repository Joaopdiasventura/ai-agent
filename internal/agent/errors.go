package agent

import "errors"

var ErrEmptyQuestion = errors.New(
	"question cannot be empty",
)
