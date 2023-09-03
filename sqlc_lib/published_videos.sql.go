// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: published_videos.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPublishedVideo = `-- name: CreatePublishedVideo :one
insert into published_videos (
    video_id,
    published_by,
    status
) values (
    $1,$2,$3
) returning id, video_id, published_by, status, created_at, updated_at
`

type CreatePublishedVideoParams struct {
	VideoID     uuid.UUID `json:"video_id"`
	PublishedBy uuid.UUID `json:"published_by"`
	Status      string    `json:"status"`
}

func (q *Queries) CreatePublishedVideo(ctx context.Context, arg CreatePublishedVideoParams) (PublishedVideos, error) {
	row := q.db.QueryRowContext(ctx, createPublishedVideo, arg.VideoID, arg.PublishedBy, arg.Status)
	var i PublishedVideos
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.PublishedBy,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePublishedVideo = `-- name: DeletePublishedVideo :execresult
delete from published_videos
where id = $1
`

func (q *Queries) DeletePublishedVideo(ctx context.Context, id uuid.UUID) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePublishedVideo, id)
}

const fetchAllPublishedVideos = `-- name: FetchAllPublishedVideos :many
select pv.status,
pv.created_at as published_at,
pv.id as published_id,
pv.video_id as video_id,
va.title as video_title,
va.file_address as video_address,
va.created_at as verified_at
from published_videos as pv
inner join video_by_admin as va
on pv.video_id = va.id
where pv.status='VIDEO_PUBLISHED'
order by pv.created_at desc
limit $1
offset $2
`

type FetchAllPublishedVideosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type FetchAllPublishedVideosRow struct {
	Status       string    `json:"status"`
	PublishedAt  time.Time `json:"published_at"`
	PublishedID  uuid.UUID `json:"published_id"`
	VideoID      uuid.UUID `json:"video_id"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	VerifiedAt   time.Time `json:"verified_at"`
}

func (q *Queries) FetchAllPublishedVideos(ctx context.Context, arg FetchAllPublishedVideosParams) ([]FetchAllPublishedVideosRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchAllPublishedVideos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchAllPublishedVideosRow{}
	for rows.Next() {
		var i FetchAllPublishedVideosRow
		if err := rows.Scan(
			&i.Status,
			&i.PublishedAt,
			&i.PublishedID,
			&i.VideoID,
			&i.VideoTitle,
			&i.VideoAddress,
			&i.VerifiedAt,
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

const fetchAllUnPublishedVideos = `-- name: FetchAllUnPublishedVideos :many
select pv.status,
pv.created_at as published_at,
pv.id as published_id,
pv.video_id as video_id,
va.title as video_title,
va.file_address as video_address,
va.created_at as verified_at
from published_videos as pv
inner join video_by_admin as va
on pv.video_id = va.id
where pv.status='VIDEO_UNPUBLISHED'
order by pv.created_at desc
limit $1
offset $2
`

type FetchAllUnPublishedVideosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type FetchAllUnPublishedVideosRow struct {
	Status       string    `json:"status"`
	PublishedAt  time.Time `json:"published_at"`
	PublishedID  uuid.UUID `json:"published_id"`
	VideoID      uuid.UUID `json:"video_id"`
	VideoTitle   string    `json:"video_title"`
	VideoAddress string    `json:"video_address"`
	VerifiedAt   time.Time `json:"verified_at"`
}

func (q *Queries) FetchAllUnPublishedVideos(ctx context.Context, arg FetchAllUnPublishedVideosParams) ([]FetchAllUnPublishedVideosRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchAllUnPublishedVideos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchAllUnPublishedVideosRow{}
	for rows.Next() {
		var i FetchAllUnPublishedVideosRow
		if err := rows.Scan(
			&i.Status,
			&i.PublishedAt,
			&i.PublishedID,
			&i.VideoID,
			&i.VideoTitle,
			&i.VideoAddress,
			&i.VerifiedAt,
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

const updatePublishedVideoStatus = `-- name: UpdatePublishedVideoStatus :one
update published_videos
set status = $2
where video_id = $1
returning id, video_id, published_by, status, created_at, updated_at
`

type UpdatePublishedVideoStatusParams struct {
	VideoID uuid.UUID `json:"video_id"`
	Status  string    `json:"status"`
}

func (q *Queries) UpdatePublishedVideoStatus(ctx context.Context, arg UpdatePublishedVideoStatusParams) (PublishedVideos, error) {
	row := q.db.QueryRowContext(ctx, updatePublishedVideoStatus, arg.VideoID, arg.Status)
	var i PublishedVideos
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.PublishedBy,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
