package constants

import "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/types"

const (
	DefaultSourceID = "maestro"
)

var FileSyncerEventDataType = types.CloudEventsDataType{
	Group:    "io.open-cluster-management.works",
	Version:  "v1alpha1",
	Resource: "filesyncers",
}
