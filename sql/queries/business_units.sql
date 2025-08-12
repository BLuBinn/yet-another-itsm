-- name: GetAllBusinessUnitsInTenant :many
SELECT 
    id,
    domain_name,
    tenant_id,
    name,
    status,
    created_at,
    updated_at,
    deleted_at
FROM business_units 
WHERE tenant_id = $1
ORDER BY created_at DESC;

-- name: GetBusinessUnitByID :one
SELECT 
    id,
    domain_name,
    tenant_id,
    name,
    status,
    created_at,
    updated_at,
    deleted_at
FROM business_units 
WHERE id = $1;

-- name: GetBusinessUnitByDomainName :one
SELECT 
    id,
    domain_name,
    tenant_id,
    name,
    status,
    created_at,
    updated_at,
    deleted_at
FROM business_units 
WHERE domain_name = $1;

-- name: GetAllDepartmentsInBusinessUnit :many
SELECT 
    id,
    business_unit_id,
    name,
    status,
    created_at,
    updated_at,
    deleted_at
FROM departments 
WHERE business_unit_id = $1
ORDER BY created_at DESC;
