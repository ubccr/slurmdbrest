# Dbv0037AssociationsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Dbv0037Error**](Dbv0037Error.md) | Slurm errors | [optional] 
**Associations** | Pointer to [**[]Dbv0037Association**](Dbv0037Association.md) | Array of associations | [optional] 

## Methods

### NewDbv0037AssociationsInfo

`func NewDbv0037AssociationsInfo() *Dbv0037AssociationsInfo`

NewDbv0037AssociationsInfo instantiates a new Dbv0037AssociationsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037AssociationsInfoWithDefaults

`func NewDbv0037AssociationsInfoWithDefaults() *Dbv0037AssociationsInfo`

NewDbv0037AssociationsInfoWithDefaults instantiates a new Dbv0037AssociationsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Dbv0037AssociationsInfo) GetErrors() []Dbv0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Dbv0037AssociationsInfo) GetErrorsOk() (*[]Dbv0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Dbv0037AssociationsInfo) SetErrors(v []Dbv0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Dbv0037AssociationsInfo) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0037AssociationsInfo) GetAssociations() []Dbv0037Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0037AssociationsInfo) GetAssociationsOk() (*[]Dbv0037Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0037AssociationsInfo) SetAssociations(v []Dbv0037Association)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0037AssociationsInfo) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


