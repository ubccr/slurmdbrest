# Dbv0037QosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Qos** | Pointer to [**[]Dbv0037Qos**](Dbv0037Qos.md) | Array of QOS | [optional] 

## Methods

### NewDbv0037QosInfo

`func NewDbv0037QosInfo() *Dbv0037QosInfo`

NewDbv0037QosInfo instantiates a new Dbv0037QosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037QosInfoWithDefaults

`func NewDbv0037QosInfoWithDefaults() *Dbv0037QosInfo`

NewDbv0037QosInfoWithDefaults instantiates a new Dbv0037QosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037QosInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037QosInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037QosInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037QosInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0037QosInfo) GetQos() []Dbv0037Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0037QosInfo) GetQosOk() (*[]Dbv0037Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0037QosInfo) SetQos(v []Dbv0037Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0037QosInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


