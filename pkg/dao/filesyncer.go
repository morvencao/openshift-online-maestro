package dao

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/openshift-online/maestro/pkg/api"
	"github.com/openshift-online/maestro/pkg/db"
)

type FileSyncerDao interface {
	Get(ctx context.Context, id string) (*api.FileSyncer, error)
	Create(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, error)
	Update(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, error)
	Delete(ctx context.Context, id string, unscoped bool) error
	FindByIDs(ctx context.Context, ids []string) (api.FileSyncerList, error)
	All(ctx context.Context) (api.FileSyncerList, error)
}

var _ FileSyncerDao = &sqlFileSyncerDao{}

type sqlFileSyncerDao struct {
	sessionFactory *db.SessionFactory
}

func NewFileSyncerDao(sessionFactory *db.SessionFactory) FileSyncerDao {
	return &sqlFileSyncerDao{sessionFactory: sessionFactory}
}

func (d *sqlFileSyncerDao) Get(ctx context.Context, id string) (*api.FileSyncer, error) {
	g2 := (*d.sessionFactory).New(ctx)
	var fileSyncer api.FileSyncer
	if err := g2.Take(&fileSyncer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &fileSyncer, nil
}

func (d *sqlFileSyncerDao) Create(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, error) {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Create(fileSyncer).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, err
	}
	return fileSyncer, nil
}

func (d *sqlFileSyncerDao) Update(ctx context.Context, fileSyncer *api.FileSyncer) (*api.FileSyncer, error) {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Updates(fileSyncer).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, err
	}
	return fileSyncer, nil
}

func (d *sqlFileSyncerDao) Delete(ctx context.Context, id string, unscoped bool) error {
	g2 := (*d.sessionFactory).New(ctx)
	if unscoped {
		// Unscoped is used to permanently delete the record
		g2 = g2.Unscoped()
	}
	if err := g2.Omit(clause.Associations).Delete(&api.FileSyncer{Meta: api.Meta{ID: id}}).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return err
	}
	return nil
}

func (d *sqlFileSyncerDao) FindByIDs(ctx context.Context, ids []string) (api.FileSyncerList, error) {
	g2 := (*d.sessionFactory).New(ctx)
	fileSyncers := api.FileSyncerList{}
	if err := g2.Where("id in (?)", ids).Find(&fileSyncers).Error; err != nil {
		return nil, err
	}
	return fileSyncers, nil
}

func (d *sqlFileSyncerDao) All(ctx context.Context) (api.FileSyncerList, error) {
	g2 := (*d.sessionFactory).New(ctx)
	fileSyncers := api.FileSyncerList{}
	if err := g2.Find(&fileSyncers).Error; err != nil {
		return nil, err
	}
	return fileSyncers, nil
}
