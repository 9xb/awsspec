package awsspec

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

type mockLambdaAPI struct {
	lambdaiface.LambdaAPI
}

var (
	functionName  = "test"
	fnWithEnv     = "testEnv"
	fnWithRole    = "testRole"
	fnWithRuntime = "testProvider"
	fnWithTimeout = "testTimeout"
	fnWithMemSize = "testMemSize"
	fnWithHandler = "testHandler"
	fnWithLayers  = "testLayers"
	fnWithSubnets = "testSubnets"
	fnWithSGs     = "testSGs"
	fnWithVPC     = "testVPC"

	qualifier   = "alias"
	envVarName  = "VAR_NAME"
	envVarValue = "var_value"
	roleName    = "test_role"
	runtime     = "test_runtime"
	timeout     = 15
	memSize     = 256
	handler     = "chandler"
	testSlice   = []string{
		"test_one",
		"test_two",
	}
	vpc       = "test_vpc"
	sourceARN = "testARN"
)

func (m mockLambdaAPI) GetFunction(input *lambda.GetFunctionInput) (o *lambda.GetFunctionOutput, err error) {
	name := aws.StringValue(input.FunctionName)
	qual := aws.StringValue(input.Qualifier)

	out := &lambda.GetFunctionOutput{
		Configuration: &lambda.FunctionConfiguration{
			FunctionName: input.FunctionName,
			Environment: &lambda.EnvironmentResponse{
				Variables: map[string]*string{
					envVarName: aws.String("nope"),
				},
			},
			Role:       aws.String("nope"),
			Runtime:    aws.String("nope"),
			Timeout:    aws.Int64(0),
			MemorySize: aws.Int64(0),
			Handler:    aws.String("nope"),
			Layers: []*lambda.Layer{
				{Arn: aws.String(testSlice[0])},
			},
			VpcConfig: &lambda.VpcConfigResponse{
				SecurityGroupIds: []*string{aws.String(testSlice[0])},
				SubnetIds:        []*string{aws.String(testSlice[0])},
				VpcId:            aws.String("nope"),
			},
		},
	}

	if name == functionName || (name == functionName && qual == qualifier) {
		return out, nil
	}

	switch name {
	case fnWithEnv:
		out.Configuration.Environment.Variables[envVarName] = aws.String(envVarValue)
		return out, nil

	case fnWithRole:
		out.Configuration.Role = aws.String(roleName)
		return out, nil

	case fnWithRuntime:
		out.Configuration.Runtime = aws.String(runtime)
		return out, nil

	case fnWithTimeout:
		out.Configuration.Timeout = aws.Int64(int64(timeout))
		return out, nil

	case fnWithMemSize:
		out.Configuration.MemorySize = aws.Int64(int64(memSize))
		return out, nil

	case fnWithHandler:
		out.Configuration.Handler = aws.String(handler)
		return out, nil

	case fnWithLayers:
		l := []*lambda.Layer{}
		for _, v := range testSlice {
			l = append(l, &lambda.Layer{Arn: aws.String(v)})
		}
		out.Configuration.Layers = l
		return out, nil

	case fnWithVPC:
		out.Configuration.VpcConfig.VpcId = aws.String(vpc)
		return out, nil

	case fnWithSubnets:
		s := []*string{}
		for _, v := range testSlice {
			s = append(s, aws.String(v))
		}
		out.Configuration.VpcConfig.SubnetIds = s
		return out, nil

	case fnWithSGs:
		s := []*string{}
		for _, v := range testSlice {
			s = append(s, aws.String(v))
		}
		out.Configuration.VpcConfig.SecurityGroupIds = s
		return out, nil
	}

	return o, awserr.New("ResourceNotFoundException", "", err)
}

func (m mockLambdaAPI) GetPolicy(input *lambda.GetPolicyInput) (o *lambda.GetPolicyOutput, err error) {
	if aws.StringValue(input.FunctionName) != functionName {
		sourceARN = sourceARN + "s"
	}

	policy := `{
		"Version": "2012-10-17",
		"Id": "default",
		"Statement": [{
			"Sid": "api-gw",
			"Effect": "Allow",
			"Principal": {
				"Service": "apigateway.amazonaws.com"
			},
			"Action": "lambda:InvokeFunction",
			"Resource": "arn:aws:lambda:eu-west-1:357027635596:function:peracto-api-qa-api:live",
			"Condition": {
				"ArnLike": {
					"AWS:SourceArn": "` + sourceARN + `"
				}
			}
		}]
	}`

	if aws.StringValue(input.FunctionName) == "nope" {
		return o, errors.New("nope")
	}

	o = &lambda.GetPolicyOutput{
		Policy: aws.String(policy),
	}

	return
}
