// Package fhirclient is a api client for the fhir server and contains fhir structs
package fhirclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/lumc/fhir-client/design"
)

const requestContentType = "application/json"

// Client manages communication with Vonk API.
type Client struct {
	baseURL *url.URL
	client  *http.Client
	logger  *logrus.Logger
	common  service // Reuse a single struct instead of allocating one for each service on the heap.
	// Services
	Patients     *PatientsService
	Observations *ObservationsService
}

type service struct {
	client *Client
}

// newRequest is a wrapper around the http.NewRequest function.
func (c *Client) newRequest(method, uri string, body interface{}) (req *http.Request, err error) {
	u, err := c.baseURL.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("uri can't be parsed: %w", err)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, fmt.Errorf("can't encode body: %w", err)
		}
	}

	req, err = http.NewRequestWithContext(context.Background(), method, u.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("can't create new request with context: %w", err)
	}

	req.Header.Set("Content-Type", requestContentType)
	req.Header.Set("Accept", requestContentType)

	return req, nil
}

// doRequest sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request failed: %w", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			c.logger.WithError(err).Warn("can't close response body")
		}
	}()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response body could not be read: %w", err)
	}

	return res, nil
}

// NewClient builds a new Client pointer using a default HTTPClient, and a default API base URL
func NewClient(baseClient *http.Client, vonkURL string) (fhir *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, err := url.Parse(vonkURL)
	if err != nil {
		return nil, fmt.Errorf("parsing url failed: %w", err)
	}

	fhir = &Client{
		baseURL: u,
		client:  baseClient,
	}

	fhir.logger = logrus.New()

	fhir.logger.WithFields(logrus.Fields{
		"base_url": u,
	})

	fhir.common.client = fhir

	// services for resources
	fhir.Patients = (*PatientsService)(&fhir.common)
	fhir.Observations = (*ObservationsService)(&fhir.common)

	return fhir, nil
}

// Call executes a api call with method, uri and payload variables when the call fails it tries to return a operation outcome struct
func (c *Client) Call(method, uri string, payload interface{}) error {
	req, err := c.newRequest(method, uri, payload)
	if err != nil {
		return fmt.Errorf("%s put failed: %w", uri, err)
	}

	res, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("%s %s failed: %w", uri, method, err)
	}
	err = json.Unmarshal(res, &payload)
	if err != nil {
		var outcome design.OperationOutcome
		err = json.Unmarshal(res, &outcome)
		if err != nil {
			return fmt.Errorf("unmarshalling response failed: %w", err)
		}

		return fmt.Errorf("%s request for %s failed: %w", method, uri, &outcome)
	}

	return nil
}
