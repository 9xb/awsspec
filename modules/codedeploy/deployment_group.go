package awsspec

import (
	iamSpec "github.com/9xb/awsspec/modules/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

// DeploymentGroupHasServiceRole returns true if the specified IAM role ARN is attached to the provided Deployment Group
func (c CodeDeploySpec) DeploymentGroupHasServiceRole(appName, id, role string) (res bool, err error) {
	svc := getCodeDeployAPI(c.Session)
	in := &codedeploy.GetDeploymentGroupInput{
		ApplicationName:     aws.String(appName),
		DeploymentGroupName: aws.String(id),
	}

	out, err := svc.GetDeploymentGroup(in)
	if err != nil {
		return
	}

	if aws.StringValue(out.DeploymentGroupInfo.ServiceRoleArn) == role {
		return true, nil
	}

	return
}

// DeploymentGroupHasDeploymentConfig returns true if the specified Deployment Configuration is associated to the provided Deployment Group
func (c CodeDeploySpec) DeploymentGroupHasDeploymentConfig(appName, id, config string) (res bool, err error) {
	svc := getCodeDeployAPI(c.Session)
	in := &codedeploy.GetDeploymentGroupInput{
		ApplicationName:     aws.String(appName),
		DeploymentGroupName: aws.String(id),
	}

	out, err := svc.GetDeploymentGroup(in)
	if err != nil {
		return
	}

	if aws.StringValue(out.DeploymentGroupInfo.DeploymentConfigName) == config {
		return true, nil
	}

	return
}

// DeploymentGroupHasAutoRollback returns true if the provided Deployment Group has auto-rollback enabled
func (c CodeDeploySpec) DeploymentGroupHasAutoRollback(appName, id string) (res bool, err error) {
	svc := getCodeDeployAPI(c.Session)
	in := &codedeploy.GetDeploymentGroupInput{
		ApplicationName:     aws.String(appName),
		DeploymentGroupName: aws.String(id),
	}

	out, err := svc.GetDeploymentGroup(in)
	if err != nil {
		return
	}

	if aws.BoolValue(out.DeploymentGroupInfo.AutoRollbackConfiguration.Enabled) {
		return true, nil
	}

	return
}

// DeploymentGroupHasAutoRollbackForEvents returns true if the provided Deployment Group is configured to auto-rollback for the specified events
func (c CodeDeploySpec) DeploymentGroupHasAutoRollbackForEvents(appName, id string, events []string) (res bool, err error) {
	svc := getCodeDeployAPI(c.Session)
	in := &codedeploy.GetDeploymentGroupInput{
		ApplicationName:     aws.String(appName),
		DeploymentGroupName: aws.String(id),
	}

	out, err := svc.GetDeploymentGroup(in)
	if err != nil {
		return
	}

	o := iamSpec.OptSlice{}
	for _, v := range aws.StringValueSlice(out.DeploymentGroupInfo.AutoRollbackConfiguration.Events) {
		o = append(o, v)
	}

	return o.Contains(events), nil
}
