-- name: GetAllScopes :many
SELECT 
    id,
    name,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM scopes 
WHERE deleted_at IS NULL
ORDER BY name ASC;

-- name: GetScopeByID :one
SELECT 
    id,
    name,
    description,
    status,
    created_at,
    updated_at,
    deleted_at
FROM scopes 
WHERE id = $1;

-- name: CreateScope :one
INSERT INTO scopes (
    id,
    name,
    description,
    status
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, name, description, status, created_at, updated_at, deleted_at;
