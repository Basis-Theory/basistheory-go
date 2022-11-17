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

// ReactRequest struct for ReactRequest
type ReactRequest struct {
	Args interface{} `json:"args,omitempty"`
}

// NewReactRequest instantiates a new ReactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactRequest() *ReactRequest {
	this := ReactRequest{}
	return &this
}

// NewReactRequestWithDefaults instantiates a new ReactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactRequestWithDefaults() *ReactRequest {
	this := ReactRequest{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactRequest) GetArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactRequest) GetArgsOk() (*interface{}, bool) {
	if o == nil || isNil(o.Args) {
		return nil, false
	}
	return &o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ReactRequest) HasArgs() bool {
	if o != nil && isNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given interface{} and assigns it to the Args field.
func (o *ReactRequest) SetArgs(v interface{}) {
	o.Args = v
}

func (o ReactRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	return json.Marshal(toSerialize)
}

type NullableReactRequest struct {
	value *ReactRequest
	isSet bool
}

func (v NullableReactRequest) Get() *ReactRequest {
	return v.value
}

func (v *NullableReactRequest) Set(val *ReactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactRequest(val *ReactRequest) *NullableReactRequest {
	return &NullableReactRequest{value: val, isSet: true}
}

func (v NullableReactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
