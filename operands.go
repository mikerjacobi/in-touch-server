package main

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"
	"github.com/mikerjacobi/in-touch-server/app"
	"github.com/mikerjacobi/in-touch-server/models"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	o := models.Operand{Left: ctx.Left, Right: ctx.Right}
	o.Sum = o.Add()
	ret, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("fail: %+v", err)
		return ctx.Err()
	}
	return ctx.OK(ret)
}
