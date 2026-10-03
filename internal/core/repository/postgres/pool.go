package core_repository_pgx

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
	timeout time.Duration
}

// Создает новый pool подключения к БД по заданному конфигу
func NewConnectionPool(ctx context.Context, config *Config) (*Pool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DB,
	)

	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pgx ping error: %w", err)
	}

	return &Pool{
		Pool:    pool,
		timeout: config.Timeout,
	}, nil
}

func (p *Pool) Timeout() time.Duration {
	return p.timeout
}
