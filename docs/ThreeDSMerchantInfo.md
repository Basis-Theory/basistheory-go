# ThreeDSMerchantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mid** | Pointer to **NullableString** |  | [optional] 
**AcquirerBin** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**CategoryCode** | Pointer to **NullableString** |  | [optional] 
**RiskInfo** | Pointer to [**ThreeDSMerchantRiskInfo**](ThreeDSMerchantRiskInfo.md) |  | [optional] 

## Methods

### NewThreeDSMerchantInfo

`func NewThreeDSMerchantInfo() *ThreeDSMerchantInfo`

NewThreeDSMerchantInfo instantiates a new ThreeDSMerchantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSMerchantInfoWithDefaults

`func NewThreeDSMerchantInfoWithDefaults() *ThreeDSMerchantInfo`

NewThreeDSMerchantInfoWithDefaults instantiates a new ThreeDSMerchantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMid

`func (o *ThreeDSMerchantInfo) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *ThreeDSMerchantInfo) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *ThreeDSMerchantInfo) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *ThreeDSMerchantInfo) HasMid() bool`

HasMid returns a boolean if a field has been set.

### SetMidNil

`func (o *ThreeDSMerchantInfo) SetMidNil(b bool)`

 SetMidNil sets the value for Mid to be an explicit nil

### UnsetMid
`func (o *ThreeDSMerchantInfo) UnsetMid()`

UnsetMid ensures that no value is present for Mid, not even an explicit nil
### GetAcquirerBin

`func (o *ThreeDSMerchantInfo) GetAcquirerBin() string`

GetAcquirerBin returns the AcquirerBin field if non-nil, zero value otherwise.

### GetAcquirerBinOk

`func (o *ThreeDSMerchantInfo) GetAcquirerBinOk() (*string, bool)`

GetAcquirerBinOk returns a tuple with the AcquirerBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBin

`func (o *ThreeDSMerchantInfo) SetAcquirerBin(v string)`

SetAcquirerBin sets AcquirerBin field to given value.

### HasAcquirerBin

`func (o *ThreeDSMerchantInfo) HasAcquirerBin() bool`

HasAcquirerBin returns a boolean if a field has been set.

### SetAcquirerBinNil

`func (o *ThreeDSMerchantInfo) SetAcquirerBinNil(b bool)`

 SetAcquirerBinNil sets the value for AcquirerBin to be an explicit nil

### UnsetAcquirerBin
`func (o *ThreeDSMerchantInfo) UnsetAcquirerBin()`

UnsetAcquirerBin ensures that no value is present for AcquirerBin, not even an explicit nil
### GetName

`func (o *ThreeDSMerchantInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreeDSMerchantInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreeDSMerchantInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreeDSMerchantInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ThreeDSMerchantInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ThreeDSMerchantInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCountryCode

`func (o *ThreeDSMerchantInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ThreeDSMerchantInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ThreeDSMerchantInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ThreeDSMerchantInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *ThreeDSMerchantInfo) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *ThreeDSMerchantInfo) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetCategoryCode

`func (o *ThreeDSMerchantInfo) GetCategoryCode() string`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *ThreeDSMerchantInfo) GetCategoryCodeOk() (*string, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *ThreeDSMerchantInfo) SetCategoryCode(v string)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *ThreeDSMerchantInfo) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### SetCategoryCodeNil

`func (o *ThreeDSMerchantInfo) SetCategoryCodeNil(b bool)`

 SetCategoryCodeNil sets the value for CategoryCode to be an explicit nil

### UnsetCategoryCode
`func (o *ThreeDSMerchantInfo) UnsetCategoryCode()`

UnsetCategoryCode ensures that no value is present for CategoryCode, not even an explicit nil
### GetRiskInfo

`func (o *ThreeDSMerchantInfo) GetRiskInfo() ThreeDSMerchantRiskInfo`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *ThreeDSMerchantInfo) GetRiskInfoOk() (*ThreeDSMerchantRiskInfo, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *ThreeDSMerchantInfo) SetRiskInfo(v ThreeDSMerchantRiskInfo)`

SetRiskInfo sets RiskInfo field to given value.

### HasRiskInfo

`func (o *ThreeDSMerchantInfo) HasRiskInfo() bool`

HasRiskInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


