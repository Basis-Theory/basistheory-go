# CreateProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DestinationUrl** | **string** |  | 
**RequestReactorId** | Pointer to **NullableString** |  | [optional] 
**ResponseReactorId** | Pointer to **NullableString** |  | [optional] 
**RequireAuth** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateProxyRequest

`func NewCreateProxyRequest(name string, destinationUrl string, ) *CreateProxyRequest`

NewCreateProxyRequest instantiates a new CreateProxyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProxyRequestWithDefaults

`func NewCreateProxyRequestWithDefaults() *CreateProxyRequest`

NewCreateProxyRequestWithDefaults instantiates a new CreateProxyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProxyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProxyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProxyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDestinationUrl

`func (o *CreateProxyRequest) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *CreateProxyRequest) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *CreateProxyRequest) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.


### GetRequestReactorId

`func (o *CreateProxyRequest) GetRequestReactorId() string`

GetRequestReactorId returns the RequestReactorId field if non-nil, zero value otherwise.

### GetRequestReactorIdOk

`func (o *CreateProxyRequest) GetRequestReactorIdOk() (*string, bool)`

GetRequestReactorIdOk returns a tuple with the RequestReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReactorId

`func (o *CreateProxyRequest) SetRequestReactorId(v string)`

SetRequestReactorId sets RequestReactorId field to given value.

### HasRequestReactorId

`func (o *CreateProxyRequest) HasRequestReactorId() bool`

HasRequestReactorId returns a boolean if a field has been set.

### SetRequestReactorIdNil

`func (o *CreateProxyRequest) SetRequestReactorIdNil(b bool)`

 SetRequestReactorIdNil sets the value for RequestReactorId to be an explicit nil

### UnsetRequestReactorId
`func (o *CreateProxyRequest) UnsetRequestReactorId()`

UnsetRequestReactorId ensures that no value is present for RequestReactorId, not even an explicit nil
### GetResponseReactorId

`func (o *CreateProxyRequest) GetResponseReactorId() string`

GetResponseReactorId returns the ResponseReactorId field if non-nil, zero value otherwise.

### GetResponseReactorIdOk

`func (o *CreateProxyRequest) GetResponseReactorIdOk() (*string, bool)`

GetResponseReactorIdOk returns a tuple with the ResponseReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseReactorId

`func (o *CreateProxyRequest) SetResponseReactorId(v string)`

SetResponseReactorId sets ResponseReactorId field to given value.

### HasResponseReactorId

`func (o *CreateProxyRequest) HasResponseReactorId() bool`

HasResponseReactorId returns a boolean if a field has been set.

### SetResponseReactorIdNil

`func (o *CreateProxyRequest) SetResponseReactorIdNil(b bool)`

 SetResponseReactorIdNil sets the value for ResponseReactorId to be an explicit nil

### UnsetResponseReactorId
`func (o *CreateProxyRequest) UnsetResponseReactorId()`

UnsetResponseReactorId ensures that no value is present for ResponseReactorId, not even an explicit nil
### GetRequireAuth

`func (o *CreateProxyRequest) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *CreateProxyRequest) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *CreateProxyRequest) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *CreateProxyRequest) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### SetRequireAuthNil

`func (o *CreateProxyRequest) SetRequireAuthNil(b bool)`

 SetRequireAuthNil sets the value for RequireAuth to be an explicit nil

### UnsetRequireAuth
`func (o *CreateProxyRequest) UnsetRequireAuth()`

UnsetRequireAuth ensures that no value is present for RequireAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


