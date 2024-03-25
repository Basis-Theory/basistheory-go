# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Keys** | Pointer to [**[]ApplicationKey**](ApplicationKey.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Rules** | Pointer to [**[]AccessRule**](AccessRule.md) |  | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *Application) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Application) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Application) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Application) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Application) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Application) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Application) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *Application) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Application) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Application) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Application) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *Application) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *Application) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeys

`func (o *Application) GetKeys() []ApplicationKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *Application) GetKeysOk() (*[]ApplicationKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *Application) SetKeys(v []ApplicationKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *Application) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *Application) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *Application) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil
### GetType

`func (o *Application) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Application) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Application) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Application) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Application) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Application) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCreatedBy

`func (o *Application) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Application) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Application) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Application) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Application) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Application) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *Application) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Application) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Application) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Application) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Application) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Application) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetModifiedBy

`func (o *Application) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Application) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Application) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Application) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *Application) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *Application) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetModifiedAt

`func (o *Application) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Application) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Application) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Application) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *Application) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Application) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetExpiresAt

`func (o *Application) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Application) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Application) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Application) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *Application) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Application) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetPermissions

`func (o *Application) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Application) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Application) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Application) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *Application) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *Application) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRules

`func (o *Application) GetRules() []AccessRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Application) GetRulesOk() (*[]AccessRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Application) SetRules(v []AccessRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Application) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *Application) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *Application) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


