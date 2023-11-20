/*
maestro API

maestro API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ErrorListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorListAllOf{}

// ErrorListAllOf struct for ErrorListAllOf
type ErrorListAllOf struct {
	Items []Error `json:"items,omitempty"`
}

// NewErrorListAllOf instantiates a new ErrorListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorListAllOf() *ErrorListAllOf {
	this := ErrorListAllOf{}
	return &this
}

// NewErrorListAllOfWithDefaults instantiates a new ErrorListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorListAllOfWithDefaults() *ErrorListAllOf {
	this := ErrorListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ErrorListAllOf) GetItems() []Error {
	if o == nil || IsNil(o.Items) {
		var ret []Error
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorListAllOf) GetItemsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ErrorListAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Error and assigns it to the Items field.
func (o *ErrorListAllOf) SetItems(v []Error) {
	o.Items = v
}

func (o ErrorListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableErrorListAllOf struct {
	value *ErrorListAllOf
	isSet bool
}

func (v NullableErrorListAllOf) Get() *ErrorListAllOf {
	return v.value
}

func (v *NullableErrorListAllOf) Set(val *ErrorListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorListAllOf(val *ErrorListAllOf) *NullableErrorListAllOf {
	return &NullableErrorListAllOf{value: val, isSet: true}
}

func (v NullableErrorListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}