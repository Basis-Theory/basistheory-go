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

// checks if the ReactorPaginatedList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactorPaginatedList{}

// ReactorPaginatedList struct for ReactorPaginatedList
type ReactorPaginatedList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []Reactor   `json:"data,omitempty"`
}

// NewReactorPaginatedList instantiates a new ReactorPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactorPaginatedList() *ReactorPaginatedList {
	this := ReactorPaginatedList{}
	return &this
}

// NewReactorPaginatedListWithDefaults instantiates a new ReactorPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactorPaginatedListWithDefaults() *ReactorPaginatedList {
	this := ReactorPaginatedList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ReactorPaginatedList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactorPaginatedList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field is not nil.
func (o *ReactorPaginatedList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ReactorPaginatedList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorPaginatedList) GetData() []Reactor {
	if o == nil {
		var ret []Reactor
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorPaginatedList) GetDataOk() ([]Reactor, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field is not nil.
func (o *ReactorPaginatedList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Reactor and assigns it to the Data field.
func (o *ReactorPaginatedList) SetData(v []Reactor) {
	o.Data = v
}

func (o ReactorPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactorPaginatedList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableReactorPaginatedList struct {
	value *ReactorPaginatedList
	isSet bool
}

func (v NullableReactorPaginatedList) Get() *ReactorPaginatedList {
	return v.value
}

func (v *NullableReactorPaginatedList) Set(val *ReactorPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableReactorPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableReactorPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactorPaginatedList(val *ReactorPaginatedList) *NullableReactorPaginatedList {
	return &NullableReactorPaginatedList{value: val, isSet: true}
}

func (v NullableReactorPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactorPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
