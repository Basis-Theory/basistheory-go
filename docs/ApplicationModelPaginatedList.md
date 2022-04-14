# ApplicationModelPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]ApplicationModel**](ApplicationModel.md) |  | [optional] 

## Methods

### NewApplicationModelPaginatedList

`func NewApplicationModelPaginatedList() *ApplicationModelPaginatedList`

NewApplicationModelPaginatedList instantiates a new ApplicationModelPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationModelPaginatedListWithDefaults

`func NewApplicationModelPaginatedListWithDefaults() *ApplicationModelPaginatedList`

NewApplicationModelPaginatedListWithDefaults instantiates a new ApplicationModelPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ApplicationModelPaginatedList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ApplicationModelPaginatedList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ApplicationModelPaginatedList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ApplicationModelPaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *ApplicationModelPaginatedList) GetData() []ApplicationModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationModelPaginatedList) GetDataOk() (*[]ApplicationModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationModelPaginatedList) SetData(v []ApplicationModel)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationModelPaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ApplicationModelPaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ApplicationModelPaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


