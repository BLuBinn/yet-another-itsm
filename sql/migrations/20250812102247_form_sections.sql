-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS form_sections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    form_template_id UUID NOT NULL REFERENCES form_templates(id) ON DELETE CASCADE,
    section_name VARCHAR(255) NOT NULL,
    section_order INTEGER NOT NULL CHECK (section_order > 0),
    description TEXT,
    -- conditional_logic JSONB,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(form_template_id, section_name),
    UNIQUE(form_template_id, section_order)
);

-- Add indexes
CREATE INDEX idx_form_sections_template ON form_sections(form_template_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_form_sections_status ON form_sections(status) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS form_sections;
-- +goose StatementEnd
