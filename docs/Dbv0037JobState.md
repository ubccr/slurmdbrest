# Dbv0037JobState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **string** | Current state of job | [optional] 
**Reason** | Pointer to **string** | Last reason job didn&#39;t run | [optional] 

## Methods

### NewDbv0037JobState

`func NewDbv0037JobState() *Dbv0037JobState`

NewDbv0037JobState instantiates a new Dbv0037JobState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobStateWithDefaults

`func NewDbv0037JobStateWithDefaults() *Dbv0037JobState`

NewDbv0037JobStateWithDefaults instantiates a new Dbv0037JobState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *Dbv0037JobState) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Dbv0037JobState) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Dbv0037JobState) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *Dbv0037JobState) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetReason

`func (o *Dbv0037JobState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Dbv0037JobState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Dbv0037JobState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Dbv0037JobState) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


