package fhirclient

import (
	"net/http"

	"github.com/lumc/fhir-client/design"
)

// MedicationAdministrationsService instance operates over MedicationAdministration resources
type MedicationAdministrationsService service

// Put can be used to update some or all details of an existing observation.
func (os *MedicationAdministrationsService) Put(payload *design.MedicationAdministration) error { // No id in input since it is or should be contained in the Observation.
	err := os.client.Call(http.MethodPut, "MedicationAdministration", payload)
	if err != nil {
		return err
	}

	return nil
}
