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

// checks if the ApplicationTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationTemplate{}

// ApplicationTemplate struct for ApplicationTemplate
type ApplicationTemplate struct {
	Id              *string        `json:"id,omitempty"`
	Name            NullableString `json:"name,omitempty"`
	Description     NullableString `json:"description,omitempty"`
	ApplicationType NullableString `json:"application_type,omitempty"`
	TemplateType    NullableString `json:"template_type,omitempty"`
	IsStarter       *bool          `json:"is_starter,omitempty"`
	Rules           []AccessRule   `json:"rules,omitempty"`
	Permissions     []string       `json:"permissions,omitempty"`
}

// NewApplicationTemplate instantiates a new ApplicationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationTemplate() *ApplicationTemplate {
	this := ApplicationTemplate{}
	return &this
}

// NewApplicationTemplateWithDefaults instantiates a new ApplicationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationTemplateWithDefaults() *ApplicationTemplate {
	this := ApplicationTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationTemplate) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApplicationTemplate) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ApplicationTemplate) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApplicationTemplate) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApplicationTemplate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApplicationTemplate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApplicationTemplate) UnsetDescription() {
	o.Description.Unset()
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationTemplate) GetApplicationType() string {
	if o == nil || IsNil(o.ApplicationType.Get()) {
		var ret string
		return ret
	}
	return *o.ApplicationType.Get()
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationTemplate) GetApplicationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationType.Get(), o.ApplicationType.IsSet()
}

// HasApplicationType returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasApplicationType() bool {
	if o != nil && !IsNil(o.ApplicationType) {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given NullableString and assigns it to the ApplicationType field.
func (o *ApplicationTemplate) SetApplicationType(v string) {
	o.ApplicationType.Set(&v)
}

// SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil
func (o *ApplicationTemplate) SetApplicationTypeNil() {
	o.ApplicationType.Set(nil)
}

// UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
func (o *ApplicationTemplate) UnsetApplicationType() {
	o.ApplicationType.Unset()
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationTemplate) GetTemplateType() string {
	if o == nil || IsNil(o.TemplateType.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateType.Get()
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationTemplate) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateType.Get(), o.TemplateType.IsSet()
}

// HasTemplateType returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasTemplateType() bool {
	if o != nil && !IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given NullableString and assigns it to the TemplateType field.
func (o *ApplicationTemplate) SetTemplateType(v string) {
	o.TemplateType.Set(&v)
}

// SetTemplateTypeNil sets the value for TemplateType to be an explicit nil
func (o *ApplicationTemplate) SetTemplateTypeNil() {
	o.TemplateType.Set(nil)
}

// UnsetTemplateType ensures that no value is present for TemplateType, not even an explicit nil
func (o *ApplicationTemplate) UnsetTemplateType() {
	o.TemplateType.Unset()
}

// GetIsStarter returns the IsStarter field value if set, zero value otherwise.
func (o *ApplicationTemplate) GetIsStarter() bool {
	if o == nil || IsNil(o.IsStarter) {
		var ret bool
		return ret
	}
	return *o.IsStarter
}

// GetIsStarterOk returns a tuple with the IsStarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTemplate) GetIsStarterOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStarter) {
		return nil, false
	}
	return o.IsStarter, true
}

// HasIsStarter returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasIsStarter() bool {
	if o != nil && !IsNil(o.IsStarter) {
		return true
	}

	return false
}

// SetIsStarter gets a reference to the given bool and assigns it to the IsStarter field.
func (o *ApplicationTemplate) SetIsStarter(v bool) {
	o.IsStarter = &v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationTemplate) GetRules() []AccessRule {
	if o == nil {
		var ret []AccessRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationTemplate) GetRulesOk() ([]AccessRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []AccessRule and assigns it to the Rules field.
func (o *ApplicationTemplate) SetRules(v []AccessRule) {
	o.Rules = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationTemplate) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationTemplate) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field is not nil.
func (o *ApplicationTemplate) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *ApplicationTemplate) SetPermissions(v []string) {
	o.Permissions = v
}

func (o ApplicationTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ApplicationType.IsSet() {
		toSerialize["application_type"] = o.ApplicationType.Get()
	}
	if o.TemplateType.IsSet() {
		toSerialize["template_type"] = o.TemplateType.Get()
	}
	if !IsNil(o.IsStarter) {
		toSerialize["is_starter"] = o.IsStarter
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableApplicationTemplate struct {
	value *ApplicationTemplate
	isSet bool
}

func (v NullableApplicationTemplate) Get() *ApplicationTemplate {
	return v.value
}

func (v *NullableApplicationTemplate) Set(val *ApplicationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationTemplate(val *ApplicationTemplate) *NullableApplicationTemplate {
	return &NullableApplicationTemplate{value: val, isSet: true}
}

func (v NullableApplicationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
