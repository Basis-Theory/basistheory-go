# UpdateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CanCreateExpiringApplications** | Pointer to **NullableBool** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Rules** | Pointer to [**[]AccessRule**](AccessRule.md) |  | [optional] 

## Methods

### NewUpdateApplicationRequest

`func NewUpdateApplicationRequest(name string, ) *UpdateApplicationRequest`

NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationRequestWithDefaults

`func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest`

NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCanCreateExpiringApplications

`func (o *UpdateApplicationRequest) GetCanCreateExpiringApplications() bool`

GetCanCreateExpiringApplications returns the CanCreateExpiringApplications field if non-nil, zero value otherwise.

### GetCanCreateExpiringApplicationsOk

`func (o *UpdateApplicationRequest) GetCanCreateExpiringApplicationsOk() (*bool, bool)`

GetCanCreateExpiringApplicationsOk returns a tuple with the CanCreateExpiringApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateExpiringApplications

`func (o *UpdateApplicationRequest) SetCanCreateExpiringApplications(v bool)`

SetCanCreateExpiringApplications sets CanCreateExpiringApplications field to given value.

### HasCanCreateExpiringApplications

`func (o *UpdateApplicationRequest) HasCanCreateExpiringApplications() bool`

HasCanCreateExpiringApplications returns a boolean if a field has been set.

### SetCanCreateExpiringApplicationsNil

`func (o *UpdateApplicationRequest) SetCanCreateExpiringApplicationsNil(b bool)`

 SetCanCreateExpiringApplicationsNil sets the value for CanCreateExpiringApplications to be an explicit nil

### UnsetCanCreateExpiringApplications
`func (o *UpdateApplicationRequest) UnsetCanCreateExpiringApplications()`

UnsetCanCreateExpiringApplications ensures that no value is present for CanCreateExpiringApplications, not even an explicit nil
### GetPermissions

`func (o *UpdateApplicationRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateApplicationRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateApplicationRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateApplicationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *UpdateApplicationRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *UpdateApplicationRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRules

`func (o *UpdateApplicationRequest) GetRules() []AccessRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateApplicationRequest) GetRulesOk() (*[]AccessRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateApplicationRequest) SetRules(v []AccessRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateApplicationRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *UpdateApplicationRequest) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *UpdateApplicationRequest) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


