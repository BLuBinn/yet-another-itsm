-- name: GetPermissionsByRole :many
SELECT 
    rp.id,
    rp.role_id,
    rp.permission_id,
    rp.scope_id,
    rp.status,
    rp.created_at,
    rp.updated_at,
    rp.deleted_at,
    r.name as role_name,
    p.name as permission_name,
    p.resource,
    p.action,
    s.name as scope_name
FROM role_permissions rp
JOIN roles r ON rp.role_id = r.id
JOIN permissions p ON rp.permission_id = p.id
LEFT JOIN scopes s ON rp.scope_id = s.id
WHERE rp.role_id = $1 AND rp.deleted_at IS NULL
ORDER BY p.resource, p.action;

-- name: GetRolePermissionByID :one
SELECT 
    rp.id,
    rp.role_id,
    rp.permission_id,
    rp.scope_id,
    rp.status,
    rp.created_at,
    rp.updated_at,
    rp.deleted_at,
    r.name as role_name,
    p.name as permission_name,
    p.resource,
    p.action,
    s.name as scope_name
FROM role_permissions rp
JOIN roles r ON rp.role_id = r.id
JOIN permissions p ON rp.permission_id = p.id
LEFT JOIN scopes s ON rp.scope_id = s.id
WHERE rp.id = $1;

-- name: CreateRolePermission :one
INSERT INTO role_permissions (
    role_id,
    permission_id,
    scope_id,
    status
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, role_id, permission_id, scope_id, status, created_at, updated_at, deleted_at;

