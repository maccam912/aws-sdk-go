// +build integration

//Package simpledb provides gucumber integration tests support.
package simpledb

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/simpledb"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@simpledb", func() {
		gucumber.World["client"] = simpledb.New(smoke.Session)
	})
}
