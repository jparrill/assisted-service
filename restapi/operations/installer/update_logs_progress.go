// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateLogsProgressHandlerFunc turns a function with the right signature into a update logs progress handler
type UpdateLogsProgressHandlerFunc func(UpdateLogsProgressParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateLogsProgressHandlerFunc) Handle(params UpdateLogsProgressParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateLogsProgressHandler interface for that can handle valid update logs progress params
type UpdateLogsProgressHandler interface {
	Handle(UpdateLogsProgressParams, interface{}) middleware.Responder
}

// NewUpdateLogsProgress creates a new http.Handler for the update logs progress operation
func NewUpdateLogsProgress(ctx *middleware.Context, handler UpdateLogsProgressHandler) *UpdateLogsProgress {
	return &UpdateLogsProgress{Context: ctx, Handler: handler}
}

/*UpdateLogsProgress swagger:route PUT /clusters/{cluster_id}/logs_progress installer updateLogsProgress

Update log collection state and progress.

*/
type UpdateLogsProgress struct {
	Context *middleware.Context
	Handler UpdateLogsProgressHandler
}

func (o *UpdateLogsProgress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateLogsProgressParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
