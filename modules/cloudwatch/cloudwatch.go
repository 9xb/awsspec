package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

type CWSpec struct {
	Session *session.Session
}

var getCWAPI = func(sess *session.Session) (client cloudwatchiface.CloudWatchAPI) {
	return cloudwatch.New(sess)
}

func New(s *session.Session) CWSpec {
	return CWSpec{
		Session: s,
	}
}

func (c CWSpec) AlarmHasState(alarmName, state string) (res bool, err error) {
	svc := getCWAPI(c.Session)
	input := &cloudwatch.DescribeAlarmsInput{
		AlarmNames: []*string{aws.String(alarmName)},
		StateValue: aws.String(state),
	}
	out, err := svc.DescribeAlarms(input)
	if err != nil {
		return
	}

	if len(out.MetricAlarms) == 0 {
		return
	}

	res = true
	return
}
