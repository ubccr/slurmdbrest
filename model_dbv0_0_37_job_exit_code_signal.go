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

// Dbv0037JobExitCodeSignal Signal details (if signaled)
type Dbv0037JobExitCodeSignal struct {
	// Signal number process received
	SignalId *int32 `json:"signal_id,omitempty"`
	// Name of signal received
	Name *string `json:"name,omitempty"`
}

// NewDbv0037JobExitCodeSignal instantiates a new Dbv0037JobExitCodeSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobExitCodeSignal() *Dbv0037JobExitCodeSignal {
	this := Dbv0037JobExitCodeSignal{}
	return &this
}

// NewDbv0037JobExitCodeSignalWithDefaults instantiates a new Dbv0037JobExitCodeSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobExitCodeSignalWithDefaults() *Dbv0037JobExitCodeSignal {
	this := Dbv0037JobExitCodeSignal{}
	return &this
}

// GetSignalId returns the SignalId field value if set, zero value otherwise.
func (o *Dbv0037JobExitCodeSignal) GetSignalId() int32 {
	if o == nil || o.SignalId == nil {
		var ret int32
		return ret
	}
	return *o.SignalId
}

// GetSignalIdOk returns a tuple with the SignalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobExitCodeSignal) GetSignalIdOk() (*int32, bool) {
	if o == nil || o.SignalId == nil {
		return nil, false
	}
	return o.SignalId, true
}

// HasSignalId returns a boolean if a field has been set.
func (o *Dbv0037JobExitCodeSignal) HasSignalId() bool {
	if o != nil && o.SignalId != nil {
		return true
	}

	return false
}

// SetSignalId gets a reference to the given int32 and assigns it to the SignalId field.
func (o *Dbv0037JobExitCodeSignal) SetSignalId(v int32) {
	o.SignalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037JobExitCodeSignal) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobExitCodeSignal) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037JobExitCodeSignal) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0037JobExitCodeSignal) SetName(v string) {
	o.Name = &v
}

func (o Dbv0037JobExitCodeSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SignalId != nil {
		toSerialize["signal_id"] = o.SignalId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobExitCodeSignal struct {
	value *Dbv0037JobExitCodeSignal
	isSet bool
}

func (v NullableDbv0037JobExitCodeSignal) Get() *Dbv0037JobExitCodeSignal {
	return v.value
}

func (v *NullableDbv0037JobExitCodeSignal) Set(val *Dbv0037JobExitCodeSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobExitCodeSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobExitCodeSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobExitCodeSignal(val *Dbv0037JobExitCodeSignal) *NullableDbv0037JobExitCodeSignal {
	return &NullableDbv0037JobExitCodeSignal{value: val, isSet: true}
}

func (v NullableDbv0037JobExitCodeSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobExitCodeSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


