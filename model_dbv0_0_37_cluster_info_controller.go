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

// Dbv0037ClusterInfoController Information about controller
type Dbv0037ClusterInfoController struct {
	// Hostname
	Host *string `json:"host,omitempty"`
	// Port number
	Port *int32 `json:"port,omitempty"`
}

// NewDbv0037ClusterInfoController instantiates a new Dbv0037ClusterInfoController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037ClusterInfoController() *Dbv0037ClusterInfoController {
	this := Dbv0037ClusterInfoController{}
	return &this
}

// NewDbv0037ClusterInfoControllerWithDefaults instantiates a new Dbv0037ClusterInfoController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037ClusterInfoControllerWithDefaults() *Dbv0037ClusterInfoController {
	this := Dbv0037ClusterInfoController{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfoController) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfoController) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfoController) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Dbv0037ClusterInfoController) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Dbv0037ClusterInfoController) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037ClusterInfoController) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Dbv0037ClusterInfoController) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *Dbv0037ClusterInfoController) SetPort(v int32) {
	o.Port = &v
}

func (o Dbv0037ClusterInfoController) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037ClusterInfoController struct {
	value *Dbv0037ClusterInfoController
	isSet bool
}

func (v NullableDbv0037ClusterInfoController) Get() *Dbv0037ClusterInfoController {
	return v.value
}

func (v *NullableDbv0037ClusterInfoController) Set(val *Dbv0037ClusterInfoController) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037ClusterInfoController) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037ClusterInfoController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037ClusterInfoController(val *Dbv0037ClusterInfoController) *NullableDbv0037ClusterInfoController {
	return &NullableDbv0037ClusterInfoController{value: val, isSet: true}
}

func (v NullableDbv0037ClusterInfoController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037ClusterInfoController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

