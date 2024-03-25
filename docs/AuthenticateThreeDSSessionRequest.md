# AuthenticateThreeDSSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCategory** | **string** |  | 
**AuthenticationType** | **string** |  | 
**ChallengePreference** | Pointer to **NullableString** |  | [optional] 
**PurchaseInfo** | Pointer to [**ThreeDSPurchaseInfo**](ThreeDSPurchaseInfo.md) |  | [optional] 
**MerchantInfo** | Pointer to [**ThreeDSMerchantInfo**](ThreeDSMerchantInfo.md) |  | [optional] 
**RequestorInfo** | [**ThreeDSRequestorInfo**](ThreeDSRequestorInfo.md) |  | 
**CardholderInfo** | Pointer to [**ThreeDSCardholderInfo**](ThreeDSCardholderInfo.md) |  | [optional] 
**BroadcastInfo** | Pointer to **interface{}** |  | [optional] 
**MessageExtensions** | Pointer to [**[]ThreeDSMessageExtension**](ThreeDSMessageExtension.md) |  | [optional] 

## Methods

### NewAuthenticateThreeDSSessionRequest

`func NewAuthenticateThreeDSSessionRequest(authenticationCategory string, authenticationType string, requestorInfo ThreeDSRequestorInfo, ) *AuthenticateThreeDSSessionRequest`

NewAuthenticateThreeDSSessionRequest instantiates a new AuthenticateThreeDSSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticateThreeDSSessionRequestWithDefaults

`func NewAuthenticateThreeDSSessionRequestWithDefaults() *AuthenticateThreeDSSessionRequest`

NewAuthenticateThreeDSSessionRequestWithDefaults instantiates a new AuthenticateThreeDSSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCategory

`func (o *AuthenticateThreeDSSessionRequest) GetAuthenticationCategory() string`

GetAuthenticationCategory returns the AuthenticationCategory field if non-nil, zero value otherwise.

### GetAuthenticationCategoryOk

`func (o *AuthenticateThreeDSSessionRequest) GetAuthenticationCategoryOk() (*string, bool)`

GetAuthenticationCategoryOk returns a tuple with the AuthenticationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCategory

`func (o *AuthenticateThreeDSSessionRequest) SetAuthenticationCategory(v string)`

SetAuthenticationCategory sets AuthenticationCategory field to given value.


### GetAuthenticationType

`func (o *AuthenticateThreeDSSessionRequest) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AuthenticateThreeDSSessionRequest) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AuthenticateThreeDSSessionRequest) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetChallengePreference

`func (o *AuthenticateThreeDSSessionRequest) GetChallengePreference() string`

GetChallengePreference returns the ChallengePreference field if non-nil, zero value otherwise.

### GetChallengePreferenceOk

`func (o *AuthenticateThreeDSSessionRequest) GetChallengePreferenceOk() (*string, bool)`

GetChallengePreferenceOk returns a tuple with the ChallengePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengePreference

`func (o *AuthenticateThreeDSSessionRequest) SetChallengePreference(v string)`

SetChallengePreference sets ChallengePreference field to given value.

### HasChallengePreference

`func (o *AuthenticateThreeDSSessionRequest) HasChallengePreference() bool`

HasChallengePreference returns a boolean if a field has been set.

### SetChallengePreferenceNil

`func (o *AuthenticateThreeDSSessionRequest) SetChallengePreferenceNil(b bool)`

 SetChallengePreferenceNil sets the value for ChallengePreference to be an explicit nil

### UnsetChallengePreference
`func (o *AuthenticateThreeDSSessionRequest) UnsetChallengePreference()`

UnsetChallengePreference ensures that no value is present for ChallengePreference, not even an explicit nil
### GetPurchaseInfo

`func (o *AuthenticateThreeDSSessionRequest) GetPurchaseInfo() ThreeDSPurchaseInfo`

GetPurchaseInfo returns the PurchaseInfo field if non-nil, zero value otherwise.

### GetPurchaseInfoOk

`func (o *AuthenticateThreeDSSessionRequest) GetPurchaseInfoOk() (*ThreeDSPurchaseInfo, bool)`

GetPurchaseInfoOk returns a tuple with the PurchaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseInfo

`func (o *AuthenticateThreeDSSessionRequest) SetPurchaseInfo(v ThreeDSPurchaseInfo)`

SetPurchaseInfo sets PurchaseInfo field to given value.

### HasPurchaseInfo

`func (o *AuthenticateThreeDSSessionRequest) HasPurchaseInfo() bool`

HasPurchaseInfo returns a boolean if a field has been set.

### GetMerchantInfo

`func (o *AuthenticateThreeDSSessionRequest) GetMerchantInfo() ThreeDSMerchantInfo`

GetMerchantInfo returns the MerchantInfo field if non-nil, zero value otherwise.

### GetMerchantInfoOk

`func (o *AuthenticateThreeDSSessionRequest) GetMerchantInfoOk() (*ThreeDSMerchantInfo, bool)`

GetMerchantInfoOk returns a tuple with the MerchantInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInfo

`func (o *AuthenticateThreeDSSessionRequest) SetMerchantInfo(v ThreeDSMerchantInfo)`

SetMerchantInfo sets MerchantInfo field to given value.

### HasMerchantInfo

`func (o *AuthenticateThreeDSSessionRequest) HasMerchantInfo() bool`

HasMerchantInfo returns a boolean if a field has been set.

### GetRequestorInfo

`func (o *AuthenticateThreeDSSessionRequest) GetRequestorInfo() ThreeDSRequestorInfo`

GetRequestorInfo returns the RequestorInfo field if non-nil, zero value otherwise.

### GetRequestorInfoOk

`func (o *AuthenticateThreeDSSessionRequest) GetRequestorInfoOk() (*ThreeDSRequestorInfo, bool)`

GetRequestorInfoOk returns a tuple with the RequestorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorInfo

`func (o *AuthenticateThreeDSSessionRequest) SetRequestorInfo(v ThreeDSRequestorInfo)`

SetRequestorInfo sets RequestorInfo field to given value.


### GetCardholderInfo

`func (o *AuthenticateThreeDSSessionRequest) GetCardholderInfo() ThreeDSCardholderInfo`

GetCardholderInfo returns the CardholderInfo field if non-nil, zero value otherwise.

### GetCardholderInfoOk

`func (o *AuthenticateThreeDSSessionRequest) GetCardholderInfoOk() (*ThreeDSCardholderInfo, bool)`

GetCardholderInfoOk returns a tuple with the CardholderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderInfo

`func (o *AuthenticateThreeDSSessionRequest) SetCardholderInfo(v ThreeDSCardholderInfo)`

SetCardholderInfo sets CardholderInfo field to given value.

### HasCardholderInfo

`func (o *AuthenticateThreeDSSessionRequest) HasCardholderInfo() bool`

HasCardholderInfo returns a boolean if a field has been set.

### GetBroadcastInfo

`func (o *AuthenticateThreeDSSessionRequest) GetBroadcastInfo() interface{}`

GetBroadcastInfo returns the BroadcastInfo field if non-nil, zero value otherwise.

### GetBroadcastInfoOk

`func (o *AuthenticateThreeDSSessionRequest) GetBroadcastInfoOk() (*interface{}, bool)`

GetBroadcastInfoOk returns a tuple with the BroadcastInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastInfo

`func (o *AuthenticateThreeDSSessionRequest) SetBroadcastInfo(v interface{})`

SetBroadcastInfo sets BroadcastInfo field to given value.

### HasBroadcastInfo

`func (o *AuthenticateThreeDSSessionRequest) HasBroadcastInfo() bool`

HasBroadcastInfo returns a boolean if a field has been set.

### SetBroadcastInfoNil

`func (o *AuthenticateThreeDSSessionRequest) SetBroadcastInfoNil(b bool)`

 SetBroadcastInfoNil sets the value for BroadcastInfo to be an explicit nil

### UnsetBroadcastInfo
`func (o *AuthenticateThreeDSSessionRequest) UnsetBroadcastInfo()`

UnsetBroadcastInfo ensures that no value is present for BroadcastInfo, not even an explicit nil
### GetMessageExtensions

`func (o *AuthenticateThreeDSSessionRequest) GetMessageExtensions() []ThreeDSMessageExtension`

GetMessageExtensions returns the MessageExtensions field if non-nil, zero value otherwise.

### GetMessageExtensionsOk

`func (o *AuthenticateThreeDSSessionRequest) GetMessageExtensionsOk() (*[]ThreeDSMessageExtension, bool)`

GetMessageExtensionsOk returns a tuple with the MessageExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageExtensions

`func (o *AuthenticateThreeDSSessionRequest) SetMessageExtensions(v []ThreeDSMessageExtension)`

SetMessageExtensions sets MessageExtensions field to given value.

### HasMessageExtensions

`func (o *AuthenticateThreeDSSessionRequest) HasMessageExtensions() bool`

HasMessageExtensions returns a boolean if a field has been set.

### SetMessageExtensionsNil

`func (o *AuthenticateThreeDSSessionRequest) SetMessageExtensionsNil(b bool)`

 SetMessageExtensionsNil sets the value for MessageExtensions to be an explicit nil

### UnsetMessageExtensions
`func (o *AuthenticateThreeDSSessionRequest) UnsetMessageExtensions()`

UnsetMessageExtensions ensures that no value is present for MessageExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


