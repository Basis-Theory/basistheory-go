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

// AccessRule struct for AccessRule
type AccessRule struct {
	Description NullableString `json:"description,omitempty"`
	Priority    NullableInt32  `json:"priority,omitempty"`
	Container   NullableString `json:"container,omitempty"`
	Transform   NullableString `json:"transform,omitempty"`
	Conditions  []Condition    `json:"conditions,omitempty"`
	Permissions []string       `json:"permissions,omitempty"`
}

// NewAccessRule instantiates a new AccessRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRule() *AccessRule {
	this := AccessRule{}
	return &this
}

// NewAccessRuleWithDefaults instantiates a new AccessRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRuleWithDefaults() *AccessRule {
	this := AccessRule{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRule) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessRule) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessRule) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessRule) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessRule) UnsetDescription() {
	o.Description.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRule) GetPriority() int32 {
	if o == nil || isNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRule) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *AccessRule) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *AccessRule) SetPriority(v int32) {
	o.Priority.Set(&v)
}

// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *AccessRule) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *AccessRule) UnsetPriority() {
	o.Priority.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRule) GetContainer() string {
	if o == nil || isNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRule) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *AccessRule) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *AccessRule) SetContainer(v string) {
	o.Container.Set(&v)
}

// SetContainerNil sets the value for Container to be an explicit nil
func (o *AccessRule) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *AccessRule) UnsetContainer() {
	o.Container.Unset()
}

// GetTransform returns the Transform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRule) GetTransform() string {
	if o == nil || isNil(o.Transform.Get()) {
		var ret string
		return ret
	}
	return *o.Transform.Get()
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRule) GetTransformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transform.Get(), o.Transform.IsSet()
}

// HasTransform returns a boolean if a field has been set.
func (o *AccessRule) HasTransform() bool {
	if o != nil && o.Transform.IsSet() {
		return true
	}

	return false
}

// SetTransform gets a reference to the given NullableString and assigns it to the Transform field.
func (o *AccessRule) SetTransform(v string) {
	o.Transform.Set(&v)
}

// SetTransformNil sets the value for Transform to be an explicit nil
func (o *AccessRule) SetTransformNil() {
	o.Transform.Set(nil)
}

// UnsetTransform ensures that no value is present for Transform, not even an explicit nil
func (o *AccessRule) UnsetTransform() {
	o.Transform.Unset()
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRule) GetConditions() []Condition {
	if o == nil {
		var ret []Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRule) GetConditionsOk() ([]Condition, bool) {
	if o == nil || isNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AccessRule) HasConditions() bool {
	if o != nil && isNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []Condition and assigns it to the Conditions field.
func (o *AccessRule) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRule) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRule) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AccessRule) HasPermissions() bool {
	if o != nil && isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *AccessRule) SetPermissions(v []string) {
	o.Permissions = v
}

func (o AccessRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.Container.IsSet() {
		toSerialize["container"] = o.Container.Get()
	}
	if o.Transform.IsSet() {
		toSerialize["transform"] = o.Transform.Get()
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAccessRule struct {
	value *AccessRule
	isSet bool
}

func (v NullableAccessRule) Get() *AccessRule {
	return v.value
}

func (v *NullableAccessRule) Set(val *AccessRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRule(val *AccessRule) *NullableAccessRule {
	return &NullableAccessRule{value: val, isSet: true}
}

func (v NullableAccessRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
