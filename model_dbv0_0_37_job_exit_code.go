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

// Dbv0037JobExitCode struct for Dbv0037JobExitCode
type Dbv0037JobExitCode struct {
	// Job exit status
	Status *string `json:"status,omitempty"`
	// Return code from parent process
	ReturnCode *int32 `json:"return_code,omitempty"`
	Signal *Dbv0037JobExitCodeSignal `json:"signal,omitempty"`
}

// NewDbv0037JobExitCode instantiates a new Dbv0037JobExitCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobExitCode() *Dbv0037JobExitCode {
	this := Dbv0037JobExitCode{}
	return &this
}

// NewDbv0037JobExitCodeWithDefaults instantiates a new Dbv0037JobExitCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobExitCodeWithDefaults() *Dbv0037JobExitCode {
	this := Dbv0037JobExitCode{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Dbv0037JobExitCode) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobExitCode) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Dbv0037JobExitCode) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Dbv0037JobExitCode) SetStatus(v string) {
	o.Status = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *Dbv0037JobExitCode) GetReturnCode() int32 {
	if o == nil || o.ReturnCode == nil {
		var ret int32
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobExitCode) GetReturnCodeOk() (*int32, bool) {
	if o == nil || o.ReturnCode == nil {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *Dbv0037JobExitCode) HasReturnCode() bool {
	if o != nil && o.ReturnCode != nil {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given int32 and assigns it to the ReturnCode field.
func (o *Dbv0037JobExitCode) SetReturnCode(v int32) {
	o.ReturnCode = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *Dbv0037JobExitCode) GetSignal() Dbv0037JobExitCodeSignal {
	if o == nil || o.Signal == nil {
		var ret Dbv0037JobExitCodeSignal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobExitCode) GetSignalOk() (*Dbv0037JobExitCodeSignal, bool) {
	if o == nil || o.Signal == nil {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *Dbv0037JobExitCode) HasSignal() bool {
	if o != nil && o.Signal != nil {
		return true
	}

	return false
}

// SetSignal gets a reference to the given Dbv0037JobExitCodeSignal and assigns it to the Signal field.
func (o *Dbv0037JobExitCode) SetSignal(v Dbv0037JobExitCodeSignal) {
	o.Signal = &v
}

func (o Dbv0037JobExitCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ReturnCode != nil {
		toSerialize["return_code"] = o.ReturnCode
	}
	if o.Signal != nil {
		toSerialize["signal"] = o.Signal
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobExitCode struct {
	value *Dbv0037JobExitCode
	isSet bool
}

func (v NullableDbv0037JobExitCode) Get() *Dbv0037JobExitCode {
	return v.value
}

func (v *NullableDbv0037JobExitCode) Set(val *Dbv0037JobExitCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobExitCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobExitCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobExitCode(val *Dbv0037JobExitCode) *NullableDbv0037JobExitCode {
	return &NullableDbv0037JobExitCode{value: val, isSet: true}
}

func (v NullableDbv0037JobExitCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobExitCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


