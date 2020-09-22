package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

var (
	metricName     = "TestMetric"
	namespace      = "TestNamespace"
	actionsEnabled = true
	alarmAction    = "action1"
	resourceARN    = "testARN"
	tags           = []*cloudwatch.Tag{
		{
			Key:   aws.String("testKey"),
			Value: aws.String("testValue"),
		},
		{
			Key:   aws.String("nope"),
			Value: aws.String("nope"),
		},
	}
)

type mockCloudWatchAPI struct {
	cloudwatchiface.CloudWatchAPI
}

func (m mockCloudWatchAPI) DescribeAlarms(in *cloudwatch.DescribeAlarmsInput) (o *cloudwatch.DescribeAlarmsOutput, err error) {
	if aws.StringValue(in.StateValue) == "ALARM" {
		o = &cloudwatch.DescribeAlarmsOutput{
			MetricAlarms: []*cloudwatch.MetricAlarm{},
		}

		return
	}

	o = &cloudwatch.DescribeAlarmsOutput{
		MetricAlarms: []*cloudwatch.MetricAlarm{
			{
				MetricName:     aws.String(metricName),
				Namespace:      aws.String(namespace),
				ActionsEnabled: aws.Bool(actionsEnabled),
				AlarmActions: []*string{
					aws.String(alarmAction),
				},
				StateValue: aws.String("OK"),
			},
		},
	}

	return
}

func (m mockCloudWatchAPI) ListTagsForResource(input *cloudwatch.ListTagsForResourceInput) (o *cloudwatch.ListTagsForResourceOutput, err error) {
	if aws.StringValue(input.ResourceARN) == resourceARN {
		t := &cloudwatch.ListTagsForResourceOutput{
			Tags: tags[0:1],
		}

		return t, nil
	}

	t := &cloudwatch.ListTagsForResourceOutput{
		Tags: tags[1:1],
	}

	return t, nil

}
