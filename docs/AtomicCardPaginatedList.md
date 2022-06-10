# AtomicCardPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]AtomicCard**](AtomicCard.md) |  | [optional] 

## Methods

### NewAtomicCardPaginatedList

`func NewAtomicCardPaginatedList() *AtomicCardPaginatedList`

NewAtomicCardPaginatedList instantiates a new AtomicCardPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtomicCardPaginatedListWithDefaults

`func NewAtomicCardPaginatedListWithDefaults() *AtomicCardPaginatedList`

NewAtomicCardPaginatedListWithDefaults instantiates a new AtomicCardPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AtomicCardPaginatedList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AtomicCardPaginatedList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AtomicCardPaginatedList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AtomicCardPaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AtomicCardPaginatedList) GetData() []AtomicCard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AtomicCardPaginatedList) GetDataOk() (*[]AtomicCard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AtomicCardPaginatedList) SetData(v []AtomicCard)`

SetData sets Data field to given value.

### HasData

`func (o *AtomicCardPaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AtomicCardPaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AtomicCardPaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


