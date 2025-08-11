package testcontainer

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()
	req := tc.ContainerRequest{
		Image:        "docker.io/library/redis:6-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := tc.GenericContainer(
		ctx,
		tc.GenericContainerRequest{
			ProviderType:     tc.ProviderPodman,
			ContainerRequest: req,
			Started:          true,
		})
	defer tc.CleanupContainer(t, redisC)
	require.NoError(t, err)
	endpoint, err := redisC.Endpoint(ctx, "")
	if err != nil {
		t.Error(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	_ = client
}
