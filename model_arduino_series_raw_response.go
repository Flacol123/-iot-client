/*
 * Arduino IoT Cloud API
 *
 *  Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
import (
	"time"
)
// ArduinoSeriesRawResponse ArduinoSeriesRawResponse media type (default view)
type ArduinoSeriesRawResponse struct {
	// Total number of values in the array 'values'
	CountValues int64 `json:"count_values"`
	// From date
	FromDate time.Time `json:"from_date"`
	// If the response is different than 'ok'
	Message string `json:"message,omitempty"`
	// Query of for the data
	Query string `json:"query"`
	// Response version
	RespVersion int64 `json:"resp_version"`
	Series BatchQueryRawResponseSeriesMediaV1 `json:"series"`
	// Max of values
	SeriesLimit int64 `json:"series_limit,omitempty"`
	// Sorting
	Sort string `json:"sort"`
	// Status of the response
	Status string `json:"status"`
	// Timestamp in RFC3339
	Times []time.Time `json:"times"`
	// To date
	ToDate time.Time `json:"to_date"`
	// Values can be in Float, String, Bool, Object
	Values []interface{} `json:"values"`
}
