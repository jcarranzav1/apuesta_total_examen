package groups

import (
	"ApuestaTotal/pkg/saga"
	"github.com/labstack/echo/v4"
)

type Workflow interface {
	Resource(c *echo.Group)
}

type workflow struct {
	workflowHandler saga.WorkFlow
}

func NewWorkflow(workflowHandler saga.WorkFlow) Workflow {
	return &workflow{
		workflowHandler,
	}
}

func (groups workflow) Resource(route *echo.Group) {

	route.GET("/paymentProcess", groups.workflowHandler.PaymentSaga)

}
