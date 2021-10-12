# Dbv0037JobTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Total time elapsed | [optional] 
**Eligible** | Pointer to **int32** | Total time eligible to run | [optional] 
**End** | Pointer to **int32** | Timestamp of when job ended | [optional] 
**Start** | Pointer to **int32** | Timestamp of when job started | [optional] 
**Submission** | Pointer to **int32** | Timestamp of when job submitted | [optional] 
**Suspended** | Pointer to **int32** | Timestamp of when job last suspended | [optional] 
**System** | Pointer to [**Dbv0037JobTimeSystem**](Dbv0037JobTimeSystem.md) |  | [optional] 
**Total** | Pointer to [**Dbv0037JobTimeTotal**](Dbv0037JobTimeTotal.md) |  | [optional] 
**User** | Pointer to [**Dbv0037JobTimeUser**](Dbv0037JobTimeUser.md) |  | [optional] 
**Limit** | Pointer to **int32** | Job wall clock time limit | [optional] 

## Methods

### NewDbv0037JobTime

`func NewDbv0037JobTime() *Dbv0037JobTime`

NewDbv0037JobTime instantiates a new Dbv0037JobTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobTimeWithDefaults

`func NewDbv0037JobTimeWithDefaults() *Dbv0037JobTime`

NewDbv0037JobTimeWithDefaults instantiates a new Dbv0037JobTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *Dbv0037JobTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *Dbv0037JobTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *Dbv0037JobTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *Dbv0037JobTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEligible

`func (o *Dbv0037JobTime) GetEligible() int32`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *Dbv0037JobTime) GetEligibleOk() (*int32, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *Dbv0037JobTime) SetEligible(v int32)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *Dbv0037JobTime) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEnd

`func (o *Dbv0037JobTime) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Dbv0037JobTime) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Dbv0037JobTime) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Dbv0037JobTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *Dbv0037JobTime) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Dbv0037JobTime) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Dbv0037JobTime) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Dbv0037JobTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubmission

`func (o *Dbv0037JobTime) GetSubmission() int32`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *Dbv0037JobTime) GetSubmissionOk() (*int32, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *Dbv0037JobTime) SetSubmission(v int32)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *Dbv0037JobTime) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetSuspended

`func (o *Dbv0037JobTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Dbv0037JobTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Dbv0037JobTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Dbv0037JobTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *Dbv0037JobTime) GetSystem() Dbv0037JobTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Dbv0037JobTime) GetSystemOk() (*Dbv0037JobTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Dbv0037JobTime) SetSystem(v Dbv0037JobTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Dbv0037JobTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTotal

`func (o *Dbv0037JobTime) GetTotal() Dbv0037JobTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Dbv0037JobTime) GetTotalOk() (*Dbv0037JobTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Dbv0037JobTime) SetTotal(v Dbv0037JobTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Dbv0037JobTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *Dbv0037JobTime) GetUser() Dbv0037JobTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Dbv0037JobTime) GetUserOk() (*Dbv0037JobTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Dbv0037JobTime) SetUser(v Dbv0037JobTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Dbv0037JobTime) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLimit

`func (o *Dbv0037JobTime) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Dbv0037JobTime) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Dbv0037JobTime) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Dbv0037JobTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


