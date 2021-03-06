// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ResetClusterHandlerFunc turns a function with the right signature into a reset cluster handler
type ResetClusterHandlerFunc func(ResetClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ResetClusterHandlerFunc) Handle(params ResetClusterParams) middleware.Responder {
	return fn(params)
}

// ResetClusterHandler interface for that can handle valid reset cluster params
type ResetClusterHandler interface {
	Handle(ResetClusterParams) middleware.Responder
}

// NewResetCluster creates a new http.Handler for the reset cluster operation
func NewResetCluster(ctx *middleware.Context, handler ResetClusterHandler) *ResetCluster {
	return &ResetCluster{Context: ctx, Handler: handler}
}

/*ResetCluster swagger:route POST /clusters/{cluster_id}/actions/reset installer resetCluster

Resets a failed installation.

*/
type ResetCluster struct {
	Context *middleware.Context
	Handler ResetClusterHandler
}

func (o *ResetCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewResetClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
