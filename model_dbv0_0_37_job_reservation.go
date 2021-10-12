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

// Dbv0037JobReservation Reservation usage details
type Dbv0037JobReservation struct {
	// Database id of reservation
	Id *int32 `json:"id,omitempty"`
	// Name of reservation
	Name *int32 `json:"name,omitempty"`
}

// NewDbv0037JobReservation instantiates a new Dbv0037JobReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobReservation() *Dbv0037JobReservation {
	this := Dbv0037JobReservation{}
	return &this
}

// NewDbv0037JobReservationWithDefaults instantiates a new Dbv0037JobReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobReservationWithDefaults() *Dbv0037JobReservation {
	this := Dbv0037JobReservation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0037JobReservation) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobReservation) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0037JobReservation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Dbv0037JobReservation) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037JobReservation) GetName() int32 {
	if o == nil || o.Name == nil {
		var ret int32
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobReservation) GetNameOk() (*int32, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037JobReservation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given int32 and assigns it to the Name field.
func (o *Dbv0037JobReservation) SetName(v int32) {
	o.Name = &v
}

func (o Dbv0037JobReservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobReservation struct {
	value *Dbv0037JobReservation
	isSet bool
}

func (v NullableDbv0037JobReservation) Get() *Dbv0037JobReservation {
	return v.value
}

func (v *NullableDbv0037JobReservation) Set(val *Dbv0037JobReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobReservation(val *Dbv0037JobReservation) *NullableDbv0037JobReservation {
	return &NullableDbv0037JobReservation{value: val, isSet: true}
}

func (v NullableDbv0037JobReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

