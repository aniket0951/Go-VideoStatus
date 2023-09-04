package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/models"
	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/google/uuid"
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

// create a published video object and make status VIDEO_PUBLISHED, this will show to end users
func (adminRepo *adminRepository) CreatePublishVideo(args db.CreatePublishedVideoParams) (db.PublishedVideos, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	publish_video, err := adminRepo.db.Queries.CreatePublishedVideo(ctx, args)

	err = helper.HandleDBErr(err)
	return publish_video, err
}

// make VIDEO_VERIFY to VIDEO_PUBLISHED
func (adminRepo *adminRepository) UpdateVerifyVideoStatus(args db.UpdateVerifyVideoStatusParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()
	_, err := adminRepo.db.Queries.UpdateVerifyVideoStatus(ctx, args)

	return err
}

// rollback the published video if create publish or  update verify status failed
func (adminRepo *adminRepository) RollBackCreatedPublishVideo(id uuid.UUID) error {

	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.DeletePublishedVideo(ctx, id)

	if err != nil {
		return err
	}

	rows_affected, _ := result.RowsAffected()
	if rows_affected == 0 {
		return errors.New("failed to rollback the video")
	}

	return nil
}

// fetch all publish videos
func (adminRepo *adminRepository) FetchPublishedVideos(args db.FetchAllPublishedVideosParams) ([]db.FetchAllPublishedVideosRow, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.FetchAllPublishedVideos(ctx, args)

	err = helper.HandleDBErr(err)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}

// unpublish the video, if it has been uploaded by mistake
func (adminRepo *adminRepository) UnPublishVideo(args db.UpdatePublishedVideoStatusParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.UpdatePublishedVideoStatus(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
	}

	err = helper.HandleDBErr(err)
	if err != nil {
		return err
	}

	if result.Status != args.Status {
		return errors.New("failed to unpublish the video")
	}

	return nil
}

// fetch unpublish videos
func (adminRepo *adminRepository) FetchAllUnPublishVideo(args db.FetchAllUnPublishedVideosParams) ([]db.FetchAllUnPublishedVideosRow, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.FetchAllUnPublishedVideos(ctx, args)

	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, err
}

// video verification failed
func (adminRepo *adminRepository) MakeVerificationFailed(args db.CreateVerificationFailedParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.CreateVerificationFailed(ctx, args)

	err = helper.HandleDBErr(err)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.New("this video is already in verification failed")
		}
		return err
	}

	if result.VideoID != args.VideoID {
		return errors.New("failed to creare video verification failed")
	}

	return nil
}

// create a unpublish video object in video verification failed
func (adminRepo *adminRepository) MakeUnPublishedVideo(args db.CreateUnPublishedVideoParams) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.CreateUnPublishedVideo(ctx, args)

	err = helper.HandleDBErr(err)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.New("this video is already in verification failed")
		}
		return err
	}

	if result.VideoID != args.VideoID {
		return errors.New("failed to creare video verification failed")
	}

	return nil
}

// rollback the video verification failed process created data, it could be unpublish or video verification failed
func (adminRepo *adminRepository) RollBackCreateVerificationFailed(video_id uuid.UUID) error {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.DeleteVerificationFailed(ctx, video_id)

	if err != nil {
		return err
	}

	rows_affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows_affected == 0 {
		return errors.New("failed to rollback the data")
	}

	return nil
}

// fetch all videos from verification failed process
func (adminRepo *adminRepository) FetchAllVerificationFailedVideos(args db.FetchAllVerirficationFailedVideoParams) ([]db.FetchAllVerirficationFailedVideoRow, error) {
	ctx, cancel := adminRepo.Init()
	defer cancel()

	result, err := adminRepo.db.Queries.FetchAllVerirficationFailedVideo(ctx, args)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("video's not available")
	}

	return result, nil
}
