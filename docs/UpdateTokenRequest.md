# UpdateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Privacy** | Pointer to [**UpdatePrivacy**](UpdatePrivacy.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SearchIndexes** | Pointer to **[]string** |  | [optional] 
**FingerprintExpression** | Pointer to **NullableString** |  | [optional] 
**Mask** | Pointer to **interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **NullableString** |  | [optional] 
**DeduplicateToken** | Pointer to **NullableBool** |  | [optional] 
**Containers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateTokenRequest

`func NewUpdateTokenRequest() *UpdateTokenRequest`

NewUpdateTokenRequest instantiates a new UpdateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTokenRequestWithDefaults

`func NewUpdateTokenRequestWithDefaults() *UpdateTokenRequest`

NewUpdateTokenRequestWithDefaults instantiates a new UpdateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateTokenRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateTokenRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateTokenRequest) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UpdateTokenRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UpdateTokenRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UpdateTokenRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPrivacy

`func (o *UpdateTokenRequest) GetPrivacy() UpdatePrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *UpdateTokenRequest) GetPrivacyOk() (*UpdatePrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *UpdateTokenRequest) SetPrivacy(v UpdatePrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *UpdateTokenRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateTokenRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateTokenRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateTokenRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateTokenRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateTokenRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateTokenRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSearchIndexes

`func (o *UpdateTokenRequest) GetSearchIndexes() []string`

GetSearchIndexes returns the SearchIndexes field if non-nil, zero value otherwise.

### GetSearchIndexesOk

`func (o *UpdateTokenRequest) GetSearchIndexesOk() (*[]string, bool)`

GetSearchIndexesOk returns a tuple with the SearchIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexes

`func (o *UpdateTokenRequest) SetSearchIndexes(v []string)`

SetSearchIndexes sets SearchIndexes field to given value.

### HasSearchIndexes

`func (o *UpdateTokenRequest) HasSearchIndexes() bool`

HasSearchIndexes returns a boolean if a field has been set.

### SetSearchIndexesNil

`func (o *UpdateTokenRequest) SetSearchIndexesNil(b bool)`

 SetSearchIndexesNil sets the value for SearchIndexes to be an explicit nil

### UnsetSearchIndexes
`func (o *UpdateTokenRequest) UnsetSearchIndexes()`

UnsetSearchIndexes ensures that no value is present for SearchIndexes, not even an explicit nil
### GetFingerprintExpression

`func (o *UpdateTokenRequest) GetFingerprintExpression() string`

GetFingerprintExpression returns the FingerprintExpression field if non-nil, zero value otherwise.

### GetFingerprintExpressionOk

`func (o *UpdateTokenRequest) GetFingerprintExpressionOk() (*string, bool)`

GetFingerprintExpressionOk returns a tuple with the FingerprintExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintExpression

`func (o *UpdateTokenRequest) SetFingerprintExpression(v string)`

SetFingerprintExpression sets FingerprintExpression field to given value.

### HasFingerprintExpression

`func (o *UpdateTokenRequest) HasFingerprintExpression() bool`

HasFingerprintExpression returns a boolean if a field has been set.

### SetFingerprintExpressionNil

`func (o *UpdateTokenRequest) SetFingerprintExpressionNil(b bool)`

 SetFingerprintExpressionNil sets the value for FingerprintExpression to be an explicit nil

### UnsetFingerprintExpression
`func (o *UpdateTokenRequest) UnsetFingerprintExpression()`

UnsetFingerprintExpression ensures that no value is present for FingerprintExpression, not even an explicit nil
### GetMask

`func (o *UpdateTokenRequest) GetMask() interface{}`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *UpdateTokenRequest) GetMaskOk() (*interface{}, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *UpdateTokenRequest) SetMask(v interface{})`

SetMask sets Mask field to given value.

### HasMask

`func (o *UpdateTokenRequest) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *UpdateTokenRequest) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *UpdateTokenRequest) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetExpiresAt

`func (o *UpdateTokenRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateTokenRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateTokenRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateTokenRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *UpdateTokenRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *UpdateTokenRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetDeduplicateToken

`func (o *UpdateTokenRequest) GetDeduplicateToken() bool`

GetDeduplicateToken returns the DeduplicateToken field if non-nil, zero value otherwise.

### GetDeduplicateTokenOk

`func (o *UpdateTokenRequest) GetDeduplicateTokenOk() (*bool, bool)`

GetDeduplicateTokenOk returns a tuple with the DeduplicateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicateToken

`func (o *UpdateTokenRequest) SetDeduplicateToken(v bool)`

SetDeduplicateToken sets DeduplicateToken field to given value.

### HasDeduplicateToken

`func (o *UpdateTokenRequest) HasDeduplicateToken() bool`

HasDeduplicateToken returns a boolean if a field has been set.

### SetDeduplicateTokenNil

`func (o *UpdateTokenRequest) SetDeduplicateTokenNil(b bool)`

 SetDeduplicateTokenNil sets the value for DeduplicateToken to be an explicit nil

### UnsetDeduplicateToken
`func (o *UpdateTokenRequest) UnsetDeduplicateToken()`

UnsetDeduplicateToken ensures that no value is present for DeduplicateToken, not even an explicit nil
### GetContainers

`func (o *UpdateTokenRequest) GetContainers() []string`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *UpdateTokenRequest) GetContainersOk() (*[]string, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *UpdateTokenRequest) SetContainers(v []string)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *UpdateTokenRequest) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### SetContainersNil

`func (o *UpdateTokenRequest) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *UpdateTokenRequest) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


