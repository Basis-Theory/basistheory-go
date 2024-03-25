# CreateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Data** | **interface{}** |  | 
**Privacy** | Pointer to [**Privacy**](Privacy.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SearchIndexes** | Pointer to **[]string** |  | [optional] 
**FingerprintExpression** | Pointer to **NullableString** |  | [optional] 
**Mask** | Pointer to **interface{}** |  | [optional] 
**DeduplicateToken** | Pointer to **NullableBool** |  | [optional] 
**ExpiresAt** | Pointer to **NullableString** |  | [optional] 
**Containers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateTokenRequest

`func NewCreateTokenRequest(data interface{}, ) *CreateTokenRequest`

NewCreateTokenRequest instantiates a new CreateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenRequestWithDefaults

`func NewCreateTokenRequestWithDefaults() *CreateTokenRequest`

NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTokenRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateTokenRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateTokenRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateTokenRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *CreateTokenRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTokenRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTokenRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTokenRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateTokenRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateTokenRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *CreateTokenRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTokenRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTokenRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *CreateTokenRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CreateTokenRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPrivacy

`func (o *CreateTokenRequest) GetPrivacy() Privacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *CreateTokenRequest) GetPrivacyOk() (*Privacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *CreateTokenRequest) SetPrivacy(v Privacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *CreateTokenRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateTokenRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTokenRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTokenRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTokenRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateTokenRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateTokenRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSearchIndexes

`func (o *CreateTokenRequest) GetSearchIndexes() []string`

GetSearchIndexes returns the SearchIndexes field if non-nil, zero value otherwise.

### GetSearchIndexesOk

`func (o *CreateTokenRequest) GetSearchIndexesOk() (*[]string, bool)`

GetSearchIndexesOk returns a tuple with the SearchIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexes

`func (o *CreateTokenRequest) SetSearchIndexes(v []string)`

SetSearchIndexes sets SearchIndexes field to given value.

### HasSearchIndexes

`func (o *CreateTokenRequest) HasSearchIndexes() bool`

HasSearchIndexes returns a boolean if a field has been set.

### SetSearchIndexesNil

`func (o *CreateTokenRequest) SetSearchIndexesNil(b bool)`

 SetSearchIndexesNil sets the value for SearchIndexes to be an explicit nil

### UnsetSearchIndexes
`func (o *CreateTokenRequest) UnsetSearchIndexes()`

UnsetSearchIndexes ensures that no value is present for SearchIndexes, not even an explicit nil
### GetFingerprintExpression

`func (o *CreateTokenRequest) GetFingerprintExpression() string`

GetFingerprintExpression returns the FingerprintExpression field if non-nil, zero value otherwise.

### GetFingerprintExpressionOk

`func (o *CreateTokenRequest) GetFingerprintExpressionOk() (*string, bool)`

GetFingerprintExpressionOk returns a tuple with the FingerprintExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintExpression

`func (o *CreateTokenRequest) SetFingerprintExpression(v string)`

SetFingerprintExpression sets FingerprintExpression field to given value.

### HasFingerprintExpression

`func (o *CreateTokenRequest) HasFingerprintExpression() bool`

HasFingerprintExpression returns a boolean if a field has been set.

### SetFingerprintExpressionNil

`func (o *CreateTokenRequest) SetFingerprintExpressionNil(b bool)`

 SetFingerprintExpressionNil sets the value for FingerprintExpression to be an explicit nil

### UnsetFingerprintExpression
`func (o *CreateTokenRequest) UnsetFingerprintExpression()`

UnsetFingerprintExpression ensures that no value is present for FingerprintExpression, not even an explicit nil
### GetMask

`func (o *CreateTokenRequest) GetMask() interface{}`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *CreateTokenRequest) GetMaskOk() (*interface{}, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *CreateTokenRequest) SetMask(v interface{})`

SetMask sets Mask field to given value.

### HasMask

`func (o *CreateTokenRequest) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *CreateTokenRequest) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *CreateTokenRequest) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetDeduplicateToken

`func (o *CreateTokenRequest) GetDeduplicateToken() bool`

GetDeduplicateToken returns the DeduplicateToken field if non-nil, zero value otherwise.

### GetDeduplicateTokenOk

`func (o *CreateTokenRequest) GetDeduplicateTokenOk() (*bool, bool)`

GetDeduplicateTokenOk returns a tuple with the DeduplicateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicateToken

`func (o *CreateTokenRequest) SetDeduplicateToken(v bool)`

SetDeduplicateToken sets DeduplicateToken field to given value.

### HasDeduplicateToken

`func (o *CreateTokenRequest) HasDeduplicateToken() bool`

HasDeduplicateToken returns a boolean if a field has been set.

### SetDeduplicateTokenNil

`func (o *CreateTokenRequest) SetDeduplicateTokenNil(b bool)`

 SetDeduplicateTokenNil sets the value for DeduplicateToken to be an explicit nil

### UnsetDeduplicateToken
`func (o *CreateTokenRequest) UnsetDeduplicateToken()`

UnsetDeduplicateToken ensures that no value is present for DeduplicateToken, not even an explicit nil
### GetExpiresAt

`func (o *CreateTokenRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateTokenRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateTokenRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateTokenRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateTokenRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateTokenRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetContainers

`func (o *CreateTokenRequest) GetContainers() []string`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *CreateTokenRequest) GetContainersOk() (*[]string, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *CreateTokenRequest) SetContainers(v []string)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *CreateTokenRequest) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### SetContainersNil

`func (o *CreateTokenRequest) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *CreateTokenRequest) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


