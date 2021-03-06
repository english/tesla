/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VehicleConfig struct {
	CanActuateTrunks bool `json:"can_actuate_trunks,omitempty"`
	CanAcceptNavigationRequests bool `json:"can_accept_navigation_requests,omitempty"`
	CarSpecialType string `json:"car_special_type,omitempty"`
	CarType string `json:"car_type,omitempty"`
	ChargePortType string `json:"charge_port_type,omitempty"`
	EuVehicle bool `json:"eu_vehicle,omitempty"`
	ExteriorColor string `json:"exterior_color,omitempty"`
	HasLudicrousMode bool `json:"has_ludicrous_mode,omitempty"`
	MotorizedChargePort bool `json:"motorized_charge_port,omitempty"`
	PerfConfig string `json:"perf_config,omitempty"`
	Plg bool `json:"plg,omitempty"`
	RearSeatHeaters float32 `json:"rear_seat_heaters,omitempty"`
	RearSeatType float32 `json:"rear_seat_type,omitempty"`
	Rhd bool `json:"rhd,omitempty"`
	RoofColor string `json:"roof_color,omitempty"`
	SeatType float32 `json:"seat_type,omitempty"`
	SpoilerType string `json:"spoiler_type,omitempty"`
	SunRoofInstalled float32 `json:"sun_roof_installed,omitempty"`
	ThirdRowSeats string `json:"third_row_seats,omitempty"`
	Timestamp float32 `json:"timestamp,omitempty"`
	TrimBadging string `json:"trim_badging,omitempty"`
	WheelType string `json:"wheel_type,omitempty"`
}
