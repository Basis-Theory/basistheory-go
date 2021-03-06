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

// GetTenantInvitations struct for GetTenantInvitations
type GetTenantInvitations struct {
	Status *TenantInvitationStatus `json:"status,omitempty"`
	Page   NullableInt32           `json:"page,omitempty"`
	Size   NullableInt32           `json:"size,omitempty"`
}

// NewGetTenantInvitations instantiates a new GetTenantInvitations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTenantInvitations() *GetTenantInvitations {
	this := GetTenantInvitations{}
	return &this
}

// NewGetTenantInvitationsWithDefaults instantiates a new GetTenantInvitations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTenantInvitationsWithDefaults() *GetTenantInvitations {
	this := GetTenantInvitations{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetTenantInvitations) GetStatus() TenantInvitationStatus {
	if o == nil || o.Status == nil {
		var ret TenantInvitationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTenantInvitations) GetStatusOk() (*TenantInvitationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetTenantInvitations) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TenantInvitationStatus and assigns it to the Status field.
func (o *GetTenantInvitations) SetStatus(v TenantInvitationStatus) {
	o.Status = &v
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantInvitations) GetPage() int32 {
	if o == nil || o.Page.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantInvitations) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *GetTenantInvitations) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *GetTenantInvitations) SetPage(v int32) {
	o.Page.Set(&v)
}

// SetPageNil sets the value for Page to be an explicit nil
func (o *GetTenantInvitations) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *GetTenantInvitations) UnsetPage() {
	o.Page.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTenantInvitations) GetSize() int32 {
	if o == nil || o.Size.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTenantInvitations) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *GetTenantInvitations) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *GetTenantInvitations) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *GetTenantInvitations) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GetTenantInvitations) UnsetSize() {
	o.Size.Unset()
}

func (o GetTenantInvitations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGetTenantInvitations struct {
	value *GetTenantInvitations
	isSet bool
}

func (v NullableGetTenantInvitations) Get() *GetTenantInvitations {
	return v.value
}

func (v *NullableGetTenantInvitations) Set(val *GetTenantInvitations) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTenantInvitations) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTenantInvitations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTenantInvitations(val *GetTenantInvitations) *NullableGetTenantInvitations {
	return &NullableGetTenantInvitations{value: val, isSet: true}
}

func (v NullableGetTenantInvitations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTenantInvitations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
