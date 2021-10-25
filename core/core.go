package core

//go:generate mockgen -destination ../mocks/mocks_core.go -package=mocks github.com/ismrmrd/mrd-storage-api/core MetadataDatabase,BlobStore

import (
	"io"
	"time"

	"github.com/gofrs/uuid"
)

type BlobTags struct {
	Subject     string
	Name        *string
	Device      *string
	Session     *string
	ContentType *string
	CustomTags  map[string][]string
}

type BlobInfo struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Tags      BlobTags
}

type ContinutationToken string

func UnixTimeMsToTime(timeValueMs int64) time.Time {
	return time.Unix(timeValueMs/1000, (timeValueMs%1000)*1000000)
}

type MetadataDatabase interface {
	CreateBlobMetadata(id uuid.UUID, tags *BlobTags) error
	GetBlobMetadata(subject string, id uuid.UUID) (*BlobInfo, error)
	SearchBlobMetadata(tags map[string][]string, at *time.Time, ct *ContinutationToken, pageSize int) ([]BlobInfo, *ContinutationToken, error)
}

type BlobStore interface {
	SaveBlob(contents io.Reader, subject string, id uuid.UUID) error
	ReadBlob(writer io.Writer, subject string, id uuid.UUID) error
}