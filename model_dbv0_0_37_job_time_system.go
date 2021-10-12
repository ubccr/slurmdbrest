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

// Dbv0037JobTimeSystem System time values
type Dbv0037JobTimeSystem struct {
	// Total number of CPU-seconds used by the system on behalf of the process (in kernel mode), in seconds
	Seconds *int32 `json:"seconds,omitempty"`
	// Total number of CPU-seconds used by the system on behalf of the process (in kernel mode), in microseconds
	Microseconds *int32 `json:"microseconds,omitempty"`
}

// NewDbv0037JobTimeSystem instantiates a new Dbv0037JobTimeSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobTimeSystem() *Dbv0037JobTimeSystem {
	this := Dbv0037JobTimeSystem{}
	return &this
}

// NewDbv0037JobTimeSystemWithDefaults instantiates a new Dbv0037JobTimeSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobTimeSystemWithDefaults() *Dbv0037JobTimeSystem {
	this := Dbv0037JobTimeSystem{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *Dbv0037JobTimeSystem) GetSeconds() int32 {
	if o == nil || o.Seconds == nil {
		var ret int32
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTimeSystem) GetSecondsOk() (*int32, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *Dbv0037JobTimeSystem) HasSeconds() bool {
	if o != nil && o.Seconds != nil {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int32 and assigns it to the Seconds field.
func (o *Dbv0037JobTimeSystem) SetSeconds(v int32) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *Dbv0037JobTimeSystem) GetMicroseconds() int32 {
	if o == nil || o.Microseconds == nil {
		var ret int32
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTimeSystem) GetMicrosecondsOk() (*int32, bool) {
	if o == nil || o.Microseconds == nil {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *Dbv0037JobTimeSystem) HasMicroseconds() bool {
	if o != nil && o.Microseconds != nil {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int32 and assigns it to the Microseconds field.
func (o *Dbv0037JobTimeSystem) SetMicroseconds(v int32) {
	o.Microseconds = &v
}

func (o Dbv0037JobTimeSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}
	if o.Microseconds != nil {
		toSerialize["microseconds"] = o.Microseconds
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobTimeSystem struct {
	value *Dbv0037JobTimeSystem
	isSet bool
}

func (v NullableDbv0037JobTimeSystem) Get() *Dbv0037JobTimeSystem {
	return v.value
}

func (v *NullableDbv0037JobTimeSystem) Set(val *Dbv0037JobTimeSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobTimeSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobTimeSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobTimeSystem(val *Dbv0037JobTimeSystem) *NullableDbv0037JobTimeSystem {
	return &NullableDbv0037JobTimeSystem{value: val, isSet: true}
}

func (v NullableDbv0037JobTimeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobTimeSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

