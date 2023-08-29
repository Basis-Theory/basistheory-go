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

// checks if the TokenEnrichments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenEnrichments{}

// TokenEnrichments struct for TokenEnrichments
type TokenEnrichments struct {
	BinDetails *BinDetails `json:"bin_details,omitempty"`
}

// NewTokenEnrichments instantiates a new TokenEnrichments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenEnrichments() *TokenEnrichments {
	this := TokenEnrichments{}
	return &this
}

// NewTokenEnrichmentsWithDefaults instantiates a new TokenEnrichments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenEnrichmentsWithDefaults() *TokenEnrichments {
	this := TokenEnrichments{}
	return &this
}

// GetBinDetails returns the BinDetails field value if set, zero value otherwise.
func (o *TokenEnrichments) GetBinDetails() BinDetails {
	if o == nil || IsNil(o.BinDetails) {
		var ret BinDetails
		return ret
	}
	return *o.BinDetails
}

// GetBinDetailsOk returns a tuple with the BinDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenEnrichments) GetBinDetailsOk() (*BinDetails, bool) {
	if o == nil || IsNil(o.BinDetails) {
		return nil, false
	}
	return o.BinDetails, true
}

// HasBinDetails returns a boolean if a field is not nil.
func (o *TokenEnrichments) HasBinDetails() bool {
	if o != nil && !IsNil(o.BinDetails) {
		return true
	}

	return false
}

// SetBinDetails gets a reference to the given BinDetails and assigns it to the BinDetails field.
func (o *TokenEnrichments) SetBinDetails(v BinDetails) {
	o.BinDetails = &v
}

func (o TokenEnrichments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenEnrichments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BinDetails) {
		toSerialize["bin_details"] = o.BinDetails
	}
	return toSerialize, nil
}

type NullableTokenEnrichments struct {
	value *TokenEnrichments
	isSet bool
}

func (v NullableTokenEnrichments) Get() *TokenEnrichments {
	return v.value
}

func (v *NullableTokenEnrichments) Set(val *TokenEnrichments) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenEnrichments) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenEnrichments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenEnrichments(val *TokenEnrichments) *NullableTokenEnrichments {
	return &NullableTokenEnrichments{value: val, isSet: true}
}

func (v NullableTokenEnrichments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenEnrichments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}