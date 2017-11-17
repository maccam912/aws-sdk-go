// +build integration

//Package applicationdiscoveryservice provides gucumber integration tests support.
package applicationdiscoveryservice

import (
	"github.com/maccam912/aws-sdk-go/aws"
	"github.com/maccam912/aws-sdk-go/awstesting/integration/smoke"
	"github.com/maccam912/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@applicationdiscoveryservice", func() {
		gucumber.World["client"] = applicationdiscoveryservice.New(
			smoke.Session, &aws.Config{Region: aws.String("us-west-2")},
		)
	})
}
