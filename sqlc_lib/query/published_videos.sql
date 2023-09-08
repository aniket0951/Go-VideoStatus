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
offset $2;

-- name: UpdatePublishedVideoStatus :one
update published_videos
set status = $2
where video_id = $1
returning *;

-- name: FetchAllUnPublishedVideos :many
select pv.status,
pv.created_at as published_at,
pv.id as published_id,
pv.video_id as video_id,
vvpf.reason as reason,
va.title as video_title,
va.file_address as video_address,
va.created_at as verified_at
from published_videos as pv
inner join video_by_admin as va
on pv.video_id = va.id
inner join video_verification_process_failed as vvpf 
on pv.video_id = vvpf.video_id
where pv.status='VIDEO_UNPUBLISHED'
order by pv.created_at desc
limit $1
offset $2;

-- name: GetPublishVideoFullDetails :one
select pv.created_at as published_at,
vv.created_at as verified_at,
va.created_at as uploaded_at,
va.title as video_title,
va.file_address as video_address,
u.name as uploaded_by,
uu.name as verified_by,
pu.name as published_by
from published_videos as pv
inner join verify_videos as vv
on pv.video_id = vv.video_id
inner join video_by_admin as va
on vv.video_id = va.id
left join users as u
on va.uploaded_by = u.id
left join users as uu
on vv.verify_by = uu.id
left join users as pu
on pv.published_by = pu.id
where pv.video_id = $1;