/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore
// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// Form parameter enum test (string array)
	EnumFormStringArray []string `json:"enum_form_string_array,omitempty"`
	// Form parameter enum test (string)
	EnumFormString string `json:"enum_form_string,omitempty"`
}
