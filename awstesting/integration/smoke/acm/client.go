// +build integration

//Package acm provides gucumber integration tests support.
package acm

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/acm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@acm", func() {
		gucumber.World["client"] = acm.New(smoke.Session)
	})
}
