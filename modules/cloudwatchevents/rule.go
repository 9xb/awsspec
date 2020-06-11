package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// RuleHasTarget returns true if the specified rule has the supplied target ARN
func (c CWEventsSpec) RuleHasTarget(ruleName, targetARN string) (res bool, err error) {
	svc := getCWEventsAPI(c.Session)
	in := &cloudwatchevents.ListTargetsByRuleInput{
		Rule: aws.String(ruleName),
	}
	out, err := svc.ListTargetsByRule(in)
	if err != nil {
		return
	}

	for _, v := range out.Targets {
		if aws.StringValue(v.Arn) == targetARN {
			return true, nil
		}
	}

	return
}
