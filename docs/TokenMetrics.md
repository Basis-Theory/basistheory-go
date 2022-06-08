# TokenMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**LastCreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTokenMetrics

`func NewTokenMetrics() *TokenMetrics`

NewTokenMetrics instantiates a new TokenMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenMetricsWithDefaults

`func NewTokenMetricsWithDefaults() *TokenMetrics`

NewTokenMetricsWithDefaults instantiates a new TokenMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TokenMetrics) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TokenMetrics) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TokenMetrics) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TokenMetrics) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastCreatedAt

`func (o *TokenMetrics) GetLastCreatedAt() time.Time`

GetLastCreatedAt returns the LastCreatedAt field if non-nil, zero value otherwise.

### GetLastCreatedAtOk

`func (o *TokenMetrics) GetLastCreatedAtOk() (*time.Time, bool)`

GetLastCreatedAtOk returns a tuple with the LastCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCreatedAt

`func (o *TokenMetrics) SetLastCreatedAt(v time.Time)`

SetLastCreatedAt sets LastCreatedAt field to given value.

### HasLastCreatedAt

`func (o *TokenMetrics) HasLastCreatedAt() bool`

HasLastCreatedAt returns a boolean if a field has been set.

### SetLastCreatedAtNil

`func (o *TokenMetrics) SetLastCreatedAtNil(b bool)`

 SetLastCreatedAtNil sets the value for LastCreatedAt to be an explicit nil

### UnsetLastCreatedAt
`func (o *TokenMetrics) UnsetLastCreatedAt()`

UnsetLastCreatedAt ensures that no value is present for LastCreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


