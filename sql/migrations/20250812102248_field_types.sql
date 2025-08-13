-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS field_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type_name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    validation_schema JSONB,
    -- default_config JSONB,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);
-- Add indexes
CREATE INDEX idx_field_types_status ON field_types(status) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_field_types_name ON field_types(type_name) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS field_types;  
-- +goose StatementEnd
