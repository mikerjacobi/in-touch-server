package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// LoginPayload defines the object to issue a login
var LoginPayload = MediaType("application/vnd.login+json", func() {
	UseTrait("jsonapi")
	Attributes(func() {
		Attribute("username", func() {
			MinLength(1)
			Example("jacobra")
		})
		Attribute("password", func() {
			MinLength(6)
			Example("password")
		})
	})
	View("default", func() {
		Attribute("username", String, "username")
		Attribute("password", String, "password")
	})

})
