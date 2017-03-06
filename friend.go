package main

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"
	"github.com/mikerjacobi/in-touch-server/app"
	"github.com/mikerjacobi/in-touch-server/models"
)

// FriendController implements the friend resource.
type FriendController struct {
	*goa.Controller
}

// NewFriendController creates a friend controller.
func NewFriendController(service *goa.Service) *FriendController {
	return &FriendController{Controller: service.NewController("FriendController")}
}

// Get runs the get action.
func (c *FriendController) Get(ctx *app.GetFriendContext) error {
	query := map[string]string{"friend_id": ctx.FriendID}
	friends, err := models.LoadFriends(query)
	if err != nil {
		fmt.Printf("fail: %+v", err)
		return ctx.Err()
	}

	ret, err := json.Marshal(friends[0])
	switch err {
	case nil:
		break
	case models.ErrNotFound:
		fmt.Printf("not found: %s", ctx.FriendID)
	default:
		fmt.Printf("fail: %+v", err)
		return ctx.Err()
	}

	return ctx.OK(ret)
}
