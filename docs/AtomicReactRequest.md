# AtomicReactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactorId** | **string** |  | 
**RequestParameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAtomicReactRequest

`func NewAtomicReactRequest(reactorId string, ) *AtomicReactRequest`

NewAtomicReactRequest instantiates a new AtomicReactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtomicReactRequestWithDefaults

`func NewAtomicReactRequestWithDefaults() *AtomicReactRequest`

NewAtomicReactRequestWithDefaults instantiates a new AtomicReactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactorId

`func (o *AtomicReactRequest) GetReactorId() string`

GetReactorId returns the ReactorId field if non-nil, zero value otherwise.

### GetReactorIdOk

`func (o *AtomicReactRequest) GetReactorIdOk() (*string, bool)`

GetReactorIdOk returns a tuple with the ReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactorId

`func (o *AtomicReactRequest) SetReactorId(v string)`

SetReactorId sets ReactorId field to given value.


### GetRequestParameters

`func (o *AtomicReactRequest) GetRequestParameters() map[string]interface{}`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *AtomicReactRequest) GetRequestParametersOk() (*map[string]interface{}, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *AtomicReactRequest) SetRequestParameters(v map[string]interface{})`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *AtomicReactRequest) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### SetRequestParametersNil

`func (o *AtomicReactRequest) SetRequestParametersNil(b bool)`

 SetRequestParametersNil sets the value for RequestParameters to be an explicit nil

### UnsetRequestParameters
`func (o *AtomicReactRequest) UnsetRequestParameters()`

UnsetRequestParameters ensures that no value is present for RequestParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


