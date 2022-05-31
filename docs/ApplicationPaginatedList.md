# ApplicationPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]Application**](Application.md) |  | [optional] 

## Methods

### NewApplicationPaginatedList

`func NewApplicationPaginatedList() *ApplicationPaginatedList`

NewApplicationPaginatedList instantiates a new ApplicationPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPaginatedListWithDefaults

`func NewApplicationPaginatedListWithDefaults() *ApplicationPaginatedList`

NewApplicationPaginatedListWithDefaults instantiates a new ApplicationPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ApplicationPaginatedList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ApplicationPaginatedList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ApplicationPaginatedList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ApplicationPaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *ApplicationPaginatedList) GetData() []Application`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationPaginatedList) GetDataOk() (*[]Application, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationPaginatedList) SetData(v []Application)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationPaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ApplicationPaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ApplicationPaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


