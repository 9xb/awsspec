package awsspec

import (
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/stretchr/testify/assert"
)

var policyVer = "v1"

type mockIAMAPI struct {
	iamiface.IAMAPI
}

func (m mockIAMAPI) GetPolicyVersion(*iam.GetPolicyVersionInput) (o *iam.GetPolicyVersionOutput, err error) {
	policy := `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Action": [
				"action1",
				"action2"
			],
			"Resource": "resource1",
			"Effect": "Allow"
		},
		{
			"Action": [
				"action3",
				"action4"
			],
			"Resource": ["resource2", "resource3"],
			"Effect": "Allow"
		}
	]
}`
	o = &iam.GetPolicyVersionOutput{
		PolicyVersion: &iam.PolicyVersion{
			Document: aws.String(url.QueryEscape(policy)),
		},
	}
	return
}

func (m mockIAMAPI) ListPolicyVersions(i *iam.ListPolicyVersionsInput) (o *iam.ListPolicyVersionsOutput, err error) {
	ver := []*iam.PolicyVersion{}

	switch aws.StringValue(i.PolicyArn) {
	case "arn":
		ver = []*iam.PolicyVersion{
			{
				VersionId:        aws.String(policyVer),
				IsDefaultVersion: aws.Bool(true),
			},
			{
				VersionId:        aws.String("v2"),
				IsDefaultVersion: aws.Bool(false),
			},
		}
	case "arnFalse":
		ver = []*iam.PolicyVersion{
			{
				VersionId:        aws.String(policyVer),
				IsDefaultVersion: aws.Bool(false),
			},
			{
				VersionId:        aws.String("v2"),
				IsDefaultVersion: aws.Bool(true),
			},
		}
	}
	o = &iam.ListPolicyVersionsOutput{
		Versions: ver,
	}
	return
}

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
