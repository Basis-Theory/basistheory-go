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

// checks if the ThreeDSRequestorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreeDSRequestorInfo{}

// ThreeDSRequestorInfo struct for ThreeDSRequestorInfo
type ThreeDSRequestorInfo struct {
	Id   NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Url  NullableString `json:"url,omitempty"`
}

// NewThreeDSRequestorInfo instantiates a new ThreeDSRequestorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSRequestorInfo() *ThreeDSRequestorInfo {
	this := ThreeDSRequestorInfo{}
	return &this
}

// NewThreeDSRequestorInfoWithDefaults instantiates a new ThreeDSRequestorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSRequestorInfoWithDefaults() *ThreeDSRequestorInfo {
	this := ThreeDSRequestorInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSRequestorInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSRequestorInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field is not nil.
func (o *ThreeDSRequestorInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ThreeDSRequestorInfo) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *ThreeDSRequestorInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ThreeDSRequestorInfo) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSRequestorInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSRequestorInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *ThreeDSRequestorInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ThreeDSRequestorInfo) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ThreeDSRequestorInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ThreeDSRequestorInfo) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSRequestorInfo) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSRequestorInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field is not nil.
func (o *ThreeDSRequestorInfo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ThreeDSRequestorInfo) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *ThreeDSRequestorInfo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ThreeDSRequestorInfo) UnsetUrl() {
	o.Url.Unset()
}

func (o ThreeDSRequestorInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSRequestorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableThreeDSRequestorInfo struct {
	value *ThreeDSRequestorInfo
	isSet bool
}

func (v NullableThreeDSRequestorInfo) Get() *ThreeDSRequestorInfo {
	return v.value
}

func (v *NullableThreeDSRequestorInfo) Set(val *ThreeDSRequestorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSRequestorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSRequestorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSRequestorInfo(val *ThreeDSRequestorInfo) *NullableThreeDSRequestorInfo {
	return &NullableThreeDSRequestorInfo{value: val, isSet: true}
}

func (v NullableThreeDSRequestorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSRequestorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
