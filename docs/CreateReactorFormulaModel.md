# CreateReactorFormulaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Configuration** | Pointer to [**[]ReactorFormulaConfigurationModel**](ReactorFormulaConfigurationModel.md) |  | [optional] 
**RequestParameters** | Pointer to [**[]ReactorFormulaRequestParameterModel**](ReactorFormulaRequestParameterModel.md) |  | [optional] 

## Methods

### NewCreateReactorFormulaModel

`func NewCreateReactorFormulaModel() *CreateReactorFormulaModel`

NewCreateReactorFormulaModel instantiates a new CreateReactorFormulaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReactorFormulaModelWithDefaults

`func NewCreateReactorFormulaModelWithDefaults() *CreateReactorFormulaModel`

NewCreateReactorFormulaModelWithDefaults instantiates a new CreateReactorFormulaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReactorFormulaModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReactorFormulaModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReactorFormulaModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateReactorFormulaModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateReactorFormulaModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateReactorFormulaModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *CreateReactorFormulaModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateReactorFormulaModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateReactorFormulaModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateReactorFormulaModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateReactorFormulaModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateReactorFormulaModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *CreateReactorFormulaModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReactorFormulaModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReactorFormulaModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateReactorFormulaModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateReactorFormulaModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateReactorFormulaModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateReactorFormulaModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReactorFormulaModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReactorFormulaModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateReactorFormulaModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateReactorFormulaModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateReactorFormulaModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *CreateReactorFormulaModel) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateReactorFormulaModel) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateReactorFormulaModel) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateReactorFormulaModel) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateReactorFormulaModel) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateReactorFormulaModel) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetCode

`func (o *CreateReactorFormulaModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateReactorFormulaModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateReactorFormulaModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateReactorFormulaModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CreateReactorFormulaModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CreateReactorFormulaModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetConfiguration

`func (o *CreateReactorFormulaModel) GetConfiguration() []ReactorFormulaConfigurationModel`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateReactorFormulaModel) GetConfigurationOk() (*[]ReactorFormulaConfigurationModel, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateReactorFormulaModel) SetConfiguration(v []ReactorFormulaConfigurationModel)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateReactorFormulaModel) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *CreateReactorFormulaModel) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *CreateReactorFormulaModel) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetRequestParameters

`func (o *CreateReactorFormulaModel) GetRequestParameters() []ReactorFormulaRequestParameterModel`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *CreateReactorFormulaModel) GetRequestParametersOk() (*[]ReactorFormulaRequestParameterModel, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *CreateReactorFormulaModel) SetRequestParameters(v []ReactorFormulaRequestParameterModel)`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *CreateReactorFormulaModel) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### SetRequestParametersNil

`func (o *CreateReactorFormulaModel) SetRequestParametersNil(b bool)`

 SetRequestParametersNil sets the value for RequestParameters to be an explicit nil

### UnsetRequestParameters
`func (o *CreateReactorFormulaModel) UnsetRequestParameters()`

UnsetRequestParameters ensures that no value is present for RequestParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


