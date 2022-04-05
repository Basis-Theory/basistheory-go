# ReactorModelPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]ReactorModel**](ReactorModel.md) |  | [optional] 

## Methods

### NewReactorModelPaginatedList

`func NewReactorModelPaginatedList() *ReactorModelPaginatedList`

NewReactorModelPaginatedList instantiates a new ReactorModelPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactorModelPaginatedListWithDefaults

`func NewReactorModelPaginatedListWithDefaults() *ReactorModelPaginatedList`

NewReactorModelPaginatedListWithDefaults instantiates a new ReactorModelPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ReactorModelPaginatedList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ReactorModelPaginatedList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ReactorModelPaginatedList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ReactorModelPaginatedList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *ReactorModelPaginatedList) GetData() []ReactorModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReactorModelPaginatedList) GetDataOk() (*[]ReactorModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReactorModelPaginatedList) SetData(v []ReactorModel)`

SetData sets Data field to given value.

### HasData

`func (o *ReactorModelPaginatedList) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ReactorModelPaginatedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ReactorModelPaginatedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

