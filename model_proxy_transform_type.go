/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"fmt"
)

// ProxyTransformType the model 'ProxyTransformType'
type ProxyTransformType string

// List of ProxyTransformType
const (
	CODE ProxyTransformType = "CODE"
	MASK ProxyTransformType = "MASK"
)

// All allowed values of ProxyTransformType enum
var AllowedProxyTransformTypeEnumValues = []ProxyTransformType{
	"CODE",
	"MASK",
}

func (v *ProxyTransformType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProxyTransformType(value)
	for _, existing := range AllowedProxyTransformTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProxyTransformType", value)
}

// NewProxyTransformTypeFromValue returns a pointer to a valid ProxyTransformType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxyTransformTypeFromValue(v string) (*ProxyTransformType, error) {
	ev := ProxyTransformType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProxyTransformType: valid values are %v", v, AllowedProxyTransformTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxyTransformType) IsValid() bool {
	for _, existing := range AllowedProxyTransformTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxyTransformType value
func (v ProxyTransformType) Ptr() *ProxyTransformType {
	return &v
}

type NullableProxyTransformType struct {
	value *ProxyTransformType
	isSet bool
}

func (v NullableProxyTransformType) Get() *ProxyTransformType {
	return v.value
}

func (v *NullableProxyTransformType) Set(val *ProxyTransformType) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyTransformType) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyTransformType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyTransformType(val *ProxyTransformType) *NullableProxyTransformType {
	return &NullableProxyTransformType{value: val, isSet: true}
}

func (v NullableProxyTransformType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyTransformType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
