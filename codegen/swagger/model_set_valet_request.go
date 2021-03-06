/*
 * Tesla Model S JSON API
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SetValetRequest struct {
	// Whether to enable or disable valet mode.
	On bool `json:"on,omitempty"`
	// (optional) A 4 digit PIN code to unlock the car.
	Password string `json:"password,omitempty"`
}
