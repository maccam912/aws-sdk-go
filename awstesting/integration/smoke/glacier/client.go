// +build integration

//Package glacier provides gucumber integration tests support.
package glacier

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/glacier"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@glacier", func() {
		gucumber.World["client"] = glacier.New(smoke.Session)
	})
}
