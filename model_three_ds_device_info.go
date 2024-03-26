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

// checks if the ThreeDSDeviceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreeDSDeviceInfo{}

// ThreeDSDeviceInfo struct for ThreeDSDeviceInfo
type ThreeDSDeviceInfo struct {
	BrowserAcceptHeader      NullableString                 `json:"browser_accept_header,omitempty"`
	BrowserIp                NullableString                 `json:"browser_ip,omitempty"`
	BrowserJavascriptEnabled NullableBool                   `json:"browser_javascript_enabled,omitempty"`
	BrowserJavaEnabled       NullableBool                   `json:"browser_java_enabled,omitempty"`
	BrowserLanguage          NullableString                 `json:"browser_language,omitempty"`
	BrowserColorDepth        NullableString                 `json:"browser_color_depth,omitempty"`
	BrowserScreenHeight      NullableString                 `json:"browser_screen_height,omitempty"`
	BrowserScreenWidth       NullableString                 `json:"browser_screen_width,omitempty"`
	BrowserTz                NullableString                 `json:"browser_tz,omitempty"`
	BrowserUserAgent         NullableString                 `json:"browser_user_agent,omitempty"`
	SdkTransactionId         NullableString                 `json:"sdk_transaction_id,omitempty"`
	SdkApplicationId         NullableString                 `json:"sdk_application_id,omitempty"`
	SdkEncryptionData        NullableString                 `json:"sdk_encryption_data,omitempty"`
	SdkEphemeralPublicKey    NullableString                 `json:"sdk_ephemeral_public_key,omitempty"`
	SdkMaxTimeout            NullableString                 `json:"sdk_max_timeout,omitempty"`
	SdkReferenceNumber       NullableString                 `json:"sdk_reference_number,omitempty"`
	SdkRenderOptions         *ThreeDSMobileSdkRenderOptions `json:"sdk_render_options,omitempty"`
}

// NewThreeDSDeviceInfo instantiates a new ThreeDSDeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSDeviceInfo() *ThreeDSDeviceInfo {
	this := ThreeDSDeviceInfo{}
	return &this
}

// NewThreeDSDeviceInfoWithDefaults instantiates a new ThreeDSDeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSDeviceInfoWithDefaults() *ThreeDSDeviceInfo {
	this := ThreeDSDeviceInfo{}
	return &this
}

// GetBrowserAcceptHeader returns the BrowserAcceptHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserAcceptHeader() string {
	if o == nil || IsNil(o.BrowserAcceptHeader.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserAcceptHeader.Get()
}

// GetBrowserAcceptHeaderOk returns a tuple with the BrowserAcceptHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserAcceptHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserAcceptHeader.Get(), o.BrowserAcceptHeader.IsSet()
}

// HasBrowserAcceptHeader returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserAcceptHeader() bool {
	if o != nil && !IsNil(o.BrowserAcceptHeader) {
		return true
	}

	return false
}

// SetBrowserAcceptHeader gets a reference to the given NullableString and assigns it to the BrowserAcceptHeader field.
func (o *ThreeDSDeviceInfo) SetBrowserAcceptHeader(v string) {
	o.BrowserAcceptHeader.Set(&v)
}

// SetBrowserAcceptHeaderNil sets the value for BrowserAcceptHeader to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserAcceptHeaderNil() {
	o.BrowserAcceptHeader.Set(nil)
}

// UnsetBrowserAcceptHeader ensures that no value is present for BrowserAcceptHeader, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserAcceptHeader() {
	o.BrowserAcceptHeader.Unset()
}

// GetBrowserIp returns the BrowserIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserIp() string {
	if o == nil || IsNil(o.BrowserIp.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserIp.Get()
}

// GetBrowserIpOk returns a tuple with the BrowserIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserIp.Get(), o.BrowserIp.IsSet()
}

// HasBrowserIp returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserIp() bool {
	if o != nil && !IsNil(o.BrowserIp) {
		return true
	}

	return false
}

// SetBrowserIp gets a reference to the given NullableString and assigns it to the BrowserIp field.
func (o *ThreeDSDeviceInfo) SetBrowserIp(v string) {
	o.BrowserIp.Set(&v)
}

// SetBrowserIpNil sets the value for BrowserIp to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserIpNil() {
	o.BrowserIp.Set(nil)
}

// UnsetBrowserIp ensures that no value is present for BrowserIp, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserIp() {
	o.BrowserIp.Unset()
}

// GetBrowserJavascriptEnabled returns the BrowserJavascriptEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserJavascriptEnabled() bool {
	if o == nil || IsNil(o.BrowserJavascriptEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.BrowserJavascriptEnabled.Get()
}

// GetBrowserJavascriptEnabledOk returns a tuple with the BrowserJavascriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserJavascriptEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserJavascriptEnabled.Get(), o.BrowserJavascriptEnabled.IsSet()
}

// HasBrowserJavascriptEnabled returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserJavascriptEnabled() bool {
	if o != nil && !IsNil(o.BrowserJavascriptEnabled) {
		return true
	}

	return false
}

// SetBrowserJavascriptEnabled gets a reference to the given NullableBool and assigns it to the BrowserJavascriptEnabled field.
func (o *ThreeDSDeviceInfo) SetBrowserJavascriptEnabled(v bool) {
	o.BrowserJavascriptEnabled.Set(&v)
}

// SetBrowserJavascriptEnabledNil sets the value for BrowserJavascriptEnabled to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserJavascriptEnabledNil() {
	o.BrowserJavascriptEnabled.Set(nil)
}

// UnsetBrowserJavascriptEnabled ensures that no value is present for BrowserJavascriptEnabled, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserJavascriptEnabled() {
	o.BrowserJavascriptEnabled.Unset()
}

// GetBrowserJavaEnabled returns the BrowserJavaEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserJavaEnabled() bool {
	if o == nil || IsNil(o.BrowserJavaEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.BrowserJavaEnabled.Get()
}

// GetBrowserJavaEnabledOk returns a tuple with the BrowserJavaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserJavaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserJavaEnabled.Get(), o.BrowserJavaEnabled.IsSet()
}

// HasBrowserJavaEnabled returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserJavaEnabled() bool {
	if o != nil && !IsNil(o.BrowserJavaEnabled) {
		return true
	}

	return false
}

// SetBrowserJavaEnabled gets a reference to the given NullableBool and assigns it to the BrowserJavaEnabled field.
func (o *ThreeDSDeviceInfo) SetBrowserJavaEnabled(v bool) {
	o.BrowserJavaEnabled.Set(&v)
}

// SetBrowserJavaEnabledNil sets the value for BrowserJavaEnabled to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserJavaEnabledNil() {
	o.BrowserJavaEnabled.Set(nil)
}

// UnsetBrowserJavaEnabled ensures that no value is present for BrowserJavaEnabled, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserJavaEnabled() {
	o.BrowserJavaEnabled.Unset()
}

// GetBrowserLanguage returns the BrowserLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserLanguage() string {
	if o == nil || IsNil(o.BrowserLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserLanguage.Get()
}

// GetBrowserLanguageOk returns a tuple with the BrowserLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserLanguage.Get(), o.BrowserLanguage.IsSet()
}

// HasBrowserLanguage returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserLanguage() bool {
	if o != nil && !IsNil(o.BrowserLanguage) {
		return true
	}

	return false
}

// SetBrowserLanguage gets a reference to the given NullableString and assigns it to the BrowserLanguage field.
func (o *ThreeDSDeviceInfo) SetBrowserLanguage(v string) {
	o.BrowserLanguage.Set(&v)
}

// SetBrowserLanguageNil sets the value for BrowserLanguage to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserLanguageNil() {
	o.BrowserLanguage.Set(nil)
}

// UnsetBrowserLanguage ensures that no value is present for BrowserLanguage, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserLanguage() {
	o.BrowserLanguage.Unset()
}

// GetBrowserColorDepth returns the BrowserColorDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserColorDepth() string {
	if o == nil || IsNil(o.BrowserColorDepth.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserColorDepth.Get()
}

// GetBrowserColorDepthOk returns a tuple with the BrowserColorDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserColorDepthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserColorDepth.Get(), o.BrowserColorDepth.IsSet()
}

// HasBrowserColorDepth returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserColorDepth() bool {
	if o != nil && !IsNil(o.BrowserColorDepth) {
		return true
	}

	return false
}

// SetBrowserColorDepth gets a reference to the given NullableString and assigns it to the BrowserColorDepth field.
func (o *ThreeDSDeviceInfo) SetBrowserColorDepth(v string) {
	o.BrowserColorDepth.Set(&v)
}

// SetBrowserColorDepthNil sets the value for BrowserColorDepth to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserColorDepthNil() {
	o.BrowserColorDepth.Set(nil)
}

// UnsetBrowserColorDepth ensures that no value is present for BrowserColorDepth, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserColorDepth() {
	o.BrowserColorDepth.Unset()
}

// GetBrowserScreenHeight returns the BrowserScreenHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserScreenHeight() string {
	if o == nil || IsNil(o.BrowserScreenHeight.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserScreenHeight.Get()
}

// GetBrowserScreenHeightOk returns a tuple with the BrowserScreenHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserScreenHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserScreenHeight.Get(), o.BrowserScreenHeight.IsSet()
}

// HasBrowserScreenHeight returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserScreenHeight() bool {
	if o != nil && !IsNil(o.BrowserScreenHeight) {
		return true
	}

	return false
}

// SetBrowserScreenHeight gets a reference to the given NullableString and assigns it to the BrowserScreenHeight field.
func (o *ThreeDSDeviceInfo) SetBrowserScreenHeight(v string) {
	o.BrowserScreenHeight.Set(&v)
}

// SetBrowserScreenHeightNil sets the value for BrowserScreenHeight to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserScreenHeightNil() {
	o.BrowserScreenHeight.Set(nil)
}

// UnsetBrowserScreenHeight ensures that no value is present for BrowserScreenHeight, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserScreenHeight() {
	o.BrowserScreenHeight.Unset()
}

// GetBrowserScreenWidth returns the BrowserScreenWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserScreenWidth() string {
	if o == nil || IsNil(o.BrowserScreenWidth.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserScreenWidth.Get()
}

// GetBrowserScreenWidthOk returns a tuple with the BrowserScreenWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserScreenWidthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserScreenWidth.Get(), o.BrowserScreenWidth.IsSet()
}

// HasBrowserScreenWidth returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserScreenWidth() bool {
	if o != nil && !IsNil(o.BrowserScreenWidth) {
		return true
	}

	return false
}

// SetBrowserScreenWidth gets a reference to the given NullableString and assigns it to the BrowserScreenWidth field.
func (o *ThreeDSDeviceInfo) SetBrowserScreenWidth(v string) {
	o.BrowserScreenWidth.Set(&v)
}

// SetBrowserScreenWidthNil sets the value for BrowserScreenWidth to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserScreenWidthNil() {
	o.BrowserScreenWidth.Set(nil)
}

// UnsetBrowserScreenWidth ensures that no value is present for BrowserScreenWidth, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserScreenWidth() {
	o.BrowserScreenWidth.Unset()
}

// GetBrowserTz returns the BrowserTz field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserTz() string {
	if o == nil || IsNil(o.BrowserTz.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserTz.Get()
}

// GetBrowserTzOk returns a tuple with the BrowserTz field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserTzOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserTz.Get(), o.BrowserTz.IsSet()
}

// HasBrowserTz returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserTz() bool {
	if o != nil && !IsNil(o.BrowserTz) {
		return true
	}

	return false
}

// SetBrowserTz gets a reference to the given NullableString and assigns it to the BrowserTz field.
func (o *ThreeDSDeviceInfo) SetBrowserTz(v string) {
	o.BrowserTz.Set(&v)
}

// SetBrowserTzNil sets the value for BrowserTz to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserTzNil() {
	o.BrowserTz.Set(nil)
}

// UnsetBrowserTz ensures that no value is present for BrowserTz, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserTz() {
	o.BrowserTz.Unset()
}

// GetBrowserUserAgent returns the BrowserUserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetBrowserUserAgent() string {
	if o == nil || IsNil(o.BrowserUserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.BrowserUserAgent.Get()
}

// GetBrowserUserAgentOk returns a tuple with the BrowserUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetBrowserUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserUserAgent.Get(), o.BrowserUserAgent.IsSet()
}

// HasBrowserUserAgent returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasBrowserUserAgent() bool {
	if o != nil && !IsNil(o.BrowserUserAgent) {
		return true
	}

	return false
}

// SetBrowserUserAgent gets a reference to the given NullableString and assigns it to the BrowserUserAgent field.
func (o *ThreeDSDeviceInfo) SetBrowserUserAgent(v string) {
	o.BrowserUserAgent.Set(&v)
}

// SetBrowserUserAgentNil sets the value for BrowserUserAgent to be an explicit nil
func (o *ThreeDSDeviceInfo) SetBrowserUserAgentNil() {
	o.BrowserUserAgent.Set(nil)
}

// UnsetBrowserUserAgent ensures that no value is present for BrowserUserAgent, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetBrowserUserAgent() {
	o.BrowserUserAgent.Unset()
}

// GetSdkTransactionId returns the SdkTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetSdkTransactionId() string {
	if o == nil || IsNil(o.SdkTransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.SdkTransactionId.Get()
}

// GetSdkTransactionIdOk returns a tuple with the SdkTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetSdkTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkTransactionId.Get(), o.SdkTransactionId.IsSet()
}

// HasSdkTransactionId returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkTransactionId() bool {
	if o != nil && !IsNil(o.SdkTransactionId) {
		return true
	}

	return false
}

// SetSdkTransactionId gets a reference to the given NullableString and assigns it to the SdkTransactionId field.
func (o *ThreeDSDeviceInfo) SetSdkTransactionId(v string) {
	o.SdkTransactionId.Set(&v)
}

// SetSdkTransactionIdNil sets the value for SdkTransactionId to be an explicit nil
func (o *ThreeDSDeviceInfo) SetSdkTransactionIdNil() {
	o.SdkTransactionId.Set(nil)
}

// UnsetSdkTransactionId ensures that no value is present for SdkTransactionId, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetSdkTransactionId() {
	o.SdkTransactionId.Unset()
}

// GetSdkApplicationId returns the SdkApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetSdkApplicationId() string {
	if o == nil || IsNil(o.SdkApplicationId.Get()) {
		var ret string
		return ret
	}
	return *o.SdkApplicationId.Get()
}

// GetSdkApplicationIdOk returns a tuple with the SdkApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetSdkApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkApplicationId.Get(), o.SdkApplicationId.IsSet()
}

// HasSdkApplicationId returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkApplicationId() bool {
	if o != nil && !IsNil(o.SdkApplicationId) {
		return true
	}

	return false
}

// SetSdkApplicationId gets a reference to the given NullableString and assigns it to the SdkApplicationId field.
func (o *ThreeDSDeviceInfo) SetSdkApplicationId(v string) {
	o.SdkApplicationId.Set(&v)
}

// SetSdkApplicationIdNil sets the value for SdkApplicationId to be an explicit nil
func (o *ThreeDSDeviceInfo) SetSdkApplicationIdNil() {
	o.SdkApplicationId.Set(nil)
}

// UnsetSdkApplicationId ensures that no value is present for SdkApplicationId, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetSdkApplicationId() {
	o.SdkApplicationId.Unset()
}

// GetSdkEncryptionData returns the SdkEncryptionData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetSdkEncryptionData() string {
	if o == nil || IsNil(o.SdkEncryptionData.Get()) {
		var ret string
		return ret
	}
	return *o.SdkEncryptionData.Get()
}

// GetSdkEncryptionDataOk returns a tuple with the SdkEncryptionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetSdkEncryptionDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkEncryptionData.Get(), o.SdkEncryptionData.IsSet()
}

// HasSdkEncryptionData returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkEncryptionData() bool {
	if o != nil && !IsNil(o.SdkEncryptionData) {
		return true
	}

	return false
}

// SetSdkEncryptionData gets a reference to the given NullableString and assigns it to the SdkEncryptionData field.
func (o *ThreeDSDeviceInfo) SetSdkEncryptionData(v string) {
	o.SdkEncryptionData.Set(&v)
}

// SetSdkEncryptionDataNil sets the value for SdkEncryptionData to be an explicit nil
func (o *ThreeDSDeviceInfo) SetSdkEncryptionDataNil() {
	o.SdkEncryptionData.Set(nil)
}

// UnsetSdkEncryptionData ensures that no value is present for SdkEncryptionData, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetSdkEncryptionData() {
	o.SdkEncryptionData.Unset()
}

// GetSdkEphemeralPublicKey returns the SdkEphemeralPublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetSdkEphemeralPublicKey() string {
	if o == nil || IsNil(o.SdkEphemeralPublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.SdkEphemeralPublicKey.Get()
}

// GetSdkEphemeralPublicKeyOk returns a tuple with the SdkEphemeralPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetSdkEphemeralPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkEphemeralPublicKey.Get(), o.SdkEphemeralPublicKey.IsSet()
}

// HasSdkEphemeralPublicKey returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkEphemeralPublicKey() bool {
	if o != nil && !IsNil(o.SdkEphemeralPublicKey) {
		return true
	}

	return false
}

// SetSdkEphemeralPublicKey gets a reference to the given NullableString and assigns it to the SdkEphemeralPublicKey field.
func (o *ThreeDSDeviceInfo) SetSdkEphemeralPublicKey(v string) {
	o.SdkEphemeralPublicKey.Set(&v)
}

// SetSdkEphemeralPublicKeyNil sets the value for SdkEphemeralPublicKey to be an explicit nil
func (o *ThreeDSDeviceInfo) SetSdkEphemeralPublicKeyNil() {
	o.SdkEphemeralPublicKey.Set(nil)
}

// UnsetSdkEphemeralPublicKey ensures that no value is present for SdkEphemeralPublicKey, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetSdkEphemeralPublicKey() {
	o.SdkEphemeralPublicKey.Unset()
}

// GetSdkMaxTimeout returns the SdkMaxTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetSdkMaxTimeout() string {
	if o == nil || IsNil(o.SdkMaxTimeout.Get()) {
		var ret string
		return ret
	}
	return *o.SdkMaxTimeout.Get()
}

// GetSdkMaxTimeoutOk returns a tuple with the SdkMaxTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetSdkMaxTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkMaxTimeout.Get(), o.SdkMaxTimeout.IsSet()
}

// HasSdkMaxTimeout returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkMaxTimeout() bool {
	if o != nil && !IsNil(o.SdkMaxTimeout) {
		return true
	}

	return false
}

// SetSdkMaxTimeout gets a reference to the given NullableString and assigns it to the SdkMaxTimeout field.
func (o *ThreeDSDeviceInfo) SetSdkMaxTimeout(v string) {
	o.SdkMaxTimeout.Set(&v)
}

// SetSdkMaxTimeoutNil sets the value for SdkMaxTimeout to be an explicit nil
func (o *ThreeDSDeviceInfo) SetSdkMaxTimeoutNil() {
	o.SdkMaxTimeout.Set(nil)
}

// UnsetSdkMaxTimeout ensures that no value is present for SdkMaxTimeout, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetSdkMaxTimeout() {
	o.SdkMaxTimeout.Unset()
}

// GetSdkReferenceNumber returns the SdkReferenceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSDeviceInfo) GetSdkReferenceNumber() string {
	if o == nil || IsNil(o.SdkReferenceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SdkReferenceNumber.Get()
}

// GetSdkReferenceNumberOk returns a tuple with the SdkReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSDeviceInfo) GetSdkReferenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkReferenceNumber.Get(), o.SdkReferenceNumber.IsSet()
}

// HasSdkReferenceNumber returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkReferenceNumber() bool {
	if o != nil && !IsNil(o.SdkReferenceNumber) {
		return true
	}

	return false
}

// SetSdkReferenceNumber gets a reference to the given NullableString and assigns it to the SdkReferenceNumber field.
func (o *ThreeDSDeviceInfo) SetSdkReferenceNumber(v string) {
	o.SdkReferenceNumber.Set(&v)
}

// SetSdkReferenceNumberNil sets the value for SdkReferenceNumber to be an explicit nil
func (o *ThreeDSDeviceInfo) SetSdkReferenceNumberNil() {
	o.SdkReferenceNumber.Set(nil)
}

// UnsetSdkReferenceNumber ensures that no value is present for SdkReferenceNumber, not even an explicit nil
func (o *ThreeDSDeviceInfo) UnsetSdkReferenceNumber() {
	o.SdkReferenceNumber.Unset()
}

// GetSdkRenderOptions returns the SdkRenderOptions field value if set, zero value otherwise.
func (o *ThreeDSDeviceInfo) GetSdkRenderOptions() ThreeDSMobileSdkRenderOptions {
	if o == nil || IsNil(o.SdkRenderOptions) {
		var ret ThreeDSMobileSdkRenderOptions
		return ret
	}
	return *o.SdkRenderOptions
}

// GetSdkRenderOptionsOk returns a tuple with the SdkRenderOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSDeviceInfo) GetSdkRenderOptionsOk() (*ThreeDSMobileSdkRenderOptions, bool) {
	if o == nil || IsNil(o.SdkRenderOptions) {
		return nil, false
	}
	return o.SdkRenderOptions, true
}

// HasSdkRenderOptions returns a boolean if a field is not nil.
func (o *ThreeDSDeviceInfo) HasSdkRenderOptions() bool {
	if o != nil && !IsNil(o.SdkRenderOptions) {
		return true
	}

	return false
}

// SetSdkRenderOptions gets a reference to the given ThreeDSMobileSdkRenderOptions and assigns it to the SdkRenderOptions field.
func (o *ThreeDSDeviceInfo) SetSdkRenderOptions(v ThreeDSMobileSdkRenderOptions) {
	o.SdkRenderOptions = &v
}

func (o ThreeDSDeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSDeviceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BrowserAcceptHeader.IsSet() {
		toSerialize["browser_accept_header"] = o.BrowserAcceptHeader.Get()
	}
	if o.BrowserIp.IsSet() {
		toSerialize["browser_ip"] = o.BrowserIp.Get()
	}
	if o.BrowserJavascriptEnabled.IsSet() {
		toSerialize["browser_javascript_enabled"] = o.BrowserJavascriptEnabled.Get()
	}
	if o.BrowserJavaEnabled.IsSet() {
		toSerialize["browser_java_enabled"] = o.BrowserJavaEnabled.Get()
	}
	if o.BrowserLanguage.IsSet() {
		toSerialize["browser_language"] = o.BrowserLanguage.Get()
	}
	if o.BrowserColorDepth.IsSet() {
		toSerialize["browser_color_depth"] = o.BrowserColorDepth.Get()
	}
	if o.BrowserScreenHeight.IsSet() {
		toSerialize["browser_screen_height"] = o.BrowserScreenHeight.Get()
	}
	if o.BrowserScreenWidth.IsSet() {
		toSerialize["browser_screen_width"] = o.BrowserScreenWidth.Get()
	}
	if o.BrowserTz.IsSet() {
		toSerialize["browser_tz"] = o.BrowserTz.Get()
	}
	if o.BrowserUserAgent.IsSet() {
		toSerialize["browser_user_agent"] = o.BrowserUserAgent.Get()
	}
	if o.SdkTransactionId.IsSet() {
		toSerialize["sdk_transaction_id"] = o.SdkTransactionId.Get()
	}
	if o.SdkApplicationId.IsSet() {
		toSerialize["sdk_application_id"] = o.SdkApplicationId.Get()
	}
	if o.SdkEncryptionData.IsSet() {
		toSerialize["sdk_encryption_data"] = o.SdkEncryptionData.Get()
	}
	if o.SdkEphemeralPublicKey.IsSet() {
		toSerialize["sdk_ephemeral_public_key"] = o.SdkEphemeralPublicKey.Get()
	}
	if o.SdkMaxTimeout.IsSet() {
		toSerialize["sdk_max_timeout"] = o.SdkMaxTimeout.Get()
	}
	if o.SdkReferenceNumber.IsSet() {
		toSerialize["sdk_reference_number"] = o.SdkReferenceNumber.Get()
	}
	if !IsNil(o.SdkRenderOptions) {
		toSerialize["sdk_render_options"] = o.SdkRenderOptions
	}
	return toSerialize, nil
}

type NullableThreeDSDeviceInfo struct {
	value *ThreeDSDeviceInfo
	isSet bool
}

func (v NullableThreeDSDeviceInfo) Get() *ThreeDSDeviceInfo {
	return v.value
}

func (v *NullableThreeDSDeviceInfo) Set(val *ThreeDSDeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSDeviceInfo(val *ThreeDSDeviceInfo) *NullableThreeDSDeviceInfo {
	return &NullableThreeDSDeviceInfo{value: val, isSet: true}
}

func (v NullableThreeDSDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}