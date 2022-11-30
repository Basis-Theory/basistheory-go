# UpdateProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DestinationUrl** | **string** |  | 
**RequestReactorId** | Pointer to **NullableString** |  | [optional] 
**ResponseReactorId** | Pointer to **NullableString** |  | [optional] 
**RequestTransform** | Pointer to [**ProxyTransform**](ProxyTransform.md) |  | [optional] 
**ResponseTransform** | Pointer to [**ProxyTransform**](ProxyTransform.md) |  | [optional] 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 
**RequireAuth** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateProxyRequest

`func NewUpdateProxyRequest(name string, destinationUrl string, ) *UpdateProxyRequest`

NewUpdateProxyRequest instantiates a new UpdateProxyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProxyRequestWithDefaults

`func NewUpdateProxyRequestWithDefaults() *UpdateProxyRequest`

NewUpdateProxyRequestWithDefaults instantiates a new UpdateProxyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProxyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProxyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProxyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDestinationUrl

`func (o *UpdateProxyRequest) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *UpdateProxyRequest) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *UpdateProxyRequest) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.


### GetRequestReactorId

`func (o *UpdateProxyRequest) GetRequestReactorId() string`

GetRequestReactorId returns the RequestReactorId field if non-nil, zero value otherwise.

### GetRequestReactorIdOk

`func (o *UpdateProxyRequest) GetRequestReactorIdOk() (*string, bool)`

GetRequestReactorIdOk returns a tuple with the RequestReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReactorId

`func (o *UpdateProxyRequest) SetRequestReactorId(v string)`

SetRequestReactorId sets RequestReactorId field to given value.

### HasRequestReactorId

`func (o *UpdateProxyRequest) HasRequestReactorId() bool`

HasRequestReactorId returns a boolean if a field has been set.

### SetRequestReactorIdNil

`func (o *UpdateProxyRequest) SetRequestReactorIdNil(b bool)`

 SetRequestReactorIdNil sets the value for RequestReactorId to be an explicit nil

### UnsetRequestReactorId
`func (o *UpdateProxyRequest) UnsetRequestReactorId()`

UnsetRequestReactorId ensures that no value is present for RequestReactorId, not even an explicit nil
### GetResponseReactorId

`func (o *UpdateProxyRequest) GetResponseReactorId() string`

GetResponseReactorId returns the ResponseReactorId field if non-nil, zero value otherwise.

### GetResponseReactorIdOk

`func (o *UpdateProxyRequest) GetResponseReactorIdOk() (*string, bool)`

GetResponseReactorIdOk returns a tuple with the ResponseReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseReactorId

`func (o *UpdateProxyRequest) SetResponseReactorId(v string)`

SetResponseReactorId sets ResponseReactorId field to given value.

### HasResponseReactorId

`func (o *UpdateProxyRequest) HasResponseReactorId() bool`

HasResponseReactorId returns a boolean if a field has been set.

### SetResponseReactorIdNil

`func (o *UpdateProxyRequest) SetResponseReactorIdNil(b bool)`

 SetResponseReactorIdNil sets the value for ResponseReactorId to be an explicit nil

### UnsetResponseReactorId
`func (o *UpdateProxyRequest) UnsetResponseReactorId()`

UnsetResponseReactorId ensures that no value is present for ResponseReactorId, not even an explicit nil
### GetRequestTransform

`func (o *UpdateProxyRequest) GetRequestTransform() ProxyTransform`

GetRequestTransform returns the RequestTransform field if non-nil, zero value otherwise.

### GetRequestTransformOk

`func (o *UpdateProxyRequest) GetRequestTransformOk() (*ProxyTransform, bool)`

GetRequestTransformOk returns a tuple with the RequestTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTransform

`func (o *UpdateProxyRequest) SetRequestTransform(v ProxyTransform)`

SetRequestTransform sets RequestTransform field to given value.

### HasRequestTransform

`func (o *UpdateProxyRequest) HasRequestTransform() bool`

HasRequestTransform returns a boolean if a field has been set.

### GetResponseTransform

`func (o *UpdateProxyRequest) GetResponseTransform() ProxyTransform`

GetResponseTransform returns the ResponseTransform field if non-nil, zero value otherwise.

### GetResponseTransformOk

`func (o *UpdateProxyRequest) GetResponseTransformOk() (*ProxyTransform, bool)`

GetResponseTransformOk returns a tuple with the ResponseTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTransform

`func (o *UpdateProxyRequest) SetResponseTransform(v ProxyTransform)`

SetResponseTransform sets ResponseTransform field to given value.

### HasResponseTransform

`func (o *UpdateProxyRequest) HasResponseTransform() bool`

HasResponseTransform returns a boolean if a field has been set.

### GetApplication

`func (o *UpdateProxyRequest) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *UpdateProxyRequest) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *UpdateProxyRequest) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *UpdateProxyRequest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetConfiguration

`func (o *UpdateProxyRequest) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateProxyRequest) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateProxyRequest) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateProxyRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *UpdateProxyRequest) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *UpdateProxyRequest) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetRequireAuth

`func (o *UpdateProxyRequest) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *UpdateProxyRequest) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *UpdateProxyRequest) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *UpdateProxyRequest) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### SetRequireAuthNil

`func (o *UpdateProxyRequest) SetRequireAuthNil(b bool)`

 SetRequireAuthNil sets the value for RequireAuth to be an explicit nil

### UnsetRequireAuth
`func (o *UpdateProxyRequest) UnsetRequireAuth()`

UnsetRequireAuth ensures that no value is present for RequireAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


