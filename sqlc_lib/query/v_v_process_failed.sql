-- name: CreateVerificationFailed :one
insert into video_verification_process_failed (
    video_id,
    verification_failed_by,
    status,
    reason
) values (
    $1,$2,$3,$4
) returning *;

-- name: CreateUnPublishedVideo :one 
insert into video_verification_process_failed (
    video_id,
    unpublished_by,
    status,
    reason
) values (
    $1,$2,$3,$4
) returning *;


-- name: DeleteVerificationFailed :execresult
delete from video_verification_process_failed
where video_id = $1;