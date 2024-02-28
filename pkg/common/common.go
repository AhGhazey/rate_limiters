package common

type ContextKey string

const (
	TraceIDKey ContextKey = "Trace-ID"
	IPKey      ContextKey = "X-Forwarded-From"
	UserIDKey  ContextKey = "User-ID"
)
