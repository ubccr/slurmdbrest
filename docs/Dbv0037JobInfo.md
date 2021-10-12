# Dbv0037JobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Jobs** | Pointer to [**[]Dbv0037Job**](Dbv0037Job.md) | Array of jobs | [optional] 

## Methods

### NewDbv0037JobInfo

`func NewDbv0037JobInfo() *Dbv0037JobInfo`

NewDbv0037JobInfo instantiates a new Dbv0037JobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037JobInfoWithDefaults

`func NewDbv0037JobInfoWithDefaults() *Dbv0037JobInfo`

NewDbv0037JobInfoWithDefaults instantiates a new Dbv0037JobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037JobInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037JobInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037JobInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037JobInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJobs

`func (o *Dbv0037JobInfo) GetJobs() []Dbv0037Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *Dbv0037JobInfo) GetJobsOk() (*[]Dbv0037Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *Dbv0037JobInfo) SetJobs(v []Dbv0037Job)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *Dbv0037JobInfo) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


