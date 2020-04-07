/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
)

// AdditionalPropertiesNumber struct for AdditionalPropertiesNumber
type AdditionalPropertiesNumber struct {
	Name *string `json:"name,omitempty"`
}

// NewAdditionalPropertiesNumber instantiates a new AdditionalPropertiesNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalPropertiesNumber() *AdditionalPropertiesNumber {
	this := AdditionalPropertiesNumber{}
	return &this
}

// NewAdditionalPropertiesNumberWithDefaults instantiates a new AdditionalPropertiesNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalPropertiesNumberWithDefaults() *AdditionalPropertiesNumber {
	this := AdditionalPropertiesNumber{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalPropertiesNumber) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesNumber) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalPropertiesNumber) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalPropertiesNumber) SetName(v string) {
	o.Name = &v
}

func (o AdditionalPropertiesNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalPropertiesNumber struct {
	value *AdditionalPropertiesNumber
	isSet bool
}

func (v NullableAdditionalPropertiesNumber) Get() *AdditionalPropertiesNumber {
	return v.value
}

func (v *NullableAdditionalPropertiesNumber) Set(val *AdditionalPropertiesNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalPropertiesNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalPropertiesNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalPropertiesNumber(val *AdditionalPropertiesNumber) *NullableAdditionalPropertiesNumber {
	return &NullableAdditionalPropertiesNumber{value: val, isSet: true}
}

func (v NullableAdditionalPropertiesNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalPropertiesNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
