// +build integration

//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/storagegateway"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@storagegateway", func() {
		gucumber.World["client"] = storagegateway.New(smoke.Session)
	})
}
