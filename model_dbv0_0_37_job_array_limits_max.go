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

// Dbv0037JobArrayLimitsMax Limits on array settings
type Dbv0037JobArrayLimitsMax struct {
	Running *Dbv0037JobArrayLimitsMaxRunning `json:"running,omitempty"`
}

// NewDbv0037JobArrayLimitsMax instantiates a new Dbv0037JobArrayLimitsMax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobArrayLimitsMax() *Dbv0037JobArrayLimitsMax {
	this := Dbv0037JobArrayLimitsMax{}
	return &this
}

// NewDbv0037JobArrayLimitsMaxWithDefaults instantiates a new Dbv0037JobArrayLimitsMax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobArrayLimitsMaxWithDefaults() *Dbv0037JobArrayLimitsMax {
	this := Dbv0037JobArrayLimitsMax{}
	return &this
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *Dbv0037JobArrayLimitsMax) GetRunning() Dbv0037JobArrayLimitsMaxRunning {
	if o == nil || o.Running == nil {
		var ret Dbv0037JobArrayLimitsMaxRunning
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobArrayLimitsMax) GetRunningOk() (*Dbv0037JobArrayLimitsMaxRunning, bool) {
	if o == nil || o.Running == nil {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *Dbv0037JobArrayLimitsMax) HasRunning() bool {
	if o != nil && o.Running != nil {
		return true
	}

	return false
}

// SetRunning gets a reference to the given Dbv0037JobArrayLimitsMaxRunning and assigns it to the Running field.
func (o *Dbv0037JobArrayLimitsMax) SetRunning(v Dbv0037JobArrayLimitsMaxRunning) {
	o.Running = &v
}

func (o Dbv0037JobArrayLimitsMax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Running != nil {
		toSerialize["running"] = o.Running
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobArrayLimitsMax struct {
	value *Dbv0037JobArrayLimitsMax
	isSet bool
}

func (v NullableDbv0037JobArrayLimitsMax) Get() *Dbv0037JobArrayLimitsMax {
	return v.value
}

func (v *NullableDbv0037JobArrayLimitsMax) Set(val *Dbv0037JobArrayLimitsMax) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobArrayLimitsMax) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobArrayLimitsMax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobArrayLimitsMax(val *Dbv0037JobArrayLimitsMax) *NullableDbv0037JobArrayLimitsMax {
	return &NullableDbv0037JobArrayLimitsMax{value: val, isSet: true}
}

func (v NullableDbv0037JobArrayLimitsMax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobArrayLimitsMax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

