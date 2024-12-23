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

// checks if the CreateThreeDSSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThreeDSSessionRequest{}

// CreateThreeDSSessionRequest struct for CreateThreeDSSessionRequest
type CreateThreeDSSessionRequest struct {
	// Deprecated
	Pan           NullableString     `json:"pan,omitempty"`
	TokenId       NullableString     `json:"token_id,omitempty"`
	TokenIntentId NullableString     `json:"token_intent_id,omitempty"`
	Type          NullableString     `json:"type,omitempty"`
	Device        NullableString     `json:"device,omitempty"`
	DeviceInfo    *ThreeDSDeviceInfo `json:"device_info,omitempty"`
}

// NewCreateThreeDSSessionRequest instantiates a new CreateThreeDSSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThreeDSSessionRequest() *CreateThreeDSSessionRequest {
	this := CreateThreeDSSessionRequest{}
	return &this
}

// NewCreateThreeDSSessionRequestWithDefaults instantiates a new CreateThreeDSSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThreeDSSessionRequestWithDefaults() *CreateThreeDSSessionRequest {
	this := CreateThreeDSSessionRequest{}
	return &this
}

// GetPan returns the Pan field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *CreateThreeDSSessionRequest) GetPan() string {
	if o == nil || IsNil(o.Pan.Get()) {
		var ret string
		return ret
	}
	return *o.Pan.Get()
}

// GetPanOk returns a tuple with the Pan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *CreateThreeDSSessionRequest) GetPanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pan.Get(), o.Pan.IsSet()
}

// HasPan returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionRequest) HasPan() bool {
	if o != nil && !IsNil(o.Pan) {
		return true
	}

	return false
}

// SetPan gets a reference to the given NullableString and assigns it to the Pan field.
// Deprecated
func (o *CreateThreeDSSessionRequest) SetPan(v string) {
	o.Pan.Set(&v)
}

// SetPanNil sets the value for Pan to be an explicit nil
func (o *CreateThreeDSSessionRequest) SetPanNil() {
	o.Pan.Set(nil)
}

// UnsetPan ensures that no value is present for Pan, not even an explicit nil
func (o *CreateThreeDSSessionRequest) UnsetPan() {
	o.Pan.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionRequest) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionRequest) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionRequest) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *CreateThreeDSSessionRequest) SetTokenId(v string) {
	o.TokenId.Set(&v)
}

// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *CreateThreeDSSessionRequest) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *CreateThreeDSSessionRequest) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetTokenIntentId returns the TokenIntentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionRequest) GetTokenIntentId() string {
	if o == nil || IsNil(o.TokenIntentId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenIntentId.Get()
}

// GetTokenIntentIdOk returns a tuple with the TokenIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionRequest) GetTokenIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenIntentId.Get(), o.TokenIntentId.IsSet()
}

// HasTokenIntentId returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionRequest) HasTokenIntentId() bool {
	if o != nil && !IsNil(o.TokenIntentId) {
		return true
	}

	return false
}

// SetTokenIntentId gets a reference to the given NullableString and assigns it to the TokenIntentId field.
func (o *CreateThreeDSSessionRequest) SetTokenIntentId(v string) {
	o.TokenIntentId.Set(&v)
}

// SetTokenIntentIdNil sets the value for TokenIntentId to be an explicit nil
func (o *CreateThreeDSSessionRequest) SetTokenIntentIdNil() {
	o.TokenIntentId.Set(nil)
}

// UnsetTokenIntentId ensures that no value is present for TokenIntentId, not even an explicit nil
func (o *CreateThreeDSSessionRequest) UnsetTokenIntentId() {
	o.TokenIntentId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionRequest) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CreateThreeDSSessionRequest) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateThreeDSSessionRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateThreeDSSessionRequest) UnsetType() {
	o.Type.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionRequest) GetDevice() string {
	if o == nil || IsNil(o.Device.Get()) {
		var ret string
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionRequest) GetDeviceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableString and assigns it to the Device field.
func (o *CreateThreeDSSessionRequest) SetDevice(v string) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *CreateThreeDSSessionRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *CreateThreeDSSessionRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetDeviceInfo returns the DeviceInfo field value if set, zero value otherwise.
func (o *CreateThreeDSSessionRequest) GetDeviceInfo() ThreeDSDeviceInfo {
	if o == nil || IsNil(o.DeviceInfo) {
		var ret ThreeDSDeviceInfo
		return ret
	}
	return *o.DeviceInfo
}

// GetDeviceInfoOk returns a tuple with the DeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThreeDSSessionRequest) GetDeviceInfoOk() (*ThreeDSDeviceInfo, bool) {
	if o == nil || IsNil(o.DeviceInfo) {
		return nil, false
	}
	return o.DeviceInfo, true
}

// HasDeviceInfo returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionRequest) HasDeviceInfo() bool {
	if o != nil && !IsNil(o.DeviceInfo) {
		return true
	}

	return false
}

// SetDeviceInfo gets a reference to the given ThreeDSDeviceInfo and assigns it to the DeviceInfo field.
func (o *CreateThreeDSSessionRequest) SetDeviceInfo(v ThreeDSDeviceInfo) {
	o.DeviceInfo = &v
}

func (o CreateThreeDSSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThreeDSSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Pan.IsSet() {
		toSerialize["pan"] = o.Pan.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}
	if o.TokenIntentId.IsSet() {
		toSerialize["token_intent_id"] = o.TokenIntentId.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if !IsNil(o.DeviceInfo) {
		toSerialize["device_info"] = o.DeviceInfo
	}
	return toSerialize, nil
}

type NullableCreateThreeDSSessionRequest struct {
	value *CreateThreeDSSessionRequest
	isSet bool
}

func (v NullableCreateThreeDSSessionRequest) Get() *CreateThreeDSSessionRequest {
	return v.value
}

func (v *NullableCreateThreeDSSessionRequest) Set(val *CreateThreeDSSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThreeDSSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThreeDSSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThreeDSSessionRequest(val *CreateThreeDSSessionRequest) *NullableCreateThreeDSSessionRequest {
	return &NullableCreateThreeDSSessionRequest{value: val, isSet: true}
}

func (v NullableCreateThreeDSSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThreeDSSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
