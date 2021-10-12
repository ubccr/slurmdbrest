# Dbv0037CoordinatorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of user | [optional] 
**Direct** | Pointer to **int32** | If user is coordinator of this account directly or coordinator status was inheirted from a higher account in the tree | [optional] 

## Methods

### NewDbv0037CoordinatorInfo

`func NewDbv0037CoordinatorInfo() *Dbv0037CoordinatorInfo`

NewDbv0037CoordinatorInfo instantiates a new Dbv0037CoordinatorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbv0037CoordinatorInfoWithDefaults

`func NewDbv0037CoordinatorInfoWithDefaults() *Dbv0037CoordinatorInfo`

NewDbv0037CoordinatorInfoWithDefaults instantiates a new Dbv0037CoordinatorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Dbv0037CoordinatorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dbv0037CoordinatorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dbv0037CoordinatorInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dbv0037CoordinatorInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirect

`func (o *Dbv0037CoordinatorInfo) GetDirect() int32`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *Dbv0037CoordinatorInfo) GetDirectOk() (*int32, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *Dbv0037CoordinatorInfo) SetDirect(v int32)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *Dbv0037CoordinatorInfo) HasDirect() bool`

HasDirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


