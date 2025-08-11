// Package testcontainer
package testcontainer

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
	tcredis "github.com/testcontainers/testcontainers-go/modules/redis"
)

func TestWithRedisModule(t *testing.T) {
	ctx := context.Background()

	redisContainer, err := tcredis.Run(
		ctx,
		"docker.io/library/redis:6-alpine",
		tcredis.WithSnapshotting(10, 1),
		tcredis.WithLogLevel(tcredis.LogLevelVerbose),
		// tcredis.WithConfigFile(filepath.Join("testdata", "redis7.conf")),
	)
	assert.Nil(t, err, "failed to start container")
	defer func() {
		err := tc.TerminateContainer(redisContainer)
		assert.Nil(t, err, "failed to terminate container")
	}()

	endpoint, err := redisContainer.Endpoint(ctx, "")
	assert.Nil(t, err, "fail to get endpoint")

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	_ = client
}
