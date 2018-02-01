package clients

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGitlabClient_GetGitlabResponseTime(t *testing.T) {

	Convey("When making a request to the Gitlab Base URL", t, func() {

		responseTime, err := GetGitlabRequestResponseTime()
		So(err, ShouldBeNil)

		Convey("Then the response time should be greater " +
			"than 0.0", func() {

			So(responseTime, ShouldBeGreaterThan, 0.0)
		})
	})
}
