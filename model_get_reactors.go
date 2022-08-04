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

// GetReactors struct for GetReactors
type GetReactors struct {
	ReactorIds []string       `json:"reactorIds,omitempty"`
	Name       NullableString `json:"name,omitempty"`
	Page       NullableInt32  `json:"page,omitempty"`
	Size       NullableInt32  `json:"size,omitempty"`
}

// NewGetReactors instantiates a new GetReactors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReactors() *GetReactors {
	this := GetReactors{}
	return &this
}

// NewGetReactorsWithDefaults instantiates a new GetReactors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReactorsWithDefaults() *GetReactors {
	this := GetReactors{}
	return &this
}

// GetReactorIds returns the ReactorIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactors) GetReactorIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ReactorIds
}

// GetReactorIdsOk returns a tuple with the ReactorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactors) GetReactorIdsOk() ([]string, bool) {
	if o == nil || o.ReactorIds == nil {
		return nil, false
	}
	return o.ReactorIds, true
}

// HasReactorIds returns a boolean if a field has been set.
func (o *GetReactors) HasReactorIds() bool {
	if o != nil && o.ReactorIds != nil {
		return true
	}

	return false
}

// SetReactorIds gets a reference to the given []string and assigns it to the ReactorIds field.
func (o *GetReactors) SetReactorIds(v []string) {
	o.ReactorIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactors) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactors) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GetReactors) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GetReactors) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *GetReactors) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GetReactors) UnsetName() {
	o.Name.Unset()
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactors) GetPage() int32 {
	if o == nil || o.Page.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactors) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *GetReactors) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *GetReactors) SetPage(v int32) {
	o.Page.Set(&v)
}

// SetPageNil sets the value for Page to be an explicit nil
func (o *GetReactors) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *GetReactors) UnsetPage() {
	o.Page.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactors) GetSize() int32 {
	if o == nil || o.Size.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactors) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *GetReactors) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *GetReactors) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *GetReactors) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GetReactors) UnsetSize() {
	o.Size.Unset()
}

func (o GetReactors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReactorIds != nil {
		toSerialize["reactorIds"] = o.ReactorIds
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetReactors struct {
	value *GetReactors
	isSet bool
}

func (v NullableGetReactors) Get() *GetReactors {
	return v.value
}

func (v *NullableGetReactors) Set(val *GetReactors) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReactors) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReactors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReactors(val *GetReactors) *NullableGetReactors {
	return &NullableGetReactors{value: val, isSet: true}
}

func (v NullableGetReactors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReactors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}