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
