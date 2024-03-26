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

// checks if the ThreeDSMerchantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreeDSMerchantInfo{}

// ThreeDSMerchantInfo struct for ThreeDSMerchantInfo
type ThreeDSMerchantInfo struct {
	Mid          NullableString           `json:"mid,omitempty"`
	AcquirerBin  NullableString           `json:"acquirer_bin,omitempty"`
	Name         NullableString           `json:"name,omitempty"`
	CountryCode  NullableString           `json:"country_code,omitempty"`
	CategoryCode NullableString           `json:"category_code,omitempty"`
	RiskInfo     *ThreeDSMerchantRiskInfo `json:"risk_info,omitempty"`
}

// NewThreeDSMerchantInfo instantiates a new ThreeDSMerchantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSMerchantInfo() *ThreeDSMerchantInfo {
	this := ThreeDSMerchantInfo{}
	return &this
}

// NewThreeDSMerchantInfoWithDefaults instantiates a new ThreeDSMerchantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSMerchantInfoWithDefaults() *ThreeDSMerchantInfo {
	this := ThreeDSMerchantInfo{}
	return &this
}

// GetMid returns the Mid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantInfo) GetMid() string {
	if o == nil || IsNil(o.Mid.Get()) {
		var ret string
		return ret
	}
	return *o.Mid.Get()
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantInfo) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mid.Get(), o.Mid.IsSet()
}

// HasMid returns a boolean if a field is not nil.
func (o *ThreeDSMerchantInfo) HasMid() bool {
	if o != nil && !IsNil(o.Mid) {
		return true
	}

	return false
}

// SetMid gets a reference to the given NullableString and assigns it to the Mid field.
func (o *ThreeDSMerchantInfo) SetMid(v string) {
	o.Mid.Set(&v)
}

// SetMidNil sets the value for Mid to be an explicit nil
func (o *ThreeDSMerchantInfo) SetMidNil() {
	o.Mid.Set(nil)
}

// UnsetMid ensures that no value is present for Mid, not even an explicit nil
func (o *ThreeDSMerchantInfo) UnsetMid() {
	o.Mid.Unset()
}

// GetAcquirerBin returns the AcquirerBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantInfo) GetAcquirerBin() string {
	if o == nil || IsNil(o.AcquirerBin.Get()) {
		var ret string
		return ret
	}
	return *o.AcquirerBin.Get()
}

// GetAcquirerBinOk returns a tuple with the AcquirerBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantInfo) GetAcquirerBinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcquirerBin.Get(), o.AcquirerBin.IsSet()
}

// HasAcquirerBin returns a boolean if a field is not nil.
func (o *ThreeDSMerchantInfo) HasAcquirerBin() bool {
	if o != nil && !IsNil(o.AcquirerBin) {
		return true
	}

	return false
}

// SetAcquirerBin gets a reference to the given NullableString and assigns it to the AcquirerBin field.
func (o *ThreeDSMerchantInfo) SetAcquirerBin(v string) {
	o.AcquirerBin.Set(&v)
}

// SetAcquirerBinNil sets the value for AcquirerBin to be an explicit nil
func (o *ThreeDSMerchantInfo) SetAcquirerBinNil() {
	o.AcquirerBin.Set(nil)
}

// UnsetAcquirerBin ensures that no value is present for AcquirerBin, not even an explicit nil
func (o *ThreeDSMerchantInfo) UnsetAcquirerBin() {
	o.AcquirerBin.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *ThreeDSMerchantInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ThreeDSMerchantInfo) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ThreeDSMerchantInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ThreeDSMerchantInfo) UnsetName() {
	o.Name.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field is not nil.
func (o *ThreeDSMerchantInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *ThreeDSMerchantInfo) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *ThreeDSMerchantInfo) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *ThreeDSMerchantInfo) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetCategoryCode returns the CategoryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantInfo) GetCategoryCode() string {
	if o == nil || IsNil(o.CategoryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryCode.Get()
}

// GetCategoryCodeOk returns a tuple with the CategoryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantInfo) GetCategoryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryCode.Get(), o.CategoryCode.IsSet()
}

// HasCategoryCode returns a boolean if a field is not nil.
func (o *ThreeDSMerchantInfo) HasCategoryCode() bool {
	if o != nil && !IsNil(o.CategoryCode) {
		return true
	}

	return false
}

// SetCategoryCode gets a reference to the given NullableString and assigns it to the CategoryCode field.
func (o *ThreeDSMerchantInfo) SetCategoryCode(v string) {
	o.CategoryCode.Set(&v)
}

// SetCategoryCodeNil sets the value for CategoryCode to be an explicit nil
func (o *ThreeDSMerchantInfo) SetCategoryCodeNil() {
	o.CategoryCode.Set(nil)
}

// UnsetCategoryCode ensures that no value is present for CategoryCode, not even an explicit nil
func (o *ThreeDSMerchantInfo) UnsetCategoryCode() {
	o.CategoryCode.Unset()
}

// GetRiskInfo returns the RiskInfo field value if set, zero value otherwise.
func (o *ThreeDSMerchantInfo) GetRiskInfo() ThreeDSMerchantRiskInfo {
	if o == nil || IsNil(o.RiskInfo) {
		var ret ThreeDSMerchantRiskInfo
		return ret
	}
	return *o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSMerchantInfo) GetRiskInfoOk() (*ThreeDSMerchantRiskInfo, bool) {
	if o == nil || IsNil(o.RiskInfo) {
		return nil, false
	}
	return o.RiskInfo, true
}

// HasRiskInfo returns a boolean if a field is not nil.
func (o *ThreeDSMerchantInfo) HasRiskInfo() bool {
	if o != nil && !IsNil(o.RiskInfo) {
		return true
	}

	return false
}

// SetRiskInfo gets a reference to the given ThreeDSMerchantRiskInfo and assigns it to the RiskInfo field.
func (o *ThreeDSMerchantInfo) SetRiskInfo(v ThreeDSMerchantRiskInfo) {
	o.RiskInfo = &v
}

func (o ThreeDSMerchantInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSMerchantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mid.IsSet() {
		toSerialize["mid"] = o.Mid.Get()
	}
	if o.AcquirerBin.IsSet() {
		toSerialize["acquirer_bin"] = o.AcquirerBin.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.CategoryCode.IsSet() {
		toSerialize["category_code"] = o.CategoryCode.Get()
	}
	if !IsNil(o.RiskInfo) {
		toSerialize["risk_info"] = o.RiskInfo
	}
	return toSerialize, nil
}

type NullableThreeDSMerchantInfo struct {
	value *ThreeDSMerchantInfo
	isSet bool
}

func (v NullableThreeDSMerchantInfo) Get() *ThreeDSMerchantInfo {
	return v.value
}

func (v *NullableThreeDSMerchantInfo) Set(val *ThreeDSMerchantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSMerchantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSMerchantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSMerchantInfo(val *ThreeDSMerchantInfo) *NullableThreeDSMerchantInfo {
	return &NullableThreeDSMerchantInfo{value: val, isSet: true}
}

func (v NullableThreeDSMerchantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSMerchantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}