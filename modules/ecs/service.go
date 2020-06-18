package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// ServiceHasRunningTasks determines whether the number of running tasks in the provided ECS servicematches the specified count
func (e ECSSpec) ServiceHasRunningTasks(cluster, service string, count int) (res bool, err error) {
	svc := getECSAPI(e.Session)
	in := &ecs.ListTasksInput{
		Cluster:       aws.String(cluster),
		ServiceName:   aws.String(service),
		DesiredStatus: aws.String("RUNNING"),
	}

	out, err := svc.ListTasks(in)
	if err != nil {
		return
	}

	if len(out.TaskArns) == count {
		return true, nil
	}

	return
}
