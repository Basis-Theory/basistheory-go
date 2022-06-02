# CreateProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DestinationUrl** | **string** |  | 
**RequestReactorId** | **string** |  | 

## Methods

### NewCreateProxyRequest

`func NewCreateProxyRequest(name string, destinationUrl string, requestReactorId string, ) *CreateProxyRequest`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


