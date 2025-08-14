-- name: GetFormSections :many
SELECT * FROM form_sections
WHERE form_template_id = $1 AND status = 'active' AND deleted_at IS NULL
ORDER BY section_order;

-- name: GetFormSectionByID :one
SELECT * FROM form_sections
WHERE id = $1 AND status = 'active' AND deleted_at IS NULL;

-- name: CreateFormSection :one
INSERT INTO form_sections (
    form_template_id, section_name, section_order,
    description 
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateFormSection :one
UPDATE form_sections
SET 
    section_name = COALESCE($2, section_name),
    section_order = COALESCE($3, section_order),
    description = COALESCE($4, description),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteFormSection :exec
UPDATE form_sections
SET 
    status = 'inactive',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
