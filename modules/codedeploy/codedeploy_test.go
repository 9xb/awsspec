package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
)

var (
	appName          = "test"
	serviceRole      = "serviceRoleARN"
	deploymentConfig = "deploymentConfig"
	events           = []string{"TEST_EVENT_1", "TEST_EVENT_2"}
)

type mockCodeDeployAPI struct {
	codedeployiface.CodeDeployAPI
}

func (m mockCodeDeployAPI) GetDeploymentGroup(input *codedeploy.GetDeploymentGroupInput) (o *codedeploy.GetDeploymentGroupOutput, err error) {
	o = &codedeploy.GetDeploymentGroupOutput{
		DeploymentGroupInfo: &codedeploy.DeploymentGroupInfo{
			ServiceRoleArn:       aws.String(serviceRole),
			DeploymentConfigName: aws.String(deploymentConfig),
			AutoRollbackConfiguration: &codedeploy.AutoRollbackConfiguration{
				Enabled: aws.Bool(true),
				Events:  aws.StringSlice(events),
			},
		},
	}

	if aws.StringValue(input.DeploymentGroupName) == "fail" {
		o = &codedeploy.GetDeploymentGroupOutput{
			DeploymentGroupInfo: &codedeploy.DeploymentGroupInfo{
				ServiceRoleArn:       aws.String("nope"),
				DeploymentConfigName: aws.String("nope"),
				AutoRollbackConfiguration: &codedeploy.AutoRollbackConfiguration{
					Enabled: aws.Bool(false),
					Events:  aws.StringSlice([]string{"NOPE_EVENT"}),
				},
			},
		}
	}
	return
}
