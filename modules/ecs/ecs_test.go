package awsspec

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

type mockECSAPI struct {
	ecsiface.ECSAPI
}

func (m mockECSAPI) ListTasks(input *ecs.ListTasksInput) (o *ecs.ListTasksOutput, err error) {
	if aws.StringValue(input.Cluster) == "test" {
		return &ecs.ListTasksOutput{
			TaskArns: []*string{aws.String("task1")},
		}, nil
	}

	if aws.StringValue(input.Cluster) == "testMulti" {
		return &ecs.ListTasksOutput{
			TaskArns: []*string{
				aws.String("task1"),
				aws.String("task2"),
			},
		}, nil
	}

	return &ecs.ListTasksOutput{TaskArns: []*string{}}, nil
}
