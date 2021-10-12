# Dbv0037ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**Dbv0037ClusterInfoController**](Dbv0037ClusterInfoController.md) |  | [optional] 
**Flags** | Pointer to **[]string** | List of properties of cluster | [optional] 
**Name** | Pointer to **string** | Cluster name | [optional] 
**Nodes** | Pointer to **string** | Assigned nodes | [optional] 
**SelectPlugin** | Pointer to **string** | Configured select plugin | [optional] 
**Associations** | Pointer to [**Dbv0037ClusterInfoAssociations**](Dbv0037ClusterInfoAssociations.md) |  | [optional] 
**RpcVersion** | Pointer to **int32** | Number rpc version | [optional] 
**Tres** | Pointer to [**[]Dbv0037ResponseTres**](Dbv0037ResponseTres.md) | List of TRES in cluster | [optional] 

## Methods

### NewDbv0037ClusterInfo

`func NewDbv0037ClusterInfo() *Dbv0037ClusterInfo`

NewDbv0037ClusterInfo instantiates a new Dbv0037ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037ClusterInfoWithDefaults

`func NewDbv0037ClusterInfoWithDefaults() *Dbv0037ClusterInfo`

NewDbv0037ClusterInfoWithDefaults instantiates a new Dbv0037ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *Dbv0037ClusterInfo) GetController() Dbv0037ClusterInfoController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *Dbv0037ClusterInfo) GetControllerOk() (*Dbv0037ClusterInfoController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *Dbv0037ClusterInfo) SetController(v Dbv0037ClusterInfoController)`

SetController sets Controller field to given value.

### HasController

`func (o *Dbv0037ClusterInfo) HasController() bool`

HasController returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0037ClusterInfo) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0037ClusterInfo) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0037ClusterInfo) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0037ClusterInfo) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *Dbv0037ClusterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0037ClusterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0037ClusterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0037ClusterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0037ClusterInfo) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0037ClusterInfo) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Dbv0037ClusterInfo) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0037ClusterInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSelectPlugin

`func (o *Dbv0037ClusterInfo) GetSelectPlugin() string`

GetSelectPlugin returns the SelectPlugin field if non-nil, zero value otherwise.

### GetSelectPluginOk

`func (o *Dbv0037ClusterInfo) GetSelectPluginOk() (*string, bool)`

GetSelectPluginOk returns a tuple with the SelectPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPlugin

`func (o *Dbv0037ClusterInfo) SetSelectPlugin(v string)`

SetSelectPlugin sets SelectPlugin field to given value.

### HasSelectPlugin

`func (o *Dbv0037ClusterInfo) HasSelectPlugin() bool`

HasSelectPlugin returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0037ClusterInfo) GetAssociations() Dbv0037ClusterInfoAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0037ClusterInfo) GetAssociationsOk() (*Dbv0037ClusterInfoAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0037ClusterInfo) SetAssociations(v Dbv0037ClusterInfoAssociations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0037ClusterInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetRpcVersion

`func (o *Dbv0037ClusterInfo) GetRpcVersion() int32`

GetRpcVersion returns the RpcVersion field if non-nil, zero value otherwise.

### GetRpcVersionOk

`func (o *Dbv0037ClusterInfo) GetRpcVersionOk() (*int32, bool)`

GetRpcVersionOk returns a tuple with the RpcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcVersion

`func (o *Dbv0037ClusterInfo) SetRpcVersion(v int32)`

SetRpcVersion sets RpcVersion field to given value.

### HasRpcVersion

`func (o *Dbv0037ClusterInfo) HasRpcVersion() bool`

HasRpcVersion returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0037ClusterInfo) GetTres() []Dbv0037ResponseTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0037ClusterInfo) GetTresOk() (*[]Dbv0037ResponseTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0037ClusterInfo) SetTres(v []Dbv0037ResponseTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0037ClusterInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


