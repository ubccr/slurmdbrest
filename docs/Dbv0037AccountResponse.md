# Dbv0037AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 

## Methods

### NewDbv0037AccountResponse

`func NewDbv0037AccountResponse() *Dbv0037AccountResponse`

NewDbv0037AccountResponse instantiates a new Dbv0037AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037AccountResponseWithDefaults

`func NewDbv0037AccountResponseWithDefaults() *Dbv0037AccountResponse`

NewDbv0037AccountResponseWithDefaults instantiates a new Dbv0037AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037AccountResponse) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037AccountResponse) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037AccountResponse) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037AccountResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


