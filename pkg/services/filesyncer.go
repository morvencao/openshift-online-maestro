package services

import (
	"context"
	e "errors"
	"reflect"

	"github.com/openshift-online/maestro/pkg/dao"
	"github.com/openshift-online/maestro/pkg/db"
	"gorm.io/gorm"

	"github.com/openshift-online/maestro/pkg/api"
	"github.com/openshift-online/maestro/pkg/errors"
)

type FileSyncerService interface {
	Get(ctx context.Context, id string) (*api.FileSyncer, *errors.ServiceError)
	Create(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, *errors.ServiceError)
	Update(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, *errors.ServiceError)
	UpdateStatus(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, bool, *errors.ServiceError)
	Delete(ctx context.Context, id string) *errors.ServiceError
	All(ctx context.Context) (api.FileSyncerList, *errors.ServiceError)
	FindByIDs(ctx context.Context, ids []string) (api.FileSyncerList, *errors.ServiceError)
}

func NewFileSyncerService(lockFactory db.LockFactory, fileSyncerDao dao.FileSyncerDao, events EventService) FileSyncerService {
	return &sqlFileSyncerService{
		lockFactory:   lockFactory,
		fileSyncerDao: fileSyncerDao,
		events:        events,
	}
}

var _ FileSyncerService = &sqlFileSyncerService{}

type sqlFileSyncerService struct {
	lockFactory   db.LockFactory
	fileSyncerDao dao.FileSyncerDao
	events        EventService
}

func (s *sqlFileSyncerService) Get(ctx context.Context, id string) (*api.FileSyncer, *errors.ServiceError) {
	fileSyncer, err := s.fileSyncerDao.Get(ctx, id)
	if err != nil {
		return nil, handleGetError("FileSyncer", "id", id, err)
	}
	return fileSyncer, nil
}

func (s *sqlFileSyncerService) Create(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, *errors.ServiceError) {
	// status may come first from agent before applying spec.
	// TODO: Add validation for the fileSyncer status.
	if len(fileSyncer.Status) != 0 {
		fileSyncer, err := s.fileSyncerDao.Create(ctx, fileSyncer)
		if err != nil {
			return nil, handleCreateError("FileSyncer", err)
		}

		return fileSyncer, nil
	}

	// TODO: Add validation for the fileSyncer spec.
	_, err := s.fileSyncerDao.Get(ctx, fileSyncer.ID)
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			// The filesyncer does not exist, create it.
			fileSyncer, err := s.fileSyncerDao.Create(ctx, fileSyncer)
			if err != nil {
				return nil, handleCreateError("FileSyncer", err)
			}

			_, eErr := s.events.Create(ctx, &api.Event{
				Source:    "FileSyncers",
				SourceID:  fileSyncer.ID,
				EventType: api.CreateEventType,
			})
			if eErr != nil {
				return nil, handleCreateError("FileSyncer", err)
			}

			return fileSyncer, nil
		}

		return nil, handleGetError("FileSyncer", "id", fileSyncer.ID, err)
	}

	// The filesyncer exists, update it.
	updated, err := s.fileSyncerDao.Update(ctx, fileSyncer)
	if err != nil {
		return nil, handleUpdateError("FileSyncer", err)
	}

	_, eErr := s.events.Create(ctx, &api.Event{
		Source:    "FileSyncers",
		SourceID:  fileSyncer.ID,
		EventType: api.UpdateEventType,
	})
	if eErr != nil {
		return nil, handleUpdateError("FileSyncer", err)
	}

	return updated, nil
}

func (s *sqlFileSyncerService) Update(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, *errors.ServiceError) {
	// Updates the filesyncer spec only when its spec changes.
	// If there are multiple requests at the same time, it will cause the race conditions among these
	// requests (read–modify–write), the advisory lock is used here to prevent the race conditions.
	lockOwnerID, err := s.lockFactory.NewAdvisoryLock(ctx, fileSyncer.ID, db.FileSyncers)
	// Ensure that the transaction related to this lock always end.
	defer s.lockFactory.Unlock(ctx, lockOwnerID)
	if err != nil {
		return nil, errors.DatabaseAdvisoryLock(err)
	}

	found, err := s.fileSyncerDao.Get(ctx, fileSyncer.ID)
	if err != nil {
		return nil, handleGetError("FileSyncer", "id", fileSyncer.ID, err)
	}

	// if !found.DeletedAt.Time.IsZero() {
	// 	return nil, errors.Conflict("the filesyncer is under deletion, id: %s", fileSyncer.ID)
	// }

	// Make sure the requested resource version is consistent with its database version.
	if found.Version >= fileSyncer.Version {
		return nil, errors.Conflict("try to update filesyncer with stale version, id: %s", fileSyncer.ID)
	}

	// New spec is not changed, the update action is not needed.
	if reflect.DeepEqual(fileSyncer.Spec, found.Spec) {
		return found, nil
	}

	// TODO: Add validation for the fileSyncer spec update.

	found.Version = fileSyncer.Version
	found.Spec = fileSyncer.Spec

	updated, err := s.fileSyncerDao.Update(ctx, found)
	if err != nil {
		return nil, handleUpdateError("FileSyncer", err)
	}

	if _, err := s.events.Create(ctx, &api.Event{
		Source:    "FileSyncers",
		SourceID:  updated.ID,
		EventType: api.UpdateEventType,
	}); err != nil {
		return nil, handleUpdateError("FileSyncer", err)
	}

	return updated, nil
}

func (s *sqlFileSyncerService) UpdateStatus(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, bool, *errors.ServiceError) {
	// logger := logger.NewOCMLogger(ctx)
	// Updates the filesyncer status only when its status changes.
	// If there are multiple requests at the same time, it will cause the race conditions among these
	// requests (read–modify–write), the advisory lock is used here to prevent the race conditions.
	lockOwnerID, err := s.lockFactory.NewAdvisoryLock(ctx, fileSyncer.ID, db.FileSyncerStatus)
	// Ensure that the transaction related to this lock always end.
	defer s.lockFactory.Unlock(ctx, lockOwnerID)
	if err != nil {
		return nil, false, errors.DatabaseAdvisoryLock(err)
	}

	found, err := s.fileSyncerDao.Get(ctx, fileSyncer.ID)
	if err != nil {
		return nil, false, handleGetError("FileSyncer", "id", fileSyncer.ID, err)
	}

	// Make sure the requested resource version is consistent with its database version.
	// if found.Version != fileSyncer.Version {
	// 	logger.Warning(fmt.Sprintf("Updating status for stale filesyncer; disregard it: id=%s, foundVersion=%d, wantedVersion=%d",
	// 		fileSyncer.ID, found.Version, fileSyncer.Version))
	// 	return found, false, nil
	// }

	// New status is not changed, the update status action is not needed.
	if reflect.DeepEqual(fileSyncer.Status, found.Status) {
		return found, false, nil
	}

	// TODO: compare sequence number and update the status only if the sequence number is greater than the current one.

	found.Status = fileSyncer.Status
	updated, err := s.fileSyncerDao.Update(ctx, found)
	if err != nil {
		return nil, false, handleUpdateError("FileSyncer", err)
	}

	return updated, true, nil
}

func (s *sqlFileSyncerService) Delete(ctx context.Context, id string) *errors.ServiceError {
	if err := s.fileSyncerDao.Delete(ctx, id, true); err != nil {
		return handleDeleteError("FileSyncer", errors.GeneralError("Unable to delete filesyncer: %s", err))
	}

	return nil
}

func (s *sqlFileSyncerService) FindByIDs(ctx context.Context, ids []string) (api.FileSyncerList, *errors.ServiceError) {
	filesyncers, err := s.fileSyncerDao.FindByIDs(ctx, ids)
	if err != nil {
		return nil, handleGetError("FileSyncer", "id", ids, err)
	}
	return filesyncers, nil
}

func (s *sqlFileSyncerService) All(ctx context.Context) (api.FileSyncerList, *errors.ServiceError) {
	filesyncers, err := s.fileSyncerDao.All(ctx)
	if err != nil {
		return nil, errors.GeneralError("Unable to get all filesyncers: %s", err)
	}
	return filesyncers, nil
}
