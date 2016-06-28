// +build integration

//Package cloudfront provides gucumber integration tests support.
package cloudfront

import (
	"github.com/aws/aws-sdk-go/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	. "github.com/gucumber/gucumber"
)

func init() {
	Before("@cloudfront", func() {
		World["client"] = cloudfront.New(smoke.Session)
	})
}
