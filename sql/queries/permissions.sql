-- name: GetAllPermissions :many
SELECT 
    id,
    name,
    description,
    resource,
    action,
    status,
    created_at,
    updated_at,
    deleted_at
FROM permissions 
WHERE deleted_at IS NULL
ORDER BY resource, action;

-- name: GetPermissionByID :one
SELECT 
    id,
    name,
    description,
    resource,
    action,
    status,
    created_at,
    updated_at,
    deleted_at
FROM permissions 
WHERE id = $1;

-- name: GetPermissionsByResource :many
SELECT 
    id,
    name,
    description,
    resource,
    action,
    status,
    created_at,
    updated_at,
    deleted_at
FROM permissions 
WHERE resource = $1 AND deleted_at IS NULL
ORDER BY action;

-- name: GetPermissionsByResourceAndAction :one
SELECT 
    id,
    name,
    description,
    resource,
    action,
    status,
    created_at,
    updated_at,
    deleted_at
FROM permissions 
WHERE resource = $1 AND action = $2 AND deleted_at IS NULL;

-- name: GetActivePermissions :many
SELECT 
    id,
    name,
    description,
    resource,
    action,
    status,
    created_at,
    updated_at,
    deleted_at
FROM permissions 
WHERE is_active = true AND deleted_at IS NULL
ORDER BY resource, action;

-- name: CreatePermission :one
INSERT INTO permissions (
    id,
    name,
    description,
    resource,
    action,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, name, description, resource, action, status, created_at, updated_at, deleted_at;

-- name: UpdatePermission :one
UPDATE permissions 
SET 
    name = $2,
    description = $3,
    resource = $4,
    action = $5,
    status = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, description, resource, action, status, created_at, updated_at, deleted_at;

-- name: DeletePermission :exec
UPDATE permissions 
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;