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

// Dbv0037QosLimitsMaxTresMinutesPer Max TRES minutes per settings
type Dbv0037QosLimitsMaxTresMinutesPer struct {
	// TRES list of attributes
	Job *[]map[string]interface{} `json:"job,omitempty"`
	// TRES list of attributes
	Account *[]map[string]interface{} `json:"account,omitempty"`
	// TRES list of attributes
	User *[]map[string]interface{} `json:"user,omitempty"`
}

// NewDbv0037QosLimitsMaxTresMinutesPer instantiates a new Dbv0037QosLimitsMaxTresMinutesPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037QosLimitsMaxTresMinutesPer() *Dbv0037QosLimitsMaxTresMinutesPer {
	this := Dbv0037QosLimitsMaxTresMinutesPer{}
	return &this
}

// NewDbv0037QosLimitsMaxTresMinutesPerWithDefaults instantiates a new Dbv0037QosLimitsMaxTresMinutesPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037QosLimitsMaxTresMinutesPerWithDefaults() *Dbv0037QosLimitsMaxTresMinutesPer {
	this := Dbv0037QosLimitsMaxTresMinutesPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) GetJob() []map[string]interface{} {
	if o == nil || o.Job == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) GetJobOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given []map[string]interface{} and assigns it to the Job field.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) SetJob(v []map[string]interface{}) {
	o.Job = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) GetAccount() []map[string]interface{} {
	if o == nil || o.Account == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) GetAccountOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []map[string]interface{} and assigns it to the Account field.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) SetAccount(v []map[string]interface{}) {
	o.Account = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) GetUser() []map[string]interface{} {
	if o == nil || o.User == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) GetUserOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given []map[string]interface{} and assigns it to the User field.
func (o *Dbv0037QosLimitsMaxTresMinutesPer) SetUser(v []map[string]interface{}) {
	o.User = &v
}

func (o Dbv0037QosLimitsMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037QosLimitsMaxTresMinutesPer struct {
	value *Dbv0037QosLimitsMaxTresMinutesPer
	isSet bool
}

func (v NullableDbv0037QosLimitsMaxTresMinutesPer) Get() *Dbv0037QosLimitsMaxTresMinutesPer {
	return v.value
}

func (v *NullableDbv0037QosLimitsMaxTresMinutesPer) Set(val *Dbv0037QosLimitsMaxTresMinutesPer) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037QosLimitsMaxTresMinutesPer) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037QosLimitsMaxTresMinutesPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037QosLimitsMaxTresMinutesPer(val *Dbv0037QosLimitsMaxTresMinutesPer) *NullableDbv0037QosLimitsMaxTresMinutesPer {
	return &NullableDbv0037QosLimitsMaxTresMinutesPer{value: val, isSet: true}
}

func (v NullableDbv0037QosLimitsMaxTresMinutesPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037QosLimitsMaxTresMinutesPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


