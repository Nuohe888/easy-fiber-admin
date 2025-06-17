package redis

import (
	"context"
	"easy-fiber-admin/pkg/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRedis(t *testing.T) {
	// Using a test-specific Redis configuration or a mock server would be ideal.
	// For this example, we'll use a common local Redis setup.
	// Ensure you have a Redis server running on localhost:6379 for this test to pass.
	cfg := config.RedisConfig{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "", // No password for local dev instance
		DB:       0,  // Default DB
	}

	client, err := InitRedis(cfg)
	assert.NoError(t, err, "InitRedis should not return an error")
	assert.NotNil(t, client, "Redis client should not be nil")

	if client != nil {
		// Ping the server to confirm connectivity
		_, err := client.Ping(context.Background()).Result()
		assert.NoError(t, err, "Pinging Redis server should not return an error")

		// Clean up: Close the client connection
		err = client.Close()
		assert.NoError(t, err, "Closing Redis client should not return an error")
	}
}
