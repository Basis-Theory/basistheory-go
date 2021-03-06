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

// UpdatePrivacy struct for UpdatePrivacy
type UpdatePrivacy struct {
	ImpactLevel       NullableString `json:"impact_level,omitempty"`
	RestrictionPolicy NullableString `json:"restriction_policy,omitempty"`
}

// NewUpdatePrivacy instantiates a new UpdatePrivacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePrivacy() *UpdatePrivacy {
	this := UpdatePrivacy{}
	return &this
}

// NewUpdatePrivacyWithDefaults instantiates a new UpdatePrivacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePrivacyWithDefaults() *UpdatePrivacy {
	this := UpdatePrivacy{}
	return &this
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePrivacy) GetImpactLevel() string {
	if o == nil || o.ImpactLevel.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImpactLevel.Get()
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePrivacy) GetImpactLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImpactLevel.Get(), o.ImpactLevel.IsSet()
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *UpdatePrivacy) HasImpactLevel() bool {
	if o != nil && o.ImpactLevel.IsSet() {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given NullableString and assigns it to the ImpactLevel field.
func (o *UpdatePrivacy) SetImpactLevel(v string) {
	o.ImpactLevel.Set(&v)
}

// SetImpactLevelNil sets the value for ImpactLevel to be an explicit nil
func (o *UpdatePrivacy) SetImpactLevelNil() {
	o.ImpactLevel.Set(nil)
}

// UnsetImpactLevel ensures that no value is present for ImpactLevel, not even an explicit nil
func (o *UpdatePrivacy) UnsetImpactLevel() {
	o.ImpactLevel.Unset()
}

// GetRestrictionPolicy returns the RestrictionPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePrivacy) GetRestrictionPolicy() string {
	if o == nil || o.RestrictionPolicy.Get() == nil {
		var ret string
		return ret
	}
	return *o.RestrictionPolicy.Get()
}

// GetRestrictionPolicyOk returns a tuple with the RestrictionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePrivacy) GetRestrictionPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestrictionPolicy.Get(), o.RestrictionPolicy.IsSet()
}

// HasRestrictionPolicy returns a boolean if a field has been set.
func (o *UpdatePrivacy) HasRestrictionPolicy() bool {
	if o != nil && o.RestrictionPolicy.IsSet() {
		return true
	}

	return false
}

// SetRestrictionPolicy gets a reference to the given NullableString and assigns it to the RestrictionPolicy field.
func (o *UpdatePrivacy) SetRestrictionPolicy(v string) {
	o.RestrictionPolicy.Set(&v)
}

// SetRestrictionPolicyNil sets the value for RestrictionPolicy to be an explicit nil
func (o *UpdatePrivacy) SetRestrictionPolicyNil() {
	o.RestrictionPolicy.Set(nil)
}

// UnsetRestrictionPolicy ensures that no value is present for RestrictionPolicy, not even an explicit nil
func (o *UpdatePrivacy) UnsetRestrictionPolicy() {
	o.RestrictionPolicy.Unset()
}

func (o UpdatePrivacy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImpactLevel.IsSet() {
		toSerialize["impact_level"] = o.ImpactLevel.Get()
	}
	if o.RestrictionPolicy.IsSet() {
		toSerialize["restriction_policy"] = o.RestrictionPolicy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePrivacy struct {
	value *UpdatePrivacy
	isSet bool
}

func (v NullableUpdatePrivacy) Get() *UpdatePrivacy {
	return v.value
}

func (v *NullableUpdatePrivacy) Set(val *UpdatePrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePrivacy(val *UpdatePrivacy) *NullableUpdatePrivacy {
	return &NullableUpdatePrivacy{value: val, isSet: true}
}

func (v NullableUpdatePrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
