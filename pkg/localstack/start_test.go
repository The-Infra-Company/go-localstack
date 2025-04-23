package localstack_test

import (
	"context"
	"testing"

	"github.com/The-Infra-Company/localstack-helpers/pkg/localstack"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestRunner_Start(t *testing.T) {
	ctx := context.Background()

	// initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	assert.NoError(t, err)
	defer cli.Close()

	// create a Runner and start LocalStack
	runner, err := localstack.NewRunner(cli)
	assert.NoError(t, err)

	containerID, err := runner.Start(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, containerID)

	// cleanup: stop & remove container
	err = cli.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true})
	assert.NoError(t, err)

	// cleanup: remove image
	_, err = cli.ImageRemove(ctx, runner.Image, image.RemoveOptions{Force: true})
	assert.NoError(t, err)
}
