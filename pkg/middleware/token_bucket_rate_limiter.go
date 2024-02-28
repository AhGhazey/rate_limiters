package middleware

import (
	"github.com/ahghazey/rate_limiter/pkg/common"
	"github.com/ahghazey/rate_limiter/pkg/limiters"
	"net/http"
	"strconv"
)

func TokenBucketRateLimiter(rateLimiter limiters.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			config := getBucketConfig(r.URL.Path)
			rateLimiter.(*limiters.TokenBucketRateLimiterService).ApplyConfig(config)
			ip := r.Context().Value(common.IPKey).(string)
			userID := r.Context().Value(common.UserIDKey).(string)
			checkResult, err := rateLimiter.Check(r.Context(), userID, ip)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Remaining-Requests", strconv.Itoa(checkResult.NumberOfRemainingRequests))

			if !checkResult.IsAllowed {
				w.Header().Set("Next-Allowed-Time", checkResult.NextAllowanceTime.String())
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getBucketConfig(path string) *limiters.TokenBucketConfig {
	if path == "/health" {
		return &limiters.TokenBucketConfig{
			BucketSize:                 4,
			BucketRefileRateEachSecond: 4,
		}
	}
	return &limiters.TokenBucketConfig{
		BucketSize:                 1,
		BucketRefileRateEachSecond: 1,
	}
}
