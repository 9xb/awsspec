package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
	"github.com/stretchr/testify/assert"
)

func TestDeploymentConfigExists(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigExists("test")
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigExists("noexist")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigHasComputePlatform(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigHasComputePlatform("test", "Test")
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigHasComputePlatform("test", "fail")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigIsCanary(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigIsCanary("test")
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigIsCanary("testLinear")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigIsLinear(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigIsLinear("testLinear")
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigIsLinear("test")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigHasCanaryInterval(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigHasCanaryInterval("test", interval)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigHasCanaryInterval("test", interval+1)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = c.DeploymentConfigHasCanaryInterval("testLinear", interval)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigHasCanaryPercentage(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigHasCanaryPercentage("test", percent)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigHasCanaryPercentage("test", percent+1)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = c.DeploymentConfigHasCanaryPercentage("testLinear", percent)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigHasLinearInterval(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigHasLinearInterval("testLinear", interval)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigHasLinearInterval("testLinear", interval+1)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = c.DeploymentConfigHasLinearInterval("test", interval)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestDeploymentConfigHasLinearPercentage(t *testing.T) {
	sess, _ := session.NewSession()
	getCodeDeployAPI = func(s *session.Session) codedeployiface.CodeDeployAPI {
		return mockCodeDeployAPI{}
	}

	c := New(sess)
	res, err := c.DeploymentConfigHasLinearPercentage("testLinear", percent)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = c.DeploymentConfigHasLinearPercentage("testLinear", percent+1)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = c.DeploymentConfigHasLinearPercentage("test", percent)
	assert.Nil(t, err)
	assert.False(t, res)
}
