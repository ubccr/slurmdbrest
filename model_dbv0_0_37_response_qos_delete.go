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

// Dbv0037ResponseQosDelete struct for Dbv0037ResponseQosDelete
type Dbv0037ResponseQosDelete struct {
	// Slurm errors
	Errors *[]Dbv0037Error `json:"errors,omitempty"`
}

// NewDbv0037ResponseQosDelete instantiates a new Dbv0037ResponseQosDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ResponseQosDelete() *Dbv0037ResponseQosDelete {
	this := Dbv0037ResponseQosDelete{}
	return &this
}

// NewDbv0037ResponseQosDeleteWithDefaults instantiates a new Dbv0037ResponseQosDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ResponseQosDeleteWithDefaults() *Dbv0037ResponseQosDelete {
	this := Dbv0037ResponseQosDelete{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037ResponseQosDelete) GetErrors() []Dbv0037Error {
	if o == nil || o.Errors == nil {
		var ret []Dbv0037Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ResponseQosDelete) GetErrorsOk() (*[]Dbv0037Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037ResponseQosDelete) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037ResponseQosDelete) SetErrors(v []Dbv0037Error) {
	o.Errors = &v
}

func (o Dbv0037ResponseQosDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037ResponseQosDelete struct {
	value *Dbv0037ResponseQosDelete
	isSet bool
}

func (v NullableDbv0037ResponseQosDelete) Get() *Dbv0037ResponseQosDelete {
	return v.value
}

func (v *NullableDbv0037ResponseQosDelete) Set(val *Dbv0037ResponseQosDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ResponseQosDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ResponseQosDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ResponseQosDelete(val *Dbv0037ResponseQosDelete) *NullableDbv0037ResponseQosDelete {
	return &NullableDbv0037ResponseQosDelete{value: val, isSet: true}
}

func (v NullableDbv0037ResponseQosDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ResponseQosDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


