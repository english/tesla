/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NearbyChargerResponseResponse struct {
	CongestionSyncTimeUtcSecs float32 `json:"congestion_sync_time_utc_secs,omitempty"`
	DestinationCharging []DestinationCharger `json:"destination_charging,omitempty"`
	Superchargers []Supercharger `json:"superchargers,omitempty"`
}
