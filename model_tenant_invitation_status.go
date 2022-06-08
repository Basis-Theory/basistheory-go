/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Server to Server Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"fmt"
)

// TenantInvitationStatus the model 'TenantInvitationStatus'
type TenantInvitationStatus string

// List of TenantInvitationStatus
const (
	PENDING TenantInvitationStatus = "PENDING"
	EXPIRED TenantInvitationStatus = "EXPIRED"
)

// All allowed values of TenantInvitationStatus enum
var AllowedTenantInvitationStatusEnumValues = []TenantInvitationStatus{
	"PENDING",
	"EXPIRED",
}

func (v *TenantInvitationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenantInvitationStatus(value)
	for _, existing := range AllowedTenantInvitationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenantInvitationStatus", value)
}

// NewTenantInvitationStatusFromValue returns a pointer to a valid TenantInvitationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenantInvitationStatusFromValue(v string) (*TenantInvitationStatus, error) {
	ev := TenantInvitationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenantInvitationStatus: valid values are %v", v, AllowedTenantInvitationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenantInvitationStatus) IsValid() bool {
	for _, existing := range AllowedTenantInvitationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenantInvitationStatus value
func (v TenantInvitationStatus) Ptr() *TenantInvitationStatus {
	return &v
}

type NullableTenantInvitationStatus struct {
	value *TenantInvitationStatus
	isSet bool
}

func (v NullableTenantInvitationStatus) Get() *TenantInvitationStatus {
	return v.value
}

func (v *NullableTenantInvitationStatus) Set(val *TenantInvitationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantInvitationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantInvitationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantInvitationStatus(val *TenantInvitationStatus) *NullableTenantInvitationStatus {
	return &NullableTenantInvitationStatus{value: val, isSet: true}
}

func (v NullableTenantInvitationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantInvitationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

