package cloudevents

import "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/types"

var JSONObjEventDataType = types.CloudEventsDataType{
	Group:    "io.open-cluster-management.works",
	Version:  "v1alpha1",
	Resource: "jsonobjects",
}
