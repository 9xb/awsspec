package awsspec

import (
	"encoding/json"
	"errors"
	"net/url"

	iamSpec "github.com/9xb/awsspec/modules/iam"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// FunctionExists returns true if the specified Lambda function exists. An optional qualifier nay be provided for versioned functions or aliases.
func (l LambdaSpec) FunctionExists(name, qualifier string) (res bool, err error) {
	svc := getLambdaAPI(l.Session)
	in := &lambda.GetFunctionInput{
		FunctionName: aws.String(name),
	}

	if qualifier != "" {
		in.Qualifier = aws.String(qualifier)
	}

	_, err = svc.GetFunction(in)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "ResourceNotFoundException" {
				return false, nil
			}
		}
		return
	}

	return true, nil
}

// FunctionHasEnvVarValue returns true if the specified Lambda function has the provided environment variable name and value. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasEnvVarValue(name, qualifier, variable, value string) (res bool, err error) {
	env, err := getFunctionConfig(name, qualifier, "environment", l.Session)
	if err != nil {
		return
	}

	if val, ok := env.(*lambda.EnvironmentResponse).Variables[variable]; ok {
		if aws.StringValue(val) == value {
			return true, nil
		}
	}
	return
}

// FunctionHasRole returns true if the specified Lambda function has the provided role. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasRole(name, qualifier, role string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "role", l.Session)
	if err != nil {
		return
	}

	if aws.StringValue(r.(*string)) == role {
		return true, nil
	}

	return
}

// FunctionHasRuntime returns true if the specified Lambda function has the provided runtime. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasRuntime(name, qualifier, runtime string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "runtime", l.Session)
	if err != nil {
		return
	}

	if aws.StringValue(r.(*string)) == runtime {
		return true, nil
	}

	return
}

// FunctionHasTimeout returns true if the specified Lambda function has the provided runtime. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasTimeout(name, qualifier string, timeout int) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "timeout", l.Session)
	if err != nil {
		return
	}

	if aws.Int64Value(r.(*int64)) == int64(timeout) {
		return true, nil
	}

	return
}

// FunctionHasMemorySize returns true if the specified Lambda function has the provided memory size. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasMemorySize(name, qualifier string, memory int) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "memory_size", l.Session)
	if err != nil {
		return
	}

	if aws.Int64Value(r.(*int64)) == int64(memory) {
		return true, nil
	}

	return
}

// FunctionHasHandler returns true if the specified Lambda function has the provided handler. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasHandler(name, qualifier, handler string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "handler", l.Session)
	if err != nil {
		return
	}

	if aws.StringValue(r.(*string)) == handler {
		return true, nil
	}

	return
}

// FunctionHasLayers returns true if the specified Lambda function has exactly the provided layers. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasLayers(name, qualifier string, layers []string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "layers", l.Session)
	if err != nil {
		return
	}

	if len(r.([]*lambda.Layer)) != len(layers) {
		return
	}

	ls := []string{}
	for _, v := range r.([]*lambda.Layer) {
		ls = append(ls, aws.StringValue(v.Arn))
	}

	if sameStringSlice(ls, layers) {
		return true, nil
	}

	return
}

// FunctionHasVPCID returns true if the specified Lambda function has the provided VPC ID. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasVPCID(name, qualifier, vpc string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "vpc", l.Session)
	if err != nil {
		return
	}

	if aws.StringValue(r.(*lambda.VpcConfigResponse).VpcId) == vpc {
		return true, nil
	}

	return
}

// FunctionHasVPCWithSubnets returns true if the specified Lambda function has exactly the provided subnets. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasVPCWithSubnets(name, qualifier string, subnets []string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "vpc", l.Session)
	if err != nil {
		return
	}

	s := r.(*lambda.VpcConfigResponse).SubnetIds
	if len(s) != len(subnets) {
		return
	}

	ls := []string{}
	for _, v := range s {
		ls = append(ls, aws.StringValue(v))
	}

	if sameStringSlice(ls, subnets) {
		return true, nil
	}

	return
}

// FunctionHasVPCWithSecurityGroups returns true if the specified Lambda function has the provided security groups. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasVPCWithSecurityGroups(name, qualifier string, sgs []string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "vpc", l.Session)
	if err != nil {
		return
	}

	s := r.(*lambda.VpcConfigResponse).SecurityGroupIds
	if len(s) != len(sgs) {
		return
	}

	ls := []string{}
	for _, v := range s {
		ls = append(ls, aws.StringValue(v))
	}

	if sameStringSlice(ls, sgs) {
		return true, nil
	}

	return
}

// FunctionHasPermissions returns true if the specified source ARN can execute the provided Lambda function. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasPermissions(name, qualifier, source string) (res bool, err error) {
	svc := getLambdaAPI(l.Session)
	in := &lambda.GetPolicyInput{
		FunctionName: aws.String(name),
	}

	if qualifier != "" {
		in.Qualifier = aws.String(qualifier)
	}

	out, err := svc.GetPolicy(in)
	if err != nil {
		return
	}

	doc := iamSpec.PolicyDocument{}
	p, err := url.QueryUnescape(aws.StringValue(out.Policy))
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(p), &doc)
	if err != nil {
		return
	}

	if doc.Statement[0].Condition[iamSpec.ConditionArnLike][iamSpec.VarSourceArn][0] == source {
		return true, nil
	}

	return
}

// FunctionAliasHasVersion returns true if the provided Lambda function alias is associated to the specified version
func (l LambdaSpec) FunctionAliasHasVersion(name, qualifier, version string) (res bool, err error) {
	r, err := getFunctionConfig(name, qualifier, "version", l.Session)
	if err != nil {
		return
	}

	if aws.StringValue(r.(*string)) == version {
		return true, nil
	}

	return
}

// FunctionHasReservedConcurrency returns true if the provided Lambda function has the provided reserved concurrency. It throws an error if no function was found.
func (l LambdaSpec) FunctionHasReservedConcurrency(name string, concurrency int) (res bool, err error) {
	svc := getLambdaAPI(l.Session)
	in := &lambda.GetFunctionConcurrencyInput{
		FunctionName: aws.String(name),
	}

	out, err := svc.GetFunctionConcurrency(in)
	if err != nil {
		return
	}

	if aws.Int64Value(out.ReservedConcurrentExecutions) == int64(concurrency) {
		res = true
	}

	return
}

func getFunctionConfig(name, qualifier, cfgName string, s *session.Session) (cfg interface{}, err error) {
	svc := getLambdaAPI(s)
	in := &lambda.GetFunctionInput{
		FunctionName: aws.String(name),
	}

	if qualifier != "" {
		in.Qualifier = aws.String(qualifier)
	}

	out, err := svc.GetFunction(in)
	if err != nil {
		return
	}

	switch cfgName {
	case "environment":
		cfg = out.Configuration.Environment
	case "role":
		cfg = out.Configuration.Role
	case "runtime":
		cfg = out.Configuration.Runtime
	case "timeout":
		cfg = out.Configuration.Timeout
	case "memory_size":
		cfg = out.Configuration.MemorySize
	case "handler":
		cfg = out.Configuration.Handler
	case "layers":
		cfg = out.Configuration.Layers
	case "vpc":
		cfg = out.Configuration.VpcConfig
	case "version":
		cfg = out.Configuration.Version
	}

	if cfg == nil {
		return cfg, errors.New("Configuration not found")
	}

	return
}

// Thanks, StackOverflow!
// https://stackoverflow.com/a/36000696/543423
func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
