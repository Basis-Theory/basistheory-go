# TokenCursorPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**CursorPagination**](CursorPagination.md) |  | [optional] 
**Data** | Pointer to [**[]Token**](Token.md) |  | [optional] 

## Methods

### NewTokenCursorPaginatedList

`func NewTokenCursorPaginatedList() *TokenCursorPaginatedList`

NewTokenCursorPaginatedList instantiates a new TokenCursorPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCursorPaginatedListWithDefaults

`func NewTokenCursorPaginatedListWithDefaults() *TokenCursorPaginatedList`

NewTokenCursorPaginatedListWithDefaults instantiates a new TokenCursorPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *TokenCursorPaginatedList) GetPagination() CursorPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenCursorPaginatedList) GetPaginationOk() (*CursorPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenCursorPaginatedList) SetPagination(v CursorPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TokenCursorPaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *TokenCursorPaginatedList) GetData() []Token`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenCursorPaginatedList) GetDataOk() (*[]Token, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenCursorPaginatedList) SetData(v []Token)`

SetData sets Data field to given value.

### HasData

`func (o *TokenCursorPaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *TokenCursorPaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TokenCursorPaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


