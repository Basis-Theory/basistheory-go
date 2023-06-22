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

// checks if the GetApplications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApplications{}

// GetApplications struct for GetApplications
type GetApplications struct {
	Id   []string      `json:"id,omitempty"`
	Type []string      `json:"type,omitempty"`
	Page NullableInt32 `json:"page,omitempty"`
	Size NullableInt32 `json:"size,omitempty"`
}

// NewGetApplications instantiates a new GetApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApplications() *GetApplications {
	this := GetApplications{}
	return &this
}

// NewGetApplicationsWithDefaults instantiates a new GetApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApplicationsWithDefaults() *GetApplications {
	this := GetApplications{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetApplications) GetId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetApplications) GetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field is not nil.
func (o *GetApplications) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *GetApplications) SetId(v []string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetApplications) GetType() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetApplications) GetTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field is not nil.
func (o *GetApplications) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *GetApplications) SetType(v []string) {
	o.Type = v
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetApplications) GetPage() int32 {
	if o == nil || IsNil(o.Page.Get()) {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetApplications) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field is not nil.
func (o *GetApplications) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *GetApplications) SetPage(v int32) {
	o.Page.Set(&v)
}

// SetPageNil sets the value for Page to be an explicit nil
func (o *GetApplications) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *GetApplications) UnsetPage() {
	o.Page.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetApplications) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetApplications) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field is not nil.
func (o *GetApplications) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *GetApplications) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *GetApplications) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GetApplications) UnsetSize() {
	o.Size.Unset()
}

func (o GetApplications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApplications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return toSerialize, nil
}

type NullableGetApplications struct {
	value *GetApplications
	isSet bool
}

func (v NullableGetApplications) Get() *GetApplications {
	return v.value
}

func (v *NullableGetApplications) Set(val *GetApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApplications(val *GetApplications) *NullableGetApplications {
	return &NullableGetApplications{value: val, isSet: true}
}

func (v NullableGetApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
