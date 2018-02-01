package clients

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGitlabClient_makeHttpRequest(t *testing.T) {

	Convey("Given an existing test URL", t, func() {

		testUrl := "http://www.google.com"
		So(testUrl, ShouldNotBeNil)

		Convey("When making a request with the existing test " +
			"URL", func() {

			resp, err := makeHttpRequest(testUrl)
			So(err, ShouldBeNil)

			Convey("Then the response code should be 200", func() {

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})
		})
	})

	Convey("Given a non-existing test URL", t, func() {

		testUrl := "http://this.url.is.not.good.com/random/123412341knkjdfg"
		So(testUrl, ShouldNotBeNil)

		Convey("When making a request with the non-existing test " +
			"URL", func() {

			resp, err := makeHttpRequest(testUrl)

			Convey("Then the HTTP response should be bil and the " +
				"error should not be nil", func() {

				So(resp, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
