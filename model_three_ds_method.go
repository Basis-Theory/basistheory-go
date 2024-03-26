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

// checks if the ThreeDSMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreeDSMethod{}

// ThreeDSMethod struct for ThreeDSMethod
type ThreeDSMethod struct {
	MethodUrl                 NullableString `json:"method_url,omitempty"`
	MethodCompletionIndicator NullableString `json:"method_completion_indicator,omitempty"`
}

// NewThreeDSMethod instantiates a new ThreeDSMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSMethod() *ThreeDSMethod {
	this := ThreeDSMethod{}
	return &this
}

// NewThreeDSMethodWithDefaults instantiates a new ThreeDSMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSMethodWithDefaults() *ThreeDSMethod {
	this := ThreeDSMethod{}
	return &this
}

// GetMethodUrl returns the MethodUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMethod) GetMethodUrl() string {
	if o == nil || IsNil(o.MethodUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MethodUrl.Get()
}

// GetMethodUrlOk returns a tuple with the MethodUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMethod) GetMethodUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MethodUrl.Get(), o.MethodUrl.IsSet()
}

// HasMethodUrl returns a boolean if a field is not nil.
func (o *ThreeDSMethod) HasMethodUrl() bool {
	if o != nil && !IsNil(o.MethodUrl) {
		return true
	}

	return false
}

// SetMethodUrl gets a reference to the given NullableString and assigns it to the MethodUrl field.
func (o *ThreeDSMethod) SetMethodUrl(v string) {
	o.MethodUrl.Set(&v)
}

// SetMethodUrlNil sets the value for MethodUrl to be an explicit nil
func (o *ThreeDSMethod) SetMethodUrlNil() {
	o.MethodUrl.Set(nil)
}

// UnsetMethodUrl ensures that no value is present for MethodUrl, not even an explicit nil
func (o *ThreeDSMethod) UnsetMethodUrl() {
	o.MethodUrl.Unset()
}

// GetMethodCompletionIndicator returns the MethodCompletionIndicator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMethod) GetMethodCompletionIndicator() string {
	if o == nil || IsNil(o.MethodCompletionIndicator.Get()) {
		var ret string
		return ret
	}
	return *o.MethodCompletionIndicator.Get()
}

// GetMethodCompletionIndicatorOk returns a tuple with the MethodCompletionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMethod) GetMethodCompletionIndicatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MethodCompletionIndicator.Get(), o.MethodCompletionIndicator.IsSet()
}

// HasMethodCompletionIndicator returns a boolean if a field is not nil.
func (o *ThreeDSMethod) HasMethodCompletionIndicator() bool {
	if o != nil && !IsNil(o.MethodCompletionIndicator) {
		return true
	}

	return false
}

// SetMethodCompletionIndicator gets a reference to the given NullableString and assigns it to the MethodCompletionIndicator field.
func (o *ThreeDSMethod) SetMethodCompletionIndicator(v string) {
	o.MethodCompletionIndicator.Set(&v)
}

// SetMethodCompletionIndicatorNil sets the value for MethodCompletionIndicator to be an explicit nil
func (o *ThreeDSMethod) SetMethodCompletionIndicatorNil() {
	o.MethodCompletionIndicator.Set(nil)
}

// UnsetMethodCompletionIndicator ensures that no value is present for MethodCompletionIndicator, not even an explicit nil
func (o *ThreeDSMethod) UnsetMethodCompletionIndicator() {
	o.MethodCompletionIndicator.Unset()
}

func (o ThreeDSMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MethodUrl.IsSet() {
		toSerialize["method_url"] = o.MethodUrl.Get()
	}
	if o.MethodCompletionIndicator.IsSet() {
		toSerialize["method_completion_indicator"] = o.MethodCompletionIndicator.Get()
	}
	return toSerialize, nil
}

type NullableThreeDSMethod struct {
	value *ThreeDSMethod
	isSet bool
}

func (v NullableThreeDSMethod) Get() *ThreeDSMethod {
	return v.value
}

func (v *NullableThreeDSMethod) Set(val *ThreeDSMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSMethod(val *ThreeDSMethod) *NullableThreeDSMethod {
	return &NullableThreeDSMethod{value: val, isSet: true}
}

func (v NullableThreeDSMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
