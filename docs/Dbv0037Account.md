# Dbv0037Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]Dbv0037AssociationShortInfo**](Dbv0037AssociationShortInfo.md) | List of assigned associations | [optional] 
**Coordinators** | Pointer to [**[]Dbv0037CoordinatorInfo**](Dbv0037CoordinatorInfo.md) | List of assigned coordinators | [optional] 
**Description** | Pointer to **string** | Description of account | [optional] 
**Name** | Pointer to **string** | Name of account | [optional] 
**Organization** | Pointer to **string** | Assigned organization of account | [optional] 
**Flags** | Pointer to **[]string** | List of properties of account | [optional] 

## Methods

### NewDbv0037Account

`func NewDbv0037Account() *Dbv0037Account`

NewDbv0037Account instantiates a new Dbv0037Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037AccountWithDefaults

`func NewDbv0037AccountWithDefaults() *Dbv0037Account`

NewDbv0037AccountWithDefaults instantiates a new Dbv0037Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *Dbv0037Account) GetAssociations() []Dbv0037AssociationShortInfo`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0037Account) GetAssociationsOk() (*[]Dbv0037AssociationShortInfo, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0037Account) SetAssociations(v []Dbv0037AssociationShortInfo)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0037Account) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *Dbv0037Account) GetCoordinators() []Dbv0037CoordinatorInfo`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *Dbv0037Account) GetCoordinatorsOk() (*[]Dbv0037CoordinatorInfo, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *Dbv0037Account) SetCoordinators(v []Dbv0037CoordinatorInfo)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *Dbv0037Account) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDescription

`func (o *Dbv0037Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dbv0037Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dbv0037Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dbv0037Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Dbv0037Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0037Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0037Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0037Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *Dbv0037Account) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Dbv0037Account) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Dbv0037Account) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Dbv0037Account) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetFlags

`func (o *Dbv0037Account) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbv0037Account) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbv0037Account) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbv0037Account) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


