package awsspec

import (
	"testing"

	s3Spec "github.com/9xb/awsspec/modules/s3"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/stretchr/testify/assert"
)

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

func TestAlarmBelongsToMetric(t *testing.T) {
	sess, _ := session.NewSession()
	getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
		return mockCloudWatchAPI{}
	}

	c := New(sess)

	res, err := c.AlarmBelongsToMetric("test", Metric{Name: metricName, Namespace: namespace})
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.AlarmBelongsToMetric("test", Metric{Name: metricName, Namespace: "test"})
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestAlarmHasAction(t *testing.T) {
	sess, _ := session.NewSession()
	getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
		return mockCloudWatchAPI{}
	}

	c := New(sess)

	res, err := c.AlarmHasAction("test", alarmAction)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.AlarmHasAction("test", "test")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestAlarmActionsEnabled(t *testing.T) {
	sess, _ := session.NewSession()
	getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
		return mockCloudWatchAPI{}
	}

	c := New(sess)

	res, err := c.AlarmActionsEnabled("test")
	assert.Nil(t, err)
	assert.True(t, res)

	actionsEnabled = false
	res, err = c.AlarmActionsEnabled("test")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestAlarmHasTag(t *testing.T) {
	sess, _ := session.NewSession()
	getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
		return mockCloudWatchAPI{}
	}

	c := New(sess)

	tag := s3Spec.Tag{
		Key:   aws.StringValue(tags[0].Key),
		Value: aws.StringValue(tags[0].Value),
	}
	res, err := c.AlarmHasTag(resourceARN, tag)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.AlarmHasTag("nope", tag)
	assert.Nil(t, err)
	assert.False(t, res)
}
