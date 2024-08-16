# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalItems** | Pointer to **NullableInt32** |  | [optional] 
**PageNumber** | Pointer to **NullableInt32** |  | [optional] 
**PageSize** | Pointer to **NullableInt32** |  | [optional] 
**TotalPages** | Pointer to **NullableInt32** |  | [optional] 
**After** | Pointer to **NullableString** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalItems

`func (o *Pagination) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *Pagination) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *Pagination) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *Pagination) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *Pagination) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *Pagination) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
### GetPageNumber

`func (o *Pagination) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *Pagination) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *Pagination) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *Pagination) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### SetPageNumberNil

`func (o *Pagination) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *Pagination) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageSize

`func (o *Pagination) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Pagination) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Pagination) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Pagination) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *Pagination) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *Pagination) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetTotalPages

`func (o *Pagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *Pagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *Pagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *Pagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### SetTotalPagesNil

`func (o *Pagination) SetTotalPagesNil(b bool)`

 SetTotalPagesNil sets the value for TotalPages to be an explicit nil

### UnsetTotalPages
`func (o *Pagination) UnsetTotalPages()`

UnsetTotalPages ensures that no value is present for TotalPages, not even an explicit nil
### GetAfter

`func (o *Pagination) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Pagination) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Pagination) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Pagination) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *Pagination) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *Pagination) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetNext

`func (o *Pagination) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Pagination) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Pagination) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Pagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *Pagination) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *Pagination) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


