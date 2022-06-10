# EncryptionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Prov** | Pointer to **NullableString** |  | [optional] 
**Alg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEncryptionKey

`func NewEncryptionKey(key string, ) *EncryptionKey`

NewEncryptionKey instantiates a new EncryptionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyWithDefaults

`func NewEncryptionKeyWithDefaults() *EncryptionKey`

NewEncryptionKeyWithDefaults instantiates a new EncryptionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EncryptionKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EncryptionKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EncryptionKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetProv

`func (o *EncryptionKey) GetProv() string`

GetProv returns the Prov field if non-nil, zero value otherwise.

### GetProvOk

`func (o *EncryptionKey) GetProvOk() (*string, bool)`

GetProvOk returns a tuple with the Prov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProv

`func (o *EncryptionKey) SetProv(v string)`

SetProv sets Prov field to given value.

### HasProv

`func (o *EncryptionKey) HasProv() bool`

HasProv returns a boolean if a field has been set.

### SetProvNil

`func (o *EncryptionKey) SetProvNil(b bool)`

 SetProvNil sets the value for Prov to be an explicit nil

### UnsetProv
`func (o *EncryptionKey) UnsetProv()`

UnsetProv ensures that no value is present for Prov, not even an explicit nil
### GetAlg

`func (o *EncryptionKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *EncryptionKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *EncryptionKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *EncryptionKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### SetAlgNil

`func (o *EncryptionKey) SetAlgNil(b bool)`

 SetAlgNil sets the value for Alg to be an explicit nil

### UnsetAlg
`func (o *EncryptionKey) UnsetAlg()`

UnsetAlg ensures that no value is present for Alg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


