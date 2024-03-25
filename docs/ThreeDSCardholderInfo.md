# ThreeDSCardholderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AccountType** | Pointer to **NullableString** |  | [optional] 
**AccountInfo** | Pointer to [**ThreeDSCardholderAccountInfo**](ThreeDSCardholderAccountInfo.md) |  | [optional] 
**AuthenticationInfo** | Pointer to [**ThreeDSCardholderAuthenticationInfo**](ThreeDSCardholderAuthenticationInfo.md) |  | [optional] 
**PriorAuthenticationInfo** | Pointer to [**ThreeDSPriorAuthenticationInfo**](ThreeDSPriorAuthenticationInfo.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to [**ThreeDSCardholderPhoneNumber**](ThreeDSCardholderPhoneNumber.md) |  | [optional] 
**MobilePhoneNumber** | Pointer to [**ThreeDSCardholderPhoneNumber**](ThreeDSCardholderPhoneNumber.md) |  | [optional] 
**WorkPhoneNumber** | Pointer to [**ThreeDSCardholderPhoneNumber**](ThreeDSCardholderPhoneNumber.md) |  | [optional] 
**BillingShippingAddressMatch** | Pointer to **NullableString** |  | [optional] 
**BillingAddress** | Pointer to [**ThreeDSAddress**](ThreeDSAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**ThreeDSAddress**](ThreeDSAddress.md) |  | [optional] 

## Methods

### NewThreeDSCardholderInfo

`func NewThreeDSCardholderInfo() *ThreeDSCardholderInfo`

NewThreeDSCardholderInfo instantiates a new ThreeDSCardholderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSCardholderInfoWithDefaults

`func NewThreeDSCardholderInfoWithDefaults() *ThreeDSCardholderInfo`

NewThreeDSCardholderInfoWithDefaults instantiates a new ThreeDSCardholderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ThreeDSCardholderInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ThreeDSCardholderInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ThreeDSCardholderInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ThreeDSCardholderInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *ThreeDSCardholderInfo) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *ThreeDSCardholderInfo) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountType

`func (o *ThreeDSCardholderInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ThreeDSCardholderInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ThreeDSCardholderInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ThreeDSCardholderInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### SetAccountTypeNil

`func (o *ThreeDSCardholderInfo) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *ThreeDSCardholderInfo) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil
### GetAccountInfo

`func (o *ThreeDSCardholderInfo) GetAccountInfo() ThreeDSCardholderAccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *ThreeDSCardholderInfo) GetAccountInfoOk() (*ThreeDSCardholderAccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *ThreeDSCardholderInfo) SetAccountInfo(v ThreeDSCardholderAccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *ThreeDSCardholderInfo) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetAuthenticationInfo

`func (o *ThreeDSCardholderInfo) GetAuthenticationInfo() ThreeDSCardholderAuthenticationInfo`

GetAuthenticationInfo returns the AuthenticationInfo field if non-nil, zero value otherwise.

### GetAuthenticationInfoOk

`func (o *ThreeDSCardholderInfo) GetAuthenticationInfoOk() (*ThreeDSCardholderAuthenticationInfo, bool)`

GetAuthenticationInfoOk returns a tuple with the AuthenticationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationInfo

`func (o *ThreeDSCardholderInfo) SetAuthenticationInfo(v ThreeDSCardholderAuthenticationInfo)`

SetAuthenticationInfo sets AuthenticationInfo field to given value.

### HasAuthenticationInfo

`func (o *ThreeDSCardholderInfo) HasAuthenticationInfo() bool`

HasAuthenticationInfo returns a boolean if a field has been set.

### GetPriorAuthenticationInfo

`func (o *ThreeDSCardholderInfo) GetPriorAuthenticationInfo() ThreeDSPriorAuthenticationInfo`

GetPriorAuthenticationInfo returns the PriorAuthenticationInfo field if non-nil, zero value otherwise.

### GetPriorAuthenticationInfoOk

`func (o *ThreeDSCardholderInfo) GetPriorAuthenticationInfoOk() (*ThreeDSPriorAuthenticationInfo, bool)`

GetPriorAuthenticationInfoOk returns a tuple with the PriorAuthenticationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorAuthenticationInfo

`func (o *ThreeDSCardholderInfo) SetPriorAuthenticationInfo(v ThreeDSPriorAuthenticationInfo)`

SetPriorAuthenticationInfo sets PriorAuthenticationInfo field to given value.

### HasPriorAuthenticationInfo

`func (o *ThreeDSCardholderInfo) HasPriorAuthenticationInfo() bool`

HasPriorAuthenticationInfo returns a boolean if a field has been set.

### GetName

`func (o *ThreeDSCardholderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreeDSCardholderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreeDSCardholderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreeDSCardholderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ThreeDSCardholderInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ThreeDSCardholderInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *ThreeDSCardholderInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ThreeDSCardholderInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ThreeDSCardholderInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ThreeDSCardholderInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ThreeDSCardholderInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ThreeDSCardholderInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhoneNumber

`func (o *ThreeDSCardholderInfo) GetPhoneNumber() ThreeDSCardholderPhoneNumber`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ThreeDSCardholderInfo) GetPhoneNumberOk() (*ThreeDSCardholderPhoneNumber, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ThreeDSCardholderInfo) SetPhoneNumber(v ThreeDSCardholderPhoneNumber)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ThreeDSCardholderInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *ThreeDSCardholderInfo) GetMobilePhoneNumber() ThreeDSCardholderPhoneNumber`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *ThreeDSCardholderInfo) GetMobilePhoneNumberOk() (*ThreeDSCardholderPhoneNumber, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *ThreeDSCardholderInfo) SetMobilePhoneNumber(v ThreeDSCardholderPhoneNumber)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *ThreeDSCardholderInfo) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetWorkPhoneNumber

`func (o *ThreeDSCardholderInfo) GetWorkPhoneNumber() ThreeDSCardholderPhoneNumber`

GetWorkPhoneNumber returns the WorkPhoneNumber field if non-nil, zero value otherwise.

### GetWorkPhoneNumberOk

`func (o *ThreeDSCardholderInfo) GetWorkPhoneNumberOk() (*ThreeDSCardholderPhoneNumber, bool)`

GetWorkPhoneNumberOk returns a tuple with the WorkPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhoneNumber

`func (o *ThreeDSCardholderInfo) SetWorkPhoneNumber(v ThreeDSCardholderPhoneNumber)`

SetWorkPhoneNumber sets WorkPhoneNumber field to given value.

### HasWorkPhoneNumber

`func (o *ThreeDSCardholderInfo) HasWorkPhoneNumber() bool`

HasWorkPhoneNumber returns a boolean if a field has been set.

### GetBillingShippingAddressMatch

`func (o *ThreeDSCardholderInfo) GetBillingShippingAddressMatch() string`

GetBillingShippingAddressMatch returns the BillingShippingAddressMatch field if non-nil, zero value otherwise.

### GetBillingShippingAddressMatchOk

`func (o *ThreeDSCardholderInfo) GetBillingShippingAddressMatchOk() (*string, bool)`

GetBillingShippingAddressMatchOk returns a tuple with the BillingShippingAddressMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingShippingAddressMatch

`func (o *ThreeDSCardholderInfo) SetBillingShippingAddressMatch(v string)`

SetBillingShippingAddressMatch sets BillingShippingAddressMatch field to given value.

### HasBillingShippingAddressMatch

`func (o *ThreeDSCardholderInfo) HasBillingShippingAddressMatch() bool`

HasBillingShippingAddressMatch returns a boolean if a field has been set.

### SetBillingShippingAddressMatchNil

`func (o *ThreeDSCardholderInfo) SetBillingShippingAddressMatchNil(b bool)`

 SetBillingShippingAddressMatchNil sets the value for BillingShippingAddressMatch to be an explicit nil

### UnsetBillingShippingAddressMatch
`func (o *ThreeDSCardholderInfo) UnsetBillingShippingAddressMatch()`

UnsetBillingShippingAddressMatch ensures that no value is present for BillingShippingAddressMatch, not even an explicit nil
### GetBillingAddress

`func (o *ThreeDSCardholderInfo) GetBillingAddress() ThreeDSAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ThreeDSCardholderInfo) GetBillingAddressOk() (*ThreeDSAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ThreeDSCardholderInfo) SetBillingAddress(v ThreeDSAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ThreeDSCardholderInfo) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *ThreeDSCardholderInfo) GetShippingAddress() ThreeDSAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *ThreeDSCardholderInfo) GetShippingAddressOk() (*ThreeDSAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *ThreeDSCardholderInfo) SetShippingAddress(v ThreeDSAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *ThreeDSCardholderInfo) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


