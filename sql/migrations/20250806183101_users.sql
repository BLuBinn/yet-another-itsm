-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    azure_ad_object_id VARCHAR(255) NOT NULL,
    home_tenant_id UUID NOT NULL,
    department_id UUID REFERENCES departments(id) ON DELETE SET NULL,
    business_unit_id UUID REFERENCES business_units(id) ON DELETE SET NULL,
    manager_id UUID REFERENCES users(id) ON DELETE SET NULL,
    mail VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    given_name VARCHAR(100),
    sur_name VARCHAR(100),
    job_title VARCHAR(255),
    office_location VARCHAR(255),
    status status_enum DEFAULT 'active',
    last_login TIMESTAMP WITH TIME ZONE,
    locked_until TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(azure_ad_object_id, mail)
);

-- Add indexes
CREATE INDEX IF NOT EXISTS idx_users_home_tenant_id ON users(home_tenant_id);
CREATE INDEX IF NOT EXISTS idx_users_mail ON users(mail);
CREATE INDEX IF NOT EXISTS idx_users_department_id ON users(department_id);
CREATE INDEX IF NOT EXISTS idx_users_manager_id ON users(manager_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_home_tenant_id;
DROP INDEX IF EXISTS idx_users_mail;
DROP INDEX IF EXISTS idx_users_department_id;
DROP INDEX IF EXISTS idx_users_manager_id;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
