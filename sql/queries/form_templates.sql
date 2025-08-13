-- name: GetFormTemplates :many
SELECT * FROM form_templates
WHERE status = 'active' AND deleted_at IS NULL
ORDER BY created_at DESC;

-- name: GetFormTemplateByID :one
SELECT * FROM form_templates
WHERE id = $1 AND status = 'active' AND deleted_at IS NULL;

-- name: GetFormTemplatesByCategory :many
SELECT * FROM form_templates
WHERE form_category_id = $1 AND status = 'active' AND deleted_at IS NULL
ORDER BY created_at DESC;

-- name: CreateFormTemplate :one
INSERT INTO form_templates (
    name, description, form_category_id, business_unit_id,
    version, created_by
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateFormTemplate :one
UPDATE form_templates
SET 
    name = COALESCE($2, name),
    description = COALESCE($3, description),
    form_category_id = COALESCE($4, form_category_id),
    business_unit_id = COALESCE($5, business_unit_id),
    version = COALESCE($6, version),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: PublishFormTemplate :one
UPDATE form_templates
SET 
    published_at = CURRENT_TIMESTAMP,
    approved_by = $2,
    approved_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteFormTemplate :exec
UPDATE form_templates
SET 
    status = 'inactive',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

