package dto

import (
	"time"

	"github.com/google/uuid"
)

type UploadVideoByAdminRequestParam struct {
	Title      string    `form:"title" binding:"required"`
	UploadedBy uuid.UUID `form:"uploaded_by" binding:"required"`
	Status     string    `form:"status" binding:"required,oneof=VERIFICATION PUBLISH"`
}

type GetVideoByAdminRequestParams struct {
	UserId   string `uri:"user_id"`
	PageID   int32  `uri:"page_id" binding:"required"`
	PageSize int32  `uri:"page_size" binding:"required"`
}

type UpdateVideoStatusRequestParams struct {
	VideoId     string `form:"video_id" binding:"required"`
	VideoStatus string `form:"video_status" binding:"required,oneof=VIDEO_INIT VIDEO_VERIFY VIDEO_VIRIFICATION_FAILED VIDEO_PUBLISHED VIDEO_UNPUBLISHED"`
}

type FetchVerifyVideosRequestParams struct {
	PageID   int32 `uri:"page_id" binding:"required"`
	PageSize int32 `uri:"page_size" binding:"required"`
}

type PublishedVideoRequestParams struct {
	VideoId string `uri:"video_id" binding:"required"`
}

type FetchAllPublishedVideosDTO struct {
	VideoID      uuid.UUID `json:"video_id"`
	Status       string    `json:"status"`
	PublishedAt  time.Time `json:"published_at"`
	PublishedID  uuid.UUID `json:"published_id"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	VerifiedAt   time.Time `json:"verified_at"`
	Reason       string    `json:"reason,omitempty"`
}

type FetchAllVerificationFailedVideosDTO struct {
	VideoID            uuid.UUID `json:"video_id"`
	Status             string    `json:"status"`
	Reason             string    `json:"reason"`
	VerificationFailed bool      `json:"verification_failed"`
	PublishReject      bool      `json:"publish_reject"`
	FailedAt           time.Time `json:"failed_at"`
	VideoTitle         string    `json:"video_title"`
	VideoAddress       string    `json:"video_address"`
	UploadedAt         time.Time `json:"uploaded_at"`
}

type CreateVerificationFailedRequestParam struct {
	VideoID string `json:"video_id" binding:"required"`
	Status  string `json:"status" binding:"required,oneof=VIDEO_VIRIFICATION_FAILED VIDEO_UNPUBLISHED"`
	Reason  string `json:"reason" binding:"required"`
}

type FetchVerifyVideoFullDetailsDTO struct {
	VideoID            uuid.UUID `json:"video_id"`
	VideoStatus        string    `json:"video_status"`
	VerificationAt     time.Time `json:"verified_at"`
	VideoTitle         string    `json:"video_title"`
	VideoAddress       string    `json:"video_address"`
	UploadedAt         time.Time `json:"uploaded_at"`
	UploadedUserName   string    `json:"uploaded_by"`
	UploadedUserType   string    `json:"uploaded_user_type"`
	VerifiedbyUserName string    `json:"verified_by"`
}

type FetchVideoByAdminFullDetailsDTO struct {
	VideoTitle       string    `json:"video_title"`
	VideoAddress     string    `json:"video_address"`
	UploadedAt       time.Time `json:"uploaded_at"`
	VideoID          uuid.UUID `json:"video_id"`
	UploadedUserName string    `json:"uploaded_by"`
}

type FetchPublishVideoFullDetailsDTO struct {
	PublishedAt  time.Time `json:"published_at"`
	VerifiedAt   time.Time `json:"verified_at"`
	UploadedAt   time.Time `json:"uploaded_at"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	UploadedBy   string    `json:"uploaded_by"`
	VerifiedBy   string    `json:"verified_by"`
	PublishedBy  string    `json:"published_by"`
}

type GetAllVerifyVideos struct {
	Status       string    `json:"status"`
	VideoID      uuid.UUID `json:"video_id"`
	VerifyAt     time.Time `json:"verify_at"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	UploadedAt   time.Time `json:"uploaded_at"`
}
