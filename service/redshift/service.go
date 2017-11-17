// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"github.com/maccam912/aws-sdk-go/aws"
	"github.com/maccam912/aws-sdk-go/aws/client"
	"github.com/maccam912/aws-sdk-go/aws/client/metadata"
	"github.com/maccam912/aws-sdk-go/aws/request"
	"github.com/maccam912/aws-sdk-go/aws/signer/v4"
	"github.com/maccam912/aws-sdk-go/private/protocol/query"
)

// Redshift provides the API operation methods for making requests to
// Amazon Redshift. See this package's package overview docs
// for details on the service.
//
// Redshift methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Redshift struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "redshift"  // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Redshift client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Redshift client from just a session.
//     svc := redshift.New(mySession)
//
//     // Create a Redshift client with additional configuration
//     svc := redshift.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Redshift {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Redshift {
	svc := &Redshift{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Redshift operation and runs any
// custom request initialization.
func (c *Redshift) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
