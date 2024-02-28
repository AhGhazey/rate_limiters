package middleware

import (
	"context"
	"github.com/ahghazey/rate_limiter/pkg/common"
	"net/http"
)

func TraceHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		traceID := r.Header.Get(string(common.TraceIDKey))
		ip := r.Header.Get(string(common.IPKey))
		userID := r.Header.Get(string(common.UserIDKey))

		ctx := context.WithValue(r.Context(), common.TraceIDKey, traceID)
		ctx = context.WithValue(ctx, common.IPKey, ip)
		ctx = context.WithValue(ctx, common.UserIDKey, userID)

		w.Header().Add(string(common.TraceIDKey), traceID)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
