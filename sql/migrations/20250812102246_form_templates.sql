-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS form_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    form_category_id UUID REFERENCES form_categories(id),
    business_unit_id UUID REFERENCES business_units(id),
    version INTEGER DEFAULT 1 CHECK (version > 0),
    published_at TIMESTAMP WITH TIME ZONE,
    created_by UUID NOT NULL REFERENCES users(id),
    -- approved_by UUID REFERENCES users(id),
    -- approved_at TIMESTAMP WITH TIME ZONE,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(name, form_category_id, business_unit_id, version)
);

-- Add indexes
CREATE INDEX idx_form_templates_category ON form_templates(form_category_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_form_templates_business_unit ON form_templates(business_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_form_templates_status ON form_templates(status) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS form_templates;
-- +goose StatementEnd
