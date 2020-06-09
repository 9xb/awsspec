package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

// RoleHasPolicy checks that the provided policy ARN is attached to the specified IAM role
func (i IAMSpec) RoleHasPolicy(role, policyARN string) (res bool, err error) {
	svc := getIAMAPI(i.Session)
	in := &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(role),
	}
	out, err := svc.ListAttachedRolePolicies(in)
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
