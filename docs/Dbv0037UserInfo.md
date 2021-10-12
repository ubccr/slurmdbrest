# Dbv0037UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Users** | Pointer to [**[]Dbv0037User**](Dbv0037User.md) | Array of users | [optional] 

## Methods

### NewDbv0037UserInfo

`func NewDbv0037UserInfo() *Dbv0037UserInfo`

NewDbv0037UserInfo instantiates a new Dbv0037UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037UserInfoWithDefaults

`func NewDbv0037UserInfoWithDefaults() *Dbv0037UserInfo`

NewDbv0037UserInfoWithDefaults instantiates a new Dbv0037UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037UserInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037UserInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037UserInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037UserInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0037UserInfo) GetUsers() []Dbv0037User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0037UserInfo) GetUsersOk() (*[]Dbv0037User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0037UserInfo) SetUsers(v []Dbv0037User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0037UserInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


