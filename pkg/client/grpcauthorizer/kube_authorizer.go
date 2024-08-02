package grpcauthorizer

import (
	"context"
	"fmt"

	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// KubeGRPCAuthorizer is a gRPC authorizer that uses the Kubernetes RBAC API to authorize requests.
type KubeGRPCAuthorizer struct {
	kubeClient kubernetes.Interface
}

func NewKubeGRPCAuthorizer(kubeClient kubernetes.Interface) GRPCAuthorizer {
	return &KubeGRPCAuthorizer{
		kubeClient: kubeClient,
	}
}

var _ GRPCAuthorizer = &KubeGRPCAuthorizer{}

// AccessReview checks if the given user or group is allowed to perform the given action on the given resource by making a SubjectAccessReview request.
func (k *KubeGRPCAuthorizer) AccessReview(ctx context.Context, action, resourceType, resource, user, group string) (allowed bool, err error) {
	if user != "" && group != "" {
		return false, fmt.Errorf("both user and group cannot be specified")
	}

	if action != "pub" && action != "sub" {
		return false, fmt.Errorf("unsupported action: %s", action)
	}

	if resource == "" {
		return false, fmt.Errorf("resource cannot be empty")
	}

	nonResourceUrl := ""
	switch resourceType {
	case "source":
		nonResourceUrl = fmt.Sprintf("/sources/%s", resource)
	case "cluster":
		nonResourceUrl = fmt.Sprintf("/clusters/%s", resource)
	default:
		return false, fmt.Errorf("unsupported resource type: %s", resourceType)
	}

	sar, err := k.kubeClient.AuthorizationV1().SubjectAccessReviews().Create(ctx, &authorizationv1.SubjectAccessReview{
		Spec: authorizationv1.SubjectAccessReviewSpec{
			NonResourceAttributes: &authorizationv1.NonResourceAttributes{
				Path: nonResourceUrl,
				Verb: action,
			},
			User:   user,
			Groups: []string{group},
		},
	}, metav1.CreateOptions{})

	if err != nil {
		return false, err
	}

	return sar.Status.Allowed, nil
}
