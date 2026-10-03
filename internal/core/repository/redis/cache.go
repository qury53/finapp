package core_repository_redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (c *Client) Get(
	ctx context.Context,
	key string,
	dst any,
) (bool, error) {
	data, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}

		return false, fmt.Errorf("failed to get cache: %w", err)
	}

	return true, json.Unmarshal(data, dst)
}

func (c *Client) Set(
	ctx context.Context,
	key string,
	value any,
) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal data for cache: %w", err)
	}

	return c.Client.Set(ctx, key, data, c.Timeout()).Err()
}

func (c *Client) Del(
	ctx context.Context,
	key string,
) error {
	return c.Client.Del(ctx, key).Err()
}
