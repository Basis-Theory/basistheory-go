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

// CreateReactorRequest struct for CreateReactorRequest
type CreateReactorRequest struct {
	Name string `json:"name"`
	Formula *ReactorFormula `json:"formula,omitempty"`
	Application *Application `json:"application,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
}

// NewCreateReactorRequest instantiates a new CreateReactorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReactorRequest(name string) *CreateReactorRequest {
	this := CreateReactorRequest{}
	this.Name = name
	return &this
}

// NewCreateReactorRequestWithDefaults instantiates a new CreateReactorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReactorRequestWithDefaults() *CreateReactorRequest {
	this := CreateReactorRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateReactorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateReactorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateReactorRequest) SetName(v string) {
	o.Name = v
}

// GetFormula returns the Formula field value if set, zero value otherwise.
func (o *CreateReactorRequest) GetFormula() ReactorFormula {
	if o == nil || o.Formula == nil {
		var ret ReactorFormula
		return ret
	}
	return *o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReactorRequest) GetFormulaOk() (*ReactorFormula, bool) {
	if o == nil || o.Formula == nil {
		return nil, false
	}
	return o.Formula, true
}

// HasFormula returns a boolean if a field has been set.
func (o *CreateReactorRequest) HasFormula() bool {
	if o != nil && o.Formula != nil {
		return true
	}

	return false
}

// SetFormula gets a reference to the given ReactorFormula and assigns it to the Formula field.
func (o *CreateReactorRequest) SetFormula(v ReactorFormula) {
	o.Formula = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *CreateReactorRequest) GetApplication() Application {
	if o == nil || o.Application == nil {
		var ret Application
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateReactorRequest) GetApplicationOk() (*Application, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *CreateReactorRequest) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given Application and assigns it to the Application field.
func (o *CreateReactorRequest) SetApplication(v Application) {
	o.Application = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateReactorRequest) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReactorRequest) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *CreateReactorRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *CreateReactorRequest) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

func (o CreateReactorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Formula != nil {
		toSerialize["formula"] = o.Formula
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReactorRequest struct {
	value *CreateReactorRequest
	isSet bool
}

func (v NullableCreateReactorRequest) Get() *CreateReactorRequest {
	return v.value
}

func (v *NullableCreateReactorRequest) Set(val *CreateReactorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReactorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReactorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReactorRequest(val *CreateReactorRequest) *NullableCreateReactorRequest {
	return &NullableCreateReactorRequest{value: val, isSet: true}
}

func (v NullableCreateReactorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReactorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


