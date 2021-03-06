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

// Dbv0037AssociationMaxTresMinutesPer Max TRES minutes per settings
type Dbv0037AssociationMaxTresMinutesPer struct {
	// TRES list of attributes
	Job *[]map[string]interface{} `json:"job,omitempty"`
}

// NewDbv0037AssociationMaxTresMinutesPer instantiates a new Dbv0037AssociationMaxTresMinutesPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationMaxTresMinutesPer() *Dbv0037AssociationMaxTresMinutesPer {
	this := Dbv0037AssociationMaxTresMinutesPer{}
	return &this
}

// NewDbv0037AssociationMaxTresMinutesPerWithDefaults instantiates a new Dbv0037AssociationMaxTresMinutesPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationMaxTresMinutesPerWithDefaults() *Dbv0037AssociationMaxTresMinutesPer {
	this := Dbv0037AssociationMaxTresMinutesPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Dbv0037AssociationMaxTresMinutesPer) GetJob() []map[string]interface{} {
	if o == nil || o.Job == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationMaxTresMinutesPer) GetJobOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Dbv0037AssociationMaxTresMinutesPer) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given []map[string]interface{} and assigns it to the Job field.
func (o *Dbv0037AssociationMaxTresMinutesPer) SetJob(v []map[string]interface{}) {
	o.Job = &v
}

func (o Dbv0037AssociationMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037AssociationMaxTresMinutesPer struct {
	value *Dbv0037AssociationMaxTresMinutesPer
	isSet bool
}

func (v NullableDbv0037AssociationMaxTresMinutesPer) Get() *Dbv0037AssociationMaxTresMinutesPer {
	return v.value
}

func (v *NullableDbv0037AssociationMaxTresMinutesPer) Set(val *Dbv0037AssociationMaxTresMinutesPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationMaxTresMinutesPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationMaxTresMinutesPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationMaxTresMinutesPer(val *Dbv0037AssociationMaxTresMinutesPer) *NullableDbv0037AssociationMaxTresMinutesPer {
	return &NullableDbv0037AssociationMaxTresMinutesPer{value: val, isSet: true}
}

func (v NullableDbv0037AssociationMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationMaxTresMinutesPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


