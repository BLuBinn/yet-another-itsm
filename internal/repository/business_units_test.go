package repository

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func TestQueries_GetAllBusinessUnitsInTenant(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	q := New(mock)

	// Mock expected query
	rows := pgxmock.NewRows([]string{"id", "domain_name", "tenant_id", "name", "is_active", "created_at", "updated_at", "deleted_at"}).
		AddRow(
			pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true},
			"test.com",
			"test-tenant",
			"Test Business Unit",
			pgtype.Bool{Bool: true, Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Valid: false},
		)

	mock.ExpectQuery(`SELECT.*FROM business_units.*WHERE tenant_id = \$1.*ORDER BY created_at DESC`).
		WithArgs("test-tenant").
		WillReturnRows(rows)

	result, err := q.GetAllBusinessUnitsInTenant(context.Background(), "test-tenant")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test.com", result[0].DomainName)
	assert.Equal(t, "test-tenant", result[0].TenantID)
	assert.Equal(t, "Test Business Unit", result[0].Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestQueries_GetBusinessUnitByDomainName(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	q := New(mock)

	row := pgxmock.NewRows([]string{"id", "domain_name", "tenant_id", "name", "is_active", "created_at", "updated_at", "deleted_at"}).
		AddRow(
			pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true},
			"test.com",
			"test-tenant",
			"Test Business Unit",
			pgtype.Bool{Bool: true, Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Valid: false},
		)

	mock.ExpectQuery(`SELECT.*FROM business_units.*WHERE domain_name = \$1`).
		WithArgs("test.com").
		WillReturnRows(row)

	result, err := q.GetBusinessUnitByDomainName(context.Background(), "test.com")

	assert.NoError(t, err)
	assert.Equal(t, "test.com", result.DomainName)
	assert.Equal(t, "test-tenant", result.TenantID)
	assert.Equal(t, "Test Business Unit", result.Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestQueries_GetBusinessUnitByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	q := New(mock)

	testUUID := pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true}

	row := pgxmock.NewRows([]string{"id", "domain_name", "tenant_id", "name", "is_active", "created_at", "updated_at", "deleted_at"}).
		AddRow(
			testUUID,
			"test.com",
			"test-tenant",
			"Test Business Unit",
			pgtype.Bool{Bool: true, Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Valid: false},
		)

	mock.ExpectQuery(`SELECT.*FROM business_units.*WHERE id = \$1`).
		WithArgs(testUUID).
		WillReturnRows(row)

	result, err := q.GetBusinessUnitByID(context.Background(), testUUID)

	assert.NoError(t, err)
	assert.Equal(t, "test.com", result.DomainName)
	assert.Equal(t, "test-tenant", result.TenantID)
	assert.Equal(t, "Test Business Unit", result.Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestQueries_GetAllDepartmentsInBusinessUnit(t *testing.T) {
	mock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer mock.Close()

	q := New(mock)

	testUUID := pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true}

	rows := pgxmock.NewRows([]string{"id", "business_unit_id", "name", "is_active", "created_at", "updated_at", "deleted_at"}).
		AddRow(
			pgtype.UUID{Bytes: [16]byte{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, Valid: true},
			testUUID,
			"Test Department",
			pgtype.Bool{Bool: true, Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			pgtype.Timestamptz{Valid: false},
		)

	mock.ExpectQuery(`SELECT.*FROM departments.*WHERE business_unit_id = \$1.*ORDER BY created_at DESC`).
		WithArgs(testUUID).
		WillReturnRows(rows)

	result, err := q.GetAllDepartmentsInBusinessUnit(context.Background(), testUUID)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Department", result[0].Name)
	assert.Equal(t, testUUID, result[0].BusinessUnitID)

	assert.NoError(t, mock.ExpectationsWereMet())
}