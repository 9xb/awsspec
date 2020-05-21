package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/stretchr/testify/assert"
)

type mockCloudWatchAPI struct {
	cloudwatchiface.CloudWatchAPI
}

func (m mockCloudWatchAPI) DescribeAlarms(in *cloudwatch.DescribeAlarmsInput) (o *cloudwatch.DescribeAlarmsOutput, err error) {
	alarm := []*cloudwatch.MetricAlarm{}
	if aws.StringValue(in.StateValue) == "OK" {
		alarm = append(alarm, &cloudwatch.MetricAlarm{StateValue: aws.String("OK")})
	}

	o = &cloudwatch.DescribeAlarmsOutput{
		MetricAlarms: alarm,
	}

	return
}

func TestAlarmHasState(t *testing.T) {
	sess, _ := session.NewSession()
	getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
		return mockCloudWatchAPI{}
	}

	c := New(sess)

	res, err := c.AlarmHasState("test", "OK")
	assert.Nil(t, err)
	assert.True(t, res)

	res, _ = c.AlarmHasState("test", "ALARM")
	assert.False(t, res)
}
