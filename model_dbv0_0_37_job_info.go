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

// Dbv0037JobInfo struct for Dbv0037JobInfo
type Dbv0037JobInfo struct {
	// Slurm errors
	Errors *[]Dbv0037Error `json:"errors,omitempty"`
	// Array of jobs
	Jobs *[]Dbv0037Job `json:"jobs,omitempty"`
}

// NewDbv0037JobInfo instantiates a new Dbv0037JobInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobInfo() *Dbv0037JobInfo {
	this := Dbv0037JobInfo{}
	return &this
}

// NewDbv0037JobInfoWithDefaults instantiates a new Dbv0037JobInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobInfoWithDefaults() *Dbv0037JobInfo {
	this := Dbv0037JobInfo{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0037JobInfo) GetErrors() []Dbv0037Error {
	if o == nil || o.Errors == nil {
		var ret []Dbv0037Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobInfo) GetErrorsOk() (*[]Dbv0037Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0037JobInfo) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0037Error and assigns it to the Errors field.
func (o *Dbv0037JobInfo) SetErrors(v []Dbv0037Error) {
	o.Errors = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *Dbv0037JobInfo) GetJobs() []Dbv0037Job {
	if o == nil || o.Jobs == nil {
		var ret []Dbv0037Job
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobInfo) GetJobsOk() (*[]Dbv0037Job, bool) {
	if o == nil || o.Jobs == nil {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *Dbv0037JobInfo) HasJobs() bool {
	if o != nil && o.Jobs != nil {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []Dbv0037Job and assigns it to the Jobs field.
func (o *Dbv0037JobInfo) SetJobs(v []Dbv0037Job) {
	o.Jobs = &v
}

func (o Dbv0037JobInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Jobs != nil {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobInfo struct {
	value *Dbv0037JobInfo
	isSet bool
}

func (v NullableDbv0037JobInfo) Get() *Dbv0037JobInfo {
	return v.value
}

func (v *NullableDbv0037JobInfo) Set(val *Dbv0037JobInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobInfo(val *Dbv0037JobInfo) *NullableDbv0037JobInfo {
	return &NullableDbv0037JobInfo{value: val, isSet: true}
}

func (v NullableDbv0037JobInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


