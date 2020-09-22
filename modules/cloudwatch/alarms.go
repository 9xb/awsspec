package awsspec

import (
	s3Spec "github.com/9xb/awsspec/modules/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type (
	// Metric represents a CloudWatch metric
	Metric struct {
		Name      string
		Namespace string
	}
)

// AlarmHasState checks whether the indicated metric alarm is in the indicated state
func (c CWSpec) AlarmHasState(alarmName, state string) (res bool, err error) {
	svc := getCWAPI(c.Session)
	in := &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []*string{aws.String(alarmName)},
		StateValue: aws.String(state),
	}
	out, err := svc.DescribeAlarms(in)
	if err != nil {
		return
	}

	if len(out.MetricAlarms) == 0 {
		return
	}

	return true, nil
}

// AlarmBelongsToMetric checks that the specified metric alarm belongs to the indicated metric
func (c CWSpec) AlarmBelongsToMetric(alarmName string, metric Metric) (res bool, err error) {
	svc := getCWAPI(c.Session)
	in := &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []*string{aws.String(alarmName)},
	}

	out, err := svc.DescribeAlarms(in)
	if err != nil {
		return
	}

	for _, v := range out.MetricAlarms {
		if aws.StringValue(v.MetricName) == metric.Name {
			if aws.StringValue(v.Namespace) == metric.Namespace {
				return true, nil
			}
		}
	}

	return
}

// AlarmHasAction checks that the specified metric alarm has the indicated actions
func (c CWSpec) AlarmHasAction(alarmName, action string) (res bool, err error) {
	svc := getCWAPI(c.Session)
	in := &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []*string{aws.String(alarmName)},
	}

	out, err := svc.DescribeAlarms(in)
	if err != nil {
		return
	}

	for _, v := range out.MetricAlarms {
		for _, a := range v.AlarmActions {
			if aws.StringValue(a) == action {
				return true, nil
			}
		}
	}

	return
}

//AlarmActionsEnabled checks that the specified metric alarm's actions are enabled
func (c CWSpec) AlarmActionsEnabled(alarmName string) (res bool, err error) {
	svc := getCWAPI(c.Session)
	in := &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []*string{aws.String(alarmName)},
	}

	out, err := svc.DescribeAlarms(in)
	if err != nil {
		return
	}

	for _, v := range out.MetricAlarms {
		if aws.BoolValue(v.ActionsEnabled) {
			return true, nil
		}
	}

	return
}

// AlarmHasTag checks whether the indicated alarm has the indicated tags
func (c CWSpec) AlarmHasTag(alarmARN string, tag s3Spec.TagGetter) (res bool, err error) {
	svc := getCWAPI(c.Session)
	in := &cloudwatch.ListTagsForResourceInput{
		ResourceARN: aws.String(alarmARN),
	}

	out, err := svc.ListTagsForResource(in)
	if err != nil {
		return
	}

	for _, t := range out.Tags {
		if aws.StringValue(t.Key) == tag.GetKey() {
			if aws.StringValue(t.Value) == tag.GetValue() {
				return true, nil
			}
		}
	}

	return

}
