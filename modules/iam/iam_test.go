package awsspec

import (
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

type mockIAMAPI struct {
	iamiface.IAMAPI
}

var (
	policyPrefix = "arn:aws:iam::aws:policy/"
	userPolicies = []string{"Test", "AnotherTest"}
)

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

func (m mockIAMAPI) ListAttachedUserPolicies(input *iam.ListAttachedUserPoliciesInput) (o *iam.ListAttachedUserPoliciesOutput, err error) {
	policies := []*iam.AttachedPolicy{}
	for _, v := range userPolicies {
		p := &iam.AttachedPolicy{
			PolicyArn:  aws.String(policyPrefix + v),
			PolicyName: aws.String(v),
		}
		policies = append(policies, p)
	}

	o = &iam.ListAttachedUserPoliciesOutput{
		AttachedPolicies: policies,
	}
	return
}
