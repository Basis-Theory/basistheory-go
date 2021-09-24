/*
BasisTheory Vault API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
)

// UpdateReactorModel struct for UpdateReactorModel
type UpdateReactorModel struct {
	Name NullableString `json:"name,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
}

// NewUpdateReactorModel instantiates a new UpdateReactorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReactorModel() *UpdateReactorModel {
	this := UpdateReactorModel{}
	return &this
}

// NewUpdateReactorModelWithDefaults instantiates a new UpdateReactorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReactorModelWithDefaults() *UpdateReactorModel {
	this := UpdateReactorModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateReactorModel) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateReactorModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateReactorModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateReactorModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateReactorModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateReactorModel) UnsetName() {
	o.Name.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateReactorModel) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateReactorModel) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UpdateReactorModel) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *UpdateReactorModel) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

func (o UpdateReactorModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReactorModel struct {
	value *UpdateReactorModel
	isSet bool
}

func (v NullableUpdateReactorModel) Get() *UpdateReactorModel {
	return v.value
}

func (v *NullableUpdateReactorModel) Set(val *UpdateReactorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReactorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReactorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReactorModel(val *UpdateReactorModel) *NullableUpdateReactorModel {
	return &NullableUpdateReactorModel{value: val, isSet: true}
}

func (v NullableUpdateReactorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReactorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


