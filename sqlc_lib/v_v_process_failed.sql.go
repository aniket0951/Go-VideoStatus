// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: v_v_process_failed.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUnPublishedVideo = `-- name: CreateUnPublishedVideo :one
insert into video_verification_process_failed (
    video_id,
    unpublished_by,
    status,
    reason,
    is_verification_failed,
    is_unpublished
) values (
    $1,$2,$3,$4,$5,$6
) returning "Id", video_id, verification_failed_by, unpublished_by, status, reason, is_verification_failed, is_unpublished, created_at, updated_at
`

type CreateUnPublishedVideoParams struct {
	VideoID              uuid.UUID     `json:"video_id"`
	UnpublishedBy        uuid.NullUUID `json:"unpublished_by"`
	Status               string        `json:"status"`
	Reason               string        `json:"reason"`
	IsVerificationFailed sql.NullBool  `json:"is_verification_failed"`
	IsUnpublished        sql.NullBool  `json:"is_unpublished"`
}

func (q *Queries) CreateUnPublishedVideo(ctx context.Context, arg CreateUnPublishedVideoParams) (VideoVerificationProcessFailed, error) {
	row := q.db.QueryRowContext(ctx, createUnPublishedVideo,
		arg.VideoID,
		arg.UnpublishedBy,
		arg.Status,
		arg.Reason,
		arg.IsVerificationFailed,
		arg.IsUnpublished,
	)
	var i VideoVerificationProcessFailed
	err := row.Scan(
		&i.Id,
		&i.VideoID,
		&i.VerificationFailedBy,
		&i.UnpublishedBy,
		&i.Status,
		&i.Reason,
		&i.IsVerificationFailed,
		&i.IsUnpublished,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createVerificationFailed = `-- name: CreateVerificationFailed :one
insert into video_verification_process_failed (
    video_id,
    verification_failed_by,
    status,
    reason,
    is_verification_failed
) values (
    $1,$2,$3,$4,$5
) returning "Id", video_id, verification_failed_by, unpublished_by, status, reason, is_verification_failed, is_unpublished, created_at, updated_at
`

type CreateVerificationFailedParams struct {
	VideoID              uuid.UUID     `json:"video_id"`
	VerificationFailedBy uuid.NullUUID `json:"verification_failed_by"`
	Status               string        `json:"status"`
	Reason               string        `json:"reason"`
	IsVerificationFailed sql.NullBool  `json:"is_verification_failed"`
}

func (q *Queries) CreateVerificationFailed(ctx context.Context, arg CreateVerificationFailedParams) (VideoVerificationProcessFailed, error) {
	row := q.db.QueryRowContext(ctx, createVerificationFailed,
		arg.VideoID,
		arg.VerificationFailedBy,
		arg.Status,
		arg.Reason,
		arg.IsVerificationFailed,
	)
	var i VideoVerificationProcessFailed
	err := row.Scan(
		&i.Id,
		&i.VideoID,
		&i.VerificationFailedBy,
		&i.UnpublishedBy,
		&i.Status,
		&i.Reason,
		&i.IsVerificationFailed,
		&i.IsUnpublished,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteVerificationFailed = `-- name: DeleteVerificationFailed :execresult
delete from video_verification_process_failed
where video_id = $1
`

func (q *Queries) DeleteVerificationFailed(ctx context.Context, videoID uuid.UUID) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteVerificationFailed, videoID)
}

const fetchAllVerirficationFailedVideo = `-- name: FetchAllVerirficationFailedVideo :many
select vvpf.video_id as video_id,
vvpf.status as status,
vvpf.reason as reason ,vvpf.is_verification_failed as verification_failed,
vvpf.is_unpublished as publish_reject,
vvpf.created_at as failed_at,
va.title as video_title,
va.file_address as video_address,
va.created_at as uploaded_at
from video_verification_process_failed as vvpf
inner join video_by_admin as va 
on vvpf.video_id = va.id 
order by vvpf.created_at
limit $1
offset $2
`

type FetchAllVerirficationFailedVideoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type FetchAllVerirficationFailedVideoRow struct {
	VideoID            uuid.UUID    `json:"video_id"`
	Status             string       `json:"status"`
	Reason             string       `json:"reason"`
	VerificationFailed sql.NullBool `json:"verification_failed"`
	PublishReject      sql.NullBool `json:"publish_reject"`
	FailedAt           time.Time    `json:"failed_at"`
	VideoTitle         string       `json:"video_title"`
	VideoAddress       string       `json:"video_address"`
	UploadedAt         time.Time    `json:"uploaded_at"`
}

func (q *Queries) FetchAllVerirficationFailedVideo(ctx context.Context, arg FetchAllVerirficationFailedVideoParams) ([]FetchAllVerirficationFailedVideoRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchAllVerirficationFailedVideo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchAllVerirficationFailedVideoRow{}
	for rows.Next() {
		var i FetchAllVerirficationFailedVideoRow
		if err := rows.Scan(
			&i.VideoID,
			&i.Status,
			&i.Reason,
			&i.VerificationFailed,
			&i.PublishReject,
			&i.FailedAt,
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
