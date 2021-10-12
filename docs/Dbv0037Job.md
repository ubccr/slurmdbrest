# Dbv0037Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account charged by job | [optional] 
**Comment** | Pointer to [**Dbv0037JobComment**](Dbv0037JobComment.md) |  | [optional] 
**AllocationNodes** | Pointer to **string** | Nodes allocated to job | [optional] 
**Array** | Pointer to [**Dbv0037JobArray**](Dbv0037JobArray.md) |  | [optional] 
**Time** | Pointer to [**Dbv0037JobTime**](Dbv0037JobTime.md) |  | [optional] 
**Association** | Pointer to [**Dbv0037AssociationShortInfo**](Dbv0037AssociationShortInfo.md) |  | [optional] 
**Cluster** | Pointer to **string** | Assigned cluster | [optional] 
**Constraints** | Pointer to **string** | Constraints on job | [optional] 
**DerivedExitCode** | Pointer to [**Dbv0037JobExitCode**](Dbv0037JobExitCode.md) |  | [optional] 
**ExitCode** | Pointer to [**Dbv0037JobExitCode**](Dbv0037JobExitCode.md) |  | [optional] 
**Flags** | Pointer to **[]string** | List of properties of job | [optional] 
**Group** | Pointer to **string** | User&#39;s group to run job | [optional] 
**Het** | Pointer to [**Dbv0037JobHet**](Dbv0037JobHet.md) |  | [optional] 
**JobId** | Pointer to **int32** | Job id | [optional] 
**Name** | Pointer to **string** | Assigned job name | [optional] 
**Mcs** | Pointer to [**Dbv0037JobMcs**](Dbv0037JobMcs.md) |  | [optional] 
**Nodes** | Pointer to **string** | List of nodes allocated for job | [optional] 
**Partition** | Pointer to **string** | Assigned job&#39;s partition | [optional] 
**Priority** | Pointer to **int32** | Priority | [optional] 
**Qos** | Pointer to **string** | Assigned qos name | [optional] 
**Required** | Pointer to [**Dbv0037JobRequired**](Dbv0037JobRequired.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User who requested job killed | [optional] 
**Reservation** | Pointer to [**Dbv0037JobReservation**](Dbv0037JobReservation.md) |  | [optional] 
**State** | Pointer to [**Dbv0037JobState**](Dbv0037JobState.md) |  | [optional] 
**Steps** | Pointer to [**[]Dbv0037JobStep**](Dbv0037JobStep.md) | Job step description | [optional] 
**Tres** | Pointer to [**Dbv0037JobTres**](Dbv0037JobTres.md) |  | [optional] 
**User** | Pointer to **string** | Job user | [optional] 
**Wckey** | Pointer to [**Dbv0037JobWckey**](Dbv0037JobWckey.md) |  | [optional] 
**WorkingDirectory** | Pointer to **string** | Directory where job was initially started | [optional] 

## Methods

### NewDbv0037Job

`func NewDbv0037Job() *Dbv0037Job`

NewDbv0037Job instantiates a new Dbv0037Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobWithDefaults

`func NewDbv0037JobWithDefaults() *Dbv0037Job`

NewDbv0037JobWithDefaults instantiates a new Dbv0037Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Dbv0037Job) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Dbv0037Job) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Dbv0037Job) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Dbv0037Job) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetComment

`func (o *Dbv0037Job) GetComment() Dbv0037JobComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Dbv0037Job) GetCommentOk() (*Dbv0037JobComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Dbv0037Job) SetComment(v Dbv0037JobComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Dbv0037Job) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAllocationNodes

`func (o *Dbv0037Job) GetAllocationNodes() string`

GetAllocationNodes returns the AllocationNodes field if non-nil, zero value otherwise.

### GetAllocationNodesOk

`func (o *Dbv0037Job) GetAllocationNodesOk() (*string, bool)`

GetAllocationNodesOk returns a tuple with the AllocationNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodes

`func (o *Dbv0037Job) SetAllocationNodes(v string)`

SetAllocationNodes sets AllocationNodes field to given value.

### HasAllocationNodes

`func (o *Dbv0037Job) HasAllocationNodes() bool`

HasAllocationNodes returns a boolean if a field has been set.

### GetArray

`func (o *Dbv0037Job) GetArray() Dbv0037JobArray`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *Dbv0037Job) GetArrayOk() (*Dbv0037JobArray, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *Dbv0037Job) SetArray(v Dbv0037JobArray)`

SetArray sets Array field to given value.

### HasArray

`func (o *Dbv0037Job) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetTime

`func (o *Dbv0037Job) GetTime() Dbv0037JobTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Dbv0037Job) GetTimeOk() (*Dbv0037JobTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Dbv0037Job) SetTime(v Dbv0037JobTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *Dbv0037Job) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAssociation

`func (o *Dbv0037Job) GetAssociation() Dbv0037AssociationShortInfo`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *Dbv0037Job) GetAssociationOk() (*Dbv0037AssociationShortInfo, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *Dbv0037Job) SetAssociation(v Dbv0037AssociationShortInfo)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *Dbv0037Job) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCluster

`func (o *Dbv0037Job) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Dbv0037Job) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Dbv0037Job) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Dbv0037Job) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConstraints

`func (o *Dbv0037Job) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Dbv0037Job) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Dbv0037Job) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Dbv0037Job) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *Dbv0037Job) GetDerivedExitCode() Dbv0037JobExitCode`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *Dbv0037Job) GetDerivedExitCodeOk() (*Dbv0037JobExitCode, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *Dbv0037Job) SetDerivedExitCode(v Dbv0037JobExitCode)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *Dbv0037Job) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetExitCode

`func (o *Dbv0037Job) GetExitCode() Dbv0037JobExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Dbv0037Job) GetExitCodeOk() (*Dbv0037JobExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Dbv0037Job) SetExitCode(v Dbv0037JobExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Dbv0037Job) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0037Job) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0037Job) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0037Job) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0037Job) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroup

`func (o *Dbv0037Job) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Dbv0037Job) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Dbv0037Job) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Dbv0037Job) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHet

`func (o *Dbv0037Job) GetHet() Dbv0037JobHet`

GetHet returns the Het field if non-nil, zero value otherwise.

### GetHetOk

`func (o *Dbv0037Job) GetHetOk() (*Dbv0037JobHet, bool)`

GetHetOk returns a tuple with the Het field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHet

`func (o *Dbv0037Job) SetHet(v Dbv0037JobHet)`

SetHet sets Het field to given value.

### HasHet

`func (o *Dbv0037Job) HasHet() bool`

HasHet returns a boolean if a field has been set.

### GetJobId

`func (o *Dbv0037Job) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Dbv0037Job) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Dbv0037Job) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *Dbv0037Job) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetName

`func (o *Dbv0037Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0037Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0037Job) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0037Job) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMcs

`func (o *Dbv0037Job) GetMcs() Dbv0037JobMcs`

GetMcs returns the Mcs field if non-nil, zero value otherwise.

### GetMcsOk

`func (o *Dbv0037Job) GetMcsOk() (*Dbv0037JobMcs, bool)`

GetMcsOk returns a tuple with the Mcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcs

`func (o *Dbv0037Job) SetMcs(v Dbv0037JobMcs)`

SetMcs sets Mcs field to given value.

### HasMcs

`func (o *Dbv0037Job) HasMcs() bool`

HasMcs returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0037Job) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0037Job) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Dbv0037Job) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0037Job) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPartition

`func (o *Dbv0037Job) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Dbv0037Job) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Dbv0037Job) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Dbv0037Job) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *Dbv0037Job) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Dbv0037Job) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Dbv0037Job) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Dbv0037Job) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *Dbv0037Job) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *Dbv0037Job) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *Dbv0037Job) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *Dbv0037Job) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRequired

`func (o *Dbv0037Job) GetRequired() Dbv0037JobRequired`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Dbv0037Job) GetRequiredOk() (*Dbv0037JobRequired, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Dbv0037Job) SetRequired(v Dbv0037JobRequired)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Dbv0037Job) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *Dbv0037Job) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *Dbv0037Job) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *Dbv0037Job) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *Dbv0037Job) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetReservation

`func (o *Dbv0037Job) GetReservation() Dbv0037JobReservation`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Dbv0037Job) GetReservationOk() (*Dbv0037JobReservation, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Dbv0037Job) SetReservation(v Dbv0037JobReservation)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Dbv0037Job) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetState

`func (o *Dbv0037Job) GetState() Dbv0037JobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dbv0037Job) GetStateOk() (*Dbv0037JobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dbv0037Job) SetState(v Dbv0037JobState)`

SetState sets State field to given value.

### HasState

`func (o *Dbv0037Job) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *Dbv0037Job) GetSteps() []Dbv0037JobStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Dbv0037Job) GetStepsOk() (*[]Dbv0037JobStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Dbv0037Job) SetSteps(v []Dbv0037JobStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Dbv0037Job) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0037Job) GetTres() Dbv0037JobTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0037Job) GetTresOk() (*Dbv0037JobTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0037Job) SetTres(v Dbv0037JobTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0037Job) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUser

`func (o *Dbv0037Job) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Dbv0037Job) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Dbv0037Job) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Dbv0037Job) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWckey

`func (o *Dbv0037Job) GetWckey() Dbv0037JobWckey`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *Dbv0037Job) GetWckeyOk() (*Dbv0037JobWckey, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *Dbv0037Job) SetWckey(v Dbv0037JobWckey)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *Dbv0037Job) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *Dbv0037Job) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *Dbv0037Job) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *Dbv0037Job) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *Dbv0037Job) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


