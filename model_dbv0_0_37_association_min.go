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

// Dbv0037AssociationMin Min settings
type Dbv0037AssociationMin struct {
	// Min priority threshold
	PriorityThreshold *int32 `json:"priority_threshold,omitempty"`
}

// NewDbv0037AssociationMin instantiates a new Dbv0037AssociationMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationMin() *Dbv0037AssociationMin {
	this := Dbv0037AssociationMin{}
	return &this
}

// NewDbv0037AssociationMinWithDefaults instantiates a new Dbv0037AssociationMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationMinWithDefaults() *Dbv0037AssociationMin {
	this := Dbv0037AssociationMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *Dbv0037AssociationMin) GetPriorityThreshold() int32 {
	if o == nil || o.PriorityThreshold == nil {
		var ret int32
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMin) GetPriorityThresholdOk() (*int32, bool) {
	if o == nil || o.PriorityThreshold == nil {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *Dbv0037AssociationMin) HasPriorityThreshold() bool {
	if o != nil && o.PriorityThreshold != nil {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given int32 and assigns it to the PriorityThreshold field.
func (o *Dbv0037AssociationMin) SetPriorityThreshold(v int32) {
	o.PriorityThreshold = &v
}

func (o Dbv0037AssociationMin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PriorityThreshold != nil {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037AssociationMin struct {
	value *Dbv0037AssociationMin
	isSet bool
}

func (v NullableDbv0037AssociationMin) Get() *Dbv0037AssociationMin {
	return v.value
}

func (v *NullableDbv0037AssociationMin) Set(val *Dbv0037AssociationMin) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationMin) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationMin(val *Dbv0037AssociationMin) *NullableDbv0037AssociationMin {
	return &NullableDbv0037AssociationMin{value: val, isSet: true}
}

func (v NullableDbv0037AssociationMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


