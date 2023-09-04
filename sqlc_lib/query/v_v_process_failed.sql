-- name: CreateVerificationFailed :one
insert into video_verification_process_failed (
    video_id,
    verification_failed_by,
    status,
    reason,
    is_verification_failed
) values (
    $1,$2,$3,$4,$5
) returning *;

-- name: CreateUnPublishedVideo :one 
insert into video_verification_process_failed (
    video_id,
    unpublished_by,
    status,
    reason,
    is_verification_failed,
    is_unpublished
) values (
    $1,$2,$3,$4,$5,$6
) returning *;


-- name: DeleteVerificationFailed :execresult
delete from video_verification_process_failed
where video_id = $1;

-- name: FetchAllVerirficationFailedVideo :many
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
offset $2;