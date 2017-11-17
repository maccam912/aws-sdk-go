// +build integration

//Package cloudwatch provides gucumber integration tests support.
package cloudwatch

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/cloudwatch"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudwatch", func() {
		gucumber.World["client"] = cloudwatch.New(smoke.Session)
	})
}
