package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

// RecordExists verifies that the provided values (domain, hostedZOne, recordType) match the expected values
func (r Route53Spec) RecordExists(domainName, hostedZone, recordType string) (res bool, err error) {
	svc := getRoute53API(r.Session)

	in := &route53.ListResourceRecordSetsInput{
		HostedZoneId:    aws.String(hostedZone),
		StartRecordName: aws.String(domainName),
		StartRecordType: aws.String(recordType),
	}

	out, err := svc.ListResourceRecordSets(in)

	if err != nil {
		return
	}

	if len(out.ResourceRecordSets) > 0 {
		return true, nil
	}
	return

}
