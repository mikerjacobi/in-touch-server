// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/mikerjacobi/in-touch-server/design
// --out=$(GOPATH)/src/github.com/mikerjacobi/in-touch-server
// --version=v1.1.0-dirty
//
// API "adder": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"unicode/utf8"
)

// PostAuthContext provides the auth post action context.
type PostAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PostAuthPayload
}

// NewPostAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller post action.
func NewPostAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*PostAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PostAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// postAuthPayload is the auth post action payload.
type postAuthPayload struct {
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *postAuthPayload) Validate() (err error) {
	if payload.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.Password != nil {
		if utf8.RuneCountInString(*payload.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, *payload.Password, utf8.RuneCountInString(*payload.Password), 6, true))
		}
	}
	if payload.Username != nil {
		if utf8.RuneCountInString(*payload.Username) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.username`, *payload.Username, utf8.RuneCountInString(*payload.Username), 1, true))
		}
	}
	return
}

// Publicize creates PostAuthPayload from postAuthPayload
func (payload *postAuthPayload) Publicize() *PostAuthPayload {
	var pub PostAuthPayload
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.Username != nil {
		pub.Username = *payload.Username
	}
	return &pub
}

// PostAuthPayload is the auth post action payload.
type PostAuthPayload struct {
	Password string `form:"password" json:"password" xml:"password"`
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate runs the validation rules defined in the design.
func (payload *PostAuthPayload) Validate() (err error) {
	if payload.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if utf8.RuneCountInString(payload.Password) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.password`, payload.Password, utf8.RuneCountInString(payload.Password), 6, true))
	}
	if utf8.RuneCountInString(payload.Username) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.username`, payload.Username, utf8.RuneCountInString(payload.Username), 1, true))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *PostAuthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// GetFriendContext provides the friend get action context.
type GetFriendContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FriendID string
}

// NewGetFriendContext parses the incoming request URL and body, performs validations and creates the
// context used by the friend controller get action.
func NewGetFriendContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetFriendContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetFriendContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFriendID := req.Params["friend_id"]
	if len(paramFriendID) > 0 {
		rawFriendID := paramFriendID[0]
		rctx.FriendID = rawFriendID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetFriendContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Left  int
	Right int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddOperandsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddOperandsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
