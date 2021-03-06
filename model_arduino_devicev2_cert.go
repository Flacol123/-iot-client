/*
 * Arduino IoT Cloud API
 *
 *  Provides a set of endpoints to manage Arduino IoT Cloud **Devices**, **Things**, **Properties** and **Timeseries**. This API can be called just with any HTTP Client, or using one of these clients:  * [Javascript NPM package](https://www.npmjs.com/package/@arduino/arduino-iot-client)  * [Python PYPI Package](https://pypi.org/project/arduino-iot-client/)  * [Golang Module](https://github.com/arduino/iot-client-go)
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// ArduinoDevicev2Cert DeviceCertV2 describes a certificate associated to the device (default view)
type ArduinoDevicev2Cert struct {
	// The Certification Authority used to sign the certificate
	Ca string `json:"ca,omitempty"`
	Compressed ArduinoCompressedv2 `json:"compressed"`
	// The certificate in DER format
	Der string `json:"der"`
	// The unique identifier of the device
	DeviceId string `json:"device_id"`
	// Whether the certificate is enabled
	Enabled bool `json:"enabled"`
	// The api reference of this cert
	Href string `json:"href"`
	// The unique identifier of the key
	Id string `json:"id"`
	// The certificate in pem format
	Pem string `json:"pem"`
}
