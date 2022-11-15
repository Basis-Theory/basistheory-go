# AccessRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**Transform** | Pointer to **NullableString** |  | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessRule

`func NewAccessRule() *AccessRule`

NewAccessRule instantiates a new AccessRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleWithDefaults

`func NewAccessRuleWithDefaults() *AccessRule`

NewAccessRuleWithDefaults instantiates a new AccessRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AccessRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *AccessRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AccessRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AccessRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AccessRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *AccessRule) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *AccessRule) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetContainer

`func (o *AccessRule) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AccessRule) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AccessRule) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *AccessRule) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *AccessRule) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *AccessRule) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetTransform

`func (o *AccessRule) GetTransform() string`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *AccessRule) GetTransformOk() (*string, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *AccessRule) SetTransform(v string)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *AccessRule) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### SetTransformNil

`func (o *AccessRule) SetTransformNil(b bool)`

 SetTransformNil sets the value for Transform to be an explicit nil

### UnsetTransform
`func (o *AccessRule) UnsetTransform()`

UnsetTransform ensures that no value is present for Transform, not even an explicit nil
### GetConditions

`func (o *AccessRule) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AccessRule) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AccessRule) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AccessRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *AccessRule) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *AccessRule) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetPermissions

`func (o *AccessRule) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccessRule) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccessRule) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccessRule) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AccessRule) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AccessRule) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


