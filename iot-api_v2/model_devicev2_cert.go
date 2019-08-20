/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

type Devicev2Cert struct {
	// The Certification Authority you want to use
	Ca string `json:"ca,omitempty"`
	// The certificate request in pem format
	Csr string `json:"csr,omitempty"`
	// Whether the certificate is enabled
	Enabled bool `json:"enabled,omitempty"`
}