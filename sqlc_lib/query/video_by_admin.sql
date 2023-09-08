-- name: UploadVideoByAdmin :one
insert into video_by_admin (
    title,
    file_address,
    uploaded_by,
    status
) values (
    $1,$2,$3,$4
) returning *;

-- name: GetVideoByAdmin :many
select * from video_by_admin
where uploaded_by = $1 and status= $2
order by created_at desc
limit $3
offset $4;

-- name: UpdateVideoStatus :execresult
update video_by_admin
set status = $2
where id = $1;

-- name: GetVideoByAdminFullDetail :one 
select va.title as video_title,
va.file_address as video_address,
va.created_at as uploaded_at,
va.id as video_id,
u.name as uploaded_user_name
from video_by_admin as va
inner join users as u
on va.uploaded_by = u.id
where va.id = $1;