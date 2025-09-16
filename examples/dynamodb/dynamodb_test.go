package test

import (
	"context"
	"testing"

	"github.com/The-Infra-Company/go-localstack"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDBWithLocalStack(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Start LocalStack container
	runner, err := localstack.NewRunner(nil)
	assert.NoError(t, err)

	containerID, err := runner.Start(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, containerID)

	// Run Terratest with Terraform options
	tfOptions := &terraform.Options{
		TerraformDir: ".",
		Upgrade:      true,
	}

	defer terraform.Destroy(t, tfOptions)
	terraform.InitAndApply(t, tfOptions)
}
