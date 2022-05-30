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

type AllOfClusterProviderMinWorkerSize struct {
	Name string `json:"name,omitempty"`
	Cpu int32 `json:"cpu,omitempty"`
	Memory int32 `json:"memory,omitempty"`
	RootDisk int32 `json:"rootDisk,omitempty"`
	EphemeralDisk int32 `json:"ephemeralDisk,omitempty"`
	PersistentDisk int32 `json:"persistentDisk,omitempty"`
}
