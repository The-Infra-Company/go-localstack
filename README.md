<img width="1228" height="203" alt="Screenshot_2025-08-11_at_11 48 02_PM-removebg-preview" src="https://github.com/user-attachments/assets/5bcfafe2-f2b9-467a-b45e-98c12f3c7db8" />

Go client for managing LocalStack in testing environments.

## Example

```go
package test

import (
	"context"
	"testing"

	"github.com/The-Infra-Company/go-localstack"
	"github.com/docker/docker/client"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformWithLocalStack(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

    // Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	assert.NoError(t, err)
	defer func() { _ = cli.Close() }()

    // Start LocalStack
	runner, err := localstack.NewRunner(cli)
	assert.NoError(t, err)

	containerID, err := runner.Start(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, containerID)

    // Run Terratests
	tfOptions := &terraform.Options{
		TerraformDir: "../../examples/complete",
		Upgrade:      true,
		VarFiles:     []string{"fixtures.us-east-2.tfvars"},
	}

	defer terraform.Destroy(t, tfOptions)
	terraform.InitAndApply(t, tfOptions)
}
```
