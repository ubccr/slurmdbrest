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

// Dbv0037AssociationDefault Default settings
type Dbv0037AssociationDefault struct {
	// Default QOS
	Qos *string `json:"qos,omitempty"`
}

// NewDbv0037AssociationDefault instantiates a new Dbv0037AssociationDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037AssociationDefault() *Dbv0037AssociationDefault {
	this := Dbv0037AssociationDefault{}
	return &this
}

// NewDbv0037AssociationDefaultWithDefaults instantiates a new Dbv0037AssociationDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037AssociationDefaultWithDefaults() *Dbv0037AssociationDefault {
	this := Dbv0037AssociationDefault{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *Dbv0037AssociationDefault) GetQos() string {
	if o == nil || o.Qos == nil {
		var ret string
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037AssociationDefault) GetQosOk() (*string, bool) {
	if o == nil || o.Qos == nil {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *Dbv0037AssociationDefault) HasQos() bool {
	if o != nil && o.Qos != nil {
		return true
	}

	return false
}

// SetQos gets a reference to the given string and assigns it to the Qos field.
func (o *Dbv0037AssociationDefault) SetQos(v string) {
	o.Qos = &v
}

func (o Dbv0037AssociationDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Qos != nil {
		toSerialize["qos"] = o.Qos
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037AssociationDefault struct {
	value *Dbv0037AssociationDefault
	isSet bool
}

func (v NullableDbv0037AssociationDefault) Get() *Dbv0037AssociationDefault {
	return v.value
}

func (v *NullableDbv0037AssociationDefault) Set(val *Dbv0037AssociationDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037AssociationDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037AssociationDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037AssociationDefault(val *Dbv0037AssociationDefault) *NullableDbv0037AssociationDefault {
	return &NullableDbv0037AssociationDefault{value: val, isSet: true}
}

func (v NullableDbv0037AssociationDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037AssociationDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


