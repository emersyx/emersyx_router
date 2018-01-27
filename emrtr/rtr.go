package emrtr

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_apis/emrtrapi"
)

// Router is the struct type which implements the emrtrapi.Router interface. Objects of this type are used to route
// events from receptor gateways to processors.
type Router struct {
	gws    []emcomapi.Identifiable
	procs  []emcomapi.Processor
	routes map[string][]string
}

// NewRouter creates a new Router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func NewRouter(options ...func(emrtrapi.Router) error) (emrtrapi.Router, error) {
	rtr := new(Router)

	// create member arrays with default sizes
	rtr.gws = make([]emcomapi.Identifiable, 1)
	rtr.procs = make([]emcomapi.Processor, 1)
	rtr.routes = make(map[string][]string)

	// apply the configuration options received as arguments
	for _, option := range options {
		err := option(rtr)
		if err != nil {
			return nil, err
		}
	}

	return rtr, nil
}
