package grpcauthorizer

import "context"

// GRPCAuthorizer defines an interface for performing access reviews in a gRPC-based authorization.
type GRPCAuthorizer interface {
	// AccessReview checks if the specified user or groups has permission to perform a given action on a specified resource.
	//
	// Parameters:
	// - ctx: The context for managing request lifecycle.
	// - action: The action being requested, e.g., "pub" (publish) or "sub" (subscribe).
	// - resourceType: The type of resource, e.g., "source" or "cluster".
	// - resource: The specific resource name within the given resource type.
	// - user: The user requesting the action (may be empty if groups are used).
	// - groups: The groups requesting the action (may be empty if user is used).
	//
	// Returns:
	// - allowed: True if access is granted, false otherwise.
	// - err: Any error encountered during the review process.
	AccessReview(ctx context.Context, action, resourceType, resource, user string, groups []string) (allowed bool, err error)
}
