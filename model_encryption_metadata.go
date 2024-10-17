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

// checks if the EncryptionMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionMetadata{}

// EncryptionMetadata struct for EncryptionMetadata
type EncryptionMetadata struct {
	Cek *EncryptionKey `json:"cek,omitempty"`
	Kek *EncryptionKey `json:"kek,omitempty"`
}

// NewEncryptionMetadata instantiates a new EncryptionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionMetadata() *EncryptionMetadata {
	this := EncryptionMetadata{}
	return &this
}

// NewEncryptionMetadataWithDefaults instantiates a new EncryptionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionMetadataWithDefaults() *EncryptionMetadata {
	this := EncryptionMetadata{}
	return &this
}

// GetCek returns the Cek field value if set, zero value otherwise.
func (o *EncryptionMetadata) GetCek() EncryptionKey {
	if o == nil || IsNil(o.Cek) {
		var ret EncryptionKey
		return ret
	}
	return *o.Cek
}

// GetCekOk returns a tuple with the Cek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionMetadata) GetCekOk() (*EncryptionKey, bool) {
	if o == nil || IsNil(o.Cek) {
		return nil, false
	}
	return o.Cek, true
}

// HasCek returns a boolean if a field is not nil.
func (o *EncryptionMetadata) HasCek() bool {
	if o != nil && !IsNil(o.Cek) {
		return true
	}

	return false
}

// SetCek gets a reference to the given EncryptionKey and assigns it to the Cek field.
func (o *EncryptionMetadata) SetCek(v EncryptionKey) {
	o.Cek = &v
}

// GetKek returns the Kek field value if set, zero value otherwise.
func (o *EncryptionMetadata) GetKek() EncryptionKey {
	if o == nil || IsNil(o.Kek) {
		var ret EncryptionKey
		return ret
	}
	return *o.Kek
}

// GetKekOk returns a tuple with the Kek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionMetadata) GetKekOk() (*EncryptionKey, bool) {
	if o == nil || IsNil(o.Kek) {
		return nil, false
	}
	return o.Kek, true
}

// HasKek returns a boolean if a field is not nil.
func (o *EncryptionMetadata) HasKek() bool {
	if o != nil && !IsNil(o.Kek) {
		return true
	}

	return false
}

// SetKek gets a reference to the given EncryptionKey and assigns it to the Kek field.
func (o *EncryptionMetadata) SetKek(v EncryptionKey) {
	o.Kek = &v
}

func (o EncryptionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cek) {
		toSerialize["cek"] = o.Cek
	}
	if !IsNil(o.Kek) {
		toSerialize["kek"] = o.Kek
	}
	return toSerialize, nil
}

type NullableEncryptionMetadata struct {
	value *EncryptionMetadata
	isSet bool
}

func (v NullableEncryptionMetadata) Get() *EncryptionMetadata {
	return v.value
}

func (v *NullableEncryptionMetadata) Set(val *EncryptionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionMetadata(val *EncryptionMetadata) *NullableEncryptionMetadata {
	return &NullableEncryptionMetadata{value: val, isSet: true}
}

func (v NullableEncryptionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}