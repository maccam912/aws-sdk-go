// +build integration

//Package sts provides gucumber integration tests support.
package sts

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/sts"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sts", func() {
		gucumber.World["client"] = sts.New(smoke.Session)
	})
}
