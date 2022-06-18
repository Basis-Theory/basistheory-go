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

// AtomicBankPaginatedList struct for AtomicBankPaginatedList
type AtomicBankPaginatedList struct {
	Pagination *Pagination  `json:"pagination,omitempty"`
	Data       []AtomicBank `json:"data,omitempty"`
}

// NewAtomicBankPaginatedList instantiates a new AtomicBankPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtomicBankPaginatedList() *AtomicBankPaginatedList {
	this := AtomicBankPaginatedList{}
	return &this
}

// NewAtomicBankPaginatedListWithDefaults instantiates a new AtomicBankPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtomicBankPaginatedListWithDefaults() *AtomicBankPaginatedList {
	this := AtomicBankPaginatedList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AtomicBankPaginatedList) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtomicBankPaginatedList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AtomicBankPaginatedList) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *AtomicBankPaginatedList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AtomicBankPaginatedList) GetData() []AtomicBank {
	if o == nil {
		var ret []AtomicBank
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AtomicBankPaginatedList) GetDataOk() ([]AtomicBank, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AtomicBankPaginatedList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AtomicBank and assigns it to the Data field.
func (o *AtomicBankPaginatedList) SetData(v []AtomicBank) {
	o.Data = v
}

func (o AtomicBankPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAtomicBankPaginatedList struct {
	value *AtomicBankPaginatedList
	isSet bool
}

func (v NullableAtomicBankPaginatedList) Get() *AtomicBankPaginatedList {
	return v.value
}

func (v *NullableAtomicBankPaginatedList) Set(val *AtomicBankPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableAtomicBankPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableAtomicBankPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtomicBankPaginatedList(val *AtomicBankPaginatedList) *NullableAtomicBankPaginatedList {
	return &NullableAtomicBankPaginatedList{value: val, isSet: true}
}

func (v NullableAtomicBankPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtomicBankPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
