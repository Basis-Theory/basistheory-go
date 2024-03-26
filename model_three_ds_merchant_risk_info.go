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

// checks if the ThreeDSMerchantRiskInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreeDSMerchantRiskInfo{}

// ThreeDSMerchantRiskInfo struct for ThreeDSMerchantRiskInfo
type ThreeDSMerchantRiskInfo struct {
	DeliveryEmail     NullableString `json:"delivery_email,omitempty"`
	DeliveryTimeFrame NullableString `json:"delivery_time_frame,omitempty"`
	GiftCardAmount    NullableString `json:"gift_card_amount,omitempty"`
	GiftCardCount     NullableString `json:"gift_card_count,omitempty"`
	GiftCardCurrency  NullableString `json:"gift_card_currency,omitempty"`
	PreOrderPurchase  NullableBool   `json:"pre_order_purchase,omitempty"`
	PreOrderDate      NullableString `json:"pre_order_date,omitempty"`
	ReorderedPurchase NullableBool   `json:"reordered_purchase,omitempty"`
	ShippingMethod    NullableString `json:"shipping_method,omitempty"`
}

// NewThreeDSMerchantRiskInfo instantiates a new ThreeDSMerchantRiskInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSMerchantRiskInfo() *ThreeDSMerchantRiskInfo {
	this := ThreeDSMerchantRiskInfo{}
	return &this
}

// NewThreeDSMerchantRiskInfoWithDefaults instantiates a new ThreeDSMerchantRiskInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSMerchantRiskInfoWithDefaults() *ThreeDSMerchantRiskInfo {
	this := ThreeDSMerchantRiskInfo{}
	return &this
}

// GetDeliveryEmail returns the DeliveryEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetDeliveryEmail() string {
	if o == nil || IsNil(o.DeliveryEmail.Get()) {
		var ret string
		return ret
	}
	return *o.DeliveryEmail.Get()
}

// GetDeliveryEmailOk returns a tuple with the DeliveryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetDeliveryEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryEmail.Get(), o.DeliveryEmail.IsSet()
}

// HasDeliveryEmail returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasDeliveryEmail() bool {
	if o != nil && !IsNil(o.DeliveryEmail) {
		return true
	}

	return false
}

// SetDeliveryEmail gets a reference to the given NullableString and assigns it to the DeliveryEmail field.
func (o *ThreeDSMerchantRiskInfo) SetDeliveryEmail(v string) {
	o.DeliveryEmail.Set(&v)
}

// SetDeliveryEmailNil sets the value for DeliveryEmail to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetDeliveryEmailNil() {
	o.DeliveryEmail.Set(nil)
}

// UnsetDeliveryEmail ensures that no value is present for DeliveryEmail, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetDeliveryEmail() {
	o.DeliveryEmail.Unset()
}

// GetDeliveryTimeFrame returns the DeliveryTimeFrame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetDeliveryTimeFrame() string {
	if o == nil || IsNil(o.DeliveryTimeFrame.Get()) {
		var ret string
		return ret
	}
	return *o.DeliveryTimeFrame.Get()
}

// GetDeliveryTimeFrameOk returns a tuple with the DeliveryTimeFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetDeliveryTimeFrameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryTimeFrame.Get(), o.DeliveryTimeFrame.IsSet()
}

// HasDeliveryTimeFrame returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasDeliveryTimeFrame() bool {
	if o != nil && !IsNil(o.DeliveryTimeFrame) {
		return true
	}

	return false
}

// SetDeliveryTimeFrame gets a reference to the given NullableString and assigns it to the DeliveryTimeFrame field.
func (o *ThreeDSMerchantRiskInfo) SetDeliveryTimeFrame(v string) {
	o.DeliveryTimeFrame.Set(&v)
}

// SetDeliveryTimeFrameNil sets the value for DeliveryTimeFrame to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetDeliveryTimeFrameNil() {
	o.DeliveryTimeFrame.Set(nil)
}

// UnsetDeliveryTimeFrame ensures that no value is present for DeliveryTimeFrame, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetDeliveryTimeFrame() {
	o.DeliveryTimeFrame.Unset()
}

// GetGiftCardAmount returns the GiftCardAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetGiftCardAmount() string {
	if o == nil || IsNil(o.GiftCardAmount.Get()) {
		var ret string
		return ret
	}
	return *o.GiftCardAmount.Get()
}

// GetGiftCardAmountOk returns a tuple with the GiftCardAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetGiftCardAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftCardAmount.Get(), o.GiftCardAmount.IsSet()
}

// HasGiftCardAmount returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasGiftCardAmount() bool {
	if o != nil && !IsNil(o.GiftCardAmount) {
		return true
	}

	return false
}

// SetGiftCardAmount gets a reference to the given NullableString and assigns it to the GiftCardAmount field.
func (o *ThreeDSMerchantRiskInfo) SetGiftCardAmount(v string) {
	o.GiftCardAmount.Set(&v)
}

// SetGiftCardAmountNil sets the value for GiftCardAmount to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetGiftCardAmountNil() {
	o.GiftCardAmount.Set(nil)
}

// UnsetGiftCardAmount ensures that no value is present for GiftCardAmount, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetGiftCardAmount() {
	o.GiftCardAmount.Unset()
}

// GetGiftCardCount returns the GiftCardCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetGiftCardCount() string {
	if o == nil || IsNil(o.GiftCardCount.Get()) {
		var ret string
		return ret
	}
	return *o.GiftCardCount.Get()
}

// GetGiftCardCountOk returns a tuple with the GiftCardCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetGiftCardCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftCardCount.Get(), o.GiftCardCount.IsSet()
}

// HasGiftCardCount returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasGiftCardCount() bool {
	if o != nil && !IsNil(o.GiftCardCount) {
		return true
	}

	return false
}

// SetGiftCardCount gets a reference to the given NullableString and assigns it to the GiftCardCount field.
func (o *ThreeDSMerchantRiskInfo) SetGiftCardCount(v string) {
	o.GiftCardCount.Set(&v)
}

// SetGiftCardCountNil sets the value for GiftCardCount to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetGiftCardCountNil() {
	o.GiftCardCount.Set(nil)
}

// UnsetGiftCardCount ensures that no value is present for GiftCardCount, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetGiftCardCount() {
	o.GiftCardCount.Unset()
}

// GetGiftCardCurrency returns the GiftCardCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetGiftCardCurrency() string {
	if o == nil || IsNil(o.GiftCardCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.GiftCardCurrency.Get()
}

// GetGiftCardCurrencyOk returns a tuple with the GiftCardCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetGiftCardCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftCardCurrency.Get(), o.GiftCardCurrency.IsSet()
}

// HasGiftCardCurrency returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasGiftCardCurrency() bool {
	if o != nil && !IsNil(o.GiftCardCurrency) {
		return true
	}

	return false
}

// SetGiftCardCurrency gets a reference to the given NullableString and assigns it to the GiftCardCurrency field.
func (o *ThreeDSMerchantRiskInfo) SetGiftCardCurrency(v string) {
	o.GiftCardCurrency.Set(&v)
}

// SetGiftCardCurrencyNil sets the value for GiftCardCurrency to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetGiftCardCurrencyNil() {
	o.GiftCardCurrency.Set(nil)
}

// UnsetGiftCardCurrency ensures that no value is present for GiftCardCurrency, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetGiftCardCurrency() {
	o.GiftCardCurrency.Unset()
}

// GetPreOrderPurchase returns the PreOrderPurchase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetPreOrderPurchase() bool {
	if o == nil || IsNil(o.PreOrderPurchase.Get()) {
		var ret bool
		return ret
	}
	return *o.PreOrderPurchase.Get()
}

// GetPreOrderPurchaseOk returns a tuple with the PreOrderPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetPreOrderPurchaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreOrderPurchase.Get(), o.PreOrderPurchase.IsSet()
}

// HasPreOrderPurchase returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasPreOrderPurchase() bool {
	if o != nil && !IsNil(o.PreOrderPurchase) {
		return true
	}

	return false
}

// SetPreOrderPurchase gets a reference to the given NullableBool and assigns it to the PreOrderPurchase field.
func (o *ThreeDSMerchantRiskInfo) SetPreOrderPurchase(v bool) {
	o.PreOrderPurchase.Set(&v)
}

// SetPreOrderPurchaseNil sets the value for PreOrderPurchase to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetPreOrderPurchaseNil() {
	o.PreOrderPurchase.Set(nil)
}

// UnsetPreOrderPurchase ensures that no value is present for PreOrderPurchase, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetPreOrderPurchase() {
	o.PreOrderPurchase.Unset()
}

// GetPreOrderDate returns the PreOrderDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetPreOrderDate() string {
	if o == nil || IsNil(o.PreOrderDate.Get()) {
		var ret string
		return ret
	}
	return *o.PreOrderDate.Get()
}

// GetPreOrderDateOk returns a tuple with the PreOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetPreOrderDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreOrderDate.Get(), o.PreOrderDate.IsSet()
}

// HasPreOrderDate returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasPreOrderDate() bool {
	if o != nil && !IsNil(o.PreOrderDate) {
		return true
	}

	return false
}

// SetPreOrderDate gets a reference to the given NullableString and assigns it to the PreOrderDate field.
func (o *ThreeDSMerchantRiskInfo) SetPreOrderDate(v string) {
	o.PreOrderDate.Set(&v)
}

// SetPreOrderDateNil sets the value for PreOrderDate to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetPreOrderDateNil() {
	o.PreOrderDate.Set(nil)
}

// UnsetPreOrderDate ensures that no value is present for PreOrderDate, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetPreOrderDate() {
	o.PreOrderDate.Unset()
}

// GetReorderedPurchase returns the ReorderedPurchase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetReorderedPurchase() bool {
	if o == nil || IsNil(o.ReorderedPurchase.Get()) {
		var ret bool
		return ret
	}
	return *o.ReorderedPurchase.Get()
}

// GetReorderedPurchaseOk returns a tuple with the ReorderedPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetReorderedPurchaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReorderedPurchase.Get(), o.ReorderedPurchase.IsSet()
}

// HasReorderedPurchase returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasReorderedPurchase() bool {
	if o != nil && !IsNil(o.ReorderedPurchase) {
		return true
	}

	return false
}

// SetReorderedPurchase gets a reference to the given NullableBool and assigns it to the ReorderedPurchase field.
func (o *ThreeDSMerchantRiskInfo) SetReorderedPurchase(v bool) {
	o.ReorderedPurchase.Set(&v)
}

// SetReorderedPurchaseNil sets the value for ReorderedPurchase to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetReorderedPurchaseNil() {
	o.ReorderedPurchase.Set(nil)
}

// UnsetReorderedPurchase ensures that no value is present for ReorderedPurchase, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetReorderedPurchase() {
	o.ReorderedPurchase.Unset()
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSMerchantRiskInfo) GetShippingMethod() string {
	if o == nil || IsNil(o.ShippingMethod.Get()) {
		var ret string
		return ret
	}
	return *o.ShippingMethod.Get()
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSMerchantRiskInfo) GetShippingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingMethod.Get(), o.ShippingMethod.IsSet()
}

// HasShippingMethod returns a boolean if a field is not nil.
func (o *ThreeDSMerchantRiskInfo) HasShippingMethod() bool {
	if o != nil && !IsNil(o.ShippingMethod) {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given NullableString and assigns it to the ShippingMethod field.
func (o *ThreeDSMerchantRiskInfo) SetShippingMethod(v string) {
	o.ShippingMethod.Set(&v)
}

// SetShippingMethodNil sets the value for ShippingMethod to be an explicit nil
func (o *ThreeDSMerchantRiskInfo) SetShippingMethodNil() {
	o.ShippingMethod.Set(nil)
}

// UnsetShippingMethod ensures that no value is present for ShippingMethod, not even an explicit nil
func (o *ThreeDSMerchantRiskInfo) UnsetShippingMethod() {
	o.ShippingMethod.Unset()
}

func (o ThreeDSMerchantRiskInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDSMerchantRiskInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeliveryEmail.IsSet() {
		toSerialize["delivery_email"] = o.DeliveryEmail.Get()
	}
	if o.DeliveryTimeFrame.IsSet() {
		toSerialize["delivery_time_frame"] = o.DeliveryTimeFrame.Get()
	}
	if o.GiftCardAmount.IsSet() {
		toSerialize["gift_card_amount"] = o.GiftCardAmount.Get()
	}
	if o.GiftCardCount.IsSet() {
		toSerialize["gift_card_count"] = o.GiftCardCount.Get()
	}
	if o.GiftCardCurrency.IsSet() {
		toSerialize["gift_card_currency"] = o.GiftCardCurrency.Get()
	}
	if o.PreOrderPurchase.IsSet() {
		toSerialize["pre_order_purchase"] = o.PreOrderPurchase.Get()
	}
	if o.PreOrderDate.IsSet() {
		toSerialize["pre_order_date"] = o.PreOrderDate.Get()
	}
	if o.ReorderedPurchase.IsSet() {
		toSerialize["reordered_purchase"] = o.ReorderedPurchase.Get()
	}
	if o.ShippingMethod.IsSet() {
		toSerialize["shipping_method"] = o.ShippingMethod.Get()
	}
	return toSerialize, nil
}

type NullableThreeDSMerchantRiskInfo struct {
	value *ThreeDSMerchantRiskInfo
	isSet bool
}

func (v NullableThreeDSMerchantRiskInfo) Get() *ThreeDSMerchantRiskInfo {
	return v.value
}

func (v *NullableThreeDSMerchantRiskInfo) Set(val *ThreeDSMerchantRiskInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSMerchantRiskInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSMerchantRiskInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSMerchantRiskInfo(val *ThreeDSMerchantRiskInfo) *NullableThreeDSMerchantRiskInfo {
	return &NullableThreeDSMerchantRiskInfo{value: val, isSet: true}
}

func (v NullableThreeDSMerchantRiskInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSMerchantRiskInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}