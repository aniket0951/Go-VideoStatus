package repository

import (
	"context"
	"time"

	"github.com/aniket0951/video_status/apis"
	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/models"
	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/google/uuid"
)

type AdminRepository interface {
	Init() (context.Context, context.CancelFunc)

	UploadVideoByAdmin(db.UploadVideoByAdminParams) (models.VideoByAdmin, error)
	VideoByAdmin(db.GetVideoByAdminParams) ([]*models.VideoByAdmin, error)
	UpdateVideoStatus(db.UpdateVideoStatusParams) error
	CreateVerifyVideo(db.CreateVerifyVideoParams) error
	FetchVerifyVideos(db.GetAllVerifyVideosParams) ([]*dto.GetAllVerifyVideos, error)

	CreatePublishVideo(db.CreatePublishedVideoParams) (db.PublishedVideos, error)
	UpdateVerifyVideoStatus(db.UpdateVerifyVideoStatusParams) error
	RollBackCreatedPublishVideo(id uuid.UUID) error

	FetchPublishedVideos(db.FetchAllPublishedVideosParams) ([]db.FetchAllPublishedVideosRow, error)
	UnPublishVideo(db.UpdatePublishedVideoStatusParams) error
	FetchAllUnPublishVideo(db.FetchAllUnPublishedVideosParams) ([]db.FetchAllUnPublishedVideosRow, error)
	MakeVerificationFailed(db.CreateVerificationFailedParams) error
	MakeUnPublishedVideo(db.CreateUnPublishedVideoParams) error
	RollBackCreateVerificationFailed(video_id uuid.UUID) error
	FetchAllVerificationFailedVideos(db.FetchAllVerirficationFailedVideoParams) ([]db.FetchAllVerirficationFailedVideoRow, error)
}

type adminRepository struct {
	db *apis.Store
}

func NewAdminRepository(store *apis.Store) AdminRepository {
	return &adminRepository{
		db: store,
	}
}

func (adminRepo *adminRepository) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
