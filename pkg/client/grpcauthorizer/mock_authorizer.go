package grpcauthorizer

import "context"

// MockGRPCAuthorizer returns allowed=true for every request
type MockGRPCAuthorizer struct {
}

func NewMockGRPCAuthorizer() GRPCAuthorizer {
	return &MockGRPCAuthorizer{}
}

var _ GRPCAuthorizer = &MockGRPCAuthorizer{}

// SelfAccessReview returns allowed=true for every request
func (m *MockGRPCAuthorizer) AccessReview(ctx context.Context, action, resourceType, resource, user string, groups []string) (allowed bool, err error) {
	return true, nil
}
