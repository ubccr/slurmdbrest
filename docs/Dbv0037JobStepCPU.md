# Dbv0037JobStepCPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedFrequency** | Pointer to [**Dbv0037JobStepCPURequestedFrequency**](Dbv0037JobStepCPURequestedFrequency.md) |  | [optional] 
**Governor** | Pointer to **[]string** | CPU governor | [optional] 

## Methods

### NewDbv0037JobStepCPU

`func NewDbv0037JobStepCPU() *Dbv0037JobStepCPU`

NewDbv0037JobStepCPU instantiates a new Dbv0037JobStepCPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobStepCPUWithDefaults

`func NewDbv0037JobStepCPUWithDefaults() *Dbv0037JobStepCPU`

NewDbv0037JobStepCPUWithDefaults instantiates a new Dbv0037JobStepCPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedFrequency

`func (o *Dbv0037JobStepCPU) GetRequestedFrequency() Dbv0037JobStepCPURequestedFrequency`

GetRequestedFrequency returns the RequestedFrequency field if non-nil, zero value otherwise.

### GetRequestedFrequencyOk

`func (o *Dbv0037JobStepCPU) GetRequestedFrequencyOk() (*Dbv0037JobStepCPURequestedFrequency, bool)`

GetRequestedFrequencyOk returns a tuple with the RequestedFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFrequency

`func (o *Dbv0037JobStepCPU) SetRequestedFrequency(v Dbv0037JobStepCPURequestedFrequency)`

SetRequestedFrequency sets RequestedFrequency field to given value.

### HasRequestedFrequency

`func (o *Dbv0037JobStepCPU) HasRequestedFrequency() bool`

HasRequestedFrequency returns a boolean if a field has been set.

### GetGovernor

`func (o *Dbv0037JobStepCPU) GetGovernor() []string`

GetGovernor returns the Governor field if non-nil, zero value otherwise.

### GetGovernorOk

`func (o *Dbv0037JobStepCPU) GetGovernorOk() (*[]string, bool)`

GetGovernorOk returns a tuple with the Governor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernor

`func (o *Dbv0037JobStepCPU) SetGovernor(v []string)`

SetGovernor sets Governor field to given value.

### HasGovernor

`func (o *Dbv0037JobStepCPU) HasGovernor() bool`

HasGovernor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


