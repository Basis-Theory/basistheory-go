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

// UpdateTenantRequest struct for UpdateTenantRequest
type UpdateTenantRequest struct {
	Name     string            `json:"name"`
	Settings map[string]string `json:"settings,omitempty"`
}

// NewUpdateTenantRequest instantiates a new UpdateTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenantRequest(name string) *UpdateTenantRequest {
	this := UpdateTenantRequest{}
	this.Name = name
	return &this
}

// NewUpdateTenantRequestWithDefaults instantiates a new UpdateTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenantRequestWithDefaults() *UpdateTenantRequest {
	this := UpdateTenantRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateTenantRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateTenantRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateTenantRequest) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateTenantRequest) GetSettings() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateTenantRequest) GetSettingsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Settings) {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *UpdateTenantRequest) HasSettings() bool {
	if o != nil && isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]string and assigns it to the Settings field.
func (o *UpdateTenantRequest) SetSettings(v map[string]string) {
	o.Settings = v
}

func (o UpdateTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTenantRequest struct {
	value *UpdateTenantRequest
	isSet bool
}

func (v NullableUpdateTenantRequest) Get() *UpdateTenantRequest {
	return v.value
}

func (v *NullableUpdateTenantRequest) Set(val *UpdateTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenantRequest(val *UpdateTenantRequest) *NullableUpdateTenantRequest {
	return &NullableUpdateTenantRequest{value: val, isSet: true}
}

func (v NullableUpdateTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
