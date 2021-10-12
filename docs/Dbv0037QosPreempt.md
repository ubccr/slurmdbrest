# Dbv0037QosPreempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **[]string** | List of preemptable QOS | [optional] 
**Mode** | Pointer to **[]string** | List of preemption modes | [optional] 
**ExemptTime** | Pointer to **int32** | Grace period (s) before jobs can preempted | [optional] 

## Methods

### NewDbv0037QosPreempt

`func NewDbv0037QosPreempt() *Dbv0037QosPreempt`

NewDbv0037QosPreempt instantiates a new Dbv0037QosPreempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037QosPreemptWithDefaults

`func NewDbv0037QosPreemptWithDefaults() *Dbv0037QosPreempt`

NewDbv0037QosPreemptWithDefaults instantiates a new Dbv0037QosPreempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *Dbv0037QosPreempt) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *Dbv0037QosPreempt) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *Dbv0037QosPreempt) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *Dbv0037QosPreempt) HasList() bool`

HasList returns a boolean if a field has been set.

### GetMode

`func (o *Dbv0037QosPreempt) GetMode() []string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Dbv0037QosPreempt) GetModeOk() (*[]string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Dbv0037QosPreempt) SetMode(v []string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Dbv0037QosPreempt) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetExemptTime

`func (o *Dbv0037QosPreempt) GetExemptTime() int32`

GetExemptTime returns the ExemptTime field if non-nil, zero value otherwise.

### GetExemptTimeOk

`func (o *Dbv0037QosPreempt) GetExemptTimeOk() (*int32, bool)`

GetExemptTimeOk returns a tuple with the ExemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptTime

`func (o *Dbv0037QosPreempt) SetExemptTime(v int32)`

SetExemptTime sets ExemptTime field to given value.

### HasExemptTime

`func (o *Dbv0037QosPreempt) HasExemptTime() bool`

HasExemptTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


