# Dbv0037ConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Tres** | Pointer to **[][]map[string]interface{}** | Array of TRES | [optional] 
**Accounts** | Pointer to [**[]Dbv0037Account**](Dbv0037Account.md) | Array of accounts | [optional] 
**Associations** | Pointer to [**[]Dbv0037Association**](Dbv0037Association.md) | Array of associations | [optional] 
**Users** | Pointer to [**[]Dbv0037User**](Dbv0037User.md) | Array of users | [optional] 
**Qos** | Pointer to [**[]Dbv0037Qos**](Dbv0037Qos.md) | Array of qos | [optional] 
**Wckeys** | Pointer to [**[]Dbv0037Wckey**](Dbv0037Wckey.md) | Array of wckeys | [optional] 

## Methods

### NewDbv0037ConfigInfo

`func NewDbv0037ConfigInfo() *Dbv0037ConfigInfo`

NewDbv0037ConfigInfo instantiates a new Dbv0037ConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037ConfigInfoWithDefaults

`func NewDbv0037ConfigInfoWithDefaults() *Dbv0037ConfigInfo`

NewDbv0037ConfigInfoWithDefaults instantiates a new Dbv0037ConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037ConfigInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037ConfigInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037ConfigInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037ConfigInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0037ConfigInfo) GetTres() [][]map[string]interface{}`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0037ConfigInfo) GetTresOk() (*[][]map[string]interface{}, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0037ConfigInfo) SetTres(v [][]map[string]interface{})`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0037ConfigInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *Dbv0037ConfigInfo) GetAccounts() []Dbv0037Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Dbv0037ConfigInfo) GetAccountsOk() (*[]Dbv0037Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Dbv0037ConfigInfo) SetAccounts(v []Dbv0037Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Dbv0037ConfigInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0037ConfigInfo) GetAssociations() []Dbv0037Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0037ConfigInfo) GetAssociationsOk() (*[]Dbv0037Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0037ConfigInfo) SetAssociations(v []Dbv0037Association)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0037ConfigInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0037ConfigInfo) GetUsers() []Dbv0037User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0037ConfigInfo) GetUsersOk() (*[]Dbv0037User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0037ConfigInfo) SetUsers(v []Dbv0037User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0037ConfigInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0037ConfigInfo) GetQos() []Dbv0037Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0037ConfigInfo) GetQosOk() (*[]Dbv0037Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0037ConfigInfo) SetQos(v []Dbv0037Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0037ConfigInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *Dbv0037ConfigInfo) GetWckeys() []Dbv0037Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *Dbv0037ConfigInfo) GetWckeysOk() (*[]Dbv0037Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *Dbv0037ConfigInfo) SetWckeys(v []Dbv0037Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *Dbv0037ConfigInfo) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


