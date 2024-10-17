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

// checks if the CreateThreeDSSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThreeDSSessionResponse{}

// CreateThreeDSSessionResponse struct for CreateThreeDSSessionResponse
type CreateThreeDSSessionResponse struct {
	Id                    *string        `json:"id,omitempty"`
	Type                  NullableString `json:"type,omitempty"`
	CardBrand             NullableString `json:"cardBrand,omitempty"`
	MethodUrl             NullableString `json:"method_url,omitempty"`
	MethodNotificationUrl NullableString `json:"method_notification_url,omitempty"`
	DirectoryServerId     NullableString `json:"directory_server_id,omitempty"`
	RecommendedVersion    NullableString `json:"recommended_version,omitempty"`
}

// NewCreateThreeDSSessionResponse instantiates a new CreateThreeDSSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThreeDSSessionResponse() *CreateThreeDSSessionResponse {
	this := CreateThreeDSSessionResponse{}
	return &this
}

// NewCreateThreeDSSessionResponseWithDefaults instantiates a new CreateThreeDSSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThreeDSSessionResponseWithDefaults() *CreateThreeDSSessionResponse {
	this := CreateThreeDSSessionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateThreeDSSessionResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThreeDSSessionResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateThreeDSSessionResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionResponse) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CreateThreeDSSessionResponse) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateThreeDSSessionResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateThreeDSSessionResponse) UnsetType() {
	o.Type.Unset()
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionResponse) GetCardBrand() string {
	if o == nil || IsNil(o.CardBrand.Get()) {
		var ret string
		return ret
	}
	return *o.CardBrand.Get()
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionResponse) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBrand.Get(), o.CardBrand.IsSet()
}

// HasCardBrand returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasCardBrand() bool {
	if o != nil && !IsNil(o.CardBrand) {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given NullableString and assigns it to the CardBrand field.
func (o *CreateThreeDSSessionResponse) SetCardBrand(v string) {
	o.CardBrand.Set(&v)
}

// SetCardBrandNil sets the value for CardBrand to be an explicit nil
func (o *CreateThreeDSSessionResponse) SetCardBrandNil() {
	o.CardBrand.Set(nil)
}

// UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
func (o *CreateThreeDSSessionResponse) UnsetCardBrand() {
	o.CardBrand.Unset()
}

// GetMethodUrl returns the MethodUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionResponse) GetMethodUrl() string {
	if o == nil || IsNil(o.MethodUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MethodUrl.Get()
}

// GetMethodUrlOk returns a tuple with the MethodUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionResponse) GetMethodUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MethodUrl.Get(), o.MethodUrl.IsSet()
}

// HasMethodUrl returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasMethodUrl() bool {
	if o != nil && !IsNil(o.MethodUrl) {
		return true
	}

	return false
}

// SetMethodUrl gets a reference to the given NullableString and assigns it to the MethodUrl field.
func (o *CreateThreeDSSessionResponse) SetMethodUrl(v string) {
	o.MethodUrl.Set(&v)
}

// SetMethodUrlNil sets the value for MethodUrl to be an explicit nil
func (o *CreateThreeDSSessionResponse) SetMethodUrlNil() {
	o.MethodUrl.Set(nil)
}

// UnsetMethodUrl ensures that no value is present for MethodUrl, not even an explicit nil
func (o *CreateThreeDSSessionResponse) UnsetMethodUrl() {
	o.MethodUrl.Unset()
}

// GetMethodNotificationUrl returns the MethodNotificationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionResponse) GetMethodNotificationUrl() string {
	if o == nil || IsNil(o.MethodNotificationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MethodNotificationUrl.Get()
}

// GetMethodNotificationUrlOk returns a tuple with the MethodNotificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionResponse) GetMethodNotificationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MethodNotificationUrl.Get(), o.MethodNotificationUrl.IsSet()
}

// HasMethodNotificationUrl returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasMethodNotificationUrl() bool {
	if o != nil && !IsNil(o.MethodNotificationUrl) {
		return true
	}

	return false
}

// SetMethodNotificationUrl gets a reference to the given NullableString and assigns it to the MethodNotificationUrl field.
func (o *CreateThreeDSSessionResponse) SetMethodNotificationUrl(v string) {
	o.MethodNotificationUrl.Set(&v)
}

// SetMethodNotificationUrlNil sets the value for MethodNotificationUrl to be an explicit nil
func (o *CreateThreeDSSessionResponse) SetMethodNotificationUrlNil() {
	o.MethodNotificationUrl.Set(nil)
}

// UnsetMethodNotificationUrl ensures that no value is present for MethodNotificationUrl, not even an explicit nil
func (o *CreateThreeDSSessionResponse) UnsetMethodNotificationUrl() {
	o.MethodNotificationUrl.Unset()
}

// GetDirectoryServerId returns the DirectoryServerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionResponse) GetDirectoryServerId() string {
	if o == nil || IsNil(o.DirectoryServerId.Get()) {
		var ret string
		return ret
	}
	return *o.DirectoryServerId.Get()
}

// GetDirectoryServerIdOk returns a tuple with the DirectoryServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionResponse) GetDirectoryServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectoryServerId.Get(), o.DirectoryServerId.IsSet()
}

// HasDirectoryServerId returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasDirectoryServerId() bool {
	if o != nil && !IsNil(o.DirectoryServerId) {
		return true
	}

	return false
}

// SetDirectoryServerId gets a reference to the given NullableString and assigns it to the DirectoryServerId field.
func (o *CreateThreeDSSessionResponse) SetDirectoryServerId(v string) {
	o.DirectoryServerId.Set(&v)
}

// SetDirectoryServerIdNil sets the value for DirectoryServerId to be an explicit nil
func (o *CreateThreeDSSessionResponse) SetDirectoryServerIdNil() {
	o.DirectoryServerId.Set(nil)
}

// UnsetDirectoryServerId ensures that no value is present for DirectoryServerId, not even an explicit nil
func (o *CreateThreeDSSessionResponse) UnsetDirectoryServerId() {
	o.DirectoryServerId.Unset()
}

// GetRecommendedVersion returns the RecommendedVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThreeDSSessionResponse) GetRecommendedVersion() string {
	if o == nil || IsNil(o.RecommendedVersion.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendedVersion.Get()
}

// GetRecommendedVersionOk returns a tuple with the RecommendedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThreeDSSessionResponse) GetRecommendedVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendedVersion.Get(), o.RecommendedVersion.IsSet()
}

// HasRecommendedVersion returns a boolean if a field is not nil.
func (o *CreateThreeDSSessionResponse) HasRecommendedVersion() bool {
	if o != nil && !IsNil(o.RecommendedVersion) {
		return true
	}

	return false
}

// SetRecommendedVersion gets a reference to the given NullableString and assigns it to the RecommendedVersion field.
func (o *CreateThreeDSSessionResponse) SetRecommendedVersion(v string) {
	o.RecommendedVersion.Set(&v)
}

// SetRecommendedVersionNil sets the value for RecommendedVersion to be an explicit nil
func (o *CreateThreeDSSessionResponse) SetRecommendedVersionNil() {
	o.RecommendedVersion.Set(nil)
}

// UnsetRecommendedVersion ensures that no value is present for RecommendedVersion, not even an explicit nil
func (o *CreateThreeDSSessionResponse) UnsetRecommendedVersion() {
	o.RecommendedVersion.Unset()
}

func (o CreateThreeDSSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThreeDSSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.CardBrand.IsSet() {
		toSerialize["cardBrand"] = o.CardBrand.Get()
	}
	if o.MethodUrl.IsSet() {
		toSerialize["method_url"] = o.MethodUrl.Get()
	}
	if o.MethodNotificationUrl.IsSet() {
		toSerialize["method_notification_url"] = o.MethodNotificationUrl.Get()
	}
	if o.DirectoryServerId.IsSet() {
		toSerialize["directory_server_id"] = o.DirectoryServerId.Get()
	}
	if o.RecommendedVersion.IsSet() {
		toSerialize["recommended_version"] = o.RecommendedVersion.Get()
	}
	return toSerialize, nil
}

type NullableCreateThreeDSSessionResponse struct {
	value *CreateThreeDSSessionResponse
	isSet bool
}

func (v NullableCreateThreeDSSessionResponse) Get() *CreateThreeDSSessionResponse {
	return v.value
}

func (v *NullableCreateThreeDSSessionResponse) Set(val *CreateThreeDSSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThreeDSSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThreeDSSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThreeDSSessionResponse(val *CreateThreeDSSessionResponse) *NullableCreateThreeDSSessionResponse {
	return &NullableCreateThreeDSSessionResponse{value: val, isSet: true}
}

func (v NullableCreateThreeDSSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThreeDSSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}