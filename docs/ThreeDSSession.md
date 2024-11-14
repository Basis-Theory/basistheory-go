# ThreeDSSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**PanTokenId** | Pointer to **NullableString** |  | [optional] 
**TokenId** | Pointer to **NullableString** |  | [optional] 
**TokenIntentId** | Pointer to **NullableString** |  | [optional] 
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] 
**Device** | Pointer to **NullableString** |  | [optional] 
**DeviceInfo** | Pointer to [**ThreeDSDeviceInfo**](ThreeDSDeviceInfo.md) |  | [optional] 
**Version** | Pointer to [**ThreeDSVersion**](ThreeDSVersion.md) |  | [optional] 
**Method** | Pointer to [**ThreeDSMethod**](ThreeDSMethod.md) |  | [optional] 
**Authentication** | Pointer to [**ThreeDSAuthentication**](ThreeDSAuthentication.md) |  | [optional] 

## Methods

### NewThreeDSSession

`func NewThreeDSSession() *ThreeDSSession`

NewThreeDSSession instantiates a new ThreeDSSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSSessionWithDefaults

`func NewThreeDSSessionWithDefaults() *ThreeDSSession`

NewThreeDSSessionWithDefaults instantiates a new ThreeDSSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreeDSSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreeDSSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreeDSSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreeDSSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ThreeDSSession) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreeDSSession) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreeDSSession) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreeDSSession) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ThreeDSSession) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ThreeDSSession) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTenantId

`func (o *ThreeDSSession) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ThreeDSSession) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ThreeDSSession) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ThreeDSSession) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPanTokenId

`func (o *ThreeDSSession) GetPanTokenId() string`

GetPanTokenId returns the PanTokenId field if non-nil, zero value otherwise.

### GetPanTokenIdOk

`func (o *ThreeDSSession) GetPanTokenIdOk() (*string, bool)`

GetPanTokenIdOk returns a tuple with the PanTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanTokenId

`func (o *ThreeDSSession) SetPanTokenId(v string)`

SetPanTokenId sets PanTokenId field to given value.

### HasPanTokenId

`func (o *ThreeDSSession) HasPanTokenId() bool`

HasPanTokenId returns a boolean if a field has been set.

### SetPanTokenIdNil

`func (o *ThreeDSSession) SetPanTokenIdNil(b bool)`

 SetPanTokenIdNil sets the value for PanTokenId to be an explicit nil

### UnsetPanTokenId
`func (o *ThreeDSSession) UnsetPanTokenId()`

UnsetPanTokenId ensures that no value is present for PanTokenId, not even an explicit nil
### GetTokenId

`func (o *ThreeDSSession) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ThreeDSSession) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ThreeDSSession) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ThreeDSSession) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *ThreeDSSession) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *ThreeDSSession) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
### GetTokenIntentId

`func (o *ThreeDSSession) GetTokenIntentId() string`

GetTokenIntentId returns the TokenIntentId field if non-nil, zero value otherwise.

### GetTokenIntentIdOk

`func (o *ThreeDSSession) GetTokenIntentIdOk() (*string, bool)`

GetTokenIntentIdOk returns a tuple with the TokenIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntentId

`func (o *ThreeDSSession) SetTokenIntentId(v string)`

SetTokenIntentId sets TokenIntentId field to given value.

### HasTokenIntentId

`func (o *ThreeDSSession) HasTokenIntentId() bool`

HasTokenIntentId returns a boolean if a field has been set.

### SetTokenIntentIdNil

`func (o *ThreeDSSession) SetTokenIntentIdNil(b bool)`

 SetTokenIntentIdNil sets the value for TokenIntentId to be an explicit nil

### UnsetTokenIntentId
`func (o *ThreeDSSession) UnsetTokenIntentId()`

UnsetTokenIntentId ensures that no value is present for TokenIntentId, not even an explicit nil
### GetCardBrand

`func (o *ThreeDSSession) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *ThreeDSSession) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *ThreeDSSession) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *ThreeDSSession) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *ThreeDSSession) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *ThreeDSSession) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetExpirationDate

`func (o *ThreeDSSession) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ThreeDSSession) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ThreeDSSession) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ThreeDSSession) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ThreeDSSession) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ThreeDSSession) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ThreeDSSession) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ThreeDSSession) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ThreeDSSession) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ThreeDSSession) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedBy

`func (o *ThreeDSSession) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ThreeDSSession) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ThreeDSSession) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ThreeDSSession) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ThreeDSSession) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ThreeDSSession) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetModifiedDate

`func (o *ThreeDSSession) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ThreeDSSession) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ThreeDSSession) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ThreeDSSession) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ThreeDSSession) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ThreeDSSession) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedBy

`func (o *ThreeDSSession) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ThreeDSSession) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ThreeDSSession) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ThreeDSSession) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *ThreeDSSession) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *ThreeDSSession) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetDevice

`func (o *ThreeDSSession) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ThreeDSSession) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ThreeDSSession) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ThreeDSSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ThreeDSSession) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ThreeDSSession) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDeviceInfo

`func (o *ThreeDSSession) GetDeviceInfo() ThreeDSDeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *ThreeDSSession) GetDeviceInfoOk() (*ThreeDSDeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *ThreeDSSession) SetDeviceInfo(v ThreeDSDeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *ThreeDSSession) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetVersion

`func (o *ThreeDSSession) GetVersion() ThreeDSVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSSession) GetVersionOk() (*ThreeDSVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSSession) SetVersion(v ThreeDSVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSSession) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMethod

`func (o *ThreeDSSession) GetMethod() ThreeDSMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ThreeDSSession) GetMethodOk() (*ThreeDSMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ThreeDSSession) SetMethod(v ThreeDSMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ThreeDSSession) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAuthentication

`func (o *ThreeDSSession) GetAuthentication() ThreeDSAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ThreeDSSession) GetAuthenticationOk() (*ThreeDSAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ThreeDSSession) SetAuthentication(v ThreeDSAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ThreeDSSession) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


