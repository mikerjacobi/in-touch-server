package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:8080")
	Scheme("http")
	Trait("jsonapi", func() {
		ContentType("application/vnd.api+json")
	})
	Consumes("application/json", func() {
        Package("github.com/goadesign/goa/encoding/json")
    })
	Produces("application/json") 
})

var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "application/json")
	})
})

var _ = Resource("friend", func() {
	Action("get", func() {
		Routing(GET("friend/:friend_id"))
		Description("get friend by id")
		Params(func() {
			Param("friend_id", String, "friend id")
		})
		Response(OK, "application/json")
	})
})

var _ = Resource("auth", func() {
	DefaultMedia(LoginPayload)
	Action("post", func() {
		Routing(POST("login"))
		Description("get friend by id")
		Payload(LoginPayload, func() {
			Required("username", "password")
		})
		Response(OK, "application/json")
	})
})
