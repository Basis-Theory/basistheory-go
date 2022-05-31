# CreateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateApplicationRequest

`func NewCreateApplicationRequest(name string, type_ string, ) *CreateApplicationRequest`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


