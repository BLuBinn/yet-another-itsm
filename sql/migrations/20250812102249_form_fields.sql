-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS form_fields (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    form_template_id UUID NOT NULL REFERENCES form_templates(id) ON DELETE CASCADE,
    form_section_id UUID REFERENCES form_sections(id),
    field_name VARCHAR(100) NOT NULL,
    -- field_type_id UUID NOT NULL REFERENCES field_types(id),
    field_type VARCHAR(50) NOT NULL, -- name, number, text, textarea, select, multiselect, radio, checkbox, date, datetime, time, email, url, file, image, signature, barcode, qrcode, custom
    field_order INTEGER NOT NULL CHECK (field_order > 0),
    -- conditional_logic JSONB,
    -- validation_schema JSONB,
    config JSONB,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(form_template_id, form_section_id, field_name),
    UNIQUE(form_template_id, form_section_id, field_order)
);

-- id form_fields field_name field_type value

-- Add indexes
CREATE INDEX idx_form_fields_template ON form_fields(form_template_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_form_fields_section ON form_fields(form_section_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_form_fields_type ON form_fields(field_type) WHERE deleted_at IS NULL;
CREATE INDEX idx_form_fields_status ON form_fields(status) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS form_fields;
-- +goose StatementEnd
