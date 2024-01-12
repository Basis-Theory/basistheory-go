# SearchTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** |  | [optional] 
**Page** | Pointer to **NullableInt32** |  | [optional] 
**Start** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSearchTokensRequest

`func NewSearchTokensRequest() *SearchTokensRequest`

NewSearchTokensRequest instantiates a new SearchTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTokensRequestWithDefaults

`func NewSearchTokensRequestWithDefaults() *SearchTokensRequest`

NewSearchTokensRequestWithDefaults instantiates a new SearchTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchTokensRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchTokensRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchTokensRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchTokensRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *SearchTokensRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *SearchTokensRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetPage

`func (o *SearchTokensRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchTokensRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchTokensRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchTokensRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *SearchTokensRequest) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *SearchTokensRequest) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil
### GetStart

`func (o *SearchTokensRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SearchTokensRequest) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SearchTokensRequest) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *SearchTokensRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *SearchTokensRequest) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *SearchTokensRequest) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetSize

`func (o *SearchTokensRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchTokensRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchTokensRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchTokensRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *SearchTokensRequest) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *SearchTokensRequest) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


