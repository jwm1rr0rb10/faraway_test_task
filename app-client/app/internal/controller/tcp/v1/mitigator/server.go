package mitigator

import (
	"context"

	"github.com/jwm1rr0rb/boilerplate_libraries/backend/golang/logging"

	"app-client/app/internal/policy/mitigator"
)

type Policy interface {
	// SolvePoWChallenge attempts to find a valid nonce for the given challenge.
	SolvePoWChallenge(ctx context.Context, challenge mitigator.PoWChallenge) (*mitigator.PoWSolution, error)
}

// Controller is a component that is responsible for interacting with the
// policy.
// Controller handles the interaction with the TCP server.
type Controller struct {
	policy Policy // Dependency on the PoW solving policy
}

// NewController creates a new TCP controller instance.
func NewController(policy Policy) *Controller {
	if policy == nil {
		logging.Default().Error("policy cannot be nil")
	}
	return &Controller{
		policy: policy,
	}
}
