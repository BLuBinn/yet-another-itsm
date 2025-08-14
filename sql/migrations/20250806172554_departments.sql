-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- business_unit_id UUID REFERENCES business_units(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    -- UNIQUE(business_unit_id, name)
    UNIQUE(name)
);

-- Add indexes
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS departments;
-- +goose StatementEnd
