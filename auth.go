package main

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"github.com/mikerjacobi/in-touch-server/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Post runs the post action.
func (c *AuthController) Post(ctx *app.PostAuthContext) error {
	payload := ctx.Payload
	logrus.Infof("DATA IS: %+v", payload)
	_, err := json.Marshal(nil)
	if err != nil {
		fmt.Printf("failed to marshal login: %+v", err)
		return ctx.Err()
	}
	return ctx.OK([]byte("{}"))
}
