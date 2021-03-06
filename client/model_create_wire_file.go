/*
 * WIRE API
 *
 * Moov WIRE implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateWireFile struct for CreateWireFile
type CreateWireFile struct {
	// File ID
	ID             string         `json:"ID,omitempty"`
	FedWireMessage FedWireMessage `json:"fedWireMessage"`
}
