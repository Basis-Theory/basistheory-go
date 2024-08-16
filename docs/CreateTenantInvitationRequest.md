# CreateTenantInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateTenantInvitationRequest

`func NewCreateTenantInvitationRequest(email string, ) *CreateTenantInvitationRequest`

NewCreateTenantInvitationRequest instantiates a new CreateTenantInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantInvitationRequestWithDefaults

`func NewCreateTenantInvitationRequestWithDefaults() *CreateTenantInvitationRequest`

NewCreateTenantInvitationRequestWithDefaults instantiates a new CreateTenantInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateTenantInvitationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateTenantInvitationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateTenantInvitationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *CreateTenantInvitationRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateTenantInvitationRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateTenantInvitationRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateTenantInvitationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *CreateTenantInvitationRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *CreateTenantInvitationRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


