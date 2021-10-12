# Dbv0037User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorLevel** | Pointer to **string** | Description of administrator level | [optional] 
**Associations** | Pointer to [**Dbv0037UserAssociations**](Dbv0037UserAssociations.md) |  | [optional] 
**Coordinators** | Pointer to [**[]Dbv0037CoordinatorInfo**](Dbv0037CoordinatorInfo.md) | List of assigned coordinators | [optional] 
**Default** | Pointer to [**Dbv0037UserDefault**](Dbv0037UserDefault.md) |  | [optional] 
**Name** | Pointer to **string** | User name | [optional] 

## Methods

### NewDbv0037User

`func NewDbv0037User() *Dbv0037User`

NewDbv0037User instantiates a new Dbv0037User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037UserWithDefaults

`func NewDbv0037UserWithDefaults() *Dbv0037User`

NewDbv0037UserWithDefaults instantiates a new Dbv0037User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorLevel

`func (o *Dbv0037User) GetAdministratorLevel() string`

GetAdministratorLevel returns the AdministratorLevel field if non-nil, zero value otherwise.

### GetAdministratorLevelOk

`func (o *Dbv0037User) GetAdministratorLevelOk() (*string, bool)`

GetAdministratorLevelOk returns a tuple with the AdministratorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorLevel

`func (o *Dbv0037User) SetAdministratorLevel(v string)`

SetAdministratorLevel sets AdministratorLevel field to given value.

### HasAdministratorLevel

`func (o *Dbv0037User) HasAdministratorLevel() bool`

HasAdministratorLevel returns a boolean if a field has been set.

### GetAssociations

`func (o *Dbv0037User) GetAssociations() Dbv0037UserAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Dbv0037User) GetAssociationsOk() (*Dbv0037UserAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Dbv0037User) SetAssociations(v Dbv0037UserAssociations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Dbv0037User) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *Dbv0037User) GetCoordinators() []Dbv0037CoordinatorInfo`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *Dbv0037User) GetCoordinatorsOk() (*[]Dbv0037CoordinatorInfo, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *Dbv0037User) SetCoordinators(v []Dbv0037CoordinatorInfo)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *Dbv0037User) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDefault

`func (o *Dbv0037User) GetDefault() Dbv0037UserDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Dbv0037User) GetDefaultOk() (*Dbv0037UserDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Dbv0037User) SetDefault(v Dbv0037UserDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Dbv0037User) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *Dbv0037User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0037User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0037User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0037User) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


