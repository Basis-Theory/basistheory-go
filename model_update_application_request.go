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

// UpdateApplicationRequest struct for UpdateApplicationRequest
type UpdateApplicationRequest struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

// NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApplicationRequest(name string) *UpdateApplicationRequest {
	this := UpdateApplicationRequest{}
	this.Name = name
	return &this
}

// NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest {
	this := UpdateApplicationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateApplicationRequest) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApplicationRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApplicationRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UpdateApplicationRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *UpdateApplicationRequest) SetPermissions(v []string) {
	o.Permissions = v
}

func (o UpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateApplicationRequest struct {
	value *UpdateApplicationRequest
	isSet bool
}

func (v NullableUpdateApplicationRequest) Get() *UpdateApplicationRequest {
	return v.value
}

func (v *NullableUpdateApplicationRequest) Set(val *UpdateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationRequest(val *UpdateApplicationRequest) *NullableUpdateApplicationRequest {
	return &NullableUpdateApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
