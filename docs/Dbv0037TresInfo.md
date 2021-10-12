# Dbv0037TresInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Tres** | Pointer to **[][]map[string]interface{}** | Array of tres | [optional] 

## Methods

### NewDbv0037TresInfo

`func NewDbv0037TresInfo() *Dbv0037TresInfo`

NewDbv0037TresInfo instantiates a new Dbv0037TresInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037TresInfoWithDefaults

`func NewDbv0037TresInfoWithDefaults() *Dbv0037TresInfo`

NewDbv0037TresInfoWithDefaults instantiates a new Dbv0037TresInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037TresInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037TresInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037TresInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037TresInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTres

`func (o *Dbv0037TresInfo) GetTres() [][]map[string]interface{}`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *Dbv0037TresInfo) GetTresOk() (*[][]map[string]interface{}, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *Dbv0037TresInfo) SetTres(v [][]map[string]interface{})`

SetTres sets Tres field to given value.

### HasTres

`func (o *Dbv0037TresInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


