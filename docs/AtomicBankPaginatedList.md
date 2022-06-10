# AtomicBankPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]AtomicBank**](AtomicBank.md) |  | [optional] 

## Methods

### NewAtomicBankPaginatedList

`func NewAtomicBankPaginatedList() *AtomicBankPaginatedList`

NewAtomicBankPaginatedList instantiates a new AtomicBankPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtomicBankPaginatedListWithDefaults

`func NewAtomicBankPaginatedListWithDefaults() *AtomicBankPaginatedList`

NewAtomicBankPaginatedListWithDefaults instantiates a new AtomicBankPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AtomicBankPaginatedList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AtomicBankPaginatedList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AtomicBankPaginatedList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AtomicBankPaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AtomicBankPaginatedList) GetData() []AtomicBank`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AtomicBankPaginatedList) GetDataOk() (*[]AtomicBank, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AtomicBankPaginatedList) SetData(v []AtomicBank)`

SetData sets Data field to given value.

### HasData

`func (o *AtomicBankPaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AtomicBankPaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AtomicBankPaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


