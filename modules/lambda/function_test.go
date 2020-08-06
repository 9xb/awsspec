package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/stretchr/testify/assert"
)

func TestFunctionExists(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionExists(functionName, "")
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionExists("nope", "")
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionExists(functionName, qualifier)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestFunctionHasEnvVarValue(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasEnvVarValue(fnWithEnv, "", envVarName, envVarValue)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasEnvVarValue(fnWithEnv, "", envVarName, "test")
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = l.FunctionHasEnvVarValue("nope", "", envVarName, envVarValue)
	assert.NotNil(t, err)
}

func TestFunctionHasRole(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasRole(fnWithRole, "", roleName)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasRole(functionName, "", roleName)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasRole("nope", "", roleName)
	assert.NotNil(t, err)
}

func TestFunctionHasRuntime(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasRuntime(fnWithRuntime, "", runtime)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasRuntime(functionName, "", runtime)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasRuntime("nope", "", runtime)
	assert.NotNil(t, err)
}

func TestFunctionHasTimeout(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasTimeout(fnWithTimeout, "", timeout)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasTimeout(functionName, "", timeout)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasTimeout("nope", "", timeout)
	assert.NotNil(t, err)
}

func TestFunctionHasMemorySize(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasMemorySize(fnWithMemSize, "", memSize)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasMemorySize(functionName, "", memSize)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasMemorySize("nope", "", memSize)
	assert.NotNil(t, err)
}

func TestFunctionHasHandler(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasHandler(fnWithHandler, "", handler)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasHandler(functionName, "", handler)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasHandler("nope", "", handler)
	assert.NotNil(t, err)
}

func TestFunctionHasLayers(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasLayers(fnWithLayers, "", testSlice)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasLayers(functionName, "", testSlice)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasLayers("nope", "", testSlice)
	assert.NotNil(t, err)
}

func TestFunctionHasVPCID(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasVPCID(fnWithVPC, "", vpc)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasVPCID(functionName, "", vpc)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasVPCID("nope", "", vpc)
	assert.NotNil(t, err)
}

func TestFunctionHasVPCWithSubnets(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasVPCWithSubnets(fnWithSubnets, "", testSlice)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasVPCWithSubnets(functionName, "", testSlice)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasVPCWithSubnets("nope", "", testSlice)
	assert.NotNil(t, err)
}

func TestFunctionHasVPCWithSecurityGroups(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasVPCWithSecurityGroups(fnWithSGs, "", testSlice)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasVPCWithSecurityGroups(functionName, "", testSlice)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionHasVPCWithSecurityGroups("nope", "", testSlice)
	assert.NotNil(t, err)
}

func TestFunctionHasPermissions(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionHasPermissions(functionName, "", sourceARN)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionHasPermissions(functionName+"s", "", sourceARN)
	assert.Nil(t, err)
	assert.False(t, res)

	_, err = l.FunctionHasPermissions("nope", "", sourceARN)
	assert.NotNil(t, err)
}

func TestFunctionAliasHasVersion(t *testing.T) {
	sess, _ := session.NewSession()
	getLambdaAPI = func(sess *session.Session) (client lambdaiface.LambdaAPI) {
		return mockLambdaAPI{}
	}

	l := New(sess)
	res, err := l.FunctionAliasHasVersion(fnWithVersion, "", version)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = l.FunctionAliasHasVersion(functionName, "", version)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = l.FunctionAliasHasVersion("nope", "", version)
	assert.NotNil(t, err)
}
