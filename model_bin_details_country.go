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

// checks if the BinDetailsCountry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BinDetailsCountry{}

// BinDetailsCountry struct for BinDetailsCountry
type BinDetailsCountry struct {
	Alpha2  NullableString `json:"alpha2,omitempty"`
	Name    NullableString `json:"name,omitempty"`
	Numeric NullableString `json:"numeric,omitempty"`
}

// NewBinDetailsCountry instantiates a new BinDetailsCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinDetailsCountry() *BinDetailsCountry {
	this := BinDetailsCountry{}
	return &this
}

// NewBinDetailsCountryWithDefaults instantiates a new BinDetailsCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinDetailsCountryWithDefaults() *BinDetailsCountry {
	this := BinDetailsCountry{}
	return &this
}

// GetAlpha2 returns the Alpha2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BinDetailsCountry) GetAlpha2() string {
	if o == nil || IsNil(o.Alpha2.Get()) {
		var ret string
		return ret
	}
	return *o.Alpha2.Get()
}

// GetAlpha2Ok returns a tuple with the Alpha2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BinDetailsCountry) GetAlpha2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alpha2.Get(), o.Alpha2.IsSet()
}

// HasAlpha2 returns a boolean if a field is not nil.
func (o *BinDetailsCountry) HasAlpha2() bool {
	if o != nil && !IsNil(o.Alpha2) {
		return true
	}

	return false
}

// SetAlpha2 gets a reference to the given NullableString and assigns it to the Alpha2 field.
func (o *BinDetailsCountry) SetAlpha2(v string) {
	o.Alpha2.Set(&v)
}

// SetAlpha2Nil sets the value for Alpha2 to be an explicit nil
func (o *BinDetailsCountry) SetAlpha2Nil() {
	o.Alpha2.Set(nil)
}

// UnsetAlpha2 ensures that no value is present for Alpha2, not even an explicit nil
func (o *BinDetailsCountry) UnsetAlpha2() {
	o.Alpha2.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BinDetailsCountry) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BinDetailsCountry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *BinDetailsCountry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BinDetailsCountry) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BinDetailsCountry) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BinDetailsCountry) UnsetName() {
	o.Name.Unset()
}

// GetNumeric returns the Numeric field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BinDetailsCountry) GetNumeric() string {
	if o == nil || IsNil(o.Numeric.Get()) {
		var ret string
		return ret
	}
	return *o.Numeric.Get()
}

// GetNumericOk returns a tuple with the Numeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BinDetailsCountry) GetNumericOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Numeric.Get(), o.Numeric.IsSet()
}

// HasNumeric returns a boolean if a field is not nil.
func (o *BinDetailsCountry) HasNumeric() bool {
	if o != nil && !IsNil(o.Numeric) {
		return true
	}

	return false
}

// SetNumeric gets a reference to the given NullableString and assigns it to the Numeric field.
func (o *BinDetailsCountry) SetNumeric(v string) {
	o.Numeric.Set(&v)
}

// SetNumericNil sets the value for Numeric to be an explicit nil
func (o *BinDetailsCountry) SetNumericNil() {
	o.Numeric.Set(nil)
}

// UnsetNumeric ensures that no value is present for Numeric, not even an explicit nil
func (o *BinDetailsCountry) UnsetNumeric() {
	o.Numeric.Unset()
}

func (o BinDetailsCountry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BinDetailsCountry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Alpha2.IsSet() {
		toSerialize["alpha2"] = o.Alpha2.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Numeric.IsSet() {
		toSerialize["numeric"] = o.Numeric.Get()
	}
	return toSerialize, nil
}

type NullableBinDetailsCountry struct {
	value *BinDetailsCountry
	isSet bool
}

func (v NullableBinDetailsCountry) Get() *BinDetailsCountry {
	return v.value
}

func (v *NullableBinDetailsCountry) Set(val *BinDetailsCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableBinDetailsCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableBinDetailsCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinDetailsCountry(val *BinDetailsCountry) *NullableBinDetailsCountry {
	return &NullableBinDetailsCountry{value: val, isSet: true}
}

func (v NullableBinDetailsCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinDetailsCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
