# Dbv0037Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errno** | Pointer to **int32** | Error number | [optional] 
**Error** | Pointer to **string** | Error message | [optional] 

## Methods

### NewDbv0037Error

`func NewDbv0037Error() *Dbv0037Error`

NewDbv0037Error instantiates a new Dbv0037Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037ErrorWithDefaults

`func NewDbv0037ErrorWithDefaults() *Dbv0037Error`

NewDbv0037ErrorWithDefaults instantiates a new Dbv0037Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrno

`func (o *Dbv0037Error) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *Dbv0037Error) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *Dbv0037Error) SetErrno(v int32)`

SetErrno sets Errno field to given value.

### HasErrno

`func (o *Dbv0037Error) HasErrno() bool`

HasErrno returns a boolean if a field has been set.

### GetError

`func (o *Dbv0037Error) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Dbv0037Error) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Dbv0037Error) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Dbv0037Error) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


