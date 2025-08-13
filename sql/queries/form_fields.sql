-- name: GetFormFields :many
SELECT * FROM form_fields
WHERE form_template_id = $1 AND status = 'active' AND deleted_at IS NULL
ORDER BY field_order;

-- name: GetFormFieldsBySection :many
SELECT * FROM form_fields
WHERE form_template_id = $1 AND form_section_id = $2 
AND status = 'active' AND deleted_at IS NULL
ORDER BY field_order;

-- name: GetFormFieldByID :one
SELECT * FROM form_fields
WHERE id = $1 AND status = 'active' AND deleted_at IS NULL;

-- name: CreateFormField :one
INSERT INTO form_fields (
    form_template_id, form_section_id, field_name,
    field_type_id, field_order, conditional_logic
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateFormField :one
UPDATE form_fields
SET 
    field_name = COALESCE($2, field_name),
    field_type_id = COALESCE($3, field_type_id),
    field_order = COALESCE($4, field_order),
    conditional_logic = COALESCE($5, conditional_logic),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteFormField :exec
UPDATE form_fields
SET 
    status = 'inactive',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

