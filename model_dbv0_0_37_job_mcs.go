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

// Dbv0037JobMcs Multi-Category Security
type Dbv0037JobMcs struct {
	// Assigned MCS label
	Label *string `json:"label,omitempty"`
}

// NewDbv0037JobMcs instantiates a new Dbv0037JobMcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobMcs() *Dbv0037JobMcs {
	this := Dbv0037JobMcs{}
	return &this
}

// NewDbv0037JobMcsWithDefaults instantiates a new Dbv0037JobMcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobMcsWithDefaults() *Dbv0037JobMcs {
	this := Dbv0037JobMcs{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Dbv0037JobMcs) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobMcs) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Dbv0037JobMcs) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Dbv0037JobMcs) SetLabel(v string) {
	o.Label = &v
}

func (o Dbv0037JobMcs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobMcs struct {
	value *Dbv0037JobMcs
	isSet bool
}

func (v NullableDbv0037JobMcs) Get() *Dbv0037JobMcs {
	return v.value
}

func (v *NullableDbv0037JobMcs) Set(val *Dbv0037JobMcs) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobMcs) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobMcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobMcs(val *Dbv0037JobMcs) *NullableDbv0037JobMcs {
	return &NullableDbv0037JobMcs{value: val, isSet: true}
}

func (v NullableDbv0037JobMcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobMcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


