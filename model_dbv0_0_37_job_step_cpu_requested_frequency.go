/*
 * Slurm Rest API
 *
 * API to access and control Slurm DB.
 *
 * API version: dbv0.0.37
 * Contact: sales@schedmd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmdbrest

import (
	"encoding/json"
)

// Dbv0037JobStepCPURequestedFrequency CPU frequency requested
type Dbv0037JobStepCPURequestedFrequency struct {
	// Min CPU frequency
	Min *int32 `json:"min,omitempty"`
	// Max CPU frequency
	Max *int32 `json:"max,omitempty"`
}

// NewDbv0037JobStepCPURequestedFrequency instantiates a new Dbv0037JobStepCPURequestedFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepCPURequestedFrequency() *Dbv0037JobStepCPURequestedFrequency {
	this := Dbv0037JobStepCPURequestedFrequency{}
	return &this
}

// NewDbv0037JobStepCPURequestedFrequencyWithDefaults instantiates a new Dbv0037JobStepCPURequestedFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepCPURequestedFrequencyWithDefaults() *Dbv0037JobStepCPURequestedFrequency {
	this := Dbv0037JobStepCPURequestedFrequency{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Dbv0037JobStepCPURequestedFrequency) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepCPURequestedFrequency) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Dbv0037JobStepCPURequestedFrequency) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *Dbv0037JobStepCPURequestedFrequency) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Dbv0037JobStepCPURequestedFrequency) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepCPURequestedFrequency) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Dbv0037JobStepCPURequestedFrequency) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *Dbv0037JobStepCPURequestedFrequency) SetMax(v int32) {
	o.Max = &v
}

func (o Dbv0037JobStepCPURequestedFrequency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobStepCPURequestedFrequency struct {
	value *Dbv0037JobStepCPURequestedFrequency
	isSet bool
}

func (v NullableDbv0037JobStepCPURequestedFrequency) Get() *Dbv0037JobStepCPURequestedFrequency {
	return v.value
}

func (v *NullableDbv0037JobStepCPURequestedFrequency) Set(val *Dbv0037JobStepCPURequestedFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepCPURequestedFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepCPURequestedFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepCPURequestedFrequency(val *Dbv0037JobStepCPURequestedFrequency) *NullableDbv0037JobStepCPURequestedFrequency {
	return &NullableDbv0037JobStepCPURequestedFrequency{value: val, isSet: true}
}

func (v NullableDbv0037JobStepCPURequestedFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepCPURequestedFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

