-- name: GetDepartmentByID :one
SELECT 
    id,
    business_unit_id,
    name,
    is_active,
    created_at,
    updated_at,
    deleted_at
FROM departments 
WHERE id = $1;

-- name: GetDepartmentByName :one
SELECT 
    id,
    business_unit_id,
    name,
    is_active,
    created_at,
    updated_at,
    deleted_at
FROM departments 
WHERE name = $1 AND business_unit_id = $2;
