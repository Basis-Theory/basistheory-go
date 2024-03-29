/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"time"
)

// checks if the TokenMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenMetrics{}

// TokenMetrics struct for TokenMetrics
type TokenMetrics struct {
	Count         *int64       `json:"count,omitempty"`
	LastCreatedAt NullableTime `json:"last_created_at,omitempty"`
}

// NewTokenMetrics instantiates a new TokenMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenMetrics() *TokenMetrics {
	this := TokenMetrics{}
	return &this
}

// NewTokenMetricsWithDefaults instantiates a new TokenMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenMetricsWithDefaults() *TokenMetrics {
	this := TokenMetrics{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TokenMetrics) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenMetrics) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field is not nil.
func (o *TokenMetrics) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *TokenMetrics) SetCount(v int64) {
	o.Count = &v
}

// GetLastCreatedAt returns the LastCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenMetrics) GetLastCreatedAt() time.Time {
	if o == nil || IsNil(o.LastCreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastCreatedAt.Get()
}

// GetLastCreatedAtOk returns a tuple with the LastCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenMetrics) GetLastCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCreatedAt.Get(), o.LastCreatedAt.IsSet()
}

// HasLastCreatedAt returns a boolean if a field is not nil.
func (o *TokenMetrics) HasLastCreatedAt() bool {
	if o != nil && !IsNil(o.LastCreatedAt) {
		return true
	}

	return false
}

// SetLastCreatedAt gets a reference to the given NullableTime and assigns it to the LastCreatedAt field.
func (o *TokenMetrics) SetLastCreatedAt(v time.Time) {
	o.LastCreatedAt.Set(&v)
}

// SetLastCreatedAtNil sets the value for LastCreatedAt to be an explicit nil
func (o *TokenMetrics) SetLastCreatedAtNil() {
	o.LastCreatedAt.Set(nil)
}

// UnsetLastCreatedAt ensures that no value is present for LastCreatedAt, not even an explicit nil
func (o *TokenMetrics) UnsetLastCreatedAt() {
	o.LastCreatedAt.Unset()
}

func (o TokenMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.LastCreatedAt.IsSet() {
		toSerialize["last_created_at"] = o.LastCreatedAt.Get()
	}
	return toSerialize, nil
}

type NullableTokenMetrics struct {
	value *TokenMetrics
	isSet bool
}

func (v NullableTokenMetrics) Get() *TokenMetrics {
	return v.value
}

func (v *NullableTokenMetrics) Set(val *TokenMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenMetrics(val *TokenMetrics) *NullableTokenMetrics {
	return &NullableTokenMetrics{value: val, isSet: true}
}

func (v NullableTokenMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
