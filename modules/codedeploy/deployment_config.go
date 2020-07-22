package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

// DeploymentConfigExists returns true if the specified Deployment Config exist
func (c CodeDeploySpec) DeploymentConfigExists(name string) (res bool, err error) {
	_, err = getDeploymentConfig(c.Session, name)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "DeploymentConfigDoesNotExistException" {
				return false, nil
			}
		}
		return
	}

	return true, nil
}

// DeploymentConfigHasComputePlatform returns true if the compute platform for the specified Deployment Config matches the one provided
func (c CodeDeploySpec) DeploymentConfigHasComputePlatform(name, computePlatform string) (res bool, err error) {
	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	if aws.StringValue(cfg.DeploymentConfigInfo.ComputePlatform) == computePlatform {
		return true, nil
	}

	return
}

// DeploymentConfigIsCanary returns true if the traffic routing type of the specified Deployment Config is of type "TimeBasedCanary"
func (c CodeDeploySpec) DeploymentConfigIsCanary(name string) (res bool, err error) {
	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	if aws.StringValue(cfg.DeploymentConfigInfo.TrafficRoutingConfig.Type) == "TimeBasedCanary" {
		return true, nil
	}

	return
}

// DeploymentConfigIsLinear returns true if the traffic routing type of the specified Deployment Config is of type "TimeBasedLinear"
func (c CodeDeploySpec) DeploymentConfigIsLinear(name string) (res bool, err error) {
	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	if aws.StringValue(cfg.DeploymentConfigInfo.TrafficRoutingConfig.Type) == "TimeBasedLinear" {
		return true, nil
	}

	return
}

// DeploymentConfigHasCanaryInterval returns true if the canary interval of the specified Deployment Config matches the one provided
func (c CodeDeploySpec) DeploymentConfigHasCanaryInterval(name string, interval int) (res bool, err error) {
	res, err = c.DeploymentConfigIsCanary(name)
	if err != nil || !res {
		return
	}

	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	i := cfg.DeploymentConfigInfo.TrafficRoutingConfig.TimeBasedCanary.CanaryInterval
	if int(aws.Int64Value(i)) == interval {
		return true, nil
	}

	return false, nil
}

// DeploymentConfigHasCanaryPercentage returns true if the canary percentage of the specified Deployment Config matches the one provided
func (c CodeDeploySpec) DeploymentConfigHasCanaryPercentage(name string, percent int) (res bool, err error) {
	res, err = c.DeploymentConfigIsCanary(name)
	if err != nil || !res {
		return
	}

	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	i := cfg.DeploymentConfigInfo.TrafficRoutingConfig.TimeBasedCanary.CanaryPercentage
	if int(aws.Int64Value(i)) == percent {
		return true, nil
	}

	return false, nil
}

// DeploymentConfigHasLinearInterval returns true if the linear interval of the specified Deployment Config matches the one provided
func (c CodeDeploySpec) DeploymentConfigHasLinearInterval(name string, interval int) (res bool, err error) {
	res, err = c.DeploymentConfigIsLinear(name)
	if err != nil || !res {
		return
	}

	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	i := cfg.DeploymentConfigInfo.TrafficRoutingConfig.TimeBasedLinear.LinearInterval
	if int(aws.Int64Value(i)) == interval {
		return true, nil
	}

	return false, nil
}

// DeploymentConfigHasLinearPercentage returns true if the linear percentage of the specified Deployment Config matches the one provided
func (c CodeDeploySpec) DeploymentConfigHasLinearPercentage(name string, percent int) (res bool, err error) {
	res, err = c.DeploymentConfigIsLinear(name)
	if err != nil || !res {
		return
	}

	cfg, err := getDeploymentConfig(c.Session, name)
	if err != nil {
		return
	}

	i := cfg.DeploymentConfigInfo.TrafficRoutingConfig.TimeBasedLinear.LinearPercentage
	if int(aws.Int64Value(i)) == percent {
		return true, nil
	}

	return false, nil
}

func getDeploymentConfig(s *session.Session, name string) (cfg *codedeploy.GetDeploymentConfigOutput, err error) {
	svc := getCodeDeployAPI(s)
	in := &codedeploy.GetDeploymentConfigInput{
		DeploymentConfigName: aws.String(name),
	}

	cfg, err = svc.GetDeploymentConfig(in)
	if err != nil {
		return
	}

	return
}
