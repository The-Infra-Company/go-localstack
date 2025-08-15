# Terratest Example

Create your test file that instantiates the Docker client, starts the Localstack container, and runs your Terratests like the following:

```go
package main

import (
	"context"
	"testing"

	"github.com/RoseSecurity/go-localstack/localstack"
	"github.com/docker/docker/client"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketWithLocalStack(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	assert.NoError(t, err)
	defer func() { _ = cli.Close() }()

	// Start LocalStack container
	runner, err := localstack.NewRunner(cli)
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
```

- Adjust the Terraform options to reflect the directory of your tests, and [ensure your provider is configured to reach the Localstack endpoint](https://docs.localstack.cloud/aws/integrations/infrastructure-as-code/terraform/#endpoint-configuration) or the `TerraformBinary` is set to `tflocal`.

- Run your tests:

```console
❯ go test -v ./...
=== RUN   TestS3BucketWithLocalStack
=== PAUSE TestS3BucketWithLocalStack
=== CONT  TestS3BucketWithLocalStack
{"status":"Pulling from localstack/localstack","id":"latest"}
{"status":"Digest: sha256:a97920288179f03638f8934fea1b6e62b96a2bcc3b69bb8ba1209c84c70281df"}
{"status":"Status: Image is up to date for localstack/localstack:latest"}
TestS3BucketWithLocalStack 2025-08-11T23:21:33-04:00 retry.go:91: terraform [init -upgrade=true]
TestS3BucketWithLocalStack 2025-08-11T23:21:33-04:00 logger.go:67: Running command terraform with args [init -upgrade=true]
TestS3BucketWithLocalStack 2025-08-11T23:21:34-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:34-04:00 logger.go:67: Initializing the backend...
TestS3BucketWithLocalStack 2025-08-11T23:21:34-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:34-04:00 logger.go:67: Initializing provider plugins...
TestS3BucketWithLocalStack 2025-08-11T23:21:34-04:00 logger.go:67: - Finding latest version of hashicorp/aws...                                       TestS3BucketWithLocalStack 2025-08-11T23:21:34-04:00 logger.go:67: - Installing hashicorp/aws v6.8.0...
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: - Installed hashicorp/aws v6.8.0 (signed by HashiCorp)
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: Terraform has created a lock file .terraform.lock.hcl to record the provider
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: selections it made above. Include this file in your version control repository
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: so that Terraform can guarantee to make the same selections by default when
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: you run "terraform init" in the future.
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: Terraform has been successfully initialized!
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: You may now begin working with Terraform. Try running "terraform plan" to see
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: any changes that are required for your infrastructure. All Terraform commands
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: should now work.
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: If you ever set or change modules or backend configuration for Terraform,
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: rerun this command to reinitialize your working directory. If you forget, other
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: commands will detect it and remind you to do so if necessary.
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 retry.go:91: terraform [apply -input=false -auto-approve -lock=false]
TestS3BucketWithLocalStack 2025-08-11T23:21:42-04:00 logger.go:67: Running command terraform with args [apply -input=false -auto-approve -lock=false]
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67: Terraform used the selected providers to generate the following execution
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67: plan. Resource actions are indicated with the following symbols:
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:   + create
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67: Terraform will perform the following actions:
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:   # aws_s3_bucket.example will be created
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:   + resource "aws_s3_bucket" "example" {
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + acceleration_status         = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + acl                         = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + arn                         = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + bucket                      = "my-tf-test-bucket"
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + bucket_domain_name          = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + bucket_prefix               = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + bucket_region               = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + bucket_regional_domain_name = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + force_destroy               = false
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + hosted_zone_id              = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + id                          = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + object_lock_enabled         = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + policy                      = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + region                      = "us-east-1"
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + request_payer               = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + tags                        = {
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:           + "Environment" = "Dev"
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:           + "Name"        = "My bucket"
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:         }
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + tags_all                    = {
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:           + "Environment" = "Dev"
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:           + "Name"        = "My bucket"
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:         }
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + website_domain              = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:       + website_endpoint            = (known after apply)
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:     }
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:51-04:00 logger.go:67: Plan: 1 to add, 0 to change, 0 to destroy.
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 logger.go:67: aws_s3_bucket.example: Creating...
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 logger.go:67: aws_s3_bucket.example: Creation complete after 1s [id=my-tf-test-bucket]
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 logger.go:67: Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 retry.go:91: terraform [destroy -auto-approve -input=false -lock=false]
TestS3BucketWithLocalStack 2025-08-11T23:21:53-04:00 logger.go:67: Running command terraform with args [destroy -auto-approve -input=false -lock=false]
TestS3BucketWithLocalStack 2025-08-11T23:21:58-04:00 logger.go:67: aws_s3_bucket.example: Refreshing state... [id=my-tf-test-bucket]
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67: Terraform used the selected providers to generate the following execution
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67: plan. Resource actions are indicated with the following symbols:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:   - destroy
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67: Terraform will perform the following actions:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:   # aws_s3_bucket.example will be destroyed
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:   - resource "aws_s3_bucket" "example" {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - arn                         = "arn:aws:s3:::my-tf-test-bucket" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - bucket                      = "my-tf-test-bucket" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - bucket_domain_name          = "my-tf-test-bucket.s3.amazonaws.com" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - bucket_region               = "us-east-1" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - bucket_regional_domain_name = "my-tf-test-bucket.s3.us-east-1.amazonaws.com" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - force_destroy               = false -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - hosted_zone_id              = "Z3AQBSTGFYJSTF" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - id                          = "my-tf-test-bucket" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - object_lock_enabled         = false -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - region                      = "us-east-1" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - request_payer               = "BucketOwner" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - tags                        = {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - "Environment" = "Dev"
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - "Name"        = "My bucket"
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:         } -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - tags_all                    = {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - "Environment" = "Dev"
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - "Name"        = "My bucket"
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:         } -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - grant {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - id          = "75aa57f09aa0c8caeab4f8c24e99d10f8e7faeebf76c078efc7c6caea54ba06a" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - permissions = [
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:               - "FULL_CONTROL",
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:             ] -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - type        = "CanonicalUser" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:         }
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - server_side_encryption_configuration {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - rule {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:               - bucket_key_enabled = false -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:               - apply_server_side_encryption_by_default {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:                   - sse_algorithm = "AES256" -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:                 }
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:             }
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:         }
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:       - versioning {
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - enabled    = false -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:           - mfa_delete = false -> null
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:         }
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:     }
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:00-04:00 logger.go:67: Plan: 0 to add, 0 to change, 1 to destroy.
TestS3BucketWithLocalStack 2025-08-11T23:22:01-04:00 logger.go:67: aws_s3_bucket.example: Destroying... [id=my-tf-test-bucket]
TestS3BucketWithLocalStack 2025-08-11T23:22:01-04:00 logger.go:67: aws_s3_bucket.example: Destruction complete after 0s
TestS3BucketWithLocalStack 2025-08-11T23:22:01-04:00 logger.go:67:
TestS3BucketWithLocalStack 2025-08-11T23:22:01-04:00 logger.go:67: Destroy complete! Resources: 1 destroyed.
TestS3BucketWithLocalStack 2025-08-11T23:22:01-04:00 logger.go:67:
--- PASS: TestS3BucketWithLocalStack (29.40s)
PASS
ok      test    29.834s
```
