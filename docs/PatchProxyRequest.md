# PatchProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DestinationUrl** | Pointer to **NullableString** |  | [optional] 
**RequestTransform** | Pointer to [**ProxyTransform**](ProxyTransform.md) |  | [optional] 
**ResponseTransform** | Pointer to [**ProxyTransform**](ProxyTransform.md) |  | [optional] 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 
**RequireAuth** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPatchProxyRequest

`func NewPatchProxyRequest() *PatchProxyRequest`

NewPatchProxyRequest instantiates a new PatchProxyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchProxyRequestWithDefaults

`func NewPatchProxyRequestWithDefaults() *PatchProxyRequest`

NewPatchProxyRequestWithDefaults instantiates a new PatchProxyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchProxyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchProxyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchProxyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchProxyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchProxyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchProxyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDestinationUrl

`func (o *PatchProxyRequest) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *PatchProxyRequest) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *PatchProxyRequest) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *PatchProxyRequest) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrlNil

`func (o *PatchProxyRequest) SetDestinationUrlNil(b bool)`

 SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil

### UnsetDestinationUrl
`func (o *PatchProxyRequest) UnsetDestinationUrl()`

UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
### GetRequestTransform

`func (o *PatchProxyRequest) GetRequestTransform() ProxyTransform`

GetRequestTransform returns the RequestTransform field if non-nil, zero value otherwise.

### GetRequestTransformOk

`func (o *PatchProxyRequest) GetRequestTransformOk() (*ProxyTransform, bool)`

GetRequestTransformOk returns a tuple with the RequestTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTransform

`func (o *PatchProxyRequest) SetRequestTransform(v ProxyTransform)`

SetRequestTransform sets RequestTransform field to given value.

### HasRequestTransform

`func (o *PatchProxyRequest) HasRequestTransform() bool`

HasRequestTransform returns a boolean if a field has been set.

### GetResponseTransform

`func (o *PatchProxyRequest) GetResponseTransform() ProxyTransform`

GetResponseTransform returns the ResponseTransform field if non-nil, zero value otherwise.

### GetResponseTransformOk

`func (o *PatchProxyRequest) GetResponseTransformOk() (*ProxyTransform, bool)`

GetResponseTransformOk returns a tuple with the ResponseTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTransform

`func (o *PatchProxyRequest) SetResponseTransform(v ProxyTransform)`

SetResponseTransform sets ResponseTransform field to given value.

### HasResponseTransform

`func (o *PatchProxyRequest) HasResponseTransform() bool`

HasResponseTransform returns a boolean if a field has been set.

### GetApplication

`func (o *PatchProxyRequest) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PatchProxyRequest) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PatchProxyRequest) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PatchProxyRequest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetConfiguration

`func (o *PatchProxyRequest) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PatchProxyRequest) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PatchProxyRequest) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PatchProxyRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *PatchProxyRequest) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *PatchProxyRequest) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetRequireAuth

`func (o *PatchProxyRequest) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *PatchProxyRequest) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *PatchProxyRequest) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *PatchProxyRequest) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### SetRequireAuthNil

`func (o *PatchProxyRequest) SetRequireAuthNil(b bool)`

 SetRequireAuthNil sets the value for RequireAuth to be an explicit nil

### UnsetRequireAuth
`func (o *PatchProxyRequest) UnsetRequireAuth()`

UnsetRequireAuth ensures that no value is present for RequireAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


