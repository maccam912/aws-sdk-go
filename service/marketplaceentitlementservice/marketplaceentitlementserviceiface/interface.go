// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package marketplaceentitlementserviceiface provides an interface to enable mocking the AWS Marketplace Entitlement Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package marketplaceentitlementserviceiface

import (
	"github.com/maccam912/aws-sdk-go/aws"
	"github.com/maccam912/aws-sdk-go/aws/request"
	"github.com/maccam912/aws-sdk-go/service/marketplaceentitlementservice"
)

// MarketplaceEntitlementServiceAPI provides an interface to enable mocking the
// marketplaceentitlementservice.MarketplaceEntitlementService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Marketplace Entitlement Service.
//    func myFunc(svc marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI) bool {
//        // Make svc.GetEntitlements request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := marketplaceentitlementservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMarketplaceEntitlementServiceClient struct {
//        marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI
//    }
//    func (m *mockMarketplaceEntitlementServiceClient) GetEntitlements(input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMarketplaceEntitlementServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MarketplaceEntitlementServiceAPI interface {
	GetEntitlements(*marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
	GetEntitlementsWithContext(aws.Context, *marketplaceentitlementservice.GetEntitlementsInput, ...request.Option) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
	GetEntitlementsRequest(*marketplaceentitlementservice.GetEntitlementsInput) (*request.Request, *marketplaceentitlementservice.GetEntitlementsOutput)
}

var _ MarketplaceEntitlementServiceAPI = (*marketplaceentitlementservice.MarketplaceEntitlementService)(nil)
