package repository

import (
	"context"
	"time"

	"github.com/aniket0951/video_status/apis"
	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/models"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

type AdminRepository interface {
	Init() (context.Context, context.CancelFunc)

	UploadVideoByAdmin(args db.UploadVideoByAdminParams) (models.VideoByAdmin, error)
	VideoByAdmin(args db.GetVideoByAdminParams) ([]*models.VideoByAdmin, error)
	UpdateVideoStatus(args db.UpdateVideoStatusParams) error
	CreateVerifyVideo(args db.CreateVerifyVideoParams) error
	FetchVerifyVideos(args db.GetAllVerifyVideosParams) ([]*dto.GetAllVerifyVideos, error)
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