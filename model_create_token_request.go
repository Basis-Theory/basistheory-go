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

// checks if the CreateTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTokenRequest{}

// CreateTokenRequest struct for CreateTokenRequest
type CreateTokenRequest struct {
	Id                    NullableString    `json:"id,omitempty"`
	Type                  NullableString    `json:"type,omitempty"`
	Data                  interface{}       `json:"data,omitempty"`
	Privacy               *Privacy          `json:"privacy,omitempty"`
	Metadata              map[string]string `json:"metadata,omitempty"`
	SearchIndexes         []string          `json:"search_indexes,omitempty"`
	FingerprintExpression NullableString    `json:"fingerprint_expression,omitempty"`
	Mask                  interface{}       `json:"mask,omitempty"`
	DeduplicateToken      NullableBool      `json:"deduplicate_token,omitempty"`
	ExpiresAt             NullableString    `json:"expires_at,omitempty"`
	Containers            []string          `json:"containers,omitempty"`
	TokenIntentId         NullableString    `json:"token_intent_id,omitempty"`
}

// NewCreateTokenRequest instantiates a new CreateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenRequest() *CreateTokenRequest {
	this := CreateTokenRequest{}
	return &this
}

// NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenRequestWithDefaults() *CreateTokenRequest {
	this := CreateTokenRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateTokenRequest) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateTokenRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateTokenRequest) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CreateTokenRequest) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateTokenRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateTokenRequest) UnsetType() {
	o.Type.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *CreateTokenRequest) SetData(v interface{}) {
	o.Data = v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *CreateTokenRequest) GetPrivacy() Privacy {
	if o == nil || IsNil(o.Privacy) {
		var ret Privacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenRequest) GetPrivacyOk() (*Privacy, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given Privacy and assigns it to the Privacy field.
func (o *CreateTokenRequest) SetPrivacy(v Privacy) {
	o.Privacy = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateTokenRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetSearchIndexes returns the SearchIndexes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetSearchIndexes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SearchIndexes
}

// GetSearchIndexesOk returns a tuple with the SearchIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetSearchIndexesOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchIndexes) {
		return nil, false
	}
	return o.SearchIndexes, true
}

// HasSearchIndexes returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasSearchIndexes() bool {
	if o != nil && !IsNil(o.SearchIndexes) {
		return true
	}

	return false
}

// SetSearchIndexes gets a reference to the given []string and assigns it to the SearchIndexes field.
func (o *CreateTokenRequest) SetSearchIndexes(v []string) {
	o.SearchIndexes = v
}

// GetFingerprintExpression returns the FingerprintExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetFingerprintExpression() string {
	if o == nil || IsNil(o.FingerprintExpression.Get()) {
		var ret string
		return ret
	}
	return *o.FingerprintExpression.Get()
}

// GetFingerprintExpressionOk returns a tuple with the FingerprintExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetFingerprintExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FingerprintExpression.Get(), o.FingerprintExpression.IsSet()
}

// HasFingerprintExpression returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasFingerprintExpression() bool {
	if o != nil && !IsNil(o.FingerprintExpression) {
		return true
	}

	return false
}

// SetFingerprintExpression gets a reference to the given NullableString and assigns it to the FingerprintExpression field.
func (o *CreateTokenRequest) SetFingerprintExpression(v string) {
	o.FingerprintExpression.Set(&v)
}

// SetFingerprintExpressionNil sets the value for FingerprintExpression to be an explicit nil
func (o *CreateTokenRequest) SetFingerprintExpressionNil() {
	o.FingerprintExpression.Set(nil)
}

// UnsetFingerprintExpression ensures that no value is present for FingerprintExpression, not even an explicit nil
func (o *CreateTokenRequest) UnsetFingerprintExpression() {
	o.FingerprintExpression.Unset()
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetMask() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetMaskOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return &o.Mask, true
}

// HasMask returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given interface{} and assigns it to the Mask field.
func (o *CreateTokenRequest) SetMask(v interface{}) {
	o.Mask = v
}

// GetDeduplicateToken returns the DeduplicateToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetDeduplicateToken() bool {
	if o == nil || IsNil(o.DeduplicateToken.Get()) {
		var ret bool
		return ret
	}
	return *o.DeduplicateToken.Get()
}

// GetDeduplicateTokenOk returns a tuple with the DeduplicateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetDeduplicateTokenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeduplicateToken.Get(), o.DeduplicateToken.IsSet()
}

// HasDeduplicateToken returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasDeduplicateToken() bool {
	if o != nil && !IsNil(o.DeduplicateToken) {
		return true
	}

	return false
}

// SetDeduplicateToken gets a reference to the given NullableBool and assigns it to the DeduplicateToken field.
func (o *CreateTokenRequest) SetDeduplicateToken(v bool) {
	o.DeduplicateToken.Set(&v)
}

// SetDeduplicateTokenNil sets the value for DeduplicateToken to be an explicit nil
func (o *CreateTokenRequest) SetDeduplicateTokenNil() {
	o.DeduplicateToken.Set(nil)
}

// UnsetDeduplicateToken ensures that no value is present for DeduplicateToken, not even an explicit nil
func (o *CreateTokenRequest) UnsetDeduplicateToken() {
	o.DeduplicateToken.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableString and assigns it to the ExpiresAt field.
func (o *CreateTokenRequest) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *CreateTokenRequest) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *CreateTokenRequest) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetContainers returns the Containers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetContainers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetContainersOk() ([]string, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []string and assigns it to the Containers field.
func (o *CreateTokenRequest) SetContainers(v []string) {
	o.Containers = v
}

// GetTokenIntentId returns the TokenIntentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenRequest) GetTokenIntentId() string {
	if o == nil || IsNil(o.TokenIntentId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenIntentId.Get()
}

// GetTokenIntentIdOk returns a tuple with the TokenIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenRequest) GetTokenIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenIntentId.Get(), o.TokenIntentId.IsSet()
}

// HasTokenIntentId returns a boolean if a field is not nil.
func (o *CreateTokenRequest) HasTokenIntentId() bool {
	if o != nil && !IsNil(o.TokenIntentId) {
		return true
	}

	return false
}

// SetTokenIntentId gets a reference to the given NullableString and assigns it to the TokenIntentId field.
func (o *CreateTokenRequest) SetTokenIntentId(v string) {
	o.TokenIntentId.Set(&v)
}

// SetTokenIntentIdNil sets the value for TokenIntentId to be an explicit nil
func (o *CreateTokenRequest) SetTokenIntentIdNil() {
	o.TokenIntentId.Set(nil)
}

// UnsetTokenIntentId ensures that no value is present for TokenIntentId, not even an explicit nil
func (o *CreateTokenRequest) UnsetTokenIntentId() {
	o.TokenIntentId.Unset()
}

func (o CreateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.SearchIndexes != nil {
		toSerialize["search_indexes"] = o.SearchIndexes
	}
	if o.FingerprintExpression.IsSet() {
		toSerialize["fingerprint_expression"] = o.FingerprintExpression.Get()
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	if o.DeduplicateToken.IsSet() {
		toSerialize["deduplicate_token"] = o.DeduplicateToken.Get()
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.TokenIntentId.IsSet() {
		toSerialize["token_intent_id"] = o.TokenIntentId.Get()
	}
	return toSerialize, nil
}

type NullableCreateTokenRequest struct {
	value *CreateTokenRequest
	isSet bool
}

func (v NullableCreateTokenRequest) Get() *CreateTokenRequest {
	return v.value
}

func (v *NullableCreateTokenRequest) Set(val *CreateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenRequest(val *CreateTokenRequest) *NullableCreateTokenRequest {
	return &NullableCreateTokenRequest{value: val, isSet: true}
}

func (v NullableCreateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
