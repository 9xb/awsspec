package awsspec

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/stretchr/testify/assert"
)

func TestAPIExists(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.APIExists(apiID)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.APIExists("nope")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageExists(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageExists(apiID, stage)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageExists(apiID, "nope")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageHasLoggingEnabled(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageHasLoggingEnabled(apiID, stage)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageHasLoggingEnabled(apiID, stageNoLog)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageHasLogDestination(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageHasLogDestination(apiID, stage, logDestination)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageHasLogDestination(apiID, stage, "nope")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageHasLogFormat(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageHasLogFormat(apiID, stage, logFormat)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageHasLogFormat(apiID, stage, "nope")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageHasDetailedMetricsEnabled(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageHasDetailedMetricsEnabled(apiID, stage)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageHasDetailedMetricsEnabled(apiID, "nope")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageHasRateLimit(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageHasRateLimit(apiID, stage, rateLimit)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageHasRateLimit(apiID, stageNoLog, rateLimit)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = a.StageHasRateLimit(apiID, "nope", rateLimit)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestStageHasBurstLimit(t *testing.T) {
	sess, _ := session.NewSession()
	getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
		return mockAPIGatewayV2API{}
	}

	a := New(sess)

	res, err := a.StageHasBurstLimit(apiID, stage, burstLimit)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = a.StageHasBurstLimit(apiID, stageNoLog, burstLimit)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = a.StageHasBurstLimit(apiID, "nope", burstLimit)
	assert.Nil(t, err)
	assert.False(t, res)
}
