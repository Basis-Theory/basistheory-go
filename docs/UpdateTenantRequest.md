# UpdateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Settings** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateTenantRequest

`func NewUpdateTenantRequest(name string, ) *UpdateTenantRequest`

NewUpdateTenantRequest instantiates a new UpdateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantRequestWithDefaults

`func NewUpdateTenantRequestWithDefaults() *UpdateTenantRequest`

NewUpdateTenantRequestWithDefaults instantiates a new UpdateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *UpdateTenantRequest) GetSettings() map[string]string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateTenantRequest) GetSettingsOk() (*map[string]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateTenantRequest) SetSettings(v map[string]string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateTenantRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *UpdateTenantRequest) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *UpdateTenantRequest) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


