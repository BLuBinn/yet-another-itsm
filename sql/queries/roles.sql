-- name: GetAllRoles :many
SELECT 
    id,
    name,
    description,
    is_system_role,
    status,
    created_at,
    updated_at,
    deleted_at
FROM roles 
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: GetRoleByID :one
SELECT 
    id,
    name,
    description,
    is_system_role,
    status,
    created_at,
    updated_at,
    deleted_at
FROM roles 
WHERE id = $1;

-- name: GetSystemRoles :many
SELECT 
    id,
    name,
    description,
    is_system_role,
    status,
    created_at,
    updated_at,
    deleted_at
FROM roles 
WHERE is_system_role AND deleted_at IS NULL
ORDER BY name ASC;

-- name: CreateRole :one
INSERT INTO roles (
    id,
    name,
    description,
    is_system_role,
    status
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, name, description, is_system_role, status, created_at, updated_at, deleted_at;
