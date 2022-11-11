# CreateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**CanCreateExpiringApplications** | Pointer to **NullableBool** |  | [optional] 
**ExpiresAt** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Rules** | Pointer to [**[]AccessRule**](AccessRule.md) |  | [optional] 

## Methods

### NewCreateApplicationRequest

`func NewCreateApplicationRequest(type_ string, ) *CreateApplicationRequest`

NewCreateApplicationRequest instantiates a new CreateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestWithDefaults

`func NewCreateApplicationRequestWithDefaults() *CreateApplicationRequest`

NewCreateApplicationRequestWithDefaults instantiates a new CreateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateApplicationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateApplicationRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateApplicationRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *CreateApplicationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApplicationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApplicationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCanCreateExpiringApplications

`func (o *CreateApplicationRequest) GetCanCreateExpiringApplications() bool`

GetCanCreateExpiringApplications returns the CanCreateExpiringApplications field if non-nil, zero value otherwise.

### GetCanCreateExpiringApplicationsOk

`func (o *CreateApplicationRequest) GetCanCreateExpiringApplicationsOk() (*bool, bool)`

GetCanCreateExpiringApplicationsOk returns a tuple with the CanCreateExpiringApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateExpiringApplications

`func (o *CreateApplicationRequest) SetCanCreateExpiringApplications(v bool)`

SetCanCreateExpiringApplications sets CanCreateExpiringApplications field to given value.

### HasCanCreateExpiringApplications

`func (o *CreateApplicationRequest) HasCanCreateExpiringApplications() bool`

HasCanCreateExpiringApplications returns a boolean if a field has been set.

### SetCanCreateExpiringApplicationsNil

`func (o *CreateApplicationRequest) SetCanCreateExpiringApplicationsNil(b bool)`

 SetCanCreateExpiringApplicationsNil sets the value for CanCreateExpiringApplications to be an explicit nil

### UnsetCanCreateExpiringApplications
`func (o *CreateApplicationRequest) UnsetCanCreateExpiringApplications()`

UnsetCanCreateExpiringApplications ensures that no value is present for CanCreateExpiringApplications, not even an explicit nil
### GetExpiresAt

`func (o *CreateApplicationRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateApplicationRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateApplicationRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateApplicationRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateApplicationRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateApplicationRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetPermissions

`func (o *CreateApplicationRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateApplicationRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateApplicationRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateApplicationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *CreateApplicationRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *CreateApplicationRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRules

`func (o *CreateApplicationRequest) GetRules() []AccessRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateApplicationRequest) GetRulesOk() (*[]AccessRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateApplicationRequest) SetRules(v []AccessRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateApplicationRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *CreateApplicationRequest) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *CreateApplicationRequest) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


