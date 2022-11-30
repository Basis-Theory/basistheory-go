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

// ReactResponse struct for ReactResponse
type ReactResponse struct {
	Tokens  interface{} `json:"tokens,omitempty"`
	Raw     interface{} `json:"raw,omitempty"`
	Body    interface{} `json:"body,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
}

// NewReactResponse instantiates a new ReactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactResponse() *ReactResponse {
	this := ReactResponse{}
	return &this
}

// NewReactResponseWithDefaults instantiates a new ReactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactResponseWithDefaults() *ReactResponse {
	this := ReactResponse{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactResponse) GetTokens() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactResponse) GetTokensOk() (*interface{}, bool) {
	if o == nil || isNil(o.Tokens) {
		return nil, false
	}
	return &o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *ReactResponse) HasTokens() bool {
	if o != nil && isNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given interface{} and assigns it to the Tokens field.
func (o *ReactResponse) SetTokens(v interface{}) {
	o.Tokens = v
}

// GetRaw returns the Raw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactResponse) GetRaw() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactResponse) GetRawOk() (*interface{}, bool) {
	if o == nil || isNil(o.Raw) {
		return nil, false
	}
	return &o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *ReactResponse) HasRaw() bool {
	if o != nil && isNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given interface{} and assigns it to the Raw field.
func (o *ReactResponse) SetRaw(v interface{}) {
	o.Raw = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactResponse) GetBody() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactResponse) GetBodyOk() (*interface{}, bool) {
	if o == nil || isNil(o.Body) {
		return nil, false
	}
	return &o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ReactResponse) HasBody() bool {
	if o != nil && isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given interface{} and assigns it to the Body field.
func (o *ReactResponse) SetBody(v interface{}) {
	o.Body = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactResponse) GetHeaders() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactResponse) GetHeadersOk() (*interface{}, bool) {
	if o == nil || isNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ReactResponse) HasHeaders() bool {
	if o != nil && isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given interface{} and assigns it to the Headers field.
func (o *ReactResponse) SetHeaders(v interface{}) {
	o.Headers = v
}

func (o ReactResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	return json.Marshal(toSerialize)
}

type NullableReactResponse struct {
	value *ReactResponse
	isSet bool
}

func (v NullableReactResponse) Get() *ReactResponse {
	return v.value
}

func (v *NullableReactResponse) Set(val *ReactResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReactResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactResponse(val *ReactResponse) *NullableReactResponse {
	return &NullableReactResponse{value: val, isSet: true}
}

func (v NullableReactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
