// +build integration

//Package elastictranscoder provides gucumber integration tests support.
package elastictranscoder

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/elastictranscoder"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@elastictranscoder", func() {
		gucumber.World["client"] = elastictranscoder.New(smoke.Session)
	})
}
