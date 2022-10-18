# ApplicationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ApplicationType** | Pointer to **NullableString** |  | [optional] 
**TemplateType** | Pointer to **NullableString** |  | [optional] 
**IsStarter** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]AccessRule**](AccessRule.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApplicationTemplate

`func NewApplicationTemplate() *ApplicationTemplate`

NewApplicationTemplate instantiates a new ApplicationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTemplateWithDefaults

`func NewApplicationTemplateWithDefaults() *ApplicationTemplate`

NewApplicationTemplateWithDefaults instantiates a new ApplicationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationTemplate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationTemplate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ApplicationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApplicationType

`func (o *ApplicationTemplate) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ApplicationTemplate) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ApplicationTemplate) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *ApplicationTemplate) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *ApplicationTemplate) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *ApplicationTemplate) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetTemplateType

`func (o *ApplicationTemplate) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *ApplicationTemplate) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *ApplicationTemplate) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *ApplicationTemplate) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### SetTemplateTypeNil

`func (o *ApplicationTemplate) SetTemplateTypeNil(b bool)`

 SetTemplateTypeNil sets the value for TemplateType to be an explicit nil

### UnsetTemplateType
`func (o *ApplicationTemplate) UnsetTemplateType()`

UnsetTemplateType ensures that no value is present for TemplateType, not even an explicit nil
### GetIsStarter

`func (o *ApplicationTemplate) GetIsStarter() bool`

GetIsStarter returns the IsStarter field if non-nil, zero value otherwise.

### GetIsStarterOk

`func (o *ApplicationTemplate) GetIsStarterOk() (*bool, bool)`

GetIsStarterOk returns a tuple with the IsStarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarter

`func (o *ApplicationTemplate) SetIsStarter(v bool)`

SetIsStarter sets IsStarter field to given value.

### HasIsStarter

`func (o *ApplicationTemplate) HasIsStarter() bool`

HasIsStarter returns a boolean if a field has been set.

### GetRules

`func (o *ApplicationTemplate) GetRules() []AccessRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ApplicationTemplate) GetRulesOk() (*[]AccessRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ApplicationTemplate) SetRules(v []AccessRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ApplicationTemplate) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *ApplicationTemplate) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *ApplicationTemplate) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetPermissions

`func (o *ApplicationTemplate) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApplicationTemplate) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApplicationTemplate) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApplicationTemplate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ApplicationTemplate) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ApplicationTemplate) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


