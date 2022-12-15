package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestEks(t *testing.T) {
	t.Parallel()

	// Get the path to the Terraform code that will be tested
	terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/eks")

	// Create a Terraform options struct with the required settings
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code
		TerraformDir: terraformDir,

		// Set any variables that the code requires
		Vars: map[string]interface{}{
			"cluster_name": "my-eks-cluster",
			"node_count": 3,
			"node_labels": map[string]string{
				"node-type": "worker",
			},
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

	// Use the AWS provider to get information about the EKS cluster
	eksCluster := aws.GetEksCluster(t, terraformOptions, terraformOptions.Vars["cluster_name"].(string))

	// Use the AWS provider to get information about the nodes in the EKS cluster
	eksNodeGroup := aws.GetEksNodeGroup(t, terraformOptions, terraformOptions.Vars["cluster_name"].(string), terraformOptions.Vars["node_group_name"].(string))

	// Verify that the EKS cluster has been created successfully
	assert.Equal(t, "ACTIVE", eksCluster.Status)
	assert.NotEmpty(t, eksCluster.Endpoint)
	assert.NotEmpty(t, eksCluster.CertificateAuthority.Data)

	// Verify that the correct number of nodes has been created
	assert.Equal(t, terraformOptions.Vars["node_count"], len(eksNodeGroup.Nodes))

	// Verify that the nodes have the correct labels
	for _, node := range eksNodeGroup.Nodes {
		assert.Equal(t, terraformOptions.Vars["node_labels"], node.Labels)
	}
}

