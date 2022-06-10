# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ApplicationTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Permission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Permission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Permission) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Permission) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Permission) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Permission) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDescription

`func (o *Permission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Permission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Permission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Permission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Permission) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Permission) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApplicationTypes

`func (o *Permission) GetApplicationTypes() []string`

GetApplicationTypes returns the ApplicationTypes field if non-nil, zero value otherwise.

### GetApplicationTypesOk

`func (o *Permission) GetApplicationTypesOk() (*[]string, bool)`

GetApplicationTypesOk returns a tuple with the ApplicationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTypes

`func (o *Permission) SetApplicationTypes(v []string)`

SetApplicationTypes sets ApplicationTypes field to given value.

### HasApplicationTypes

`func (o *Permission) HasApplicationTypes() bool`

HasApplicationTypes returns a boolean if a field has been set.

### SetApplicationTypesNil

`func (o *Permission) SetApplicationTypesNil(b bool)`

 SetApplicationTypesNil sets the value for ApplicationTypes to be an explicit nil

### UnsetApplicationTypes
`func (o *Permission) UnsetApplicationTypes()`

UnsetApplicationTypes ensures that no value is present for ApplicationTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


