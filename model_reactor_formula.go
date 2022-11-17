/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"time"
)

// ReactorFormula struct for ReactorFormula
type ReactorFormula struct {
	Id                *string                          `json:"id,omitempty"`
	Type              NullableString                   `json:"type,omitempty"`
	Status            NullableString                   `json:"status,omitempty"`
	Name              NullableString                   `json:"name,omitempty"`
	Description       NullableString                   `json:"description,omitempty"`
	Icon              NullableString                   `json:"icon,omitempty"`
	Code              NullableString                   `json:"code,omitempty"`
	CreatedBy         NullableString                   `json:"created_by,omitempty"`
	CreatedAt         NullableTime                     `json:"created_at,omitempty"`
	ModifiedBy        NullableString                   `json:"modified_by,omitempty"`
	ModifiedAt        NullableTime                     `json:"modified_at,omitempty"`
	Configuration     []ReactorFormulaConfiguration    `json:"configuration,omitempty"`
	RequestParameters []ReactorFormulaRequestParameter `json:"request_parameters,omitempty"`
}

// NewReactorFormula instantiates a new ReactorFormula object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactorFormula() *ReactorFormula {
	this := ReactorFormula{}
	return &this
}

// NewReactorFormulaWithDefaults instantiates a new ReactorFormula object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactorFormulaWithDefaults() *ReactorFormula {
	this := ReactorFormula{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReactorFormula) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactorFormula) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReactorFormula) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReactorFormula) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetType() string {
	if o == nil || isNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ReactorFormula) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ReactorFormula) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *ReactorFormula) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ReactorFormula) UnsetType() {
	o.Type.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetStatus() string {
	if o == nil || isNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ReactorFormula) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ReactorFormula) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *ReactorFormula) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ReactorFormula) UnsetStatus() {
	o.Status.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReactorFormula) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReactorFormula) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ReactorFormula) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReactorFormula) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ReactorFormula) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ReactorFormula) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ReactorFormula) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ReactorFormula) UnsetDescription() {
	o.Description.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetIcon() string {
	if o == nil || isNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ReactorFormula) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ReactorFormula) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *ReactorFormula) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ReactorFormula) UnsetIcon() {
	o.Icon.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetCode() string {
	if o == nil || isNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ReactorFormula) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *ReactorFormula) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *ReactorFormula) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ReactorFormula) UnsetCode() {
	o.Code.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ReactorFormula) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ReactorFormula) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ReactorFormula) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ReactorFormula) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReactorFormula) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ReactorFormula) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ReactorFormula) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ReactorFormula) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetModifiedBy() string {
	if o == nil || isNil(o.ModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ReactorFormula) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableString and assigns it to the ModifiedBy field.
func (o *ReactorFormula) SetModifiedBy(v string) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *ReactorFormula) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *ReactorFormula) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ReactorFormula) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *ReactorFormula) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *ReactorFormula) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *ReactorFormula) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetConfiguration() []ReactorFormulaConfiguration {
	if o == nil {
		var ret []ReactorFormulaConfiguration
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetConfigurationOk() ([]ReactorFormulaConfiguration, bool) {
	if o == nil || isNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ReactorFormula) HasConfiguration() bool {
	if o != nil && isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given []ReactorFormulaConfiguration and assigns it to the Configuration field.
func (o *ReactorFormula) SetConfiguration(v []ReactorFormulaConfiguration) {
	o.Configuration = v
}

// GetRequestParameters returns the RequestParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReactorFormula) GetRequestParameters() []ReactorFormulaRequestParameter {
	if o == nil {
		var ret []ReactorFormulaRequestParameter
		return ret
	}
	return o.RequestParameters
}

// GetRequestParametersOk returns a tuple with the RequestParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReactorFormula) GetRequestParametersOk() ([]ReactorFormulaRequestParameter, bool) {
	if o == nil || isNil(o.RequestParameters) {
		return nil, false
	}
	return o.RequestParameters, true
}

// HasRequestParameters returns a boolean if a field has been set.
func (o *ReactorFormula) HasRequestParameters() bool {
	if o != nil && isNil(o.RequestParameters) {
		return true
	}

	return false
}

// SetRequestParameters gets a reference to the given []ReactorFormulaRequestParameter and assigns it to the RequestParameters field.
func (o *ReactorFormula) SetRequestParameters(v []ReactorFormulaRequestParameter) {
	o.RequestParameters = v
}

func (o ReactorFormula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modified_by"] = o.ModifiedBy.Get()
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.RequestParameters != nil {
		toSerialize["request_parameters"] = o.RequestParameters
	}
	return json.Marshal(toSerialize)
}

type NullableReactorFormula struct {
	value *ReactorFormula
	isSet bool
}

func (v NullableReactorFormula) Get() *ReactorFormula {
	return v.value
}

func (v *NullableReactorFormula) Set(val *ReactorFormula) {
	v.value = val
	v.isSet = true
}

func (v NullableReactorFormula) IsSet() bool {
	return v.isSet
}

func (v *NullableReactorFormula) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactorFormula(val *ReactorFormula) *NullableReactorFormula {
	return &NullableReactorFormula{value: val, isSet: true}
}

func (v NullableReactorFormula) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactorFormula) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
