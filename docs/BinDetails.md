# BinDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Prepaid** | Pointer to **NullableBool** |  | [optional] 
**CardSegmentType** | Pointer to **NullableString** |  | [optional] 
**Bank** | Pointer to [**BinDetailsBank**](BinDetailsBank.md) |  | [optional] 
**Product** | Pointer to [**BinDetailsProduct**](BinDetailsProduct.md) |  | [optional] 
**Country** | Pointer to [**BinDetailsCountry**](BinDetailsCountry.md) |  | [optional] 
**Reloadable** | Pointer to **NullableBool** |  | [optional] 
**PanOrToken** | Pointer to **NullableString** |  | [optional] 
**AccountUpdater** | Pointer to **NullableBool** |  | [optional] 
**Alm** | Pointer to **NullableBool** |  | [optional] 
**DomesticOnly** | Pointer to **NullableBool** |  | [optional] 
**GamblingBlocked** | Pointer to **NullableBool** |  | [optional] 
**Level2** | Pointer to **NullableBool** |  | [optional] 
**Level3** | Pointer to **NullableBool** |  | [optional] 
**IssuerCurrency** | Pointer to **NullableString** |  | [optional] 
**ComboCard** | Pointer to **NullableString** |  | [optional] 
**BinLength** | Pointer to **NullableInt32** |  | [optional] 
**Authentication** | Pointer to **interface{}** |  | [optional] 
**Cost** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewBinDetails

`func NewBinDetails() *BinDetails`

NewBinDetails instantiates a new BinDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinDetailsWithDefaults

`func NewBinDetailsWithDefaults() *BinDetails`

NewBinDetailsWithDefaults instantiates a new BinDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBrand

`func (o *BinDetails) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *BinDetails) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *BinDetails) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *BinDetails) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *BinDetails) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *BinDetails) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetType

`func (o *BinDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BinDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BinDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BinDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BinDetails) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BinDetails) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPrepaid

`func (o *BinDetails) GetPrepaid() bool`

GetPrepaid returns the Prepaid field if non-nil, zero value otherwise.

### GetPrepaidOk

`func (o *BinDetails) GetPrepaidOk() (*bool, bool)`

GetPrepaidOk returns a tuple with the Prepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaid

`func (o *BinDetails) SetPrepaid(v bool)`

SetPrepaid sets Prepaid field to given value.

### HasPrepaid

`func (o *BinDetails) HasPrepaid() bool`

HasPrepaid returns a boolean if a field has been set.

### SetPrepaidNil

`func (o *BinDetails) SetPrepaidNil(b bool)`

 SetPrepaidNil sets the value for Prepaid to be an explicit nil

### UnsetPrepaid
`func (o *BinDetails) UnsetPrepaid()`

UnsetPrepaid ensures that no value is present for Prepaid, not even an explicit nil
### GetCardSegmentType

`func (o *BinDetails) GetCardSegmentType() string`

GetCardSegmentType returns the CardSegmentType field if non-nil, zero value otherwise.

### GetCardSegmentTypeOk

`func (o *BinDetails) GetCardSegmentTypeOk() (*string, bool)`

GetCardSegmentTypeOk returns a tuple with the CardSegmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSegmentType

`func (o *BinDetails) SetCardSegmentType(v string)`

SetCardSegmentType sets CardSegmentType field to given value.

### HasCardSegmentType

`func (o *BinDetails) HasCardSegmentType() bool`

HasCardSegmentType returns a boolean if a field has been set.

### SetCardSegmentTypeNil

`func (o *BinDetails) SetCardSegmentTypeNil(b bool)`

 SetCardSegmentTypeNil sets the value for CardSegmentType to be an explicit nil

### UnsetCardSegmentType
`func (o *BinDetails) UnsetCardSegmentType()`

UnsetCardSegmentType ensures that no value is present for CardSegmentType, not even an explicit nil
### GetBank

`func (o *BinDetails) GetBank() BinDetailsBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *BinDetails) GetBankOk() (*BinDetailsBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *BinDetails) SetBank(v BinDetailsBank)`

SetBank sets Bank field to given value.

### HasBank

`func (o *BinDetails) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetProduct

`func (o *BinDetails) GetProduct() BinDetailsProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BinDetails) GetProductOk() (*BinDetailsProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BinDetails) SetProduct(v BinDetailsProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BinDetails) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCountry

`func (o *BinDetails) GetCountry() BinDetailsCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BinDetails) GetCountryOk() (*BinDetailsCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BinDetails) SetCountry(v BinDetailsCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BinDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetReloadable

`func (o *BinDetails) GetReloadable() bool`

GetReloadable returns the Reloadable field if non-nil, zero value otherwise.

### GetReloadableOk

`func (o *BinDetails) GetReloadableOk() (*bool, bool)`

GetReloadableOk returns a tuple with the Reloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloadable

`func (o *BinDetails) SetReloadable(v bool)`

SetReloadable sets Reloadable field to given value.

### HasReloadable

`func (o *BinDetails) HasReloadable() bool`

HasReloadable returns a boolean if a field has been set.

### SetReloadableNil

`func (o *BinDetails) SetReloadableNil(b bool)`

 SetReloadableNil sets the value for Reloadable to be an explicit nil

### UnsetReloadable
`func (o *BinDetails) UnsetReloadable()`

UnsetReloadable ensures that no value is present for Reloadable, not even an explicit nil
### GetPanOrToken

`func (o *BinDetails) GetPanOrToken() string`

GetPanOrToken returns the PanOrToken field if non-nil, zero value otherwise.

### GetPanOrTokenOk

`func (o *BinDetails) GetPanOrTokenOk() (*string, bool)`

GetPanOrTokenOk returns a tuple with the PanOrToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanOrToken

`func (o *BinDetails) SetPanOrToken(v string)`

SetPanOrToken sets PanOrToken field to given value.

### HasPanOrToken

`func (o *BinDetails) HasPanOrToken() bool`

HasPanOrToken returns a boolean if a field has been set.

### SetPanOrTokenNil

`func (o *BinDetails) SetPanOrTokenNil(b bool)`

 SetPanOrTokenNil sets the value for PanOrToken to be an explicit nil

### UnsetPanOrToken
`func (o *BinDetails) UnsetPanOrToken()`

UnsetPanOrToken ensures that no value is present for PanOrToken, not even an explicit nil
### GetAccountUpdater

`func (o *BinDetails) GetAccountUpdater() bool`

GetAccountUpdater returns the AccountUpdater field if non-nil, zero value otherwise.

### GetAccountUpdaterOk

`func (o *BinDetails) GetAccountUpdaterOk() (*bool, bool)`

GetAccountUpdaterOk returns a tuple with the AccountUpdater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdater

`func (o *BinDetails) SetAccountUpdater(v bool)`

SetAccountUpdater sets AccountUpdater field to given value.

### HasAccountUpdater

`func (o *BinDetails) HasAccountUpdater() bool`

HasAccountUpdater returns a boolean if a field has been set.

### SetAccountUpdaterNil

`func (o *BinDetails) SetAccountUpdaterNil(b bool)`

 SetAccountUpdaterNil sets the value for AccountUpdater to be an explicit nil

### UnsetAccountUpdater
`func (o *BinDetails) UnsetAccountUpdater()`

UnsetAccountUpdater ensures that no value is present for AccountUpdater, not even an explicit nil
### GetAlm

`func (o *BinDetails) GetAlm() bool`

GetAlm returns the Alm field if non-nil, zero value otherwise.

### GetAlmOk

`func (o *BinDetails) GetAlmOk() (*bool, bool)`

GetAlmOk returns a tuple with the Alm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlm

`func (o *BinDetails) SetAlm(v bool)`

SetAlm sets Alm field to given value.

### HasAlm

`func (o *BinDetails) HasAlm() bool`

HasAlm returns a boolean if a field has been set.

### SetAlmNil

`func (o *BinDetails) SetAlmNil(b bool)`

 SetAlmNil sets the value for Alm to be an explicit nil

### UnsetAlm
`func (o *BinDetails) UnsetAlm()`

UnsetAlm ensures that no value is present for Alm, not even an explicit nil
### GetDomesticOnly

`func (o *BinDetails) GetDomesticOnly() bool`

GetDomesticOnly returns the DomesticOnly field if non-nil, zero value otherwise.

### GetDomesticOnlyOk

`func (o *BinDetails) GetDomesticOnlyOk() (*bool, bool)`

GetDomesticOnlyOk returns a tuple with the DomesticOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticOnly

`func (o *BinDetails) SetDomesticOnly(v bool)`

SetDomesticOnly sets DomesticOnly field to given value.

### HasDomesticOnly

`func (o *BinDetails) HasDomesticOnly() bool`

HasDomesticOnly returns a boolean if a field has been set.

### SetDomesticOnlyNil

`func (o *BinDetails) SetDomesticOnlyNil(b bool)`

 SetDomesticOnlyNil sets the value for DomesticOnly to be an explicit nil

### UnsetDomesticOnly
`func (o *BinDetails) UnsetDomesticOnly()`

UnsetDomesticOnly ensures that no value is present for DomesticOnly, not even an explicit nil
### GetGamblingBlocked

`func (o *BinDetails) GetGamblingBlocked() bool`

GetGamblingBlocked returns the GamblingBlocked field if non-nil, zero value otherwise.

### GetGamblingBlockedOk

`func (o *BinDetails) GetGamblingBlockedOk() (*bool, bool)`

GetGamblingBlockedOk returns a tuple with the GamblingBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamblingBlocked

`func (o *BinDetails) SetGamblingBlocked(v bool)`

SetGamblingBlocked sets GamblingBlocked field to given value.

### HasGamblingBlocked

`func (o *BinDetails) HasGamblingBlocked() bool`

HasGamblingBlocked returns a boolean if a field has been set.

### SetGamblingBlockedNil

`func (o *BinDetails) SetGamblingBlockedNil(b bool)`

 SetGamblingBlockedNil sets the value for GamblingBlocked to be an explicit nil

### UnsetGamblingBlocked
`func (o *BinDetails) UnsetGamblingBlocked()`

UnsetGamblingBlocked ensures that no value is present for GamblingBlocked, not even an explicit nil
### GetLevel2

`func (o *BinDetails) GetLevel2() bool`

GetLevel2 returns the Level2 field if non-nil, zero value otherwise.

### GetLevel2Ok

`func (o *BinDetails) GetLevel2Ok() (*bool, bool)`

GetLevel2Ok returns a tuple with the Level2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel2

`func (o *BinDetails) SetLevel2(v bool)`

SetLevel2 sets Level2 field to given value.

### HasLevel2

`func (o *BinDetails) HasLevel2() bool`

HasLevel2 returns a boolean if a field has been set.

### SetLevel2Nil

`func (o *BinDetails) SetLevel2Nil(b bool)`

 SetLevel2Nil sets the value for Level2 to be an explicit nil

### UnsetLevel2
`func (o *BinDetails) UnsetLevel2()`

UnsetLevel2 ensures that no value is present for Level2, not even an explicit nil
### GetLevel3

`func (o *BinDetails) GetLevel3() bool`

GetLevel3 returns the Level3 field if non-nil, zero value otherwise.

### GetLevel3Ok

`func (o *BinDetails) GetLevel3Ok() (*bool, bool)`

GetLevel3Ok returns a tuple with the Level3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel3

`func (o *BinDetails) SetLevel3(v bool)`

SetLevel3 sets Level3 field to given value.

### HasLevel3

`func (o *BinDetails) HasLevel3() bool`

HasLevel3 returns a boolean if a field has been set.

### SetLevel3Nil

`func (o *BinDetails) SetLevel3Nil(b bool)`

 SetLevel3Nil sets the value for Level3 to be an explicit nil

### UnsetLevel3
`func (o *BinDetails) UnsetLevel3()`

UnsetLevel3 ensures that no value is present for Level3, not even an explicit nil
### GetIssuerCurrency

`func (o *BinDetails) GetIssuerCurrency() string`

GetIssuerCurrency returns the IssuerCurrency field if non-nil, zero value otherwise.

### GetIssuerCurrencyOk

`func (o *BinDetails) GetIssuerCurrencyOk() (*string, bool)`

GetIssuerCurrencyOk returns a tuple with the IssuerCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCurrency

`func (o *BinDetails) SetIssuerCurrency(v string)`

SetIssuerCurrency sets IssuerCurrency field to given value.

### HasIssuerCurrency

`func (o *BinDetails) HasIssuerCurrency() bool`

HasIssuerCurrency returns a boolean if a field has been set.

### SetIssuerCurrencyNil

`func (o *BinDetails) SetIssuerCurrencyNil(b bool)`

 SetIssuerCurrencyNil sets the value for IssuerCurrency to be an explicit nil

### UnsetIssuerCurrency
`func (o *BinDetails) UnsetIssuerCurrency()`

UnsetIssuerCurrency ensures that no value is present for IssuerCurrency, not even an explicit nil
### GetComboCard

`func (o *BinDetails) GetComboCard() string`

GetComboCard returns the ComboCard field if non-nil, zero value otherwise.

### GetComboCardOk

`func (o *BinDetails) GetComboCardOk() (*string, bool)`

GetComboCardOk returns a tuple with the ComboCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComboCard

`func (o *BinDetails) SetComboCard(v string)`

SetComboCard sets ComboCard field to given value.

### HasComboCard

`func (o *BinDetails) HasComboCard() bool`

HasComboCard returns a boolean if a field has been set.

### SetComboCardNil

`func (o *BinDetails) SetComboCardNil(b bool)`

 SetComboCardNil sets the value for ComboCard to be an explicit nil

### UnsetComboCard
`func (o *BinDetails) UnsetComboCard()`

UnsetComboCard ensures that no value is present for ComboCard, not even an explicit nil
### GetBinLength

`func (o *BinDetails) GetBinLength() int32`

GetBinLength returns the BinLength field if non-nil, zero value otherwise.

### GetBinLengthOk

`func (o *BinDetails) GetBinLengthOk() (*int32, bool)`

GetBinLengthOk returns a tuple with the BinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinLength

`func (o *BinDetails) SetBinLength(v int32)`

SetBinLength sets BinLength field to given value.

### HasBinLength

`func (o *BinDetails) HasBinLength() bool`

HasBinLength returns a boolean if a field has been set.

### SetBinLengthNil

`func (o *BinDetails) SetBinLengthNil(b bool)`

 SetBinLengthNil sets the value for BinLength to be an explicit nil

### UnsetBinLength
`func (o *BinDetails) UnsetBinLength()`

UnsetBinLength ensures that no value is present for BinLength, not even an explicit nil
### GetAuthentication

`func (o *BinDetails) GetAuthentication() interface{}`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *BinDetails) GetAuthenticationOk() (*interface{}, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *BinDetails) SetAuthentication(v interface{})`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *BinDetails) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *BinDetails) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *BinDetails) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil
### GetCost

`func (o *BinDetails) GetCost() interface{}`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BinDetails) GetCostOk() (*interface{}, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BinDetails) SetCost(v interface{})`

SetCost sets Cost field to given value.

### HasCost

`func (o *BinDetails) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *BinDetails) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *BinDetails) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


