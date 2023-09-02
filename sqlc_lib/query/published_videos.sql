-- name: CreatePublishedVideo :one
insert into published_videos (
    video_id,
    published_by,
    status
) values (
    $1,$2,$3
) returning *;


-- name: DeletePublishedVideo :execresult
delete from published_videos
where id = $1;

-- name: FetchAllPublishedVideos :many  
select pv.status,
pv.created_at as published_at,
pv.id as published_id,
va.title as video_title,
va.file_address as video_address,
va.created_at as verified_at
from published_videos as pv
inner join video_by_admin as va
on pv.video_id = va.id
where pv.status='VIDEO_PUBLISHED'
order by pv.created_at desc
limit $1
offset $2;