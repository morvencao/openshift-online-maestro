package grpcauthorizer

import "context"

// GRPCAuthorizer defines an interface for performing access reviews in a gRPC-based authorization.
type GRPCAuthorizer interface {
	// AccessReview checks if the specified user or group has permission to perform a given action on a specified resource.
	//
	// Parameters:
	// - ctx: The context for managing request lifecycle.
	// - action: The action being requested, e.g., "pub" (publish) or "sub" (subscribe).
	// - resourceType: The type of resource, e.g., "source" or "cluster".
	// - resource: The specific resource name within the given resource type.
	// - user: The username requesting the action (may be empty if group is used).
	// - group: The group name requesting the action (may be empty if user is used).
	//
	// Returns:
	// - allowed: True if access is granted, false otherwise.
	// - err: Any error encountered during the review process.
	AccessReview(ctx context.Context, action, resourceType, resource, user, group string) (allowed bool, err error)
}
