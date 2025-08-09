package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type BaseModel struct {
	ID        pgtype.UUID        `json:"id"`
	IsActive  pgtype.Bool        `json:"is_active"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
