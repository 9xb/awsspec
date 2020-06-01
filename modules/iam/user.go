package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

// UserHasPolicy checks that the provided policy ARN is attached to the specified IAM user
func (i IAMSpec) UserHasPolicy(user, policyARN string) (res bool, err error) {
	svc := getIAMAPI(i.Session)
	in := &iam.ListAttachedUserPoliciesInput{
		UserName: aws.String(user),
	}
	out, err := svc.ListAttachedUserPolicies(in)
	if err != nil {
		return
	}

	for _, v := range out.AttachedPolicies {
		if aws.StringValue(v.PolicyArn) == policyARN {
			return true, nil
		}
	}

	return
}
