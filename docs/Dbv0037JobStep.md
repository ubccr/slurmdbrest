# Dbv0037JobStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**Dbv0037JobStepTime**](Dbv0037JobStepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**Dbv0037JobExitCode**](Dbv0037JobExitCode.md) |  | [optional] 
**Nodes** | Pointer to [**Dbv0037JobStepNodes**](Dbv0037JobStepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**Dbv0037JobStepTasks**](Dbv0037JobStepTasks.md) |  | [optional] 
**Pid** | Pointer to **string** | First process PID | [optional] 
**CPU** | Pointer to [**Dbv0037JobStepCPU**](Dbv0037JobStepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User who requested job killed | [optional] 
**State** | Pointer to **string** | State of job step | [optional] 
**Statistics** | Pointer to [**Dbv0037JobStepStatistics**](Dbv0037JobStepStatistics.md) |  | [optional] 
**Step** | Pointer to [**Dbv0037JobStepStep**](Dbv0037JobStepStep.md) |  | [optional] 

## Methods

### NewDbv0037JobStep

`func NewDbv0037JobStep() *Dbv0037JobStep`

NewDbv0037JobStep instantiates a new Dbv0037JobStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobStepWithDefaults

`func NewDbv0037JobStepWithDefaults() *Dbv0037JobStep`

NewDbv0037JobStepWithDefaults instantiates a new Dbv0037JobStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *Dbv0037JobStep) GetTime() Dbv0037JobStepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Dbv0037JobStep) GetTimeOk() (*Dbv0037JobStepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Dbv0037JobStep) SetTime(v Dbv0037JobStepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *Dbv0037JobStep) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *Dbv0037JobStep) GetExitCode() Dbv0037JobExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Dbv0037JobStep) GetExitCodeOk() (*Dbv0037JobExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Dbv0037JobStep) SetExitCode(v Dbv0037JobExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Dbv0037JobStep) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *Dbv0037JobStep) GetNodes() Dbv0037JobStepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Dbv0037JobStep) GetNodesOk() (*Dbv0037JobStepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Dbv0037JobStep) SetNodes(v Dbv0037JobStepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Dbv0037JobStep) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *Dbv0037JobStep) GetTasks() Dbv0037JobStepTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Dbv0037JobStep) GetTasksOk() (*Dbv0037JobStepTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Dbv0037JobStep) SetTasks(v Dbv0037JobStepTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Dbv0037JobStep) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *Dbv0037JobStep) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Dbv0037JobStep) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Dbv0037JobStep) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *Dbv0037JobStep) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *Dbv0037JobStep) GetCPU() Dbv0037JobStepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *Dbv0037JobStep) GetCPUOk() (*Dbv0037JobStepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *Dbv0037JobStep) SetCPU(v Dbv0037JobStepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *Dbv0037JobStep) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *Dbv0037JobStep) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *Dbv0037JobStep) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *Dbv0037JobStep) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *Dbv0037JobStep) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *Dbv0037JobStep) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Dbv0037JobStep) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Dbv0037JobStep) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Dbv0037JobStep) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *Dbv0037JobStep) GetStatistics() Dbv0037JobStepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Dbv0037JobStep) GetStatisticsOk() (*Dbv0037JobStepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Dbv0037JobStep) SetStatistics(v Dbv0037JobStepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Dbv0037JobStep) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *Dbv0037JobStep) GetStep() Dbv0037JobStepStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *Dbv0037JobStep) GetStepOk() (*Dbv0037JobStepStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *Dbv0037JobStep) SetStep(v Dbv0037JobStepStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *Dbv0037JobStep) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


