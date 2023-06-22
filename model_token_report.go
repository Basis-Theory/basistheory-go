/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
)

// checks if the TokenReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenReport{}

// TokenReport struct for TokenReport
type TokenReport struct {
	IncludedMonthlyActiveTokens *int64                      `json:"included_monthly_active_tokens,omitempty"`
	MonthlyActiveTokens         *int64                      `json:"monthly_active_tokens,omitempty"`
	MetricsByType               map[string]TokenMetrics     `json:"metrics_by_type,omitempty"`
	MonthlyActiveTokenHistory   []MonthlyActiveTokenHistory `json:"monthly_active_token_history,omitempty"`
}

// NewTokenReport instantiates a new TokenReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenReport() *TokenReport {
	this := TokenReport{}
	return &this
}

// NewTokenReportWithDefaults instantiates a new TokenReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenReportWithDefaults() *TokenReport {
	this := TokenReport{}
	return &this
}

// GetIncludedMonthlyActiveTokens returns the IncludedMonthlyActiveTokens field value if set, zero value otherwise.
func (o *TokenReport) GetIncludedMonthlyActiveTokens() int64 {
	if o == nil || IsNil(o.IncludedMonthlyActiveTokens) {
		var ret int64
		return ret
	}
	return *o.IncludedMonthlyActiveTokens
}

// GetIncludedMonthlyActiveTokensOk returns a tuple with the IncludedMonthlyActiveTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenReport) GetIncludedMonthlyActiveTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.IncludedMonthlyActiveTokens) {
		return nil, false
	}
	return o.IncludedMonthlyActiveTokens, true
}

// HasIncludedMonthlyActiveTokens returns a boolean if a field is not nil.
func (o *TokenReport) HasIncludedMonthlyActiveTokens() bool {
	if o != nil && !IsNil(o.IncludedMonthlyActiveTokens) {
		return true
	}

	return false
}

// SetIncludedMonthlyActiveTokens gets a reference to the given int64 and assigns it to the IncludedMonthlyActiveTokens field.
func (o *TokenReport) SetIncludedMonthlyActiveTokens(v int64) {
	o.IncludedMonthlyActiveTokens = &v
}

// GetMonthlyActiveTokens returns the MonthlyActiveTokens field value if set, zero value otherwise.
func (o *TokenReport) GetMonthlyActiveTokens() int64 {
	if o == nil || IsNil(o.MonthlyActiveTokens) {
		var ret int64
		return ret
	}
	return *o.MonthlyActiveTokens
}

// GetMonthlyActiveTokensOk returns a tuple with the MonthlyActiveTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenReport) GetMonthlyActiveTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.MonthlyActiveTokens) {
		return nil, false
	}
	return o.MonthlyActiveTokens, true
}

// HasMonthlyActiveTokens returns a boolean if a field is not nil.
func (o *TokenReport) HasMonthlyActiveTokens() bool {
	if o != nil && !IsNil(o.MonthlyActiveTokens) {
		return true
	}

	return false
}

// SetMonthlyActiveTokens gets a reference to the given int64 and assigns it to the MonthlyActiveTokens field.
func (o *TokenReport) SetMonthlyActiveTokens(v int64) {
	o.MonthlyActiveTokens = &v
}

// GetMetricsByType returns the MetricsByType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenReport) GetMetricsByType() map[string]TokenMetrics {
	if o == nil {
		var ret map[string]TokenMetrics
		return ret
	}
	return o.MetricsByType
}

// GetMetricsByTypeOk returns a tuple with the MetricsByType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenReport) GetMetricsByTypeOk() (*map[string]TokenMetrics, bool) {
	if o == nil || IsNil(o.MetricsByType) {
		return nil, false
	}
	return &o.MetricsByType, true
}

// HasMetricsByType returns a boolean if a field is not nil.
func (o *TokenReport) HasMetricsByType() bool {
	if o != nil && !IsNil(o.MetricsByType) {
		return true
	}

	return false
}

// SetMetricsByType gets a reference to the given map[string]TokenMetrics and assigns it to the MetricsByType field.
func (o *TokenReport) SetMetricsByType(v map[string]TokenMetrics) {
	o.MetricsByType = v
}

// GetMonthlyActiveTokenHistory returns the MonthlyActiveTokenHistory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenReport) GetMonthlyActiveTokenHistory() []MonthlyActiveTokenHistory {
	if o == nil {
		var ret []MonthlyActiveTokenHistory
		return ret
	}
	return o.MonthlyActiveTokenHistory
}

// GetMonthlyActiveTokenHistoryOk returns a tuple with the MonthlyActiveTokenHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenReport) GetMonthlyActiveTokenHistoryOk() ([]MonthlyActiveTokenHistory, bool) {
	if o == nil || IsNil(o.MonthlyActiveTokenHistory) {
		return nil, false
	}
	return o.MonthlyActiveTokenHistory, true
}

// HasMonthlyActiveTokenHistory returns a boolean if a field is not nil.
func (o *TokenReport) HasMonthlyActiveTokenHistory() bool {
	if o != nil && !IsNil(o.MonthlyActiveTokenHistory) {
		return true
	}

	return false
}

// SetMonthlyActiveTokenHistory gets a reference to the given []MonthlyActiveTokenHistory and assigns it to the MonthlyActiveTokenHistory field.
func (o *TokenReport) SetMonthlyActiveTokenHistory(v []MonthlyActiveTokenHistory) {
	o.MonthlyActiveTokenHistory = v
}

func (o TokenReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludedMonthlyActiveTokens) {
		toSerialize["included_monthly_active_tokens"] = o.IncludedMonthlyActiveTokens
	}
	if !IsNil(o.MonthlyActiveTokens) {
		toSerialize["monthly_active_tokens"] = o.MonthlyActiveTokens
	}
	if o.MetricsByType != nil {
		toSerialize["metrics_by_type"] = o.MetricsByType
	}
	if o.MonthlyActiveTokenHistory != nil {
		toSerialize["monthly_active_token_history"] = o.MonthlyActiveTokenHistory
	}
	return toSerialize, nil
}

type NullableTokenReport struct {
	value *TokenReport
	isSet bool
}

func (v NullableTokenReport) Get() *TokenReport {
	return v.value
}

func (v *NullableTokenReport) Set(val *TokenReport) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenReport) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenReport(val *TokenReport) *NullableTokenReport {
	return &NullableTokenReport{value: val, isSet: true}
}

func (v NullableTokenReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
