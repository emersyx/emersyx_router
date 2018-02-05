package emrtr

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"sync"
)

// router is the struct type which implements the emcomapi.Router interface. Objects of this type are used to route
// events from receptor gateways to processors.
type router struct {
	gws       []emcomapi.Identifiable
	procs     []emcomapi.Processor
	routes    map[string][]string
	isRunning bool
	mutex     sync.Mutex
}

// NewRouter creates a new router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func NewRouter() (emcomapi.Router, error) {
	rtr := new(router)

	// the router is initially not running
	// this member is set to true once the router.Run method is called
	rtr.isRunning = false

	// create member arrays with default sizes
	rtr.gws = make([]emcomapi.Identifiable, 1)
	rtr.procs = make([]emcomapi.Processor, 1)
	rtr.routes = make(map[string][]string)

	return rtr, nil
}
