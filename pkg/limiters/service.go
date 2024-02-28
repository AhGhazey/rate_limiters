package limiters

import (
	"context"
	"time"
)

type RateLimiter interface {
	Check(ctx context.Context, userId, ip string) (*CheckResult, error)
}

type CheckResult struct {
	IsAllowed                 bool
	NumberOfRemainingRequests int
	NextAllowanceTime         time.Time
}
