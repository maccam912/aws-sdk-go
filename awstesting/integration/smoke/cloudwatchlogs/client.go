// +build integration

//Package cloudwatchlogs provides gucumber integration tests support.
package cloudwatchlogs

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudwatchlogs", func() {
		gucumber.World["client"] = cloudwatchlogs.New(smoke.Session)
	})
}
