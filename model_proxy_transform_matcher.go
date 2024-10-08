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

// ProxyTransformMatcher the model 'ProxyTransformMatcher'
type ProxyTransformMatcher string

// List of ProxyTransformMatcher
const (
	REGEX             ProxyTransformMatcher = "REGEX"
	CHASE_STRATUS_PAN ProxyTransformMatcher = "CHASE_STRATUS_PAN"
)

// All allowed values of ProxyTransformMatcher enum
var AllowedProxyTransformMatcherEnumValues = []ProxyTransformMatcher{
	"REGEX",
	"CHASE_STRATUS_PAN",
}

func (v *ProxyTransformMatcher) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProxyTransformMatcher(value)
	for _, existing := range AllowedProxyTransformMatcherEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProxyTransformMatcher", value)
}

// NewProxyTransformMatcherFromValue returns a pointer to a valid ProxyTransformMatcher
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxyTransformMatcherFromValue(v string) (*ProxyTransformMatcher, error) {
	ev := ProxyTransformMatcher(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProxyTransformMatcher: valid values are %v", v, AllowedProxyTransformMatcherEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxyTransformMatcher) IsValid() bool {
	for _, existing := range AllowedProxyTransformMatcherEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxyTransformMatcher value
func (v ProxyTransformMatcher) Ptr() *ProxyTransformMatcher {
	return &v
}

type NullableProxyTransformMatcher struct {
	value *ProxyTransformMatcher
	isSet bool
}

func (v NullableProxyTransformMatcher) Get() *ProxyTransformMatcher {
	return v.value
}

func (v *NullableProxyTransformMatcher) Set(val *ProxyTransformMatcher) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyTransformMatcher) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyTransformMatcher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyTransformMatcher(val *ProxyTransformMatcher) *NullableProxyTransformMatcher {
	return &NullableProxyTransformMatcher{value: val, isSet: true}
}

func (v NullableProxyTransformMatcher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyTransformMatcher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
