# UpdateApplicationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateApplicationModel

`func NewUpdateApplicationModel() *UpdateApplicationModel`

NewUpdateApplicationModel instantiates a new UpdateApplicationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationModelWithDefaults

`func NewUpdateApplicationModelWithDefaults() *UpdateApplicationModel`

NewUpdateApplicationModelWithDefaults instantiates a new UpdateApplicationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApplicationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplicationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApplicationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateApplicationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateApplicationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateApplicationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *UpdateApplicationModel) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateApplicationModel) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateApplicationModel) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateApplicationModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *UpdateApplicationModel) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *UpdateApplicationModel) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


