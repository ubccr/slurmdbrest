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

// Dbv0037JobTime Time properties
type Dbv0037JobTime struct {
	// Total time elapsed
	Elapsed *int32 `json:"elapsed,omitempty"`
	// Total time eligible to run
	Eligible *int32 `json:"eligible,omitempty"`
	// Timestamp of when job ended
	End *int32 `json:"end,omitempty"`
	// Timestamp of when job started
	Start *int32 `json:"start,omitempty"`
	// Timestamp of when job submitted
	Submission *int32 `json:"submission,omitempty"`
	// Timestamp of when job last suspended
	Suspended *int32 `json:"suspended,omitempty"`
	System *Dbv0037JobTimeSystem `json:"system,omitempty"`
	Total *Dbv0037JobTimeTotal `json:"total,omitempty"`
	User *Dbv0037JobTimeUser `json:"user,omitempty"`
	// Job wall clock time limit
	Limit *int32 `json:"limit,omitempty"`
}

// NewDbv0037JobTime instantiates a new Dbv0037JobTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0037JobTime() *Dbv0037JobTime {
	this := Dbv0037JobTime{}
	return &this
}

// NewDbv0037JobTimeWithDefaults instantiates a new Dbv0037JobTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0037JobTimeWithDefaults() *Dbv0037JobTime {
	this := Dbv0037JobTime{}
	return &this
}

// GetElapsed returns the Elapsed field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetElapsed() int32 {
	if o == nil || o.Elapsed == nil {
		var ret int32
		return ret
	}
	return *o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetElapsedOk() (*int32, bool) {
	if o == nil || o.Elapsed == nil {
		return nil, false
	}
	return o.Elapsed, true
}

// HasElapsed returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasElapsed() bool {
	if o != nil && o.Elapsed != nil {
		return true
	}

	return false
}

// SetElapsed gets a reference to the given int32 and assigns it to the Elapsed field.
func (o *Dbv0037JobTime) SetElapsed(v int32) {
	o.Elapsed = &v
}

// GetEligible returns the Eligible field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetEligible() int32 {
	if o == nil || o.Eligible == nil {
		var ret int32
		return ret
	}
	return *o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetEligibleOk() (*int32, bool) {
	if o == nil || o.Eligible == nil {
		return nil, false
	}
	return o.Eligible, true
}

// HasEligible returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasEligible() bool {
	if o != nil && o.Eligible != nil {
		return true
	}

	return false
}

// SetEligible gets a reference to the given int32 and assigns it to the Eligible field.
func (o *Dbv0037JobTime) SetEligible(v int32) {
	o.Eligible = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *Dbv0037JobTime) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *Dbv0037JobTime) SetStart(v int32) {
	o.Start = &v
}

// GetSubmission returns the Submission field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetSubmission() int32 {
	if o == nil || o.Submission == nil {
		var ret int32
		return ret
	}
	return *o.Submission
}

// GetSubmissionOk returns a tuple with the Submission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetSubmissionOk() (*int32, bool) {
	if o == nil || o.Submission == nil {
		return nil, false
	}
	return o.Submission, true
}

// HasSubmission returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasSubmission() bool {
	if o != nil && o.Submission != nil {
		return true
	}

	return false
}

// SetSubmission gets a reference to the given int32 and assigns it to the Submission field.
func (o *Dbv0037JobTime) SetSubmission(v int32) {
	o.Submission = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetSuspended() int32 {
	if o == nil || o.Suspended == nil {
		var ret int32
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetSuspendedOk() (*int32, bool) {
	if o == nil || o.Suspended == nil {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasSuspended() bool {
	if o != nil && o.Suspended != nil {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given int32 and assigns it to the Suspended field.
func (o *Dbv0037JobTime) SetSuspended(v int32) {
	o.Suspended = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetSystem() Dbv0037JobTimeSystem {
	if o == nil || o.System == nil {
		var ret Dbv0037JobTimeSystem
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetSystemOk() (*Dbv0037JobTimeSystem, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given Dbv0037JobTimeSystem and assigns it to the System field.
func (o *Dbv0037JobTime) SetSystem(v Dbv0037JobTimeSystem) {
	o.System = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetTotal() Dbv0037JobTimeTotal {
	if o == nil || o.Total == nil {
		var ret Dbv0037JobTimeTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetTotalOk() (*Dbv0037JobTimeTotal, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given Dbv0037JobTimeTotal and assigns it to the Total field.
func (o *Dbv0037JobTime) SetTotal(v Dbv0037JobTimeTotal) {
	o.Total = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetUser() Dbv0037JobTimeUser {
	if o == nil || o.User == nil {
		var ret Dbv0037JobTimeUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetUserOk() (*Dbv0037JobTimeUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given Dbv0037JobTimeUser and assigns it to the User field.
func (o *Dbv0037JobTime) SetUser(v Dbv0037JobTimeUser) {
	o.User = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Dbv0037JobTime) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0037JobTime) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Dbv0037JobTime) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *Dbv0037JobTime) SetLimit(v int32) {
	o.Limit = &v
}

func (o Dbv0037JobTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Elapsed != nil {
		toSerialize["elapsed"] = o.Elapsed
	}
	if o.Eligible != nil {
		toSerialize["eligible"] = o.Eligible
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Submission != nil {
		toSerialize["submission"] = o.Submission
	}
	if o.Suspended != nil {
		toSerialize["suspended"] = o.Suspended
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableDbv0037JobTime struct {
	value *Dbv0037JobTime
	isSet bool
}

func (v NullableDbv0037JobTime) Get() *Dbv0037JobTime {
	return v.value
}

func (v *NullableDbv0037JobTime) Set(val *Dbv0037JobTime) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0037JobTime) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0037JobTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0037JobTime(val *Dbv0037JobTime) *NullableDbv0037JobTime {
	return &NullableDbv0037JobTime{value: val, isSet: true}
}

func (v NullableDbv0037JobTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0037JobTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


