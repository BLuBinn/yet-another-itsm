-- +goose Up
-- +goose StatementBegin
CREATE TYPE status_enum AS ENUM ('active', 'inactive', 'deleted');

CREATE TABLE IF NOT EXISTS business_units (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    domain_name VARCHAR(255) UNIQUE NOT NULL,
    tenant_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(domain_name, tenant_id, name)
);

-- Add indexes

CREATE INDEX IF NOT EXISTS idx_business_units_tenant_id ON business_units(tenant_id);
CREATE INDEX IF NOT EXISTS idx_business_units_domain_name ON business_units(domain_name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_business_units_tenant_id;
DROP TABLE IF EXISTS business_units;
-- +goose StatementEnd
