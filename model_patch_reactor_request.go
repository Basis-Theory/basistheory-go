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

// checks if the PatchReactorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchReactorRequest{}

// PatchReactorRequest struct for PatchReactorRequest
type PatchReactorRequest struct {
	Name          NullableString    `json:"name,omitempty"`
	Application   *Application      `json:"application,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
}

// NewPatchReactorRequest instantiates a new PatchReactorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchReactorRequest() *PatchReactorRequest {
	this := PatchReactorRequest{}
	return &this
}

// NewPatchReactorRequestWithDefaults instantiates a new PatchReactorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchReactorRequestWithDefaults() *PatchReactorRequest {
	this := PatchReactorRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchReactorRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchReactorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *PatchReactorRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchReactorRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchReactorRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchReactorRequest) UnsetName() {
	o.Name.Unset()
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *PatchReactorRequest) GetApplication() Application {
	if o == nil || IsNil(o.Application) {
		var ret Application
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchReactorRequest) GetApplicationOk() (*Application, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field is not nil.
func (o *PatchReactorRequest) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given Application and assigns it to the Application field.
func (o *PatchReactorRequest) SetApplication(v Application) {
	o.Application = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchReactorRequest) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchReactorRequest) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field is not nil.
func (o *PatchReactorRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *PatchReactorRequest) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

func (o PatchReactorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchReactorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullablePatchReactorRequest struct {
	value *PatchReactorRequest
	isSet bool
}

func (v NullablePatchReactorRequest) Get() *PatchReactorRequest {
	return v.value
}

func (v *NullablePatchReactorRequest) Set(val *PatchReactorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchReactorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchReactorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchReactorRequest(val *PatchReactorRequest) *NullablePatchReactorRequest {
	return &NullablePatchReactorRequest{value: val, isSet: true}
}

func (v NullablePatchReactorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchReactorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}