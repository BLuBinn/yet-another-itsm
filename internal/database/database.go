package database

import (
	"context"
	"fmt"
	"time"

	"yet-another-itsm/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	log.Info().
		Str("host", cfg.Database.Host).
		Str("port", cfg.Database.Port).
		Str("dbname", cfg.Database.DBName).
		Str("sslmode", cfg.Database.SSLMode).
		Msg("Connecting to database")

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse database config")
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// Configure connection pool
	poolConfig.MaxConns = cfg.Database.MaxConns
	poolConfig.MinConns = cfg.Database.MinConns
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = time.Minute * 30

	// Create connection pool with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create database pool")
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to ping database")
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info().
		Int32("max_conns", cfg.Database.MaxConns).
		Int32("min_conns", cfg.Database.MinConns).
		Msg("Database connection pool created successfully")

	return &Database{
		Pool: pool,
	}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		log.Info().Msg("Closing database connection pool")
		db.Pool.Close()
	}
}

func (db *Database) Health(ctx context.Context) error {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Error().Err(err).Msg("Database health check failed")
		return err
	}
	return nil
}

func (db *Database) Stats() *pgxpool.Stat {
	return db.Pool.Stat()
}
