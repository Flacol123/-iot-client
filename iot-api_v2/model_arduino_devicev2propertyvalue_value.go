/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

type ArduinoDevicev2propertyvalueValue struct {
	Payload string `json:"payload,omitempty"`
	Seqno int64 `json:"seqno,omitempty"`
	Statistics ArduinoDevicev2propertyvalueValueStatistics `json:"statistics,omitempty"`
}