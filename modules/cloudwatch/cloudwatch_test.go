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
