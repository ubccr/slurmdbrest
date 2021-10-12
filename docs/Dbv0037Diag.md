# Dbv0037Diag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Statistics** | Pointer to [**Dbv0037DiagStatistics**](Dbv0037DiagStatistics.md) |  | [optional] 

## Methods

### NewDbv0037Diag

`func NewDbv0037Diag() *Dbv0037Diag`

NewDbv0037Diag instantiates a new Dbv0037Diag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037DiagWithDefaults

`func NewDbv0037DiagWithDefaults() *Dbv0037Diag`

NewDbv0037DiagWithDefaults instantiates a new Dbv0037Diag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037Diag) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037Diag) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037Diag) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037Diag) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatistics

`func (o *Dbv0037Diag) GetStatistics() Dbv0037DiagStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Dbv0037Diag) GetStatisticsOk() (*Dbv0037DiagStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Dbv0037Diag) SetStatistics(v Dbv0037DiagStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Dbv0037Diag) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


