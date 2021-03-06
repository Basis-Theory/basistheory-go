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

// GetTenantMembers struct for GetTenantMembers
type GetTenantMembers struct {
	UserIds []string      `json:"userIds,omitempty"`
	Page    NullableInt32 `json:"page,omitempty"`
	Size    NullableInt32 `json:"size,omitempty"`
}

// NewGetTenantMembers instantiates a new GetTenantMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTenantMembers() *GetTenantMembers {
	this := GetTenantMembers{}
	return &this
}

// NewGetTenantMembersWithDefaults instantiates a new GetTenantMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTenantMembersWithDefaults() *GetTenantMembers {
	this := GetTenantMembers{}
	return &this
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantMembers) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantMembers) GetUserIdsOk() ([]string, bool) {
	if o == nil || o.UserIds == nil {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *GetTenantMembers) HasUserIds() bool {
	if o != nil && o.UserIds != nil {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *GetTenantMembers) SetUserIds(v []string) {
	o.UserIds = v
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantMembers) GetPage() int32 {
	if o == nil || o.Page.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantMembers) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *GetTenantMembers) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *GetTenantMembers) SetPage(v int32) {
	o.Page.Set(&v)
}

// SetPageNil sets the value for Page to be an explicit nil
func (o *GetTenantMembers) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *GetTenantMembers) UnsetPage() {
	o.Page.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantMembers) GetSize() int32 {
	if o == nil || o.Size.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantMembers) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *GetTenantMembers) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *GetTenantMembers) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *GetTenantMembers) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GetTenantMembers) UnsetSize() {
	o.Size.Unset()
}

func (o GetTenantMembers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetTenantMembers struct {
	value *GetTenantMembers
	isSet bool
}

func (v NullableGetTenantMembers) Get() *GetTenantMembers {
	return v.value
}

func (v *NullableGetTenantMembers) Set(val *GetTenantMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTenantMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTenantMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTenantMembers(val *GetTenantMembers) *NullableGetTenantMembers {
	return &NullableGetTenantMembers{value: val, isSet: true}
}

func (v NullableGetTenantMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTenantMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
