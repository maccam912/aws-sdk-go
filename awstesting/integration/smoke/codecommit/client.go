// +build integration

//Package codecommit provides gucumber integration tests support.
package codecommit

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/codecommit"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codecommit", func() {
		gucumber.World["client"] = codecommit.New(smoke.Session)
	})
}
