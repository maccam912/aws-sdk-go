// +build integration

//Package workspaces provides gucumber integration tests support.
package workspaces

import (
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/workspaces"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@workspaces", func() {
		gucumber.World["client"] = workspaces.New(smoke.Session)
	})
}
