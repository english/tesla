/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SpeedLimitRequest struct {
	// The current, or if activating, new, speed limit PIN
	Pin string `json:"pin,omitempty"`
}
