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

// Dbv0037JobStepStep Step details
type Dbv0037JobStepStep struct {
	// Parent job id
	JobId *int32 `json:"job_id,omitempty"`
	Het *Dbv0037JobStepStepHet `json:"het,omitempty"`
	// Step id
	Id *string `json:"id,omitempty"`
	// Step name
	Name *string `json:"name,omitempty"`
	Task *Dbv0037JobStepStepTask `json:"task,omitempty"`
	Tres *Dbv0037JobStepStepTres `json:"tres,omitempty"`
}

// NewDbv0037JobStepStep instantiates a new Dbv0037JobStepStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobStepStep() *Dbv0037JobStepStep {
	this := Dbv0037JobStepStep{}
	return &this
}

// NewDbv0037JobStepStepWithDefaults instantiates a new Dbv0037JobStepStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobStepStepWithDefaults() *Dbv0037JobStepStep {
	this := Dbv0037JobStepStep{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetJobId() int32 {
	if o == nil || o.JobId == nil {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetJobIdOk() (*int32, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *Dbv0037JobStepStep) SetJobId(v int32) {
	o.JobId = &v
}

// GetHet returns the Het field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetHet() Dbv0037JobStepStepHet {
	if o == nil || o.Het == nil {
		var ret Dbv0037JobStepStepHet
		return ret
	}
	return *o.Het
}

// GetHetOk returns a tuple with the Het field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetHetOk() (*Dbv0037JobStepStepHet, bool) {
	if o == nil || o.Het == nil {
		return nil, false
	}
	return o.Het, true
}

// HasHet returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasHet() bool {
	if o != nil && o.Het != nil {
		return true
	}

	return false
}

// SetHet gets a reference to the given Dbv0037JobStepStepHet and assigns it to the Het field.
func (o *Dbv0037JobStepStep) SetHet(v Dbv0037JobStepStepHet) {
	o.Het = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dbv0037JobStepStep) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dbv0037JobStepStep) SetName(v string) {
	o.Name = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetTask() Dbv0037JobStepStepTask {
	if o == nil || o.Task == nil {
		var ret Dbv0037JobStepStepTask
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetTaskOk() (*Dbv0037JobStepStepTask, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given Dbv0037JobStepStepTask and assigns it to the Task field.
func (o *Dbv0037JobStepStep) SetTask(v Dbv0037JobStepStepTask) {
	o.Task = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0037JobStepStep) GetTres() Dbv0037JobStepStepTres {
	if o == nil || o.Tres == nil {
		var ret Dbv0037JobStepStepTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobStepStep) GetTresOk() (*Dbv0037JobStepStepTres, bool) {
	if o == nil || o.Tres == nil {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0037JobStepStep) HasTres() bool {
	if o != nil && o.Tres != nil {
		return true
	}

	return false
}

// SetTres gets a reference to the given Dbv0037JobStepStepTres and assigns it to the Tres field.
func (o *Dbv0037JobStepStep) SetTres(v Dbv0037JobStepStepTres) {
	o.Tres = &v
}

func (o Dbv0037JobStepStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	if o.Het != nil {
		toSerialize["het"] = o.Het
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	if o.Tres != nil {
		toSerialize["tres"] = o.Tres
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobStepStep struct {
	value *Dbv0037JobStepStep
	isSet bool
}

func (v NullableDbv0037JobStepStep) Get() *Dbv0037JobStepStep {
	return v.value
}

func (v *NullableDbv0037JobStepStep) Set(val *Dbv0037JobStepStep) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobStepStep) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobStepStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobStepStep(val *Dbv0037JobStepStep) *NullableDbv0037JobStepStep {
	return &NullableDbv0037JobStepStep{value: val, isSet: true}
}

func (v NullableDbv0037JobStepStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobStepStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

