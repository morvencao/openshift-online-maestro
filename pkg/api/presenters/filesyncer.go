package presenters

import (
	"github.com/openshift-online/maestro/pkg/api"
	"github.com/openshift-online/maestro/pkg/api/openapi"
	"github.com/openshift-online/maestro/pkg/constants"
	"github.com/openshift-online/maestro/pkg/util"
	"gorm.io/datatypes"
)

func ConvertFileSyncer(fileSyncer openapi.FileSyncer) (*api.FileSyncer, error) {
	// spec, err := ConvertFileSyncerSpec(fileSyncer.Spec)
	// if err != nil {
	// 	return nil, err
	// }
	spec := ConvertFileSyncerSpec(fileSyncer.Spec)

	return &api.FileSyncer{
		Meta: api.Meta{
			ID: util.NilToEmptyString(fileSyncer.Id),
		},
		// Set the default source ID for RESTful API calls and do not allow modification
		Source:       constants.DefaultSourceID,
		ConsumerName: util.NilToEmptyString(fileSyncer.ConsumerName),
		Version:      util.NilToEmptyInt32(fileSyncer.Version),
		Spec:         spec,
	}, nil
}

func PresentFileSyncer(fileSyncer *api.FileSyncer) (*openapi.FileSyncer, error) {
	// spec, err := presentFileSyncerSpec(fileSyncer.Spec)
	// if err != nil {
	// 	return nil, err
	// }
	// status, err := presentFileSyncerStatus(fileSyncer.Status)
	// if err != nil {
	// 	return nil, err
	// }
	spec := presentFileSyncerSpec(fileSyncer.Spec)
	status := presentFileSyncerStatus(fileSyncer.Status)
	reference := PresentReference(fileSyncer.ID, fileSyncer)
	return &openapi.FileSyncer{
		Id:           reference.Id,
		Kind:         reference.Kind,
		Href:         reference.Href,
		ConsumerName: openapi.PtrString(fileSyncer.ConsumerName),
		Version:      openapi.PtrInt32(fileSyncer.Version),
		CreatedAt:    openapi.PtrTime(fileSyncer.CreatedAt),
		UpdatedAt:    openapi.PtrTime(fileSyncer.UpdatedAt),
		Spec:         spec,
		Status:       status,
	}, nil
}

// func ConvertFileSyncerSpec(spec map[string]interface{}) (api.FileSyncerSpec, error) {
// 	specJSON, err := json.Marshal(spec)
// 	if err != nil {
// 		return api.FileSyncerSpec{}, err
// 	}
// 	fileSyncerSpec := api.FileSyncerSpec{}
// 	if err := json.Unmarshal(specJSON, &fileSyncerSpec); err != nil {
// 		return api.FileSyncerSpec{}, err
// 	}

// 	return fileSyncerSpec, nil
// }

func ConvertFileSyncerSpec(spec map[string]interface{}) datatypes.JSONMap {
	return datatypes.JSONMap(spec)
}

func presentFileSyncerSpec(spec datatypes.JSONMap) map[string]interface{} {
	return map[string]interface{}(spec)
}

func presentFileSyncerStatus(status datatypes.JSONMap) map[string]interface{} {
	return map[string]interface{}(status)
}

// func presentFileSyncerSpec(spec api.FileSyncerSpec) (map[string]interface{}, error) {
// 	specJSON, err := json.Marshal(spec)
// 	if err != nil {
// 		return nil, err
// 	}
// 	specMap := map[string]interface{}{}
// 	if err := json.Unmarshal(specJSON, &specMap); err != nil {
// 		return nil, err
// 	}

// 	return specMap, nil
// }

// func presentFileSyncerStatus(status api.FileSyncerStatus) (map[string]interface{}, error) {
// 	statusJSON, err := json.Marshal(status)
// 	if err != nil {
// 		return nil, err
// 	}
// 	statusMap := map[string]interface{}{}
// 	if err := json.Unmarshal(statusJSON, &statusMap); err != nil {
// 		return nil, err
// 	}

// 	return statusMap, nil
// }
