-- name: CreateVerifyVideo :one
insert into verify_videos (
    video_id,
    verify_by,
    status
) values (
    $1,$2,$3
) returning *;

-- name: UpdateVerifyVideoStatus :one
update verify_videos
set status = $2
where video_id = $1
returning *;

-- name: GetAllVerifyVideos :many
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
offset $2;

-- name: GetVerifyVideoFullDetails :one
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
where vv.video_id= $1;
