-- name: GetAllUsersInDepartment :many
SELECT 
    id,
    azure_ad_object_id,
    home_tenant_id,
    department_id,
    business_unit_id,
    manager_id,
    mail,
    display_name,
    given_name,
    sur_name,
    job_title,
    office_location,
    status,
    last_login,
    locked_until,
    created_at,
    updated_at,
    deleted_at
FROM users 
WHERE department_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC;

-- name: GetUserByID :one
SELECT 
    id,
    azure_ad_object_id,
    home_tenant_id,
    department_id,
    business_unit_id,
    manager_id,
    mail,
    display_name,
    given_name,
    sur_name,
    job_title,
    office_location,
    status,
    last_login,
    locked_until,
    created_at,
    updated_at,
    deleted_at
FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT 
    id,
    azure_ad_object_id,
    home_tenant_id,
    department_id,
    business_unit_id,
    manager_id,
    mail,
    display_name,
    given_name,
    sur_name,
    job_title,
    office_location,
    status,
    last_login,
    locked_until,
    created_at,
    updated_at,
    deleted_at
FROM users 
WHERE mail = $1;

-- name: CreateUser :one
INSERT INTO users (
    azure_ad_object_id,
    home_tenant_id,
    department_id,
    business_unit_id,
    manager_id,
    mail,
    display_name,
    given_name,
    sur_name,
    job_title,
    office_location,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, azure_ad_object_id, home_tenant_id, department_id, business_unit_id, manager_id, mail, display_name, given_name, sur_name, job_title, office_location, status, last_login, locked_until, created_at, updated_at, deleted_at;

-- name: UpdateUserLastLogin :exec
UPDATE users 
SET 
    last_login = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE mail = $1;
