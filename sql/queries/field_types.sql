-- name: GetFieldTypes :many
SELECT * FROM field_types
WHERE status = 'active' AND deleted_at IS NULL
ORDER BY type_name;

-- name: GetFieldTypeByID :one
SELECT * FROM field_types
WHERE id = $1 AND status = 'active' AND deleted_at IS NULL;

-- name: CreateFieldType :one
INSERT INTO field_types (
    type_name, description, validation_schema
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateFieldType :one
UPDATE field_types
SET 
    type_name = COALESCE($2, type_name),
    description = COALESCE($3, description),
    validation_schema = COALESCE($4, validation_schema),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteFieldType :exec
UPDATE field_types
SET 
    status = 'inactive',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
