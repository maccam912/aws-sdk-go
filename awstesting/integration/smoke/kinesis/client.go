// +build integration

//Package kinesis provides gucumber integration tests support.
package kinesis

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/kinesis"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@kinesis", func() {
		gucumber.World["client"] = kinesis.New(smoke.Session)
	})
}
