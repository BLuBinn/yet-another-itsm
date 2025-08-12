-- name: GetUserRoleAssignments :many
SELECT 
    ra.id,
    ra.role_permissions_id,
    ra.assignee_id,
    ra.business_unit_id,
    ra.department_id,
    ra.assigned_by,
    ra.assigned_at,
    ra.expires_at,
    ra.status,
    ra.updated_at,
    ra.deleted_at,
    r.name as role_name,
    p.name as permission_name,
    p.resource,
    p.action,
    s.name as scope_name,
    bu.name as business_unit_name,
    d.name as department_name
FROM role_assignment ra
JOIN role_permissions rp ON ra.role_permissions_id = rp.id
JOIN roles r ON rp.role_id = r.id
JOIN permissions p ON rp.permission_id = p.id
LEFT JOIN scopes s ON rp.scope_id = s.id
LEFT JOIN business_units bu ON ra.business_unit_id = bu.id
LEFT JOIN departments d ON ra.department_id = d.id
WHERE ra.assignee_id = $1 AND ra.deleted_at IS NULL
ORDER BY ra.assigned_at DESC;

-- name: CreateRoleAssignment :one
INSERT INTO role_assignment (
    role_permissions_id,
    assignee_id,
    business_unit_id,
    department_id,
    assigned_by,
    expires_at,
    status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, role_permissions_id, assignee_id, business_unit_id, department_id, assigned_by, assigned_at, expires_at, status, updated_at, deleted_at;

-- name: CheckUserPermission :one
SELECT COUNT(*) > 0 as hasPermission
FROM role_assignment ra
JOIN role_permissions rp ON ra.role_permissions_id = rp.id
JOIN permissions p ON rp.permission_id = p.id
WHERE ra.assignee_id = $1 
    AND p.resource = $2 
    AND p.action = $3
    AND ra.status = 'active'
    AND ra.deleted_at IS NULL
    AND rp.status = 'active'
    AND rp.deleted_at IS NULL
    AND (ra.expires_at IS NULL OR ra.expires_at > CURRENT_TIMESTAMP);