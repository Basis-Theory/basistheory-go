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

// TenantMemberResponsePaginatedList struct for TenantMemberResponsePaginatedList
type TenantMemberResponsePaginatedList struct {
	Pagination *Pagination            `json:"pagination,omitempty"`
	Data       []TenantMemberResponse `json:"data,omitempty"`
}

// NewTenantMemberResponsePaginatedList instantiates a new TenantMemberResponsePaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantMemberResponsePaginatedList() *TenantMemberResponsePaginatedList {
	this := TenantMemberResponsePaginatedList{}
	return &this
}

// NewTenantMemberResponsePaginatedListWithDefaults instantiates a new TenantMemberResponsePaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantMemberResponsePaginatedListWithDefaults() *TenantMemberResponsePaginatedList {
	this := TenantMemberResponsePaginatedList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TenantMemberResponsePaginatedList) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantMemberResponsePaginatedList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TenantMemberResponsePaginatedList) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *TenantMemberResponsePaginatedList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantMemberResponsePaginatedList) GetData() []TenantMemberResponse {
	if o == nil {
		var ret []TenantMemberResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantMemberResponsePaginatedList) GetDataOk() ([]TenantMemberResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TenantMemberResponsePaginatedList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TenantMemberResponse and assigns it to the Data field.
func (o *TenantMemberResponsePaginatedList) SetData(v []TenantMemberResponse) {
	o.Data = v
}

func (o TenantMemberResponsePaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTenantMemberResponsePaginatedList struct {
	value *TenantMemberResponsePaginatedList
	isSet bool
}

func (v NullableTenantMemberResponsePaginatedList) Get() *TenantMemberResponsePaginatedList {
	return v.value
}

func (v *NullableTenantMemberResponsePaginatedList) Set(val *TenantMemberResponsePaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantMemberResponsePaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantMemberResponsePaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantMemberResponsePaginatedList(val *TenantMemberResponsePaginatedList) *NullableTenantMemberResponsePaginatedList {
	return &NullableTenantMemberResponsePaginatedList{value: val, isSet: true}
}

func (v NullableTenantMemberResponsePaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantMemberResponsePaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
