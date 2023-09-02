package services

import (
	"database/sql"
	"errors"
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

	args := db.UploadVideoByAdminParams{
		Title:       req.Title,
		FileAddress: file_path,
		UploadedBy:  req.UploadedBy,
		Status:      req.Status,
	}

	video, err := adminServ.adminRepo.UploadVideoByAdmin(args)

	return video, err
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
	err = adminServ.updateVideoStatusForPublish(video_id)
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

// update video status in verify_video, make VIDEO_VERIFY to VIDEO_PUBLISHED
func (adminSer *adminService) updateVideoStatusForPublish(video_id uuid.UUID) error {
	args := db.UpdateVerifyVideoStatusParams{
		VideoID: video_id,
		Status:  helper.VIDEO_PUBLISHED,
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
