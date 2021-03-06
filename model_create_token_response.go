/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Server to Server Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"time"
)

// CreateTokenResponse struct for CreateTokenResponse
type CreateTokenResponse struct {
	Id                    NullableString    `json:"id,omitempty"`
	TenantId              *string           `json:"tenant_id,omitempty"`
	Type                  NullableString    `json:"type,omitempty"`
	Fingerprint           NullableString    `json:"fingerprint,omitempty"`
	FingerprintExpression NullableString    `json:"fingerprint_expression,omitempty"`
	Mask                  interface{}       `json:"mask,omitempty"`
	Data                  interface{}       `json:"data,omitempty"`
	Metadata              map[string]string `json:"metadata,omitempty"`
	Privacy               *Privacy          `json:"privacy,omitempty"`
	SearchIndexes         []string          `json:"search_indexes,omitempty"`
	CreatedBy             NullableString    `json:"created_by,omitempty"`
	CreatedAt             NullableTime      `json:"created_at,omitempty"`
	ModifiedBy            NullableString    `json:"modified_by,omitempty"`
	ModifiedAt            NullableTime      `json:"modified_at,omitempty"`
	ExpiresAt             NullableTime      `json:"expires_at,omitempty"`
}

// NewCreateTokenResponse instantiates a new CreateTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenResponse() *CreateTokenResponse {
	this := CreateTokenResponse{}
	return &this
}

// NewCreateTokenResponseWithDefaults instantiates a new CreateTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenResponseWithDefaults() *CreateTokenResponse {
	this := CreateTokenResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateTokenResponse) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateTokenResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateTokenResponse) UnsetId() {
	o.Id.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CreateTokenResponse) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CreateTokenResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CreateTokenResponse) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *CreateTokenResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CreateTokenResponse) UnsetType() {
	o.Type.Unset()
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetFingerprint() string {
	if o == nil || o.Fingerprint.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasFingerprint() bool {
	if o != nil && o.Fingerprint.IsSet() {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given NullableString and assigns it to the Fingerprint field.
func (o *CreateTokenResponse) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}

// SetFingerprintNil sets the value for Fingerprint to be an explicit nil
func (o *CreateTokenResponse) SetFingerprintNil() {
	o.Fingerprint.Set(nil)
}

// UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
func (o *CreateTokenResponse) UnsetFingerprint() {
	o.Fingerprint.Unset()
}

// GetFingerprintExpression returns the FingerprintExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetFingerprintExpression() string {
	if o == nil || o.FingerprintExpression.Get() == nil {
		var ret string
		return ret
	}
	return *o.FingerprintExpression.Get()
}

// GetFingerprintExpressionOk returns a tuple with the FingerprintExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetFingerprintExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FingerprintExpression.Get(), o.FingerprintExpression.IsSet()
}

// HasFingerprintExpression returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasFingerprintExpression() bool {
	if o != nil && o.FingerprintExpression.IsSet() {
		return true
	}

	return false
}

// SetFingerprintExpression gets a reference to the given NullableString and assigns it to the FingerprintExpression field.
func (o *CreateTokenResponse) SetFingerprintExpression(v string) {
	o.FingerprintExpression.Set(&v)
}

// SetFingerprintExpressionNil sets the value for FingerprintExpression to be an explicit nil
func (o *CreateTokenResponse) SetFingerprintExpressionNil() {
	o.FingerprintExpression.Set(nil)
}

// UnsetFingerprintExpression ensures that no value is present for FingerprintExpression, not even an explicit nil
func (o *CreateTokenResponse) UnsetFingerprintExpression() {
	o.FingerprintExpression.Unset()
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetMask() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetMaskOk() (*interface{}, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return &o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given interface{} and assigns it to the Mask field.
func (o *CreateTokenResponse) SetMask(v interface{}) {
	o.Mask = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetDataOk() (*interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *CreateTokenResponse) SetData(v interface{}) {
	o.Data = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateTokenResponse) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *CreateTokenResponse) GetPrivacy() Privacy {
	if o == nil || o.Privacy == nil {
		var ret Privacy
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTokenResponse) GetPrivacyOk() (*Privacy, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given Privacy and assigns it to the Privacy field.
func (o *CreateTokenResponse) SetPrivacy(v Privacy) {
	o.Privacy = &v
}

// GetSearchIndexes returns the SearchIndexes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetSearchIndexes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SearchIndexes
}

// GetSearchIndexesOk returns a tuple with the SearchIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetSearchIndexesOk() ([]string, bool) {
	if o == nil || o.SearchIndexes == nil {
		return nil, false
	}
	return o.SearchIndexes, true
}

// HasSearchIndexes returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasSearchIndexes() bool {
	if o != nil && o.SearchIndexes != nil {
		return true
	}

	return false
}

// SetSearchIndexes gets a reference to the given []string and assigns it to the SearchIndexes field.
func (o *CreateTokenResponse) SetSearchIndexes(v []string) {
	o.SearchIndexes = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *CreateTokenResponse) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *CreateTokenResponse) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *CreateTokenResponse) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *CreateTokenResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *CreateTokenResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *CreateTokenResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetModifiedBy() string {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableString and assigns it to the ModifiedBy field.
func (o *CreateTokenResponse) SetModifiedBy(v string) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *CreateTokenResponse) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *CreateTokenResponse) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt.IsSet() {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *CreateTokenResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *CreateTokenResponse) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *CreateTokenResponse) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTokenResponse) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTokenResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateTokenResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *CreateTokenResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *CreateTokenResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *CreateTokenResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o CreateTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Fingerprint.IsSet() {
		toSerialize["fingerprint"] = o.Fingerprint.Get()
	}
	if o.FingerprintExpression.IsSet() {
		toSerialize["fingerprint_expression"] = o.FingerprintExpression.Get()
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if o.SearchIndexes != nil {
		toSerialize["search_indexes"] = o.SearchIndexes
	}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modified_by"] = o.ModifiedBy.Get()
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokenResponse struct {
	value *CreateTokenResponse
	isSet bool
}

func (v NullableCreateTokenResponse) Get() *CreateTokenResponse {
	return v.value
}

func (v *NullableCreateTokenResponse) Set(val *CreateTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenResponse(val *CreateTokenResponse) *NullableCreateTokenResponse {
	return &NullableCreateTokenResponse{value: val, isSet: true}
}

func (v NullableCreateTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
