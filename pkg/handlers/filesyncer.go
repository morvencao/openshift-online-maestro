package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/openshift-online/maestro/pkg/api"
	"github.com/openshift-online/maestro/pkg/api/openapi"
	"github.com/openshift-online/maestro/pkg/api/presenters"
	"github.com/openshift-online/maestro/pkg/errors"
	"github.com/openshift-online/maestro/pkg/services"
)

var _ RestHandler = fileSyncerHandler{}

type fileSyncerHandler struct {
	fileSyncer services.FileSyncerService
	generic    services.GenericService
}

func NewFileSyncerHandler(fileSyncer services.FileSyncerService, generic services.GenericService) *fileSyncerHandler {
	return &fileSyncerHandler{
		fileSyncer: fileSyncer,
		generic:    generic,
	}
}

func (h fileSyncerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var fs openapi.FileSyncer
	cfg := &handlerConfig{
		&fs,
		[]validate{
			// validateEmpty(&fs, "Id", "id"),
			validateNotEmpty(&fs, "ConsumerName", "consumer_name"),
			validateNotEmpty(&fs, "Version", "version"),
			validateNotEmpty(&fs, "Spec", "spec"),
		},
		func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			fileSyncer, err := presenters.ConvertFileSyncer(fs)
			if err != nil {
				return nil, errors.GeneralError("failed to convert filesyncer: %s", err)
			}
			fileSyncer, serviceErr := h.fileSyncer.Create(ctx, fileSyncer)
			if serviceErr != nil {
				return nil, serviceErr
			}
			ret, err := presenters.PresentFileSyncer(fileSyncer)
			if err != nil {
				return nil, errors.GeneralError("failed to present filesyncer: %s", err)
			}
			return ret, nil
		},
		handleError,
	}

	handle(w, r, cfg, http.StatusCreated)
}

func (h fileSyncerHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var patch openapi.FileSyncerPatchRequest

	cfg := &handlerConfig{
		&patch,
		[]validate{
			validateNotEmpty(&patch, "Version", "version"),
			validateNotEmpty(&patch, "Spec", "spec"),
		},
		func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			id := mux.Vars(r)["id"]
			// fsSpec, err := presenters.ConvertFileSyncerSpec(patch.Spec)
			// if err != nil {
			// 	return nil, errors.GeneralError("failed to convert filesyncer spec: %s", err)
			// }
			fsSpec := presenters.ConvertFileSyncerSpec(patch.Spec)
			fileSyncer, serviceErr := h.fileSyncer.Update(ctx, &api.FileSyncer{
				Meta: api.Meta{ID: id},
				Spec: fsSpec,
			})
			if serviceErr != nil {
				return nil, serviceErr
			}
			fs, err := presenters.PresentFileSyncer(fileSyncer)
			if err != nil {
				return nil, errors.GeneralError("failed to present filesyncer: %s", err)
			}
			return fs, nil
		},
		handleError,
	}

	handle(w, r, cfg, http.StatusOK)
}

func (h fileSyncerHandler) List(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()

			listArgs := services.NewListArguments(r.URL.Query())
			fileSyncers := []api.FileSyncer{}
			paging, err := h.generic.List(ctx, "username", listArgs, &fileSyncers)
			if err != nil {
				return nil, err
			}
			fileSyncerList := openapi.FileSyncerList{
				Kind:  *presenters.ObjectKind(fileSyncers),
				Page:  int32(paging.Page),
				Size:  int32(paging.Size),
				Total: int32(paging.Total),
				Items: []openapi.FileSyncer{},
			}

			for _, fileSyncer := range fileSyncers {
				converted, err := presenters.PresentFileSyncer(&fileSyncer)
				if err != nil {
					return nil, errors.GeneralError("failed to present filesyncer: %s", err)
				}
				fileSyncerList.Items = append(fileSyncerList.Items, *converted)
			}
			if listArgs.Fields != nil {
				filteredItems, err := presenters.SliceFilter(listArgs.Fields, fileSyncerList.Items)
				if err != nil {
					return nil, err
				}
				return filteredItems, nil
			}
			return fileSyncerList, nil
		},
	}

	handleList(w, r, cfg)
}

func (h fileSyncerHandler) Get(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			ctx := r.Context()
			fs, serviceErr := h.fileSyncer.Get(ctx, id)
			if serviceErr != nil {
				return nil, serviceErr
			}

			ret, err := presenters.PresentFileSyncer(fs)
			if err != nil {
				return nil, errors.GeneralError("failed to present filesyncer: %s", err)
			}
			return ret, nil
		},
	}

	handleGet(w, r, cfg)
}

func (h fileSyncerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			if err := h.fileSyncer.Delete(r.Context(), id); err != nil {
				return nil, err
			}
			return nil, nil
		},
	}
	handleDelete(w, r, cfg, http.StatusNoContent)
}
