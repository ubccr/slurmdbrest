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

// Dbv0037ResponseWckeyDelete struct for Dbv0037ResponseWckeyDelete
type Dbv0037ResponseWckeyDelete struct {
	// Slurm errors
	Errors *[]Dbv0037Error `json:"errors,omitempty"`
}

// NewDbv0037ResponseWckeyDelete instantiates a new Dbv0037ResponseWckeyDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ResponseWckeyDelete() *Dbv0037ResponseWckeyDelete {
	this := Dbv0037ResponseWckeyDelete{}
	return &this
}

// NewDbv0037ResponseWckeyDeleteWithDefaults instantiates a new Dbv0037ResponseWckeyDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ResponseWckeyDeleteWithDefaults() *Dbv0037ResponseWckeyDelete {
	this := Dbv0037ResponseWckeyDelete{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037ResponseWckeyDelete) GetErrors() []Dbv0037Error {
	if o == nil || o.Errors == nil {
		var ret []Dbv0037Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ResponseWckeyDelete) GetErrorsOk() (*[]Dbv0037Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037ResponseWckeyDelete) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037ResponseWckeyDelete) SetErrors(v []Dbv0037Error) {
	o.Errors = &v
}

func (o Dbv0037ResponseWckeyDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037ResponseWckeyDelete struct {
	value *Dbv0037ResponseWckeyDelete
	isSet bool
}

func (v NullableDbv0037ResponseWckeyDelete) Get() *Dbv0037ResponseWckeyDelete {
	return v.value
}

func (v *NullableDbv0037ResponseWckeyDelete) Set(val *Dbv0037ResponseWckeyDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ResponseWckeyDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ResponseWckeyDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ResponseWckeyDelete(val *Dbv0037ResponseWckeyDelete) *NullableDbv0037ResponseWckeyDelete {
	return &NullableDbv0037ResponseWckeyDelete{value: val, isSet: true}
}

func (v NullableDbv0037ResponseWckeyDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ResponseWckeyDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


