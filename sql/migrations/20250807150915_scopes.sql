-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS scopes (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(id)
);

-- Add indexes
CREATE INDEX IF NOT EXISTS idx_scopes_name ON scopes(name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_scopes_name;
DROP TABLE IF EXISTS scopes;
-- +goose StatementEnd
