package limiters

import "time"

type RateLimiter interface {
	Check(userId, ip string) (*CheckResult, error)
}

type CheckResult struct {
	IsAllowed                 bool
	NumberOfRemainingRequests int
	NextAllowanceTime         time.Time
}
