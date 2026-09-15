package wmp

import "context"

type contextKey int

const (
	ctxKeySession contextKey = iota
	ctxKeySender
	ctxKeyRequestID
)

// ContextWithSession returns a context with the session attached.
func ContextWithSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, ctxKeySession, session)
}

// SessionFromContext returns the session from the context, or nil.
func SessionFromContext(ctx context.Context) *Session {
	s, _ := ctx.Value(ctxKeySession).(*Session)
	return s
}

// ContextWithSender returns a context with the sender identity attached.
func ContextWithSender(ctx context.Context, sender string) context.Context {
	return context.WithValue(ctx, ctxKeySender, sender)
}

// SenderFromContext returns the sender identity from the context.
func SenderFromContext(ctx context.Context) string {
	s, _ := ctx.Value(ctxKeySender).(string)
	return s
}

// ContextWithRequestID returns a context with the inbound JSON-RPC request id attached.
// The id is the raw (unquoted) string form of the request's "id" member.
func ContextWithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKeyRequestID, id)
}

// RequestIDFromContext returns the inbound JSON-RPC request id from the context,
// or "" if the context carries none (e.g. the request was a notification).
func RequestIDFromContext(ctx context.Context) string {
	id, _ := ctx.Value(ctxKeyRequestID).(string)
	return id
}
