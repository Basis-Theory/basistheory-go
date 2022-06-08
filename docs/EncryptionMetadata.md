# EncryptionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cek** | Pointer to [**EncryptionKey**](EncryptionKey.md) |  | [optional] 
**Kek** | Pointer to [**EncryptionKey**](EncryptionKey.md) |  | [optional] 

## Methods

### NewEncryptionMetadata

`func NewEncryptionMetadata() *EncryptionMetadata`

NewEncryptionMetadata instantiates a new EncryptionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionMetadataWithDefaults

`func NewEncryptionMetadataWithDefaults() *EncryptionMetadata`

NewEncryptionMetadataWithDefaults instantiates a new EncryptionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCek

`func (o *EncryptionMetadata) GetCek() EncryptionKey`

GetCek returns the Cek field if non-nil, zero value otherwise.

### GetCekOk

`func (o *EncryptionMetadata) GetCekOk() (*EncryptionKey, bool)`

GetCekOk returns a tuple with the Cek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCek

`func (o *EncryptionMetadata) SetCek(v EncryptionKey)`

SetCek sets Cek field to given value.

### HasCek

`func (o *EncryptionMetadata) HasCek() bool`

HasCek returns a boolean if a field has been set.

### GetKek

`func (o *EncryptionMetadata) GetKek() EncryptionKey`

GetKek returns the Kek field if non-nil, zero value otherwise.

### GetKekOk

`func (o *EncryptionMetadata) GetKekOk() (*EncryptionKey, bool)`

GetKekOk returns a tuple with the Kek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKek

`func (o *EncryptionMetadata) SetKek(v EncryptionKey)`

SetKek sets Kek field to given value.

### HasKek

`func (o *EncryptionMetadata) HasKek() bool`

HasKek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


