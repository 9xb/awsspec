package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

type mockELBV2API struct {
	elbv2iface.ELBV2API
}

var (
	resourceARN = "testARN"
	tags        = []*elbv2.Tag{
		{
			Key:   aws.String("testKey"),
			Value: aws.String("testValue"),
		},
		{
			Key:   aws.String("nope"),
			Value: aws.String("nope"),
		},
	}
)

func (m mockELBV2API) DescribeTags(input *elbv2.DescribeTagsInput) (o *elbv2.DescribeTagsOutput, err error) {
	if aws.StringValue(input.ResourceArns[0]) == resourceARN {
		t := &elbv2.DescribeTagsOutput{
			TagDescriptions: []*elbv2.TagDescription{
				{
					ResourceArn: aws.String(resourceARN),
					Tags:        tags[0:1],
				},
			},
		}

		return t, nil
	}

	t := &elbv2.DescribeTagsOutput{
		TagDescriptions: []*elbv2.TagDescription{
			{
				ResourceArn: aws.String(resourceARN),
				Tags:        tags[1:1],
			},
		},
	}

	return t, nil
}
