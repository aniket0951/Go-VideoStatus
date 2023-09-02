-- name: CreatePublishedVideo :one
insert into published_videos (
    video_id,
    published_by,
    status
) values (
    $1,$2,$3
) returning *;