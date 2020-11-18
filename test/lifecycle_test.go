package test

import (
	"testing"

	openapi "github.com/benjvi/osb-checker/autogenerated/go-client"
	. "github.com/benjvi/osb-checker/config"
	"github.com/benjvi/osb-checker/test/common"
	uuid "github.com/satori/go.uuid"
)

var (
	configFile string
)

func init() {
	// Load configuration info into global CONF variable.
	if err := Load("configs/config.yaml"); err != nil {
		panic(err)
	}
	// Initialize work
	if err := common.InitClientWithAuthCtx(); err != nil {
		panic(err)
	}
}

func TestLifeCycle(t *testing.T) {
	t.Parallel()

	common.TestGetCatalog(t)

	for _, svc := range CONF.Services {
		instanceID := uuid.NewV4().String()
		bindingID := uuid.NewV4().String()
		serviceID, organizationGUID, spaceGUID :=
			svc.ServiceID, svc.OrganizationGUID, svc.SpaceGUID

		for _, operation := range svc.Operations {
			switch operation.Type {
			case "provision":
				req := &openapi.ServiceInstanceProvisionRequest{
					ServiceId:        serviceID,
					PlanId:           operation.PlanID,
					OrganizationGuid: organizationGUID,
					SpaceGuid:        spaceGUID,
					Parameters:       operation.Parameters,
				}

				common.TestProvision(t, instanceID, req, operation.Async)
				break
			case "get_instance":
				common.TestGetInstance(t, instanceID)
				break
			case "update":
				var currentPlanID string
				if operation.PlanID != "" {
					currentPlanID = operation.PlanID
				}
				req := &openapi.ServiceInstanceUpdateRequest{
					ServiceId:  serviceID,
					PlanId:     currentPlanID,
					Parameters: operation.Parameters,
				}

				common.TestUpdateInstance(t, instanceID, req, operation.Async)
				break
			case "deprovision":
				common.TestDeprovision(t, instanceID, serviceID, operation.PlanID, operation.Async)
				break
			case "bind":
				req := &openapi.ServiceBindingRequest{
					ServiceId:  serviceID,
					PlanId:     operation.PlanID,
					Parameters: operation.Parameters,
				}

				common.TestBind(t, instanceID, bindingID, req, operation.Async)
				break
			case "get_binding":
				common.TestGetBinding(t, instanceID, bindingID)
				break
			case "unbind":
				common.TestUnbind(t, instanceID, bindingID, serviceID, operation.PlanID, operation.Async)
				break
			default:
				break
			}
		}
	}
}
