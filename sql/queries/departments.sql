-- name: GetDepartmentByID :one
SELECT 
    id,
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
    name,
    status,
    created_at,
    updated_at,
    deleted_at
FROM departments
WHERE name = $1;

-- name: CreateDepartment :one
INSERT INTO departments (
    name,
    status
) VALUES (
    $1,
    $2
) RETURNING 
    id,
    name,
    status,
    created_at,
    updated_at,
    deleted_at;
