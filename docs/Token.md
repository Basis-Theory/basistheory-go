# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Encryption** | Pointer to [**EncryptionMetadata**](EncryptionMetadata.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FingerprintExpression** | Pointer to **NullableString** |  | [optional] 
**Mask** | Pointer to **interface{}** |  | [optional] 
**Privacy** | Pointer to [**Privacy**](Privacy.md) |  | [optional] 
**SearchIndexes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Token) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Token) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Token) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Token) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Token) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Token) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Token) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Token) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTenantId

`func (o *Token) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Token) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Token) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Token) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetData

`func (o *Token) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Token) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Token) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Token) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Token) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Token) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMetadata

`func (o *Token) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Token) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Token) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Token) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Token) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Token) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetEncryption

`func (o *Token) GetEncryption() EncryptionMetadata`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Token) GetEncryptionOk() (*EncryptionMetadata, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Token) SetEncryption(v EncryptionMetadata)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Token) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Token) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Token) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Token) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Token) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Token) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Token) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *Token) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Token) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Token) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Token) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Token) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Token) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetModifiedBy

`func (o *Token) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Token) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Token) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Token) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *Token) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *Token) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetModifiedAt

`func (o *Token) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Token) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Token) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Token) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *Token) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Token) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetFingerprint

`func (o *Token) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Token) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Token) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Token) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Token) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Token) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFingerprintExpression

`func (o *Token) GetFingerprintExpression() string`

GetFingerprintExpression returns the FingerprintExpression field if non-nil, zero value otherwise.

### GetFingerprintExpressionOk

`func (o *Token) GetFingerprintExpressionOk() (*string, bool)`

GetFingerprintExpressionOk returns a tuple with the FingerprintExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintExpression

`func (o *Token) SetFingerprintExpression(v string)`

SetFingerprintExpression sets FingerprintExpression field to given value.

### HasFingerprintExpression

`func (o *Token) HasFingerprintExpression() bool`

HasFingerprintExpression returns a boolean if a field has been set.

### SetFingerprintExpressionNil

`func (o *Token) SetFingerprintExpressionNil(b bool)`

 SetFingerprintExpressionNil sets the value for FingerprintExpression to be an explicit nil

### UnsetFingerprintExpression
`func (o *Token) UnsetFingerprintExpression()`

UnsetFingerprintExpression ensures that no value is present for FingerprintExpression, not even an explicit nil
### GetMask

`func (o *Token) GetMask() interface{}`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *Token) GetMaskOk() (*interface{}, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *Token) SetMask(v interface{})`

SetMask sets Mask field to given value.

### HasMask

`func (o *Token) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *Token) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *Token) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetPrivacy

`func (o *Token) GetPrivacy() Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *Token) GetPrivacyOk() (*Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *Token) SetPrivacy(v Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *Token) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetSearchIndexes

`func (o *Token) GetSearchIndexes() []string`

GetSearchIndexes returns the SearchIndexes field if non-nil, zero value otherwise.

### GetSearchIndexesOk

`func (o *Token) GetSearchIndexesOk() (*[]string, bool)`

GetSearchIndexesOk returns a tuple with the SearchIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexes

`func (o *Token) SetSearchIndexes(v []string)`

SetSearchIndexes sets SearchIndexes field to given value.

### HasSearchIndexes

`func (o *Token) HasSearchIndexes() bool`

HasSearchIndexes returns a boolean if a field has been set.

### SetSearchIndexesNil

`func (o *Token) SetSearchIndexesNil(b bool)`

 SetSearchIndexesNil sets the value for SearchIndexes to be an explicit nil

### UnsetSearchIndexes
`func (o *Token) UnsetSearchIndexes()`

UnsetSearchIndexes ensures that no value is present for SearchIndexes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


