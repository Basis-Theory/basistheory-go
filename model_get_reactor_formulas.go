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

// checks if the GetReactorFormulas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReactorFormulas{}

// GetReactorFormulas struct for GetReactorFormulas
type GetReactorFormulas struct {
	Name NullableString `json:"name,omitempty"`
	Page NullableInt32  `json:"page,omitempty"`
	Size NullableInt32  `json:"size,omitempty"`
}

// NewGetReactorFormulas instantiates a new GetReactorFormulas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReactorFormulas() *GetReactorFormulas {
	this := GetReactorFormulas{}
	return &this
}

// NewGetReactorFormulasWithDefaults instantiates a new GetReactorFormulas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReactorFormulasWithDefaults() *GetReactorFormulas {
	this := GetReactorFormulas{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactorFormulas) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactorFormulas) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *GetReactorFormulas) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GetReactorFormulas) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *GetReactorFormulas) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GetReactorFormulas) UnsetName() {
	o.Name.Unset()
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactorFormulas) GetPage() int32 {
	if o == nil || IsNil(o.Page.Get()) {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactorFormulas) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field is not nil.
func (o *GetReactorFormulas) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *GetReactorFormulas) SetPage(v int32) {
	o.Page.Set(&v)
}

// SetPageNil sets the value for Page to be an explicit nil
func (o *GetReactorFormulas) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *GetReactorFormulas) UnsetPage() {
	o.Page.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReactorFormulas) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReactorFormulas) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field is not nil.
func (o *GetReactorFormulas) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *GetReactorFormulas) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *GetReactorFormulas) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GetReactorFormulas) UnsetSize() {
	o.Size.Unset()
}

func (o GetReactorFormulas) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReactorFormulas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return toSerialize, nil
}

type NullableGetReactorFormulas struct {
	value *GetReactorFormulas
	isSet bool
}

func (v NullableGetReactorFormulas) Get() *GetReactorFormulas {
	return v.value
}

func (v *NullableGetReactorFormulas) Set(val *GetReactorFormulas) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReactorFormulas) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReactorFormulas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReactorFormulas(val *GetReactorFormulas) *NullableGetReactorFormulas {
	return &NullableGetReactorFormulas{value: val, isSet: true}
}

func (v NullableGetReactorFormulas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReactorFormulas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
