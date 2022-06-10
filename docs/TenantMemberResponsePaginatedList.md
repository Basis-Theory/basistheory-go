# TenantMemberResponsePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]TenantMemberResponse**](TenantMemberResponse.md) |  | [optional] 

## Methods

### NewTenantMemberResponsePaginatedList

`func NewTenantMemberResponsePaginatedList() *TenantMemberResponsePaginatedList`

NewTenantMemberResponsePaginatedList instantiates a new TenantMemberResponsePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantMemberResponsePaginatedListWithDefaults

`func NewTenantMemberResponsePaginatedListWithDefaults() *TenantMemberResponsePaginatedList`

NewTenantMemberResponsePaginatedListWithDefaults instantiates a new TenantMemberResponsePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *TenantMemberResponsePaginatedList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TenantMemberResponsePaginatedList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TenantMemberResponsePaginatedList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TenantMemberResponsePaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *TenantMemberResponsePaginatedList) GetData() []TenantMemberResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TenantMemberResponsePaginatedList) GetDataOk() (*[]TenantMemberResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TenantMemberResponsePaginatedList) SetData(v []TenantMemberResponse)`

SetData sets Data field to given value.

### HasData

`func (o *TenantMemberResponsePaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *TenantMemberResponsePaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TenantMemberResponsePaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


