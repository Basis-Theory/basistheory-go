# CreateTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FingerprintExpression** | Pointer to **NullableString** |  | [optional] 
**Mask** | Pointer to **interface{}** |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Privacy** | Pointer to [**Privacy**](Privacy.md) |  | [optional] 
**SearchIndexes** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateTokenResponse

`func NewCreateTokenResponse() *CreateTokenResponse`

NewCreateTokenResponse instantiates a new CreateTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenResponseWithDefaults

`func NewCreateTokenResponseWithDefaults() *CreateTokenResponse`

NewCreateTokenResponseWithDefaults instantiates a new CreateTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateTokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTokenResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateTokenResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateTokenResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateTokenResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTenantId

`func (o *CreateTokenResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateTokenResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateTokenResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateTokenResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetType

`func (o *CreateTokenResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTokenResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTokenResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTokenResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateTokenResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateTokenResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFingerprint

`func (o *CreateTokenResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CreateTokenResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CreateTokenResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CreateTokenResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *CreateTokenResponse) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *CreateTokenResponse) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFingerprintExpression

`func (o *CreateTokenResponse) GetFingerprintExpression() string`

GetFingerprintExpression returns the FingerprintExpression field if non-nil, zero value otherwise.

### GetFingerprintExpressionOk

`func (o *CreateTokenResponse) GetFingerprintExpressionOk() (*string, bool)`

GetFingerprintExpressionOk returns a tuple with the FingerprintExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintExpression

`func (o *CreateTokenResponse) SetFingerprintExpression(v string)`

SetFingerprintExpression sets FingerprintExpression field to given value.

### HasFingerprintExpression

`func (o *CreateTokenResponse) HasFingerprintExpression() bool`

HasFingerprintExpression returns a boolean if a field has been set.

### SetFingerprintExpressionNil

`func (o *CreateTokenResponse) SetFingerprintExpressionNil(b bool)`

 SetFingerprintExpressionNil sets the value for FingerprintExpression to be an explicit nil

### UnsetFingerprintExpression
`func (o *CreateTokenResponse) UnsetFingerprintExpression()`

UnsetFingerprintExpression ensures that no value is present for FingerprintExpression, not even an explicit nil
### GetMask

`func (o *CreateTokenResponse) GetMask() interface{}`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *CreateTokenResponse) GetMaskOk() (*interface{}, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *CreateTokenResponse) SetMask(v interface{})`

SetMask sets Mask field to given value.

### HasMask

`func (o *CreateTokenResponse) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *CreateTokenResponse) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *CreateTokenResponse) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetData

`func (o *CreateTokenResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTokenResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTokenResponse) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CreateTokenResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CreateTokenResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CreateTokenResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMetadata

`func (o *CreateTokenResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTokenResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTokenResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTokenResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateTokenResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateTokenResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPrivacy

`func (o *CreateTokenResponse) GetPrivacy() Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *CreateTokenResponse) GetPrivacyOk() (*Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *CreateTokenResponse) SetPrivacy(v Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *CreateTokenResponse) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetSearchIndexes

`func (o *CreateTokenResponse) GetSearchIndexes() []string`

GetSearchIndexes returns the SearchIndexes field if non-nil, zero value otherwise.

### GetSearchIndexesOk

`func (o *CreateTokenResponse) GetSearchIndexesOk() (*[]string, bool)`

GetSearchIndexesOk returns a tuple with the SearchIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexes

`func (o *CreateTokenResponse) SetSearchIndexes(v []string)`

SetSearchIndexes sets SearchIndexes field to given value.

### HasSearchIndexes

`func (o *CreateTokenResponse) HasSearchIndexes() bool`

HasSearchIndexes returns a boolean if a field has been set.

### SetSearchIndexesNil

`func (o *CreateTokenResponse) SetSearchIndexesNil(b bool)`

 SetSearchIndexesNil sets the value for SearchIndexes to be an explicit nil

### UnsetSearchIndexes
`func (o *CreateTokenResponse) UnsetSearchIndexes()`

UnsetSearchIndexes ensures that no value is present for SearchIndexes, not even an explicit nil
### GetCreatedBy

`func (o *CreateTokenResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateTokenResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateTokenResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateTokenResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *CreateTokenResponse) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *CreateTokenResponse) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *CreateTokenResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateTokenResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateTokenResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateTokenResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CreateTokenResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CreateTokenResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetModifiedBy

`func (o *CreateTokenResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CreateTokenResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CreateTokenResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CreateTokenResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *CreateTokenResponse) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *CreateTokenResponse) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetModifiedAt

`func (o *CreateTokenResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CreateTokenResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CreateTokenResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CreateTokenResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *CreateTokenResponse) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *CreateTokenResponse) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetExpiresAt

`func (o *CreateTokenResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateTokenResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateTokenResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateTokenResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateTokenResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateTokenResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetContainer

`func (o *CreateTokenResponse) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CreateTokenResponse) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CreateTokenResponse) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *CreateTokenResponse) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *CreateTokenResponse) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *CreateTokenResponse) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


