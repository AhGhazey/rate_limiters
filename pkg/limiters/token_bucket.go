package limiters

import "time"

type TokenBucketRateLimiterService struct {
	Config *TokenBucketConfig
}

type TokenBucketConfig struct {
	BucketSize                 int
	BucketRefileRateEachSecond int
}

func NewRateLimiterService() *TokenBucketRateLimiterService {
	return &TokenBucketRateLimiterService{}
}

func (t *TokenBucketRateLimiterService) ApplyConfig(config *TokenBucketConfig) {
	t.Config = config
}

func (t *TokenBucketRateLimiterService) Check(userId, ip string) (*CheckResult, error) {
	// TODO implement me
	return &CheckResult{
		IsAllowed:                 true,
		NumberOfRemainingRequests: 0,
		NextAllowanceTime:         time.Now(),
	}, nil
}
