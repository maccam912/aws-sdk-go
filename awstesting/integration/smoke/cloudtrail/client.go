// +build integration

//Package cloudtrail provides gucumber integration tests support.
package cloudtrail

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/cloudtrail"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudtrail", func() {
		gucumber.World["client"] = cloudtrail.New(smoke.Session)
	})
}
