package services

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"

	"github.com/aniket0951/video_status/apis/dto"
	"github.com/aniket0951/video_status/apis/helper"
	"github.com/aniket0951/video_status/apis/models"
	"github.com/aniket0951/video_status/apis/repository"
	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/google/uuid"
)

type AdminService interface {
	UploadVideoByAdmin(dto.UploadVideoByAdminRequestParam, multipart.File) (models.VideoByAdmin, error)
	GetVideoByAdmin(dto.GetVideoByAdminRequestParams) ([]*models.VideoByAdmin, error)
	UpdateVideoStatus(dto.UpdateVideoStatusRequestParams) error
	FetchVerifyVideos(dto.FetchVerifyVideosRequestParams) ([]*dto.GetAllVerifyVideos, error)
	PublishVideo(req dto.PublishedVideoRequestParams) error
	FetchAllPublishedVideos(req dto.FetchVerifyVideosRequestParams) ([]dto.FetchAllPublishedVideosDTO, error)
	UnPublishVideo(req dto.PublishedVideoRequestParams) error

	FetchUnPublishVideos(req dto.FetchVerifyVideosRequestParams) ([]dto.FetchAllPublishedVideosDTO, error)
	MakeVerificationFailed(req dto.CreateVerificationFailedRequestParam) error
	MakeUnPublishedVideo(req dto.CreateVerificationFailedRequestParam) error
	FetchAllVerificationFailedVideos(req dto.FetchVerifyVideosRequestParams) ([]dto.FetchAllVerificationFailedVideosDTO, error)
	FetchVerifyVideoFullDetails(video_id uuid.UUID) (dto.FetchVerifyVideoFullDetailsDTO, error)
	FetchVideoByAdminFullDetails(video_id uuid.UUID) (dto.FetchVideoByAdminFullDetailsDTO, error)
	FetchPublishVideoFullDetails(video_id uuid.UUID) (dto.FetchPublishVideoFullDetailsDTO, error)
}

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{
		adminRepo: repo,
	}
}

func (adminServ *adminService) UploadVideoByAdmin(req dto.UploadVideoByAdminRequestParam, file multipart.File) (models.VideoByAdmin, error) {
	tempFile, err := ioutil.TempFile("static", "upload-*.mp4")

	if err != nil {
		return models.VideoByAdmin{}, err
	}

	defer tempFile.Close()

	fileBytes, fileReader := ioutil.ReadAll(file)

	if fileReader != nil {
		return models.VideoByAdmin{}, fileReader
	}

	tempFile.Write(fileBytes)
	defer file.Close()
	defer tempFile.Close()
	file_path := path.Base(tempFile.Name())

	if strings.Contains(file_path, "static\\") {
		file_path = strings.TrimLeft(file_path, "static\\")
	}

	fmt.Println("File Path Has been changed : ", file_path)

	args := db.UploadVideoByAdminParams{
		Title:       req.Title,
		FileAddress: file_path,
		UploadedBy:  req.UploadedBy,
		Status:      req.Status,
	}

	video, err := adminServ.adminRepo.UploadVideoByAdmin(args)

	return video, err
	//return models.VideoByAdmin{}, nil
}

func (adminServ *adminService) GetVideoByAdmin(req dto.GetVideoByAdminRequestParams) ([]*models.VideoByAdmin, error) {
	adminId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	args := db.GetVideoByAdminParams{
		UploadedBy: adminId,
		Status:     helper.VIDEO_INIT,
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}

	videos, err := adminServ.adminRepo.VideoByAdmin(args)

	if err != nil {
		return nil, err
	}

	if len(videos) == 0 {
		return nil, errors.New("videos not found for this admin")
	}

	err = helper.HandleDBErr(err)

	for _, video := range videos {
		//if strings.Contains(video.FileAddress, "static")
		video.SetFileAddress("http://localhost:8080/static/")
	}

	return videos, err
}

func (adminServ *adminService) UpdateVideoStatus(req dto.UpdateVideoStatusRequestParams) error {
	video_id, err := uuid.Parse(req.VideoId)

	if err != nil {
		return err
	}

	args := db.UpdateVideoStatusParams{
		ID:     video_id,
		Status: req.VideoStatus,
	}

	// first update status in video_by_admin
	err = adminServ.adminRepo.UpdateVideoStatus(args)

	if err != nil {
		return err
	}

	// sec create a new video in verify_videos
	err = adminServ.CreateVerifyVideo(video_id)

	return err
}

func (adminServ *adminService) CreateVerifyVideo(video_id uuid.UUID) error {
	verify_by, err := uuid.Parse(helper.TOKEN_ID)

	if err != nil {
		return err
	}

	verify_video_args := db.CreateVerifyVideoParams{
		VideoID:  video_id,
		VerifyBy: verify_by,
		Status:   helper.VIDEO_VERIFY,
	}

	err = adminServ.adminRepo.CreateVerifyVideo(verify_video_args)

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.New("this video is already verify")
		}
	}

	return err
}

func (adminServ *adminService) FetchVerifyVideos(req dto.FetchVerifyVideosRequestParams) ([]*dto.GetAllVerifyVideos, error) {
	args := db.GetAllVerifyVideosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	videos, err := adminServ.adminRepo.FetchVerifyVideos(args)

	if len(videos) == 0 {
		return nil, sql.ErrNoRows
	}

	for _, video := range videos {
		video.VideoAddress = "http://localhost:8080/static/" + video.VideoAddress
	}

	return videos, err
}

// publish the video
func (adminServ *adminService) PublishVideo(req dto.PublishedVideoRequestParams) error {
	video_id, err := uuid.Parse(req.VideoId)
	if err != nil {
		return err
	}

	published_by, err := uuid.Parse(helper.TOKEN_ID)

	if err != nil {
		return err
	}

	// first create publish object
	publish_args := db.CreatePublishedVideoParams{
		VideoID:     video_id,
		PublishedBy: published_by,
		Status:      helper.VIDEO_PUBLISHED,
	}

	publish_video_obj, err := adminServ.adminRepo.CreatePublishVideo(publish_args)
	if err != nil {
		return err
	}

	// secound update status
	err = adminServ.updateVideoStatusForPublish(video_id, helper.VIDEO_PUBLISHED)
	if err != nil {
		// rollback publish_video_obj from published_video
		err_ := adminServ.rollBackCreatedPublishVideo(publish_video_obj.ID)
		if err_ != nil {
			return err_
		}
		return err
	}

	return nil
}

// update video status in verify_video, make VIDEO_VERIFY ,VIDEO_PUBLISHED, VIDEO_UNPUBLISHED
func (adminSer *adminService) updateVideoStatusForPublish(video_id uuid.UUID, status string) error {
	args := db.UpdateVerifyVideoStatusParams{
		VideoID: video_id,
		Status:  status,
	}

	err := adminSer.adminRepo.UpdateVerifyVideoStatus(args)
	return err
}

// rollback db data in case create publish_video or update video status failed
func (adminSer *adminService) rollBackCreatedPublishVideo(id uuid.UUID) error {
	return adminSer.adminRepo.RollBackCreatedPublishVideo(id)
}

// fetch all publish videos
func (adminSer *adminService) FetchAllPublishedVideos(req dto.FetchVerifyVideosRequestParams) ([]dto.FetchAllPublishedVideosDTO, error) {
	args := db.FetchAllPublishedVideosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	result, err := adminSer.adminRepo.FetchPublishedVideos(args)

	if err != nil {
		return nil, err
	}

	published_video := make([]dto.FetchAllPublishedVideosDTO, len(result))

	for i, video := range result {
		published_video[i] = dto.FetchAllPublishedVideosDTO{
			VideoID:      video.VideoID,
			Status:       video.Status,
			PublishedAt:  video.PublishedAt,
			PublishedID:  video.PublishedID,
			VideoTitle:   video.VideoTitle,
			VideoAddress: "http://localhost:8080/static/" + video.VideoAddress,
			VerifiedAt:   video.VerifiedAt,
		}
	}

	return published_video, nil
}

// unpublish the video
func (adminSer *adminService) UnPublishVideo(req dto.PublishedVideoRequestParams) error {
	video_id, err := uuid.Parse(req.VideoId)

	if err != nil {
		return err
	}

	args := db.UpdatePublishedVideoStatusParams{
		VideoID: video_id,
		Status:  helper.VIDEO_UNPUBLISHED,
	}

	err = adminSer.adminRepo.UnPublishVideo(args)

	return err
}

// fetch all un-publish Video
func (adminServ *adminService) FetchUnPublishVideos(req dto.FetchVerifyVideosRequestParams) ([]dto.FetchAllPublishedVideosDTO, error) {
	args := db.FetchAllUnPublishedVideosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	result, err := adminServ.adminRepo.FetchAllUnPublishVideo(args)

	if err != nil {
		return nil, err
	}

	unpublish_videos := make([]dto.FetchAllPublishedVideosDTO, len(result))

	for i, video := range result {
		unpublish_videos[i] = dto.FetchAllPublishedVideosDTO{
			VideoID:      video.VideoID,
			Status:       video.Status,
			PublishedAt:  video.PublishedAt,
			PublishedID:  video.PublishedID,
			VideoTitle:   video.VideoTitle,
			VideoAddress: "http://localhost:8080/static/" + video.VideoAddress,
			VerifiedAt:   video.VerifiedAt,
			Reason:       video.Reason,
		}
	}

	return unpublish_videos, nil
}

// make video verification failed
func (adminSer *adminService) MakeVerificationFailed(req dto.CreateVerificationFailedRequestParam) error {
	video_id, err := uuid.Parse(req.VideoID)

	if err != nil {
		return err
	}

	verification_failed_by, err := uuid.Parse(helper.TOKEN_ID)

	if err != nil {
		return err
	}

	args := db.CreateVerificationFailedParams{
		VideoID: video_id,
		VerificationFailedBy: uuid.NullUUID{
			UUID:  verification_failed_by,
			Valid: true,
		},
		Status: req.Status,
		Reason: req.Reason,
		IsVerificationFailed: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	err = adminSer.adminRepo.MakeVerificationFailed(args)

	if err != nil {
		return err
	}

	// update video status in verify_video, make it VIDEO_INIT to VIDEO_VIRIFICATION_FAILED
	update_args := db.UpdateVideoStatusParams{
		ID:     video_id,
		Status: helper.VIDEO_VIRIFICATION_FAILED,
	}

	// first update status in video_by_admin
	err = adminSer.adminRepo.UpdateVideoStatus(update_args)

	// if err get occured when updating a status, then roolback the db transaction, means delete recent created object
	if err != nil {
		// roll back the create data
		err_ := adminSer.adminRepo.RollBackCreateVerificationFailed(video_id)
		if err_ != nil {
			return err_
		}
		return err
	}

	return err
}

// make unpublish video object
func (adminSer *adminService) MakeUnPublishedVideo(req dto.CreateVerificationFailedRequestParam) error {
	video_id, err := uuid.Parse(req.VideoID)

	if err != nil {
		return err
	}

	unpublish_by, err := uuid.Parse(helper.TOKEN_ID)

	if err != nil {
		return err
	}

	args := db.CreateUnPublishedVideoParams{
		VideoID: video_id,
		UnpublishedBy: uuid.NullUUID{
			UUID:  unpublish_by,
			Valid: true,
		},
		Status: req.Status,
		Reason: req.Reason,
		IsVerificationFailed: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
		IsUnpublished: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	err = adminSer.adminRepo.MakeUnPublishedVideo(args)

	if err != nil {
		return err
	}

	err = adminSer.updateVideoStatusForPublish(video_id, helper.VIDEO_UNPUBLISHED)
	if err != nil {
		// rollback the create obj from video_verification_process_failed
		err_ := adminSer.adminRepo.RollBackCreateVerificationFailed(video_id)
		if err_ != nil {
			return err_
		}

		return err
	}

	return err
}

// fetch video verification failed videos
func (adminSer *adminService) FetchAllVerificationFailedVideos(req dto.FetchVerifyVideosRequestParams) ([]dto.FetchAllVerificationFailedVideosDTO, error) {
	args := db.FetchAllVerirficationFailedVideoParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	result, err := adminSer.adminRepo.FetchAllVerificationFailedVideos(args)

	if err != nil {
		return nil, err
	}

	videos := make([]dto.FetchAllVerificationFailedVideosDTO, len(result))

	for i, video := range result {
		videos[i] = dto.FetchAllVerificationFailedVideosDTO{
			VideoID:            video.VideoID,
			Status:             video.Status,
			Reason:             video.Reason,
			VerificationFailed: video.VerificationFailed.Bool,
			PublishReject:      video.PublishReject.Bool,
			FailedAt:           video.FailedAt,
			VideoTitle:         video.VideoTitle,
			VideoAddress:       "http://localhost:8080/static/" + video.VideoAddress,
			UploadedAt:         video.UploadedAt,
		}
	}

	return videos, nil
}

// -------------------------------- Fetch Video Full Details --------------------------------------- //
func (adminSer *adminService) FetchVerifyVideoFullDetails(video_id uuid.UUID) (dto.FetchVerifyVideoFullDetailsDTO, error) {
	result, err := adminSer.adminRepo.FetchVerifyVideoFullDetails(video_id)

	if err != nil {
		return dto.FetchVerifyVideoFullDetailsDTO{}, err
	}

	return dto.FetchVerifyVideoFullDetailsDTO{
		VideoID:            result.VideoID,
		VideoStatus:        result.VideoStatus,
		VerificationAt:     result.VerificationAt,
		VideoTitle:         result.VideoTitle,
		VideoAddress:       "http://localhost:8080/static/" + result.VideoAddress,
		UploadedAt:         result.UploadedAt,
		UploadedUserName:   result.UploadedUserName.String,
		UploadedUserType:   result.UploadedUserType.String,
		VerifiedbyUserName: result.VerifiedbyUserName.String,
	}, nil
}

// fetch video by admin full Details
func (adminSer *adminService) FetchVideoByAdminFullDetails(video_id uuid.UUID) (dto.FetchVideoByAdminFullDetailsDTO, error) {
	result, err := adminSer.adminRepo.FetchVideoByAdminFullDetails(video_id)
	if err != nil {
		return dto.FetchVideoByAdminFullDetailsDTO{}, err
	}

	return dto.FetchVideoByAdminFullDetailsDTO{
		VideoTitle:       result.VideoTitle,
		VideoAddress:     helper.LOCAL_ADDRESS + result.VideoAddress,
		UploadedAt:       result.UploadedAt,
		VideoID:          result.VideoID,
		UploadedUserName: result.UploadedUserName,
	}, nil
}

// fetch publish video full details
func (adminSer *adminService) FetchPublishVideoFullDetails(video_id uuid.UUID) (dto.FetchPublishVideoFullDetailsDTO, error) {
	result, err := adminSer.adminRepo.FetchPublishVideoFullDetails(video_id)

	if err != nil {
		return dto.FetchPublishVideoFullDetailsDTO{}, err
	}

	return dto.FetchPublishVideoFullDetailsDTO{
		PublishedAt:  result.PublishedAt,
		VerifiedAt:   result.VerifiedAt,
		UploadedAt:   result.UploadedAt,
		VideoTitle:   result.VideoTitle,
		VideoAddress: helper.LOCAL_ADDRESS + result.VideoAddress,
		UploadedBy:   result.UploadedBy.String,
		VerifiedBy:   result.VerifiedBy.String,
		PublishedBy:  result.PublishedBy.String,
	}, nil
}
