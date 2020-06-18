package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/stretchr/testify/assert"
)

func TestServiceHasRunningTasks(t *testing.T) {
	sess, _ := session.NewSession()
	getECSAPI = func(s *session.Session) ecsiface.ECSAPI {
		return mockECSAPI{}
	}

	e := New(sess)
	res, err := e.ServiceHasRunningTasks("test", "test", 1)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = e.ServiceHasRunningTasks("testMulti", "test", 2)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = e.ServiceHasRunningTasks("test", "test", 2)
	assert.Nil(t, err)
	assert.False(t, res)
}
