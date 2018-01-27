package emrtr

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_apis/emrtrapi"
	"errors"
	"fmt"
)

// RouterOptions implements the emrtrapi.RouterOptions interface. Each method returns a function, which applies a
// specific configuration to an IRCGateway object.
type RouterOptions struct {
}

// Gateways sets the emersyx gateway instances for the router.
func (o RouterOptions) Gateways(gws ...emcomapi.Identifiable) func(emrtrapi.Router) error {
	return func(rtr emrtrapi.Router) error {
		crtr, ok := rtr.(*Router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}
		for _, gw := range gws {
			crtr.gws = append(crtr.gws, gw)
		}
		return nil
	}
}

// Processors sets the emersyx processor instances for the router.
func (o RouterOptions) Processors(procs ...emcomapi.Processor) func(emrtrapi.Router) error {
	return func(rtr emrtrapi.Router) error {
		crtr, ok := rtr.(*Router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}
		for _, proc := range procs {
			crtr.procs = append(crtr.procs, proc)
		}
		return nil
	}
}

// Routes sets the emersyx routes required to forward events between components.
func (o RouterOptions) Routes(routes map[string][]string) func(emrtrapi.Router) error {
	return func(rtr emrtrapi.Router) error {
		crtr, ok := rtr.(*Router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}
		for src, dsts := range routes {
			if len(src) == 0 {
				return errors.New("provided route with empty source is not valid")
			}
			if dsts == nil || len(dsts) == 0 {
				return fmt.Errorf("route with source \"%s\" has an invalid set of destinations", src)
			}
			crtr.routes[src] = make([]string, len(dsts))
			for _, dst := range dsts {
				if len(dst) == 0 {
					return fmt.Errorf("route with source \"%s\" has an invalid destination", src)
				}
				crtr.routes[src] = append(crtr.routes[src], dst)
			}
		}
		return nil
	}
}
