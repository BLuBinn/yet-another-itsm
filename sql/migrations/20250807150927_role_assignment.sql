-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_assignment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_permissions_id UUID NOT NULL REFERENCES role_permissions(id) ON DELETE CASCADE,
    assignee_id UUID NOT NULL,
    business_unit_id UUID REFERENCES business_units(id),
    department_id UUID REFERENCES departments(id),
    assigned_by UUID REFERENCES users(id),
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE,
    status status_enum DEFAULT 'active',
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    UNIQUE(role_permissions_id, assignee_id, business_unit_id)
);

-- Add indexes
CREATE INDEX IF NOT EXISTS idx_role_assignment_role_permissions_id ON role_assignment(role_permissions_id);
CREATE INDEX IF NOT EXISTS idx_role_assignment_assignee_id ON role_assignment(assignee_id);
CREATE INDEX IF NOT EXISTS idx_role_assignment_business_unit_id ON role_assignment(business_unit_id);
CREATE INDEX IF NOT EXISTS idx_role_assignment_department_id ON role_assignment(department_id);
CREATE INDEX IF NOT EXISTS idx_role_assignment_assigned_by ON role_assignment(assigned_by);
CREATE INDEX IF NOT EXISTS idx_role_assignment_expires_at ON role_assignment(expires_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_role_assignment_role_permissions_id;
DROP INDEX IF EXISTS idx_role_assignment_assignee_id;
DROP INDEX IF EXISTS idx_role_assignment_business_unit_id;
DROP INDEX IF EXISTS idx_role_assignment_department_id;
DROP INDEX IF EXISTS idx_role_assignment_assigned_by;
DROP INDEX IF EXISTS idx_role_assignment_expires_at;
DROP TABLE IF EXISTS role_assignment;
-- +goose StatementEnd
