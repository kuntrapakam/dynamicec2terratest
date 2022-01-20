package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformAwsHelloWorldExample(t *testing.T) {
	t.Parallel()
	awsRegion := "ap-south-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: ".",

		Vars: map[string]interface{}{
			"awsRegion": awsRegion,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	instanceId := terraform.Output(t, terraformOptions, "instancee_id")
	amiId := terraform.Output(t, terraformOptions, "amii")

	publicIp := aws.GetPublicIpOfEc2Instance(t, instanceId, awsRegion)
	privateIp := aws.GetPrivateIpOfEc2Instance(t, instanceId, awsRegion)
	getTag := aws.GetTagsForEc2Instance(t, awsRegion, instanceId)
	//amiAccess := aws.GetAmiPubliclyAccessible(t, awsRegion, amiId)
	logData := aws.GetSyslogForInstance(t, instanceId, awsRegion)
	fmt.Println(publicIp)
	fmt.Println(privateIp)
	fmt.Println(getTag)
	fmt.Println(amiId)
	fmt.Println(awsRegion)
	fmt.Println(logData)
}
