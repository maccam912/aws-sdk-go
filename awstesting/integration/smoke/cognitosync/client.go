// +build integration

//Package cognitosync provides gucumber integration tests support.
package cognitosync

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/cognitosync"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cognitosync", func() {
		gucumber.World["client"] = cognitosync.New(smoke.Session)
	})
}
