package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
	"github.com/stretchr/testify/assert"
)

func TestRuleHasTarget(t *testing.T) {
	sess, _ := session.NewSession()
	getCWEventsAPI = func(s *session.Session) cloudwatcheventsiface.CloudWatchEventsAPI {
		return mockCloudWatchEventsAPI{}
	}
	rule := "testrole"
	tgt := targetARN

	e := New(sess)
	res, err := e.RuleHasTarget(rule, tgt)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = e.RuleHasTarget(rule, "nope")
	assert.Nil(t, err)
	assert.False(t, res)
}
