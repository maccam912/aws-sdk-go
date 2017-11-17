// +build integration

//Package elasticloadbalancing provides gucumber integration tests support.
package elasticloadbalancing

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/elb"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@elasticloadbalancing", func() {
		gucumber.World["client"] = elb.New(smoke.Session)
	})
}
