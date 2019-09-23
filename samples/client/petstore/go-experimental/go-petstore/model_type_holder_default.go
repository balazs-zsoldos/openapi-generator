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
// TypeHolderDefault struct for TypeHolderDefault
type TypeHolderDefault struct {
	StringItem *string `json:"string_item,omitempty"`

	NumberItem *float32 `json:"number_item,omitempty"`

	IntegerItem *int32 `json:"integer_item,omitempty"`

	BoolItem *bool `json:"bool_item,omitempty"`

	ArrayItem *[]int32 `json:"array_item,omitempty"`

}

// GetStringItem returns the StringItem field if non-nil, zero value otherwise.
func (o *TypeHolderDefault) GetStringItem() string {
	if o == nil || o.StringItem == nil {
		var ret string
		return ret
	}
	return *o.StringItem
}

// GetStringItemOk returns a tuple with the StringItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetStringItemOk() (string, bool) {
	if o == nil || o.StringItem == nil {
		var ret string
		return ret, false
	}
	return *o.StringItem, true
}

// HasStringItem returns a boolean if a field has been set.
func (o *TypeHolderDefault) HasStringItem() bool {
	if o != nil && o.StringItem != nil {
		return true
	}

	return false
}

// SetStringItem gets a reference to the given string and assigns it to the StringItem field.
func (o *TypeHolderDefault) SetStringItem(v string) {
	o.StringItem = &v
}

// GetNumberItem returns the NumberItem field if non-nil, zero value otherwise.
func (o *TypeHolderDefault) GetNumberItem() float32 {
	if o == nil || o.NumberItem == nil {
		var ret float32
		return ret
	}
	return *o.NumberItem
}

// GetNumberItemOk returns a tuple with the NumberItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetNumberItemOk() (float32, bool) {
	if o == nil || o.NumberItem == nil {
		var ret float32
		return ret, false
	}
	return *o.NumberItem, true
}

// HasNumberItem returns a boolean if a field has been set.
func (o *TypeHolderDefault) HasNumberItem() bool {
	if o != nil && o.NumberItem != nil {
		return true
	}

	return false
}

// SetNumberItem gets a reference to the given float32 and assigns it to the NumberItem field.
func (o *TypeHolderDefault) SetNumberItem(v float32) {
	o.NumberItem = &v
}

// GetIntegerItem returns the IntegerItem field if non-nil, zero value otherwise.
func (o *TypeHolderDefault) GetIntegerItem() int32 {
	if o == nil || o.IntegerItem == nil {
		var ret int32
		return ret
	}
	return *o.IntegerItem
}

// GetIntegerItemOk returns a tuple with the IntegerItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetIntegerItemOk() (int32, bool) {
	if o == nil || o.IntegerItem == nil {
		var ret int32
		return ret, false
	}
	return *o.IntegerItem, true
}

// HasIntegerItem returns a boolean if a field has been set.
func (o *TypeHolderDefault) HasIntegerItem() bool {
	if o != nil && o.IntegerItem != nil {
		return true
	}

	return false
}

// SetIntegerItem gets a reference to the given int32 and assigns it to the IntegerItem field.
func (o *TypeHolderDefault) SetIntegerItem(v int32) {
	o.IntegerItem = &v
}

// GetBoolItem returns the BoolItem field if non-nil, zero value otherwise.
func (o *TypeHolderDefault) GetBoolItem() bool {
	if o == nil || o.BoolItem == nil {
		var ret bool
		return ret
	}
	return *o.BoolItem
}

// GetBoolItemOk returns a tuple with the BoolItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetBoolItemOk() (bool, bool) {
	if o == nil || o.BoolItem == nil {
		var ret bool
		return ret, false
	}
	return *o.BoolItem, true
}

// HasBoolItem returns a boolean if a field has been set.
func (o *TypeHolderDefault) HasBoolItem() bool {
	if o != nil && o.BoolItem != nil {
		return true
	}

	return false
}

// SetBoolItem gets a reference to the given bool and assigns it to the BoolItem field.
func (o *TypeHolderDefault) SetBoolItem(v bool) {
	o.BoolItem = &v
}

// GetArrayItem returns the ArrayItem field if non-nil, zero value otherwise.
func (o *TypeHolderDefault) GetArrayItem() []int32 {
	if o == nil || o.ArrayItem == nil {
		var ret []int32
		return ret
	}
	return *o.ArrayItem
}

// GetArrayItemOk returns a tuple with the ArrayItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetArrayItemOk() ([]int32, bool) {
	if o == nil || o.ArrayItem == nil {
		var ret []int32
		return ret, false
	}
	return *o.ArrayItem, true
}

// HasArrayItem returns a boolean if a field has been set.
func (o *TypeHolderDefault) HasArrayItem() bool {
	if o != nil && o.ArrayItem != nil {
		return true
	}

	return false
}

// SetArrayItem gets a reference to the given []int32 and assigns it to the ArrayItem field.
func (o *TypeHolderDefault) SetArrayItem(v []int32) {
	o.ArrayItem = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o TypeHolderDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StringItem == nil {
		return nil, errors.New("StringItem is required and not nullable, but was not set on TypeHolderDefault")
	}
	if o.StringItem != nil {
		toSerialize["string_item"] = o.StringItem
	}
	if o.NumberItem == nil {
		return nil, errors.New("NumberItem is required and not nullable, but was not set on TypeHolderDefault")
	}
	if o.NumberItem != nil {
		toSerialize["number_item"] = o.NumberItem
	}
	if o.IntegerItem == nil {
		return nil, errors.New("IntegerItem is required and not nullable, but was not set on TypeHolderDefault")
	}
	if o.IntegerItem != nil {
		toSerialize["integer_item"] = o.IntegerItem
	}
	if o.BoolItem == nil {
		return nil, errors.New("BoolItem is required and not nullable, but was not set on TypeHolderDefault")
	}
	if o.BoolItem != nil {
		toSerialize["bool_item"] = o.BoolItem
	}
	if o.ArrayItem == nil {
		return nil, errors.New("ArrayItem is required and not nullable, but was not set on TypeHolderDefault")
	}
	if o.ArrayItem != nil {
		toSerialize["array_item"] = o.ArrayItem
	}
	return json.Marshal(toSerialize)
}


