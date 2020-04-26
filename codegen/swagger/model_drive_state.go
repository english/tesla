/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DriveState struct {
	ShiftState string `json:"shift_state,omitempty"`
	Speed float64 `json:"speed,omitempty"`
	Power float64 `json:"power,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Heading int32 `json:"heading,omitempty"`
	GpsAsOf float32 `json:"gps_as_of,omitempty"`
	NativeLocationSupported int32 `json:"native_location_supported,omitempty"`
	NativeLatitude float64 `json:"native_latitude,omitempty"`
	NativeLongitude float64 `json:"native_longitude,omitempty"`
	NativeType string `json:"native_type,omitempty"`
	Timestamp float32 `json:"timestamp,omitempty"`
}
