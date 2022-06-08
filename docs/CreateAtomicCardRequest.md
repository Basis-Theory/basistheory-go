# CreateAtomicCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | [**Card**](Card.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateAtomicCardRequest

`func NewCreateAtomicCardRequest(card Card, ) *CreateAtomicCardRequest`

NewCreateAtomicCardRequest instantiates a new CreateAtomicCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAtomicCardRequestWithDefaults

`func NewCreateAtomicCardRequestWithDefaults() *CreateAtomicCardRequest`

NewCreateAtomicCardRequestWithDefaults instantiates a new CreateAtomicCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *CreateAtomicCardRequest) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CreateAtomicCardRequest) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CreateAtomicCardRequest) SetCard(v Card)`

SetCard sets Card field to given value.


### GetMetadata

`func (o *CreateAtomicCardRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateAtomicCardRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateAtomicCardRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateAtomicCardRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateAtomicCardRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateAtomicCardRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


