package api

import (
	"strconv"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
)

type FileSyncer struct {
	Meta
	Source       string
	ConsumerName string
	Version      int32
	Spec         datatypes.JSONMap
	Status       datatypes.JSONMap
	// Spec         FileSyncerSpec
	// Status       FileSyncerStatus
}

type FileSyncerSpec struct {
	Files []FileObject
}

type FileObject struct {
	Name         string
	Content      string
	Verification string
	Path         string
	Mode         uint32
	Overwrite    bool
	User         string
	Group        string
}

type FileSyncerStatus struct {
	ContentStatus   []FileSyncerContentStatus
	ReconcileStatus []FileSyncerReconcileStatus
}

type FileSyncerContentStatus struct {
	Name    string
	Content string
}

type FileSyncerReconcileStatus struct {
	Name       string
	Conditions []metav1.Condition
}

type FileSyncerList []*FileSyncer
type FileSyncerIndex map[string]*FileSyncer

func (l FileSyncerList) Index() FileSyncerIndex {
	index := FileSyncerIndex{}
	for _, o := range l {
		index[o.ID] = o
	}
	return index
}

func (f *FileSyncer) BeforeCreate(tx *gorm.DB) error {
	// generate a new ID if it doesn't exist
	if f.ID == "" {
		f.ID = NewID()
	}

	return nil
}

func (f *FileSyncer) GetUID() ktypes.UID {
	return ktypes.UID(f.Meta.ID)
}

func (f *FileSyncer) GetResourceVersion() string {
	return strconv.FormatInt(int64(f.Version), 10)
}

func (f *FileSyncer) GetDeletionTimestamp() *metav1.Time {
	return &metav1.Time{Time: f.Meta.DeletedAt.Time}
}

type FileSyncerPatchRequest struct {
}
