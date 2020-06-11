package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
)

type mockCloudWatchEventsAPI struct {
	cloudwatcheventsiface.CloudWatchEventsAPI
}

var targetARN = "testARN"

func (m mockCloudWatchEventsAPI) ListTargetsByRule(input *cloudwatchevents.ListTargetsByRuleInput) (out *cloudwatchevents.ListTargetsByRuleOutput, err error) {
	targets := []*cloudwatchevents.Target{
		{
			Arn: aws.String(targetARN),
		},
		{
			Arn: aws.String("arn2"),
		},
	}
	return &cloudwatchevents.ListTargetsByRuleOutput{
		Targets: targets,
	}, nil
}
