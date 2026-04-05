package mitigator

import "github.com/jwm1rr0rb/boilerplate_libraries/backend/golang/errors"

var (
	ErrPoWTimeout      = errors.New("proof-of-work challenge solving timed out")
	ErrPoWInvalidInput = errors.New("invalid input for PoW challenge")
)
