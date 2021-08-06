/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"encoding/json"
)

// SpreadTarget struct for SpreadTarget
type SpreadTarget struct {
	Percent *int32  `json:"Percent,omitempty"`
	Value   *string `json:"Value,omitempty"`
}

// NewSpreadTarget instantiates a new SpreadTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpreadTarget() *SpreadTarget {
	this := SpreadTarget{}
	return &this
}

// NewSpreadTargetWithDefaults instantiates a new SpreadTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpreadTargetWithDefaults() *SpreadTarget {
	this := SpreadTarget{}
	return &this
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *SpreadTarget) GetPercent() int32 {
	if o == nil || o.Percent == nil {
		var ret int32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpreadTarget) GetPercentOk() (*int32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *SpreadTarget) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given int32 and assigns it to the Percent field.
func (o *SpreadTarget) SetPercent(v int32) {
	o.Percent = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SpreadTarget) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpreadTarget) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SpreadTarget) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SpreadTarget) SetValue(v string) {
	o.Value = &v
}

func (o SpreadTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Percent != nil {
		toSerialize["Percent"] = o.Percent
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSpreadTarget struct {
	value *SpreadTarget
	isSet bool
}

func (v NullableSpreadTarget) Get() *SpreadTarget {
	return v.value
}

func (v *NullableSpreadTarget) Set(val *SpreadTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSpreadTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSpreadTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpreadTarget(val *SpreadTarget) *NullableSpreadTarget {
	return &NullableSpreadTarget{value: val, isSet: true}
}

func (v NullableSpreadTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpreadTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}