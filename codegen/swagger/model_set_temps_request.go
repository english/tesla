/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SetTempsRequest struct {
	// The desired temperature on the driver's side in celcius.
	DriverTemp float64 `json:"driver_temp,omitempty"`
	// The desired temperature on the passenger's side in celcius.
	PassengerTemp float64 `json:"passenger_temp,omitempty"`
}