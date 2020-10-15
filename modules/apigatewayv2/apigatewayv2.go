package awsspec

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
)

// APIGatewayV2Spec contains the AWS session
type APIGatewayV2Spec struct {
	Session *session.Session
}

var getAPIGatewayV2API = func(sess *session.Session) (client apigatewayv2iface.ApiGatewayV2API) {
	return apigatewayv2.New(sess)
}

// New returns a new APIGatewayV2Spec
func New(s *session.Session) APIGatewayV2Spec {
	return APIGatewayV2Spec{
		Session: s,
	}
}
