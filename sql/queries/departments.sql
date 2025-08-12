-- name: GetDepartmentByID :one
SELECT 
    id,
    business_unit_id,
    name,
    status,
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
    status,
    created_at,
    updated_at,
    deleted_at
FROM departments 
WHERE name = $1 AND business_unit_id = $2;
