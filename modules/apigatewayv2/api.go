package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
)

// APIExists returns true if the specified API ID exists.
func (a APIGatewayV2Spec) APIExists(apiID string) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetApiInput{
		ApiId: aws.String(apiID),
	}

	out, err := svc.GetApi(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if aws.StringValue(out.ApiId) == apiID {
		return true, nil
	}

	return
}

// StageExists returns true if the specified stage exists for the provided API ID.
func (a APIGatewayV2Spec) StageExists(apiID, stage string) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if aws.StringValue(out.StageName) == stage {
		return true, nil
	}

	return
}

// StageHasLoggingEnabled returns true if logging is enabled for the specified stage associated to the provided API ID.
func (a APIGatewayV2Spec) StageHasLoggingEnabled(apiID, stage string) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if out.AccessLogSettings != nil {
		return true, nil
	}

	return
}

// StageHasLogDestination returns true for the specified stage associated to the provided API ID logs to the specified log group ARN.
func (a APIGatewayV2Spec) StageHasLogDestination(apiID, stage, logGroupARN string) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if out.AccessLogSettings != nil {
		if aws.StringValue(out.AccessLogSettings.DestinationArn) == logGroupARN {
			return true, nil
		}
	}

	return
}

// StageHasLogFormat returns true if the specified stage associated to the provided API ID has the specified format.
func (a APIGatewayV2Spec) StageHasLogFormat(apiID, stage, format string) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if out.AccessLogSettings != nil {
		if aws.StringValue(out.AccessLogSettings.Format) == format {
			return true, nil
		}
	}

	return
}

// StageHasDetailedMetricsEnabled returns true if detailed metrics are enabled for the specified stage associated to the provided API ID.
func (a APIGatewayV2Spec) StageHasDetailedMetricsEnabled(apiID, stage string) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if out.DefaultRouteSettings.DetailedMetricsEnabled != nil {
		if aws.BoolValue(out.DefaultRouteSettings.DetailedMetricsEnabled) {
			return true, nil
		}
	}

	return
}

// StageHasRateLimit returns true if the specified stage associated to the provided API ID has rate limiting set to the provided value.
func (a APIGatewayV2Spec) StageHasRateLimit(apiID, stage string, limit int) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if out.DefaultRouteSettings.ThrottlingRateLimit != nil {
		if aws.Float64Value(out.DefaultRouteSettings.ThrottlingRateLimit) == float64(limit) {
			return true, nil
		}
	}

	return
}

// StageHasBurstLimit returns true if the specified stage associated to the provided API ID has burst limiting set to the provided value.
func (a APIGatewayV2Spec) StageHasBurstLimit(apiID, stage string, limit int) (res bool, err error) {
	svc := getAPIGatewayV2API(a.Session)
	in := &apigatewayv2.GetStageInput{
		ApiId:     aws.String(apiID),
		StageName: aws.String(stage),
	}

	out, err := svc.GetStage(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFoundException" {
				return false, nil
			}
		}
		return
	}

	if out.DefaultRouteSettings.ThrottlingBurstLimit != nil {
		if aws.Int64Value(out.DefaultRouteSettings.ThrottlingBurstLimit) == int64(limit) {
			return true, nil
		}
	}

	return
}
