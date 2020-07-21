package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
	"github.com/stretchr/testify/assert"
)

func TestDeploymentGroupHasServiceRole(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentGroupHasServiceRole(appName, "test", serviceRole)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentGroupHasServiceRole(appName, "fail", serviceRole)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentGroupHasDeploymentConfig(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentGroupHasDeploymentConfig(appName, "test", deploymentConfig)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentGroupHasDeploymentConfig(appName, "fail", deploymentConfig)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentGroupHasAutoRollback(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentGroupHasAutoRollback(appName, "test")
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentGroupHasAutoRollback(appName, "fail")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentGroupHasAutoRollbackForEvents(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentGroupHasAutoRollbackForEvents(appName, "test", events)
	assert.Nil(t, err)
	assert.True(t, res)

	assert.Nil(t, err)
	res, err = c.DeploymentGroupHasAutoRollbackForEvents(appName, "test", []string{events[0]})
	assert.True(t, res)

	assert.Nil(t, err)
	res, err = c.DeploymentGroupHasAutoRollbackForEvents(appName, "test", []string{events[1], events[0]})
	assert.True(t, res)

	res, err = c.DeploymentGroupHasAutoRollbackForEvents(appName, "fail", events)
	assert.Nil(t, err)
	assert.False(t, res)
}
