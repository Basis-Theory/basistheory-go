# CreateAtomicBankRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | [**Bank**](Bank.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateAtomicBankRequest

`func NewCreateAtomicBankRequest(bank Bank, ) *CreateAtomicBankRequest`

NewCreateAtomicBankRequest instantiates a new CreateAtomicBankRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAtomicBankRequestWithDefaults

`func NewCreateAtomicBankRequestWithDefaults() *CreateAtomicBankRequest`

NewCreateAtomicBankRequestWithDefaults instantiates a new CreateAtomicBankRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *CreateAtomicBankRequest) GetBank() Bank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *CreateAtomicBankRequest) GetBankOk() (*Bank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *CreateAtomicBankRequest) SetBank(v Bank)`

SetBank sets Bank field to given value.


### GetMetadata

`func (o *CreateAtomicBankRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateAtomicBankRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateAtomicBankRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateAtomicBankRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateAtomicBankRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateAtomicBankRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


