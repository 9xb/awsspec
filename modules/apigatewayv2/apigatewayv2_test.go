package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
)

type mockAPIGatewayV2API struct {
	apigatewayv2iface.ApiGatewayV2API
}

var (
	apiID          = "test"
	stage          = "testStage"
	stageNoLog     = "testStageNoLog"
	logDestination = "testLogs"
	logFormat      = "testFormat"
	rateLimit      = 5
	burstLimit     = 10
)

func (m mockAPIGatewayV2API) GetApi(input *apigatewayv2.GetApiInput) (o *apigatewayv2.GetApiOutput, err error) {
	if aws.StringValue(input.ApiId) == apiID {
		o := &apigatewayv2.GetApiOutput{
			ApiId: aws.String(apiID),
		}

		return o, nil
	}

	return o, awserr.New("NotFoundException", "", err)
}

func (m mockAPIGatewayV2API) GetStage(input *apigatewayv2.GetStageInput) (o *apigatewayv2.GetStageOutput, err error) {
	if aws.StringValue(input.ApiId) == apiID && aws.StringValue(input.StageName) == stage {
		o := &apigatewayv2.GetStageOutput{
			StageName: aws.String(stage),
			AccessLogSettings: &apigatewayv2.AccessLogSettings{
				DestinationArn: aws.String(logDestination),
				Format:         aws.String(logFormat),
			},
			DefaultRouteSettings: &apigatewayv2.RouteSettings{
				DetailedMetricsEnabled: aws.Bool(true),
				ThrottlingBurstLimit:   aws.Int64(int64(burstLimit)),
				ThrottlingRateLimit:    aws.Float64(float64(rateLimit)),
			},
		}

		return o, nil
	}

	if aws.StringValue(input.ApiId) == apiID && aws.StringValue(input.StageName) == stageNoLog {
		o := &apigatewayv2.GetStageOutput{
			StageName: aws.String(stage),
			DefaultRouteSettings: &apigatewayv2.RouteSettings{
				DetailedMetricsEnabled: aws.Bool(true),
			},
		}

		return o, nil
	}

	if aws.StringValue(input.ApiId) == apiID {
		o := &apigatewayv2.GetStageOutput{
			StageName: aws.String(stage),
			AccessLogSettings: &apigatewayv2.AccessLogSettings{
				DestinationArn: aws.String("nope"),
				Format:         aws.String("nope"),
			},
			DefaultRouteSettings: &apigatewayv2.RouteSettings{
				DetailedMetricsEnabled: aws.Bool(false),
				ThrottlingBurstLimit:   aws.Int64(int64(100)),
				ThrottlingRateLimit:    aws.Float64(float64(200)),
			},
		}

		return o, nil
	}

	return o, awserr.New("NotFoundException", "", err)
}
