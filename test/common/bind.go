package common

import (
	"testing"

	v2 "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
	osbclient "github.com/openservicebrokerapi/osb-checker/client"
	"github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBind(
	t *testing.T,
	instanceID, bindingID string,
	req *v2.ServiceBindingRequest,
	async, looseCheck bool,
) {
	Convey("BINDING - request syntax", t, func() {

		So(testAPIVersionHeader(config.GenerateBindingURL(instanceID, bindingID), "PUT"), ShouldBeNil)
		So(testAuthentication(config.GenerateBindingURL(instanceID, bindingID), "PUT"), ShouldBeNil)
		if async {
			So(testAsyncParameters(config.GenerateBindingURL(instanceID, bindingID), "PUT"), ShouldBeNil)
		}

		var emptyValue, fakeValue = "", "xxxx-xxxx"
		Convey("should reject if missing service_id", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			tempBody.ServiceID = &emptyValue
			code, _, _, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})

		Convey("should reject if missing plan_id", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			tempBody.PlanID = &emptyValue
			code, _, _, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})

		Convey("should reject if service_id is invalid", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			tempBody.ServiceID = &fakeValue
			code, _, _, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})

		Convey("should reject if plan_id is invalid", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			tempBody.PlanID = &fakeValue
			code, _, _, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})

		Convey("should reject if parameters are not following schema", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			tempBody.Parameters = map[string]interface{}{
				"can-not": "be-good",
			}
			if err := testCatalogSchema(&SchemaOpts{
				ServiceID:  *tempBody.ServiceID,
				PlanID:     *tempBody.PlanID,
				Parameters: tempBody.Parameters,
				SchemaType: config.TypeServiceBinding,
				Action:     config.ActionCreate,
			}); err == nil {
				return
			}
			code, _, _, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 400)
		})
	})

	Convey("BINDING - new", t, func() {
		Convey("should accept a valid binding request", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			code, syncBody, asyncBody, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			if async {
				So(code, ShouldEqual, 202)
				So(testJSONSchema(asyncBody), ShouldBeNil)
			} else {
				So(code, ShouldEqual, 201)
				So(testJSONSchema(syncBody), ShouldBeNil)
			}
		})
	})

	if async {
		Convey("BINDING - poll", t, func(c C) {
			testPollBindingLastOperation(instanceID, bindingID)

			So(pollBindingLastOperationStatus(instanceID, bindingID), ShouldBeNil)
		})
	}

	Convey("BINDING - existed", t, func() {
		Convey("should return 200 OK when binding Id with same instance Id exists with identical properties", func() {
			tempBody := new(v2.ServiceBindingRequest)
			deepCopy(req, tempBody)
			code, syncBody, _, err := osbclient.Default.Bind(instanceID, bindingID, tempBody, async)

			So(err, ShouldBeNil)
			So(code, ShouldEqual, 200)
			So(testJSONSchema(syncBody), ShouldBeNil)
		})
	})
}