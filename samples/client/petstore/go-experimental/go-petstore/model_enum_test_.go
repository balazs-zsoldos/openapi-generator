/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"encoding/json"
	"errors"
)
// EnumTest struct for EnumTest
type EnumTest struct {
	EnumString *string `json:"enum_string,omitempty"`

	EnumStringRequired *string `json:"enum_string_required,omitempty"`

	EnumInteger *int32 `json:"enum_integer,omitempty"`

	EnumNumber *float64 `json:"enum_number,omitempty"`

	OuterEnum *OuterEnum `json:"outerEnum,omitempty"`

}

// GetEnumString returns the EnumString field if non-nil, zero value otherwise.
func (o *EnumTest) GetEnumString() string {
	if o == nil || o.EnumString == nil {
		var ret string
		return ret
	}
	return *o.EnumString
}

// GetEnumStringOk returns a tuple with the EnumString field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumStringOk() (string, bool) {
	if o == nil || o.EnumString == nil {
		var ret string
		return ret, false
	}
	return *o.EnumString, true
}

// HasEnumString returns a boolean if a field has been set.
func (o *EnumTest) HasEnumString() bool {
	if o != nil && o.EnumString != nil {
		return true
	}

	return false
}

// SetEnumString gets a reference to the given string and assigns it to the EnumString field.
func (o *EnumTest) SetEnumString(v string) {
	o.EnumString = &v
}

// GetEnumStringRequired returns the EnumStringRequired field if non-nil, zero value otherwise.
func (o *EnumTest) GetEnumStringRequired() string {
	if o == nil || o.EnumStringRequired == nil {
		var ret string
		return ret
	}
	return *o.EnumStringRequired
}

// GetEnumStringRequiredOk returns a tuple with the EnumStringRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumStringRequiredOk() (string, bool) {
	if o == nil || o.EnumStringRequired == nil {
		var ret string
		return ret, false
	}
	return *o.EnumStringRequired, true
}

// HasEnumStringRequired returns a boolean if a field has been set.
func (o *EnumTest) HasEnumStringRequired() bool {
	if o != nil && o.EnumStringRequired != nil {
		return true
	}

	return false
}

// SetEnumStringRequired gets a reference to the given string and assigns it to the EnumStringRequired field.
func (o *EnumTest) SetEnumStringRequired(v string) {
	o.EnumStringRequired = &v
}

// GetEnumInteger returns the EnumInteger field if non-nil, zero value otherwise.
func (o *EnumTest) GetEnumInteger() int32 {
	if o == nil || o.EnumInteger == nil {
		var ret int32
		return ret
	}
	return *o.EnumInteger
}

// GetEnumIntegerOk returns a tuple with the EnumInteger field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumIntegerOk() (int32, bool) {
	if o == nil || o.EnumInteger == nil {
		var ret int32
		return ret, false
	}
	return *o.EnumInteger, true
}

// HasEnumInteger returns a boolean if a field has been set.
func (o *EnumTest) HasEnumInteger() bool {
	if o != nil && o.EnumInteger != nil {
		return true
	}

	return false
}

// SetEnumInteger gets a reference to the given int32 and assigns it to the EnumInteger field.
func (o *EnumTest) SetEnumInteger(v int32) {
	o.EnumInteger = &v
}

// GetEnumNumber returns the EnumNumber field if non-nil, zero value otherwise.
func (o *EnumTest) GetEnumNumber() float64 {
	if o == nil || o.EnumNumber == nil {
		var ret float64
		return ret
	}
	return *o.EnumNumber
}

// GetEnumNumberOk returns a tuple with the EnumNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumNumberOk() (float64, bool) {
	if o == nil || o.EnumNumber == nil {
		var ret float64
		return ret, false
	}
	return *o.EnumNumber, true
}

// HasEnumNumber returns a boolean if a field has been set.
func (o *EnumTest) HasEnumNumber() bool {
	if o != nil && o.EnumNumber != nil {
		return true
	}

	return false
}

// SetEnumNumber gets a reference to the given float64 and assigns it to the EnumNumber field.
func (o *EnumTest) SetEnumNumber(v float64) {
	o.EnumNumber = &v
}

// GetOuterEnum returns the OuterEnum field if non-nil, zero value otherwise.
func (o *EnumTest) GetOuterEnum() OuterEnum {
	if o == nil || o.OuterEnum == nil {
		var ret OuterEnum
		return ret
	}
	return *o.OuterEnum
}

// GetOuterEnumOk returns a tuple with the OuterEnum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetOuterEnumOk() (OuterEnum, bool) {
	if o == nil || o.OuterEnum == nil {
		var ret OuterEnum
		return ret, false
	}
	return *o.OuterEnum, true
}

// HasOuterEnum returns a boolean if a field has been set.
func (o *EnumTest) HasOuterEnum() bool {
	if o != nil && o.OuterEnum != nil {
		return true
	}

	return false
}

// SetOuterEnum gets a reference to the given OuterEnum and assigns it to the OuterEnum field.
func (o *EnumTest) SetOuterEnum(v OuterEnum) {
	o.OuterEnum = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o EnumTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnumString != nil {
		toSerialize["enum_string"] = o.EnumString
	}
	if o.EnumStringRequired == nil {
		return nil, errors.New("EnumStringRequired is required and not nullable, but was not set on EnumTest")
	}
	if o.EnumStringRequired != nil {
		toSerialize["enum_string_required"] = o.EnumStringRequired
	}
	if o.EnumInteger != nil {
		toSerialize["enum_integer"] = o.EnumInteger
	}
	if o.EnumNumber != nil {
		toSerialize["enum_number"] = o.EnumNumber
	}
	if o.OuterEnum != nil {
		toSerialize["outerEnum"] = o.OuterEnum
	}
	return json.Marshal(toSerialize)
}


