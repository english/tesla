/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ShareRequestValue struct {
	// Must match this syntax \"DATA_TO_SEND_TO_NAV_SYSTEM\\n\\nhttps://goo.gl/maps/X\"
	AndroidIntentExtraTEXT string `json:"android.intent.extra.TEXT,omitempty"`
	// Latitude of destination to send. Only send with 'long'
	Lat float64 `json:"lat,omitempty"`
	// Longitude of destination to send. Only send with 'lat'
	Long float64 `json:"long,omitempty"`
}
