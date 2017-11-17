// +build integration

//Package lambda provides gucumber integration tests support.
package lambda

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/lambda"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@lambda", func() {
		gucumber.World["client"] = lambda.New(smoke.Session)
	})
}
