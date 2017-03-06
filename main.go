//go:generate goagen bootstrap -d github.com/mikerjacobi/in-touch-server/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mikerjacobi/in-touch-server/app"
)

func main() {
	// Create service
	service := goa.New("adder")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount controllers
	app.MountFriendController(service, NewFriendController(service))
	app.MountAuthController(service, NewAuthController(service))
	app.MountOperandsController(service, NewOperandsController(service))

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
