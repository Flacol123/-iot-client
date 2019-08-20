/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

// ArduinoMqttv1 media type (default view)
type ArduinoMqttv1 struct {
	// A signed Websocket url to use to connect to things in the cloud
	SignedWebsocket string `json:"signed_websocket"`
	// The url endpoint of the cloud
	Url string `json:"url"`
}