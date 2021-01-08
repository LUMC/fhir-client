package fhirclient

import (
	"net/http"

	"github.com/lumc/fhir-client/design"
)

// ObservationsService instance operates over observation resources.
type ObservationsService service

// Put can be used to update some or all details of an existing observation.
func (os *ObservationsService) Put(payload *design.Observation) error { // No id in input since it is or should be contained in the Observation.
	err := os.client.Call(http.MethodPut, "Observation", payload)
	if err != nil {
		return err
	}

	return nil
}
