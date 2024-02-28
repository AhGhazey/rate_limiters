package limiters

import (
	"context"
	"errors"
	clients "github.com/ahghazey/rate_limiter/pkg/clients/redis"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

type TokenBucketRateLimiterService struct {
	Config *TokenBucketConfig
	cache  clients.Cache
	mu     sync.Mutex
}

type TokenBucketConfig struct {
	BucketSize                 int
	BucketRefileRateEachSecond int
}

func NewRateLimiterService(cache clients.Cache) *TokenBucketRateLimiterService {
	return &TokenBucketRateLimiterService{
		cache: cache,
	}
}

func (t *TokenBucketRateLimiterService) ApplyConfig(config *TokenBucketConfig) {
	t.Config = config
}

func (t *TokenBucketRateLimiterService) Check(ctx context.Context, userId, ip string) (*CheckResult, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	key := userId + "_" + ip

	// Get current bucket state from cache
	tokens, lastFill, err := t.getCurrentBucketState(ctx, key)
	if err != nil {
		return nil, err
	}

	// Refill tokens
	t.refillTokens(ctx, key, tokens, lastFill)

	// Check if enough tokens are available
	return t.checkTokens(ctx, key, tokens)
}

// getCurrentBucketState gets the current bucket state from the cache.
func (t *TokenBucketRateLimiterService) getCurrentBucketState(ctx context.Context, key string) (int, time.Time, error) {
	var tokens int
	var lastFill time.Time
	lastFillStr, err := t.cache.Get(ctx, key+"_last_fill")
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return 0, time.Time{}, err
		}
		tokens = t.Config.BucketSize
	} else {
		lastFill, err = time.Parse(time.RFC3339Nano, lastFillStr)
		if err != nil {
			return 0, time.Time{}, err
		}
		elapsed := time.Since(lastFill).Seconds()
		refillAmount := int(elapsed) * t.Config.BucketRefileRateEachSecond
		tokensStr, err := t.cache.Get(ctx, key+"_tokens")
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				return 0, time.Time{}, err
			}
			tokens = 0
		}
		tokens = atoi(tokensStr)
		tokens += refillAmount
		if tokens > t.Config.BucketSize {
			tokens = t.Config.BucketSize
		}
	}
	return tokens, lastFill, nil
}

// refillTokens refills tokens in the bucket.
func (t *TokenBucketRateLimiterService) refillTokens(ctx context.Context, key string, tokens int, lastFill time.Time) {
	elapsed := time.Since(lastFill).Seconds()
	refillAmount := int(elapsed) * t.Config.BucketRefileRateEachSecond
	tokens += refillAmount
	if tokens > t.Config.BucketSize {
		tokens = t.Config.BucketSize
	}
	_ = t.cache.Set(ctx, key+"_tokens", itoa(tokens), time.Second)
	_ = t.cache.Set(ctx, key+"_last_fill", time.Now().Format(time.RFC3339Nano), time.Second)
}

// checkTokens checks if enough tokens are available.
func (t *TokenBucketRateLimiterService) checkTokens(ctx context.Context, key string, tokens int) (*CheckResult, error) {
	var allowed bool
	var remainingRequests int
	var nextAllowanceTime time.Time
	if tokens > 0 {
		allowed = true
		tokens--
		remainingRequests = tokens
		nextAllowanceTime = time.Now().Add(time.Second)
		_ = t.cache.Set(ctx, key+"_tokens", itoa(tokens), time.Second)
		_ = t.cache.Set(ctx, key+"_last_fill", time.Now().Format(time.RFC3339Nano), time.Second)
	} else {
		allowed = false
		remainingRequests = 0
		nextAllowanceTimeStr, err := t.cache.Get(ctx, key+"_last_fill")
		if err != nil {
			return nil, err
		}
		lastFill, err := time.Parse(time.RFC3339Nano, nextAllowanceTimeStr)
		if err != nil {
			return nil, err
		}
		nextAllowanceTime = lastFill.Add(time.Second)
	}
	return &CheckResult{
		IsAllowed:                 allowed,
		NumberOfRemainingRequests: remainingRequests,
		NextAllowanceTime:         nextAllowanceTime,
	}, nil
}

// atoi converts a string to an integer.
func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// itoa converts an integer to a string.
func itoa(i int) string {
	return strconv.Itoa(i)
}
