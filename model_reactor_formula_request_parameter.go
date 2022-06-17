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

// ReactorFormulaRequestParameter struct for ReactorFormulaRequestParameter
type ReactorFormulaRequestParameter struct {
	Name        string         `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Type        string         `json:"type"`
	Optional    *bool          `json:"optional,omitempty"`
}

// NewReactorFormulaRequestParameter instantiates a new ReactorFormulaRequestParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactorFormulaRequestParameter(name string, type_ string) *ReactorFormulaRequestParameter {
	this := ReactorFormulaRequestParameter{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewReactorFormulaRequestParameterWithDefaults instantiates a new ReactorFormulaRequestParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactorFormulaRequestParameterWithDefaults() *ReactorFormulaRequestParameter {
	this := ReactorFormulaRequestParameter{}
	return &this
}

// GetName returns the Name field value
func (o *ReactorFormulaRequestParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReactorFormulaRequestParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReactorFormulaRequestParameter) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormulaRequestParameter) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormulaRequestParameter) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ReactorFormulaRequestParameter) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ReactorFormulaRequestParameter) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ReactorFormulaRequestParameter) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ReactorFormulaRequestParameter) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value
func (o *ReactorFormulaRequestParameter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReactorFormulaRequestParameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReactorFormulaRequestParameter) SetType(v string) {
	o.Type = v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *ReactorFormulaRequestParameter) GetOptional() bool {
	if o == nil || o.Optional == nil {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactorFormulaRequestParameter) GetOptionalOk() (*bool, bool) {
	if o == nil || o.Optional == nil {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *ReactorFormulaRequestParameter) HasOptional() bool {
	if o != nil && o.Optional != nil {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *ReactorFormulaRequestParameter) SetOptional(v bool) {
	o.Optional = &v
}

func (o ReactorFormulaRequestParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Optional != nil {
		toSerialize["optional"] = o.Optional
	}
	return json.Marshal(toSerialize)
}

type NullableReactorFormulaRequestParameter struct {
	value *ReactorFormulaRequestParameter
	isSet bool
}

func (v NullableReactorFormulaRequestParameter) Get() *ReactorFormulaRequestParameter {
	return v.value
}

func (v *NullableReactorFormulaRequestParameter) Set(val *ReactorFormulaRequestParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableReactorFormulaRequestParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableReactorFormulaRequestParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactorFormulaRequestParameter(val *ReactorFormulaRequestParameter) *NullableReactorFormulaRequestParameter {
	return &NullableReactorFormulaRequestParameter{value: val, isSet: true}
}

func (v NullableReactorFormulaRequestParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactorFormulaRequestParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
