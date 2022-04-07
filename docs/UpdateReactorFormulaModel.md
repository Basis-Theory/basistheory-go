# UpdateReactorFormulaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Configuration** | Pointer to [**[]ReactorFormulaConfigurationModel**](ReactorFormulaConfigurationModel.md) |  | [optional] 
**RequestParameters** | Pointer to [**[]ReactorFormulaRequestParameterModel**](ReactorFormulaRequestParameterModel.md) |  | [optional] 

## Methods

### NewUpdateReactorFormulaModel

`func NewUpdateReactorFormulaModel() *UpdateReactorFormulaModel`

NewUpdateReactorFormulaModel instantiates a new UpdateReactorFormulaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReactorFormulaModelWithDefaults

`func NewUpdateReactorFormulaModelWithDefaults() *UpdateReactorFormulaModel`

NewUpdateReactorFormulaModelWithDefaults instantiates a new UpdateReactorFormulaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateReactorFormulaModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateReactorFormulaModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateReactorFormulaModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateReactorFormulaModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UpdateReactorFormulaModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateReactorFormulaModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *UpdateReactorFormulaModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateReactorFormulaModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateReactorFormulaModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateReactorFormulaModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateReactorFormulaModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateReactorFormulaModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateReactorFormulaModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateReactorFormulaModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateReactorFormulaModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateReactorFormulaModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateReactorFormulaModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateReactorFormulaModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *UpdateReactorFormulaModel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UpdateReactorFormulaModel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UpdateReactorFormulaModel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UpdateReactorFormulaModel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *UpdateReactorFormulaModel) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *UpdateReactorFormulaModel) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetCode

`func (o *UpdateReactorFormulaModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateReactorFormulaModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateReactorFormulaModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateReactorFormulaModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateReactorFormulaModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateReactorFormulaModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetConfiguration

`func (o *UpdateReactorFormulaModel) GetConfiguration() []ReactorFormulaConfigurationModel`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UpdateReactorFormulaModel) GetConfigurationOk() (*[]ReactorFormulaConfigurationModel, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UpdateReactorFormulaModel) SetConfiguration(v []ReactorFormulaConfigurationModel)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UpdateReactorFormulaModel) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *UpdateReactorFormulaModel) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *UpdateReactorFormulaModel) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetRequestParameters

`func (o *UpdateReactorFormulaModel) GetRequestParameters() []ReactorFormulaRequestParameterModel`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *UpdateReactorFormulaModel) GetRequestParametersOk() (*[]ReactorFormulaRequestParameterModel, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *UpdateReactorFormulaModel) SetRequestParameters(v []ReactorFormulaRequestParameterModel)`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *UpdateReactorFormulaModel) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### SetRequestParametersNil

`func (o *UpdateReactorFormulaModel) SetRequestParametersNil(b bool)`

 SetRequestParametersNil sets the value for RequestParameters to be an explicit nil

### UnsetRequestParameters
`func (o *UpdateReactorFormulaModel) UnsetRequestParameters()`

UnsetRequestParameters ensures that no value is present for RequestParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


