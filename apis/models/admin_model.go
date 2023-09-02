package models

import (
	"time"

	"github.com/google/uuid"
)

type VideoByAdmin struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	FileAddress string    `json:"file_address"`
	UploadedBy  uuid.UUID `json:"uploaded_by"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (videoAdmin *VideoByAdmin) SetFileAddress(address string) {
	videoAdmin.FileAddress = address + videoAdmin.FileAddress
}
