package repository

import (
	"errors"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/models"
	db "github.com/aniket0951/video_status/sqlc_lib"
)

func (adminRepo *adminRepository) UploadVideoByAdmin(args db.UploadVideoByAdminParams) (models.VideoByAdmin, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	video, err := adminRepo.db.Queries.UploadVideoByAdmin(ctx, args)

	return models.VideoByAdmin(video), err
}

// fetch all verify videos
func (adminRepo *adminRepository) FetchVerifyVideos(args db.GetAllVerifyVideosParams) ([]*dto.GetAllVerifyVideos, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.GetAllVerifyVideos(ctx, args)

	err = helper.HandleDBErr(err)

	if err != nil {
		return nil, err
	}

	videos := make([]*dto.GetAllVerifyVideos, len(result))

	for i, video := range result {
		videos[i] = &dto.GetAllVerifyVideos{
			Status:       video.Status,
			VideoID:      video.VideoID,
			VideoTitle:   video.VideoTitle,
			VideoAddress: video.VideoAddress,
			VerifyAt:     video.VerifyAt,
			UploadedAt:   video.UploadedAt,
		}
	}

	return videos, err
}

// fetch only VIDEO_INIT status videos
func (adminRepo *adminRepository) VideoByAdmin(args db.GetVideoByAdminParams) ([]*models.VideoByAdmin, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.GetVideoByAdmin(ctx, args)

	videos := make([]*models.VideoByAdmin, len(result))

	for i, video := range result {
		videos[i] = &models.VideoByAdmin{
			ID:          video.ID,
			Title:       video.Title,
			FileAddress: video.FileAddress,
			UploadedBy:  video.UploadedBy,
			Status:      video.Status,
			CreatedAt:   video.CreatedAt,
			UpdatedAt:   video.UpdatedAt,
		}

	}

	return videos, err
}

// make VIDEO_INIT to VIDEO_VERIFY
func (adminRepo *adminRepository) UpdateVideoStatus(args db.UpdateVideoStatusParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()
	result, err := adminRepo.db.Queries.UpdateVideoStatus(ctx, args)

	if err != nil {
		return err
	}

	err = helper.HandleDBErr(err)
	if err != nil {
		return err
	}
	if rows_affected, _ := result.RowsAffected(); rows_affected != 1 {
		err = errors.New("failed to update the video status")
	}

	return err
}

// create a verify video object and make a status VIDEO_VERIFY, this needs to be a VIDEO_PUBLISHED
func (adminRepo *adminRepository) CreateVerifyVideo(args db.CreateVerifyVideoParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	_, err := adminRepo.db.Queries.CreateVerifyVideo(ctx, args)

	err = helper.HandleDBErr(err)
	return err
}

// make VIDEO_VERIFY to VIDEO_PUBLISHED
func (adminRepo *adminRepository) UpdateVerifyVideoStatus(args db.UpdateVerifyVideoStatusParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()
	_, err := adminRepo.db.Queries.UpdateVerifyVideoStatus(ctx, args)

	return err
}
