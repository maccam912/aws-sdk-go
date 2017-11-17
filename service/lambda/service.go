// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"github.com/maccam912/aws-sdk-go/aws"
	"github.com/maccam912/aws-sdk-go/aws/client"
	"github.com/maccam912/aws-sdk-go/aws/client/metadata"
	"github.com/maccam912/aws-sdk-go/aws/request"
	"github.com/maccam912/aws-sdk-go/aws/signer/v4"
	"github.com/maccam912/aws-sdk-go/private/protocol/restjson"
)

// Lambda provides the API operation methods for making requests to
// AWS Lambda. See this package's package overview docs
// for details on the service.
//
// Lambda methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Lambda struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "lambda"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Lambda client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Lambda client from just a session.
//     svc := lambda.New(mySession)
//
//     // Create a Lambda client with additional configuration
//     svc := lambda.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Lambda {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Lambda {
	svc := &Lambda{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-03-31",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Lambda operation and runs any
// custom request initialization.
func (c *Lambda) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
