// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: verify_videos.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createVerifyVideo = `-- name: CreateVerifyVideo :one
insert into verify_videos (
    video_id,
    verify_by,
    status
) values (
    $1,$2,$3
) returning id, video_id, verify_by, status, created_at, updated_at
`

type CreateVerifyVideoParams struct {
	VideoID  uuid.UUID `json:"video_id"`
	VerifyBy uuid.UUID `json:"verify_by"`
	Status   string    `json:"status"`
}

func (q *Queries) CreateVerifyVideo(ctx context.Context, arg CreateVerifyVideoParams) (VerifyVideos, error) {
	row := q.db.QueryRowContext(ctx, createVerifyVideo, arg.VideoID, arg.VerifyBy, arg.Status)
	var i VerifyVideos
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.VerifyBy,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllVerifyVideos = `-- name: GetAllVerifyVideos :many
select vv.status,
vv.id,
vv.video_id as video_id,
vv.created_at as verify_at,
va.title as video_title,
va.file_address as video_address,
va.created_at as uploaded_at
from verify_videos as vv 
inner join video_by_admin as va 
on vv.video_id = va.id
where vv.status='VIDEO_VERIFY'
order by vv.created_at desc 
limit $1
offset $2
`

type GetAllVerifyVideosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllVerifyVideosRow struct {
	Status       string    `json:"status"`
	ID           uuid.UUID `json:"id"`
	VideoID      uuid.UUID `json:"video_id"`
	VerifyAt     time.Time `json:"verify_at"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	UploadedAt   time.Time `json:"uploaded_at"`
}

func (q *Queries) GetAllVerifyVideos(ctx context.Context, arg GetAllVerifyVideosParams) ([]GetAllVerifyVideosRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllVerifyVideos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllVerifyVideosRow{}
	for rows.Next() {
		var i GetAllVerifyVideosRow
		if err := rows.Scan(
			&i.Status,
			&i.ID,
			&i.VideoID,
			&i.VerifyAt,
			&i.VideoTitle,
			&i.VideoAddress,
			&i.UploadedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVerifyVideoFullDetails = `-- name: GetVerifyVideoFullDetails :one
select vv.video_id as video_id,
vv.status as video_status,
vv.created_at as verification_at,
va.title as video_title,
va.file_address as video_address,
va.created_at as uploaded_at,
u.name as uploaded_user_name,
u.user_type as uploaded_user_type,
uu.name as verifiedby_user_name
from verify_videos as vv
inner join video_by_admin as va
on vv.video_id = va.id
left join users as u
on va.uploaded_by = u.id
left join users as uu
on vv.verify_by = uu.id
where vv.video_id= $1
`

type GetVerifyVideoFullDetailsRow struct {
	VideoID            uuid.UUID      `json:"video_id"`
	VideoStatus        string         `json:"video_status"`
	VerificationAt     time.Time      `json:"verification_at"`
	VideoTitle         string         `json:"video_title"`
	VideoAddress       string         `json:"video_address"`
	UploadedAt         time.Time      `json:"uploaded_at"`
	UploadedUserName   sql.NullString `json:"uploaded_user_name"`
	UploadedUserType   sql.NullString `json:"uploaded_user_type"`
	VerifiedbyUserName sql.NullString `json:"verifiedby_user_name"`
}

func (q *Queries) GetVerifyVideoFullDetails(ctx context.Context, videoID uuid.UUID) (GetVerifyVideoFullDetailsRow, error) {
	row := q.db.QueryRowContext(ctx, getVerifyVideoFullDetails, videoID)
	var i GetVerifyVideoFullDetailsRow
	err := row.Scan(
		&i.VideoID,
		&i.VideoStatus,
		&i.VerificationAt,
		&i.VideoTitle,
		&i.VideoAddress,
		&i.UploadedAt,
		&i.UploadedUserName,
		&i.UploadedUserType,
		&i.VerifiedbyUserName,
	)
	return i, err
}

const updateVerifyVideoStatus = `-- name: UpdateVerifyVideoStatus :one
update verify_videos
set status = $2
where video_id = $1
returning id, video_id, verify_by, status, created_at, updated_at
`

type UpdateVerifyVideoStatusParams struct {
	VideoID uuid.UUID `json:"video_id"`
	Status  string    `json:"status"`
}

func (q *Queries) UpdateVerifyVideoStatus(ctx context.Context, arg UpdateVerifyVideoStatusParams) (VerifyVideos, error) {
	row := q.db.QueryRowContext(ctx, updateVerifyVideoStatus, arg.VideoID, arg.Status)
	var i VerifyVideos
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.VerifyBy,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
