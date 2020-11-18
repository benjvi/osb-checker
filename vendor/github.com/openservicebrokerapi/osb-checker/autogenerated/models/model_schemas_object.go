/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SchemasObject struct {
	ServiceInstance ServiceInstanceSchemaObject `json:"service_instance,omitempty"`

	ServiceBinding ServiceBindingSchemaObject `json:"service_binding,omitempty"`
}
