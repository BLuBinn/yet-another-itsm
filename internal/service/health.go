package service

import (
	"context"
	"msn-map-api/internal/constants"
	"msn-map-api/internal/database"

	"github.com/rs/zerolog/log"
)

type HealthService interface {
	CheckHealth(ctx context.Context) (*HealthStatus, error)
}

type healthService struct {
	db *database.Database
}

type HealthStatus struct {
	Status   string            `json:"status"`
	Version  string            `json:"version"`
	Services map[string]string `json:"services"`
}

func NewHealthService(db *database.Database) HealthService {
	return &healthService{db: db}
}

func (s *healthService) CheckHealth(ctx context.Context) (*HealthStatus, error) {
	log.Info().
		Str("service", "HealthService").
		Str("endpoint", "CheckHealth").
		Msg("Checking health")

	health := &HealthStatus{
		Status:   constants.HealthStatusOK,
		Version:  "1.0.0",
		Services: make(map[string]string),
	}

	// Check database health
	if err := s.db.Health(ctx); err != nil {
		log.Error().Err(err).Msg(constants.ErrDatabaseHealthCheckFailed)
		health.Status = constants.HealthStatusDegraded
		health.Services[constants.HealthServiceDatabase] = constants.HealthServiceUnhealthy
	} else {
		health.Services[constants.HealthServiceDatabase] = constants.HealthServiceHealthy
		log.Info().Msg("Database health check passed")
	}

	return health, nil
}
