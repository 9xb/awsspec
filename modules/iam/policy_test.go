package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/stretchr/testify/assert"
)

var policyVer = "v1"

func TestFindDefaultPolicyVersion(t *testing.T) {
	sess, _ := session.NewSession()
	getIAMAPI = func(s *session.Session) iamiface.IAMAPI {
		return mockIAMAPI{}
	}

	i := New(sess)

	ver, err := i.findDefaultPolicyVersion("arn")
	assert.Nil(t, err)
	assert.Equal(t, policyVer, ver)

	ver, err = i.findDefaultPolicyVersion("arnFalse")
	assert.Nil(t, err)
	assert.NotEqual(t, policyVer, ver)

	ver, err = i.findDefaultPolicyVersion("err")
	assert.NotNil(t, err)
}

func TestPolicyAllows(t *testing.T) {
	sess, _ := session.NewSession()
	getIAMAPI = func(s *session.Session) iamiface.IAMAPI {
		return mockIAMAPI{}
	}

	i := New(sess)
	res, err := i.PolicyAllows("arn", []string{"action1", "action2"}, []string{"resource1"})
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = i.PolicyAllows("arn", []string{"action3"}, []string{"resource2"})
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = i.PolicyAllows("arn", []string{"action1"}, []string{"resource2"})
	assert.Nil(t, err)
	assert.False(t, res)
}
