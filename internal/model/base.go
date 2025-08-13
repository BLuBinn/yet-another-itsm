package model

import (
	"yet-another-itsm/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

const (
	UserStatusActive   = string(repository.StatusEnumActive)
	UserStatusInactive = string(repository.StatusEnumInactive)
	UserStatusDeleted  = string(repository.StatusEnumDeleted)
)

type BaseModel struct {
	ID        string             `json:"id"`
	Status    pgtype.Text        `json:"status"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
