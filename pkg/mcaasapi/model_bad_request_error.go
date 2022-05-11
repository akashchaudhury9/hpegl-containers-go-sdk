/*
 * MCaaS API
 *
 * APIs and object schemas for MCaaS.
 *
 * API version: v1
 * Contact: Centaur-GLHC@hpe.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcaasapi

type BadRequestError struct {
	Message string `json:"Message,omitempty"`
	Details string `json:"Details,omitempty"`
	RecommendedActions []string `json:"RecommendedActions,omitempty"`
	ErrorCode string `json:"ErrorCode,omitempty"`
}
