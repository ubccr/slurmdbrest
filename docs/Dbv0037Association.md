# Dbv0037Association

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Assigned account | [optional] 
**Cluster** | Pointer to **string** | Assigned cluster | [optional] 
**Default** | Pointer to [**Dbv0037AssociationDefault**](Dbv0037AssociationDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | List of properties of association | [optional] 
**Max** | Pointer to [**Dbv0037AssociationMax**](Dbv0037AssociationMax.md) |  | [optional] 
**Min** | Pointer to [**Dbv0037AssociationMin**](Dbv0037AssociationMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** | Parent account name | [optional] 
**Partition** | Pointer to **string** | Assigned partition | [optional] 
**Priority** | Pointer to **int32** | Assigned priority | [optional] 
**Qos** | Pointer to **[]string** | Assigned QOS | [optional] 
**SharesRaw** | Pointer to **int32** | Raw fairshare shares | [optional] 
**Usage** | Pointer to [**Dbv0037AssociationUsage**](Dbv0037AssociationUsage.md) |  | [optional] 
**User** | Pointer to **string** | Assigned user | [optional] 

## Methods

### NewDbv0037Association

`func NewDbv0037Association() *Dbv0037Association`

NewDbv0037Association instantiates a new Dbv0037Association object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037AssociationWithDefaults

`func NewDbv0037AssociationWithDefaults() *Dbv0037Association`

NewDbv0037AssociationWithDefaults instantiates a new Dbv0037Association object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Dbv0037Association) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Dbv0037Association) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Dbv0037Association) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Dbv0037Association) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *Dbv0037Association) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Dbv0037Association) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Dbv0037Association) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Dbv0037Association) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDefault

`func (o *Dbv0037Association) GetDefault() Dbv0037AssociationDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Dbv0037Association) GetDefaultOk() (*Dbv0037AssociationDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Dbv0037Association) SetDefault(v Dbv0037AssociationDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Dbv0037Association) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0037Association) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0037Association) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0037Association) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0037Association) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *Dbv0037Association) GetMax() Dbv0037AssociationMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Dbv0037Association) GetMaxOk() (*Dbv0037AssociationMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Dbv0037Association) SetMax(v Dbv0037AssociationMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *Dbv0037Association) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *Dbv0037Association) GetMin() Dbv0037AssociationMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Dbv0037Association) GetMinOk() (*Dbv0037AssociationMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Dbv0037Association) SetMin(v Dbv0037AssociationMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *Dbv0037Association) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *Dbv0037Association) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *Dbv0037Association) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *Dbv0037Association) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *Dbv0037Association) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *Dbv0037Association) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Dbv0037Association) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Dbv0037Association) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Dbv0037Association) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *Dbv0037Association) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Dbv0037Association) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Dbv0037Association) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Dbv0037Association) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0037Association) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0037Association) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0037Association) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0037Association) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *Dbv0037Association) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *Dbv0037Association) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *Dbv0037Association) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *Dbv0037Association) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUsage

`func (o *Dbv0037Association) GetUsage() Dbv0037AssociationUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Dbv0037Association) GetUsageOk() (*Dbv0037AssociationUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Dbv0037Association) SetUsage(v Dbv0037AssociationUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Dbv0037Association) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUser

`func (o *Dbv0037Association) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Dbv0037Association) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Dbv0037Association) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Dbv0037Association) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


