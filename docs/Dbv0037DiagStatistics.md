# Dbv0037DiagStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int32** | Unix timestamp of start time | [optional] 
**Rollups** | Pointer to [**[]Dbv0037DiagStatisticsRollups**](Dbv0037DiagStatisticsRollups.md) |  | [optional] 
**RPCs** | Pointer to [**[]Dbv0037DiagStatisticsRPCs**](Dbv0037DiagStatisticsRPCs.md) |  | [optional] 
**Users** | Pointer to [**[]Dbv0037DiagStatisticsUsers**](Dbv0037DiagStatisticsUsers.md) |  | [optional] 

## Methods

### NewDbv0037DiagStatistics

`func NewDbv0037DiagStatistics() *Dbv0037DiagStatistics`

NewDbv0037DiagStatistics instantiates a new Dbv0037DiagStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037DiagStatisticsWithDefaults

`func NewDbv0037DiagStatisticsWithDefaults() *Dbv0037DiagStatistics`

NewDbv0037DiagStatisticsWithDefaults instantiates a new Dbv0037DiagStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *Dbv0037DiagStatistics) GetTimeStart() int32`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *Dbv0037DiagStatistics) GetTimeStartOk() (*int32, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *Dbv0037DiagStatistics) SetTimeStart(v int32)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *Dbv0037DiagStatistics) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *Dbv0037DiagStatistics) GetRollups() []Dbv0037DiagStatisticsRollups`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *Dbv0037DiagStatistics) GetRollupsOk() (*[]Dbv0037DiagStatisticsRollups, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *Dbv0037DiagStatistics) SetRollups(v []Dbv0037DiagStatisticsRollups)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *Dbv0037DiagStatistics) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *Dbv0037DiagStatistics) GetRPCs() []Dbv0037DiagStatisticsRPCs`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *Dbv0037DiagStatistics) GetRPCsOk() (*[]Dbv0037DiagStatisticsRPCs, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *Dbv0037DiagStatistics) SetRPCs(v []Dbv0037DiagStatisticsRPCs)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *Dbv0037DiagStatistics) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *Dbv0037DiagStatistics) GetUsers() []Dbv0037DiagStatisticsUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Dbv0037DiagStatistics) GetUsersOk() (*[]Dbv0037DiagStatisticsUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Dbv0037DiagStatistics) SetUsers(v []Dbv0037DiagStatisticsUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Dbv0037DiagStatistics) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


