-- name: CreateAdminUser :one
insert into users (
    name,
    email,
    contact,
    password,
    user_type,
    is_account_active
) values (
    $1,$2,$3,$4,$5,$6
) returning *;

-- name: GetUser :one
select * from users
where ID = $1 limit 1;


-- name: GetUsers :many
select * from users
order by ID 
limit $1
offset $2;

-- name: DeleteUser :execresult
delete from users
where ID = $1;

-- name: GetUserByEmail :one 
select * from users
where email = $1 limit 1;

-- name: UpdateUserAccountActive :one
update users
set is_account_active = $2
where ID = $1
returning *; 