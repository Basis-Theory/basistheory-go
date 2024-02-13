# CreateReactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Code** | **string** |  | 
**Application** | Pointer to [**Application**](Application.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateReactorRequest

`func NewCreateReactorRequest(name string, code string, ) *CreateReactorRequest`

NewCreateReactorRequest instantiates a new CreateReactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReactorRequestWithDefaults

`func NewCreateReactorRequestWithDefaults() *CreateReactorRequest`

NewCreateReactorRequestWithDefaults instantiates a new CreateReactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateReactorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReactorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReactorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CreateReactorRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateReactorRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateReactorRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetApplication

`func (o *CreateReactorRequest) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CreateReactorRequest) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CreateReactorRequest) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *CreateReactorRequest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateReactorRequest) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateReactorRequest) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateReactorRequest) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateReactorRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *CreateReactorRequest) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *CreateReactorRequest) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


