-- name: GetFormCategories :many
SELECT * FROM form_categories
WHERE status = 'active' AND deleted_at IS NULL
ORDER BY name;

-- name: GetFormCategoryByID :one
SELECT * FROM form_categories
WHERE id = $1 AND status = 'active' AND deleted_at IS NULL;

-- name: CreateFormCategory :one
INSERT INTO form_categories (
    name, description
) VALUES ($1, $2)
RETURNING *;

-- name: UpdateFormCategory :one
UPDATE form_categories
SET 
    name = COALESCE($2, name),
    description = COALESCE($3, description),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteFormCategory :exec
UPDATE form_categories
SET 
    status = 'inactive',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
