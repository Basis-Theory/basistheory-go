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

// UpdateReactorRequest struct for UpdateReactorRequest
type UpdateReactorRequest struct {
	Name string `json:"name"`
	Application *Application `json:"application,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
}

// NewUpdateReactorRequest instantiates a new UpdateReactorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReactorRequest(name string) *UpdateReactorRequest {
	this := UpdateReactorRequest{}
	this.Name = name
	return &this
}

// NewUpdateReactorRequestWithDefaults instantiates a new UpdateReactorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReactorRequestWithDefaults() *UpdateReactorRequest {
	this := UpdateReactorRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateReactorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateReactorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateReactorRequest) SetName(v string) {
	o.Name = v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *UpdateReactorRequest) GetApplication() Application {
	if o == nil || o.Application == nil {
		var ret Application
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReactorRequest) GetApplicationOk() (*Application, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *UpdateReactorRequest) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given Application and assigns it to the Application field.
func (o *UpdateReactorRequest) SetApplication(v Application) {
	o.Application = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateReactorRequest) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateReactorRequest) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UpdateReactorRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *UpdateReactorRequest) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

func (o UpdateReactorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReactorRequest struct {
	value *UpdateReactorRequest
	isSet bool
}

func (v NullableUpdateReactorRequest) Get() *UpdateReactorRequest {
	return v.value
}

func (v *NullableUpdateReactorRequest) Set(val *UpdateReactorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReactorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReactorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReactorRequest(val *UpdateReactorRequest) *NullableUpdateReactorRequest {
	return &NullableUpdateReactorRequest{value: val, isSet: true}
}

func (v NullableUpdateReactorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReactorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


