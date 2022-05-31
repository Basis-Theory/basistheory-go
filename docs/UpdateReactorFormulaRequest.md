# UpdateReactorFormulaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Configuration** | Pointer to [**[]ReactorFormulaConfiguration**](ReactorFormulaConfiguration.md) |  | [optional] 
**RequestParameters** | Pointer to [**[]ReactorFormulaRequestParameter**](ReactorFormulaRequestParameter.md) |  | [optional] 

## Methods

### NewUpdateReactorFormulaRequest

`func NewUpdateReactorFormulaRequest(type_ string, name string, ) *UpdateReactorFormulaRequest`

NewUpdateReactorFormulaRequest instantiates a new UpdateReactorFormulaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReactorFormulaRequestWithDefaults

`func NewUpdateReactorFormulaRequestWithDefaults() *UpdateReactorFormulaRequest`

NewUpdateReactorFormulaRequestWithDefaults instantiates a new UpdateReactorFormulaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateReactorFormulaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateReactorFormulaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateReactorFormulaRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *UpdateReactorFormulaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateReactorFormulaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateReactorFormulaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateReactorFormulaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateReactorFormulaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateReactorFormulaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateReactorFormulaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateReactorFormulaRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateReactorFormulaRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *UpdateReactorFormulaRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UpdateReactorFormulaRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UpdateReactorFormulaRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UpdateReactorFormulaRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *UpdateReactorFormulaRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *UpdateReactorFormulaRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetCode

`func (o *UpdateReactorFormulaRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateReactorFormulaRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateReactorFormulaRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateReactorFormulaRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateReactorFormulaRequest) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateReactorFormulaRequest) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetConfiguration

`func (o *UpdateReactorFormulaRequest) GetConfiguration() []ReactorFormulaConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateReactorFormulaRequest) GetConfigurationOk() (*[]ReactorFormulaConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateReactorFormulaRequest) SetConfiguration(v []ReactorFormulaConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateReactorFormulaRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *UpdateReactorFormulaRequest) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *UpdateReactorFormulaRequest) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetRequestParameters

`func (o *UpdateReactorFormulaRequest) GetRequestParameters() []ReactorFormulaRequestParameter`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *UpdateReactorFormulaRequest) GetRequestParametersOk() (*[]ReactorFormulaRequestParameter, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *UpdateReactorFormulaRequest) SetRequestParameters(v []ReactorFormulaRequestParameter)`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *UpdateReactorFormulaRequest) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### SetRequestParametersNil

`func (o *UpdateReactorFormulaRequest) SetRequestParametersNil(b bool)`

 SetRequestParametersNil sets the value for RequestParameters to be an explicit nil

### UnsetRequestParameters
`func (o *UpdateReactorFormulaRequest) UnsetRequestParameters()`

UnsetRequestParameters ensures that no value is present for RequestParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


