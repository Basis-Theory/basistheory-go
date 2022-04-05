# CreateReactorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Formula** | Pointer to [**ReactorFormulaModel**](ReactorFormulaModel.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateReactorModel

`func NewCreateReactorModel() *CreateReactorModel`

NewCreateReactorModel instantiates a new CreateReactorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReactorModelWithDefaults

`func NewCreateReactorModelWithDefaults() *CreateReactorModel`

NewCreateReactorModelWithDefaults instantiates a new CreateReactorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateReactorModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReactorModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReactorModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateReactorModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateReactorModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateReactorModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFormula

`func (o *CreateReactorModel) GetFormula() ReactorFormulaModel`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *CreateReactorModel) GetFormulaOk() (*ReactorFormulaModel, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *CreateReactorModel) SetFormula(v ReactorFormulaModel)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *CreateReactorModel) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateReactorModel) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateReactorModel) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateReactorModel) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateReactorModel) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *CreateReactorModel) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *CreateReactorModel) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


