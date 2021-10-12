# Dbv0037WckeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Wckeys** | Pointer to [**[]Dbv0037Wckey**](Dbv0037Wckey.md) | List of wckeys | [optional] 

## Methods

### NewDbv0037WckeyInfo

`func NewDbv0037WckeyInfo() *Dbv0037WckeyInfo`

NewDbv0037WckeyInfo instantiates a new Dbv0037WckeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037WckeyInfoWithDefaults

`func NewDbv0037WckeyInfoWithDefaults() *Dbv0037WckeyInfo`

NewDbv0037WckeyInfoWithDefaults instantiates a new Dbv0037WckeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037WckeyInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037WckeyInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037WckeyInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037WckeyInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0037WckeyInfo) GetWckeys() []Dbv0037Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0037WckeyInfo) GetWckeysOk() (*[]Dbv0037Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0037WckeyInfo) SetWckeys(v []Dbv0037Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0037WckeyInfo) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


