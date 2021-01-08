package fhirclient

import (
	"net/http"

	"github.com/lumc/fhir-client/design"
)

// BundleService instance operates over bundle resources
type BundleService service

// Put can be used to update some or all details of an existing bundle.
func (bs *BundleService) Put(payload *design.Bundle) error {
	err := bs.client.Call(http.MethodPut, "Bundle", payload)
	if err != nil {
		return err
	}

	return nil
}
