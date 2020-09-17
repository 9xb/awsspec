package awsspec

import (
	s3Spec "github.com/9xb/awsspec/modules/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// ResourceHasTag verifies that the provided LoadBalancer resource (one of ALB, NLB, Target Group) has the indicated tag
func (e ELBV2Spec) ResourceHasTag(resourceARN string, tag s3Spec.TagGetter) (res bool, err error) {

	svc := getELBV2API(e.Session)
	in := &elbv2.DescribeTagsInput{
		ResourceArns: []*string{
			aws.String(resourceARN),
		},
	}

	out, err := svc.DescribeTags(in)
	if err != nil {
		return
	}

	tags := out.TagDescriptions

	for _, v := range tags {
		for _, t := range v.Tags {
			if aws.StringValue(t.Key) == tag.GetKey() {
				if aws.StringValue(t.Value) == tag.GetValue() {
					return true, nil
				}
			}
		}
	}

	return
}
