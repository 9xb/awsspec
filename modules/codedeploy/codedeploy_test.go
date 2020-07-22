package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
)

var (
	appName          = "test"
	serviceRole      = "serviceRoleARN"
	deploymentConfig = "deploymentConfig"
	events           = []string{"TEST_EVENT_1", "TEST_EVENT_2"}
	interval         = 10
	percent          = 30
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

func (m mockCodeDeployAPI) GetDeploymentConfig(input *codedeploy.GetDeploymentConfigInput) (o *codedeploy.GetDeploymentConfigOutput, err error) {
	if aws.StringValue(input.DeploymentConfigName) == "noexist" {
		return o, awserr.New("DeploymentConfigDoesNotExistException", "", err)
	}

	routing := &codedeploy.TrafficRoutingConfig{
		TimeBasedCanary: &codedeploy.TimeBasedCanary{
			CanaryInterval:   aws.Int64(int64(interval)),
			CanaryPercentage: aws.Int64(int64(percent)),
		},
		Type: aws.String("TimeBasedCanary"),
	}

	if aws.StringValue(input.DeploymentConfigName) == "testLinear" {
		routing = &codedeploy.TrafficRoutingConfig{
			TimeBasedLinear: &codedeploy.TimeBasedLinear{
				LinearInterval:   aws.Int64(int64(interval)),
				LinearPercentage: aws.Int64(int64(percent)),
			},
			Type: aws.String("TimeBasedLinear"),
		}
	}

	o = &codedeploy.GetDeploymentConfigOutput{
		DeploymentConfigInfo: &codedeploy.DeploymentConfigInfo{
			ComputePlatform:      aws.String("Test"),
			DeploymentConfigName: aws.String("TestConfig"),
			MinimumHealthyHosts: &codedeploy.MinimumHealthyHosts{
				Type:  aws.String("TEST"),
				Value: aws.Int64(4),
			},
			TrafficRoutingConfig: routing,
		},
	}
	return
}
