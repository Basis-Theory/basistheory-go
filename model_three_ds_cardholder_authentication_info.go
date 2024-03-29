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

// checks if the ThreeDSCardholderAuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreeDSCardholderAuthenticationInfo{}

// ThreeDSCardholderAuthenticationInfo struct for ThreeDSCardholderAuthenticationInfo
type ThreeDSCardholderAuthenticationInfo struct {
	Method    NullableString `json:"method,omitempty"`
	Timestamp NullableString `json:"timestamp,omitempty"`
	Data      NullableString `json:"data,omitempty"`
}

// NewThreeDSCardholderAuthenticationInfo instantiates a new ThreeDSCardholderAuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSCardholderAuthenticationInfo() *ThreeDSCardholderAuthenticationInfo {
	this := ThreeDSCardholderAuthenticationInfo{}
	return &this
}

// NewThreeDSCardholderAuthenticationInfoWithDefaults instantiates a new ThreeDSCardholderAuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSCardholderAuthenticationInfoWithDefaults() *ThreeDSCardholderAuthenticationInfo {
	this := ThreeDSCardholderAuthenticationInfo{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSCardholderAuthenticationInfo) GetMethod() string {
	if o == nil || IsNil(o.Method.Get()) {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSCardholderAuthenticationInfo) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field is not nil.
func (o *ThreeDSCardholderAuthenticationInfo) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *ThreeDSCardholderAuthenticationInfo) SetMethod(v string) {
	o.Method.Set(&v)
}

// SetMethodNil sets the value for Method to be an explicit nil
func (o *ThreeDSCardholderAuthenticationInfo) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *ThreeDSCardholderAuthenticationInfo) UnsetMethod() {
	o.Method.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSCardholderAuthenticationInfo) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret string
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSCardholderAuthenticationInfo) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field is not nil.
func (o *ThreeDSCardholderAuthenticationInfo) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableString and assigns it to the Timestamp field.
func (o *ThreeDSCardholderAuthenticationInfo) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *ThreeDSCardholderAuthenticationInfo) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *ThreeDSCardholderAuthenticationInfo) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSCardholderAuthenticationInfo) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSCardholderAuthenticationInfo) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field is not nil.
func (o *ThreeDSCardholderAuthenticationInfo) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *ThreeDSCardholderAuthenticationInfo) SetData(v string) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil
func (o *ThreeDSCardholderAuthenticationInfo) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ThreeDSCardholderAuthenticationInfo) UnsetData() {
	o.Data.Unset()
}

func (o ThreeDSCardholderAuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSCardholderAuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Method.IsSet() {
		toSerialize["method"] = o.Method.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableThreeDSCardholderAuthenticationInfo struct {
	value *ThreeDSCardholderAuthenticationInfo
	isSet bool
}

func (v NullableThreeDSCardholderAuthenticationInfo) Get() *ThreeDSCardholderAuthenticationInfo {
	return v.value
}

func (v *NullableThreeDSCardholderAuthenticationInfo) Set(val *ThreeDSCardholderAuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSCardholderAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSCardholderAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSCardholderAuthenticationInfo(val *ThreeDSCardholderAuthenticationInfo) *NullableThreeDSCardholderAuthenticationInfo {
	return &NullableThreeDSCardholderAuthenticationInfo{value: val, isSet: true}
}

func (v NullableThreeDSCardholderAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSCardholderAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
