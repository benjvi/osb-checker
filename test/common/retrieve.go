package common

import (
	"testing"

	osbclient "github.com/openservicebrokerapi/osb-checker/client"
	"github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInstance(
	t *testing.T,
	instanceID string,
	looseCheck bool,
) {
	Convey("RETRIEVE INSTANCE", t, func() {

		So(testAPIVersionHeader(config.GenerateInstanceURL(instanceID), "GET"), ShouldBeNil)
		So(testAuthentication(config.GenerateInstanceURL(instanceID), "GET"), ShouldBeNil)

		Convey("should accept a valid get instance request", func() {
			code, body, err := osbclient.Default.GetInstance(instanceID)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 200)
			So(testJSONSchema(body), ShouldBeNil)
		})
	})
}

func TestGetBinding(
	t *testing.T,
	instanceID, bindingID string,
	looseCheck bool,
) {
	Convey("RETRIEVE BINDING", t, func() {

		So(testAPIVersionHeader(config.GenerateBindingURL(instanceID, bindingID), "GET"), ShouldBeNil)
		So(testAuthentication(config.GenerateBindingURL(instanceID, bindingID), "GET"), ShouldBeNil)

		Convey("should accept a valid get binding request", func() {
			code, body, err := osbclient.Default.GetBinding(instanceID, bindingID)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 200)
			So(testJSONSchema(body), ShouldBeNil)
		})
	})
}