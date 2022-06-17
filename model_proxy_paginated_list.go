/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Server to Server Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
)

// ProxyPaginatedList struct for ProxyPaginatedList
type ProxyPaginatedList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []Proxy     `json:"data,omitempty"`
}

// NewProxyPaginatedList instantiates a new ProxyPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyPaginatedList() *ProxyPaginatedList {
	this := ProxyPaginatedList{}
	return &this
}

// NewProxyPaginatedListWithDefaults instantiates a new ProxyPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyPaginatedListWithDefaults() *ProxyPaginatedList {
	this := ProxyPaginatedList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ProxyPaginatedList) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyPaginatedList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ProxyPaginatedList) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ProxyPaginatedList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyPaginatedList) GetData() []Proxy {
	if o == nil {
		var ret []Proxy
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyPaginatedList) GetDataOk() ([]Proxy, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProxyPaginatedList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Proxy and assigns it to the Data field.
func (o *ProxyPaginatedList) SetData(v []Proxy) {
	o.Data = v
}

func (o ProxyPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableProxyPaginatedList struct {
	value *ProxyPaginatedList
	isSet bool
}

func (v NullableProxyPaginatedList) Get() *ProxyPaginatedList {
	return v.value
}

func (v *NullableProxyPaginatedList) Set(val *ProxyPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyPaginatedList(val *ProxyPaginatedList) *NullableProxyPaginatedList {
	return &NullableProxyPaginatedList{value: val, isSet: true}
}

func (v NullableProxyPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
