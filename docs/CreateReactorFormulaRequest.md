# CreateReactorFormulaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Configuration** | Pointer to [**[]ReactorFormulaConfiguration**](ReactorFormulaConfiguration.md) |  | [optional] 
**RequestParameters** | Pointer to [**[]ReactorFormulaRequestParameter**](ReactorFormulaRequestParameter.md) |  | [optional] 

## Methods

### NewCreateReactorFormulaRequest

`func NewCreateReactorFormulaRequest(type_ string, name string, ) *CreateReactorFormulaRequest`

NewCreateReactorFormulaRequest instantiates a new CreateReactorFormulaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReactorFormulaRequestWithDefaults

`func NewCreateReactorFormulaRequestWithDefaults() *CreateReactorFormulaRequest`

NewCreateReactorFormulaRequestWithDefaults instantiates a new CreateReactorFormulaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReactorFormulaRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReactorFormulaRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReactorFormulaRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateReactorFormulaRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateReactorFormulaRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateReactorFormulaRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *CreateReactorFormulaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateReactorFormulaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateReactorFormulaRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CreateReactorFormulaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReactorFormulaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReactorFormulaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateReactorFormulaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReactorFormulaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReactorFormulaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateReactorFormulaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateReactorFormulaRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateReactorFormulaRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *CreateReactorFormulaRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateReactorFormulaRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateReactorFormulaRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateReactorFormulaRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateReactorFormulaRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateReactorFormulaRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetCode

`func (o *CreateReactorFormulaRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateReactorFormulaRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateReactorFormulaRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateReactorFormulaRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CreateReactorFormulaRequest) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CreateReactorFormulaRequest) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetConfiguration

`func (o *CreateReactorFormulaRequest) GetConfiguration() []ReactorFormulaConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateReactorFormulaRequest) GetConfigurationOk() (*[]ReactorFormulaConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateReactorFormulaRequest) SetConfiguration(v []ReactorFormulaConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateReactorFormulaRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *CreateReactorFormulaRequest) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *CreateReactorFormulaRequest) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetRequestParameters

`func (o *CreateReactorFormulaRequest) GetRequestParameters() []ReactorFormulaRequestParameter`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *CreateReactorFormulaRequest) GetRequestParametersOk() (*[]ReactorFormulaRequestParameter, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *CreateReactorFormulaRequest) SetRequestParameters(v []ReactorFormulaRequestParameter)`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *CreateReactorFormulaRequest) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### SetRequestParametersNil

`func (o *CreateReactorFormulaRequest) SetRequestParametersNil(b bool)`

 SetRequestParametersNil sets the value for RequestParameters to be an explicit nil

### UnsetRequestParameters
`func (o *CreateReactorFormulaRequest) UnsetRequestParameters()`

UnsetRequestParameters ensures that no value is present for RequestParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


