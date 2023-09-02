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
	VideoStatus string `form:"video_status" binding:"required,oneof=VIDEO_INIT VIDEO_VERIFY VIDEO_VIRIFICATION_FAILED VIDEO_PUBLISHED"`
}

type FetchVerifyVideosRequestParams struct {
	PageID   int32 `uri:"page_id" binding:"required"`
	PageSize int32 `uri:"page_size" binding:"required"`
}

type GetAllVerifyVideos struct {
	Status       string    `json:"status"`
	VideoID      uuid.UUID `json:"video_id"`
	VerifyAt     time.Time `json:"verify_at"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	UploadedAt   time.Time `json:"uploaded_at"`
}
