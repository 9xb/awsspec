package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

var (
	domainName = "sandwich.test.com"
	hostedZone = "ASIFHNCSIJDHF"
	recordType = "CNAME"
)

type mockRoute53API struct {
	route53iface.Route53API
}

func (m mockRoute53API) ListResourceRecordSets(input *route53.ListResourceRecordSetsInput) (o *route53.ListResourceRecordSetsOutput, err error) {
	if aws.StringValue(input.StartRecordName) == domainName {
		o := &route53.ListResourceRecordSetsOutput{
			ResourceRecordSets: []*route53.ResourceRecordSet{
				{
					Name: aws.String(domainName),
				},
			},
		}
		return o, nil
	}
	o = &route53.ListResourceRecordSetsOutput{}
	return o, nil
}
