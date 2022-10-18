# TokenReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedMonthlyActiveTokens** | Pointer to **int64** |  | [optional] 
**MonthlyActiveTokens** | Pointer to **int64** |  | [optional] 
**MetricsByType** | Pointer to [**map[string]TokenMetrics**](TokenMetrics.md) |  | [optional] 
**MonthlyActiveTokenHistory** | Pointer to [**[]MonthlyActiveTokenHistory**](MonthlyActiveTokenHistory.md) |  | [optional] 

## Methods

### NewTokenReport

`func NewTokenReport() *TokenReport`

NewTokenReport instantiates a new TokenReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenReportWithDefaults

`func NewTokenReportWithDefaults() *TokenReport`

NewTokenReportWithDefaults instantiates a new TokenReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedMonthlyActiveTokens

`func (o *TokenReport) GetIncludedMonthlyActiveTokens() int64`

GetIncludedMonthlyActiveTokens returns the IncludedMonthlyActiveTokens field if non-nil, zero value otherwise.

### GetIncludedMonthlyActiveTokensOk

`func (o *TokenReport) GetIncludedMonthlyActiveTokensOk() (*int64, bool)`

GetIncludedMonthlyActiveTokensOk returns a tuple with the IncludedMonthlyActiveTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedMonthlyActiveTokens

`func (o *TokenReport) SetIncludedMonthlyActiveTokens(v int64)`

SetIncludedMonthlyActiveTokens sets IncludedMonthlyActiveTokens field to given value.

### HasIncludedMonthlyActiveTokens

`func (o *TokenReport) HasIncludedMonthlyActiveTokens() bool`

HasIncludedMonthlyActiveTokens returns a boolean if a field has been set.

### GetMonthlyActiveTokens

`func (o *TokenReport) GetMonthlyActiveTokens() int64`

GetMonthlyActiveTokens returns the MonthlyActiveTokens field if non-nil, zero value otherwise.

### GetMonthlyActiveTokensOk

`func (o *TokenReport) GetMonthlyActiveTokensOk() (*int64, bool)`

GetMonthlyActiveTokensOk returns a tuple with the MonthlyActiveTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyActiveTokens

`func (o *TokenReport) SetMonthlyActiveTokens(v int64)`

SetMonthlyActiveTokens sets MonthlyActiveTokens field to given value.

### HasMonthlyActiveTokens

`func (o *TokenReport) HasMonthlyActiveTokens() bool`

HasMonthlyActiveTokens returns a boolean if a field has been set.

### GetMetricsByType

`func (o *TokenReport) GetMetricsByType() map[string]TokenMetrics`

GetMetricsByType returns the MetricsByType field if non-nil, zero value otherwise.

### GetMetricsByTypeOk

`func (o *TokenReport) GetMetricsByTypeOk() (*map[string]TokenMetrics, bool)`

GetMetricsByTypeOk returns a tuple with the MetricsByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsByType

`func (o *TokenReport) SetMetricsByType(v map[string]TokenMetrics)`

SetMetricsByType sets MetricsByType field to given value.

### HasMetricsByType

`func (o *TokenReport) HasMetricsByType() bool`

HasMetricsByType returns a boolean if a field has been set.

### SetMetricsByTypeNil

`func (o *TokenReport) SetMetricsByTypeNil(b bool)`

 SetMetricsByTypeNil sets the value for MetricsByType to be an explicit nil

### UnsetMetricsByType
`func (o *TokenReport) UnsetMetricsByType()`

UnsetMetricsByType ensures that no value is present for MetricsByType, not even an explicit nil
### GetMonthlyActiveTokenHistory

`func (o *TokenReport) GetMonthlyActiveTokenHistory() []MonthlyActiveTokenHistory`

GetMonthlyActiveTokenHistory returns the MonthlyActiveTokenHistory field if non-nil, zero value otherwise.

### GetMonthlyActiveTokenHistoryOk

`func (o *TokenReport) GetMonthlyActiveTokenHistoryOk() (*[]MonthlyActiveTokenHistory, bool)`

GetMonthlyActiveTokenHistoryOk returns a tuple with the MonthlyActiveTokenHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyActiveTokenHistory

`func (o *TokenReport) SetMonthlyActiveTokenHistory(v []MonthlyActiveTokenHistory)`

SetMonthlyActiveTokenHistory sets MonthlyActiveTokenHistory field to given value.

### HasMonthlyActiveTokenHistory

`func (o *TokenReport) HasMonthlyActiveTokenHistory() bool`

HasMonthlyActiveTokenHistory returns a boolean if a field has been set.

### SetMonthlyActiveTokenHistoryNil

`func (o *TokenReport) SetMonthlyActiveTokenHistoryNil(b bool)`

 SetMonthlyActiveTokenHistoryNil sets the value for MonthlyActiveTokenHistory to be an explicit nil

### UnsetMonthlyActiveTokenHistory
`func (o *TokenReport) UnsetMonthlyActiveTokenHistory()`

UnsetMonthlyActiveTokenHistory ensures that no value is present for MonthlyActiveTokenHistory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


