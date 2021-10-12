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

// Dbv0037DiagStatisticsRPCs Statistics by RPC type
type Dbv0037DiagStatisticsRPCs struct {
	// RPC type
	Rpc *string `json:"rpc,omitempty"`
	// Number of RPCs
	Count *int32 `json:"count,omitempty"`
	Time *Dbv0037DiagStatisticsTime `json:"time,omitempty"`
}

// NewDbv0037DiagStatisticsRPCs instantiates a new Dbv0037DiagStatisticsRPCs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037DiagStatisticsRPCs() *Dbv0037DiagStatisticsRPCs {
	this := Dbv0037DiagStatisticsRPCs{}
	return &this
}

// NewDbv0037DiagStatisticsRPCsWithDefaults instantiates a new Dbv0037DiagStatisticsRPCs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037DiagStatisticsRPCsWithDefaults() *Dbv0037DiagStatisticsRPCs {
	this := Dbv0037DiagStatisticsRPCs{}
	return &this
}

// GetRpc returns the Rpc field value if set, zero value otherwise.
func (o *Dbv0037DiagStatisticsRPCs) GetRpc() string {
	if o == nil || o.Rpc == nil {
		var ret string
		return ret
	}
	return *o.Rpc
}

// GetRpcOk returns a tuple with the Rpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037DiagStatisticsRPCs) GetRpcOk() (*string, bool) {
	if o == nil || o.Rpc == nil {
		return nil, false
	}
	return o.Rpc, true
}

// HasRpc returns a boolean if a field has been set.
func (o *Dbv0037DiagStatisticsRPCs) HasRpc() bool {
	if o != nil && o.Rpc != nil {
		return true
	}

	return false
}

// SetRpc gets a reference to the given string and assigns it to the Rpc field.
func (o *Dbv0037DiagStatisticsRPCs) SetRpc(v string) {
	o.Rpc = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Dbv0037DiagStatisticsRPCs) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037DiagStatisticsRPCs) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Dbv0037DiagStatisticsRPCs) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Dbv0037DiagStatisticsRPCs) SetCount(v int32) {
	o.Count = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Dbv0037DiagStatisticsRPCs) GetTime() Dbv0037DiagStatisticsTime {
	if o == nil || o.Time == nil {
		var ret Dbv0037DiagStatisticsTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037DiagStatisticsRPCs) GetTimeOk() (*Dbv0037DiagStatisticsTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Dbv0037DiagStatisticsRPCs) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given Dbv0037DiagStatisticsTime and assigns it to the Time field.
func (o *Dbv0037DiagStatisticsRPCs) SetTime(v Dbv0037DiagStatisticsTime) {
	o.Time = &v
}

func (o Dbv0037DiagStatisticsRPCs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rpc != nil {
		toSerialize["rpc"] = o.Rpc
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037DiagStatisticsRPCs struct {
	value *Dbv0037DiagStatisticsRPCs
	isSet bool
}

func (v NullableDbv0037DiagStatisticsRPCs) Get() *Dbv0037DiagStatisticsRPCs {
	return v.value
}

func (v *NullableDbv0037DiagStatisticsRPCs) Set(val *Dbv0037DiagStatisticsRPCs) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037DiagStatisticsRPCs) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037DiagStatisticsRPCs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037DiagStatisticsRPCs(val *Dbv0037DiagStatisticsRPCs) *NullableDbv0037DiagStatisticsRPCs {
	return &NullableDbv0037DiagStatisticsRPCs{value: val, isSet: true}
}

func (v NullableDbv0037DiagStatisticsRPCs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037DiagStatisticsRPCs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

