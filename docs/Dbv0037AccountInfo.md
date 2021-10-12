# Dbv0037AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Accounts** | Pointer to [**[]Dbv0037Account**](Dbv0037Account.md) | List of accounts | [optional] 

## Methods

### NewDbv0037AccountInfo

`func NewDbv0037AccountInfo() *Dbv0037AccountInfo`

NewDbv0037AccountInfo instantiates a new Dbv0037AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037AccountInfoWithDefaults

`func NewDbv0037AccountInfoWithDefaults() *Dbv0037AccountInfo`

NewDbv0037AccountInfoWithDefaults instantiates a new Dbv0037AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037AccountInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037AccountInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037AccountInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037AccountInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0037AccountInfo) GetAccounts() []Dbv0037Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0037AccountInfo) GetAccountsOk() (*[]Dbv0037Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0037AccountInfo) SetAccounts(v []Dbv0037Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0037AccountInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


