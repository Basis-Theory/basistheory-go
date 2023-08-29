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

// checks if the BinDetailsProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BinDetailsProduct{}

// BinDetailsProduct struct for BinDetailsProduct
type BinDetailsProduct struct {
	Code NullableString `json:"code,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewBinDetailsProduct instantiates a new BinDetailsProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinDetailsProduct() *BinDetailsProduct {
	this := BinDetailsProduct{}
	return &this
}

// NewBinDetailsProductWithDefaults instantiates a new BinDetailsProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinDetailsProductWithDefaults() *BinDetailsProduct {
	this := BinDetailsProduct{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BinDetailsProduct) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BinDetailsProduct) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field is not nil.
func (o *BinDetailsProduct) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *BinDetailsProduct) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *BinDetailsProduct) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *BinDetailsProduct) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BinDetailsProduct) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BinDetailsProduct) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *BinDetailsProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BinDetailsProduct) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BinDetailsProduct) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BinDetailsProduct) UnsetName() {
	o.Name.Unset()
}

func (o BinDetailsProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BinDetailsProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableBinDetailsProduct struct {
	value *BinDetailsProduct
	isSet bool
}

func (v NullableBinDetailsProduct) Get() *BinDetailsProduct {
	return v.value
}

func (v *NullableBinDetailsProduct) Set(val *BinDetailsProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableBinDetailsProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableBinDetailsProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinDetailsProduct(val *BinDetailsProduct) *NullableBinDetailsProduct {
	return &NullableBinDetailsProduct{value: val, isSet: true}
}

func (v NullableBinDetailsProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinDetailsProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
