package fhirclient

import (
	"net/http"

	"github.com/lumc/fhir-client/design"
)

// PatientsService instance operates over patient resources
type PatientsService service

// Put can be used to update some or all details of an existing observation.
func (os *PatientsService) Put(payload *design.Patient) error { // No id in input since it is or should be contained in the Observation.
	err := os.client.Call(http.MethodPut, "Patient", payload)
	if err != nil {
		return err
	}

	return nil
}
