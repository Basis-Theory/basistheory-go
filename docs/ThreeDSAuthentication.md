# ThreeDSAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreedsVersion** | Pointer to **NullableString** |  | [optional] 
**AcsTransactionId** | Pointer to **NullableString** |  | [optional] 
**DsTransactionId** | Pointer to **NullableString** |  | [optional] 
**SdkTransactionId** | Pointer to **NullableString** |  | [optional] 
**AcsReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**DsReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**AuthenticationValue** | Pointer to **NullableString** |  | [optional] 
**AuthenticationStatus** | Pointer to **NullableString** |  | [optional] 
**AuthenticationStatusReason** | Pointer to **NullableString** |  | [optional] 
**Eci** | Pointer to **NullableString** |  | [optional] 
**AcsChallengeMandated** | Pointer to **NullableString** |  | [optional] 
**AcsDecoupledAuthentication** | Pointer to **NullableString** |  | [optional] 
**AuthenticationChallengeType** | Pointer to **NullableString** |  | [optional] 
**AcsRenderingType** | Pointer to [**ThreeDSAcsRenderingType**](ThreeDSAcsRenderingType.md) |  | [optional] 
**AcsSignedContent** | Pointer to **NullableString** |  | [optional] 
**AcsChallengeUrl** | Pointer to **NullableString** |  | [optional] 
**ChallengeAttempts** | Pointer to **NullableString** |  | [optional] 
**ChallengeCancelReason** | Pointer to **NullableString** |  | [optional] 
**CardholderInfo** | Pointer to **NullableString** |  | [optional] 
**WhitelistStatus** | Pointer to **NullableString** |  | [optional] 
**WhitelistStatusSource** | Pointer to **NullableString** |  | [optional] 
**MessageExtensions** | Pointer to [**[]ThreeDSMessageExtension**](ThreeDSMessageExtension.md) |  | [optional] 

## Methods

### NewThreeDSAuthentication

`func NewThreeDSAuthentication() *ThreeDSAuthentication`

NewThreeDSAuthentication instantiates a new ThreeDSAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSAuthenticationWithDefaults

`func NewThreeDSAuthenticationWithDefaults() *ThreeDSAuthentication`

NewThreeDSAuthenticationWithDefaults instantiates a new ThreeDSAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreedsVersion

`func (o *ThreeDSAuthentication) GetThreedsVersion() string`

GetThreedsVersion returns the ThreedsVersion field if non-nil, zero value otherwise.

### GetThreedsVersionOk

`func (o *ThreeDSAuthentication) GetThreedsVersionOk() (*string, bool)`

GetThreedsVersionOk returns a tuple with the ThreedsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreedsVersion

`func (o *ThreeDSAuthentication) SetThreedsVersion(v string)`

SetThreedsVersion sets ThreedsVersion field to given value.

### HasThreedsVersion

`func (o *ThreeDSAuthentication) HasThreedsVersion() bool`

HasThreedsVersion returns a boolean if a field has been set.

### SetThreedsVersionNil

`func (o *ThreeDSAuthentication) SetThreedsVersionNil(b bool)`

 SetThreedsVersionNil sets the value for ThreedsVersion to be an explicit nil

### UnsetThreedsVersion
`func (o *ThreeDSAuthentication) UnsetThreedsVersion()`

UnsetThreedsVersion ensures that no value is present for ThreedsVersion, not even an explicit nil
### GetAcsTransactionId

`func (o *ThreeDSAuthentication) GetAcsTransactionId() string`

GetAcsTransactionId returns the AcsTransactionId field if non-nil, zero value otherwise.

### GetAcsTransactionIdOk

`func (o *ThreeDSAuthentication) GetAcsTransactionIdOk() (*string, bool)`

GetAcsTransactionIdOk returns a tuple with the AcsTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsTransactionId

`func (o *ThreeDSAuthentication) SetAcsTransactionId(v string)`

SetAcsTransactionId sets AcsTransactionId field to given value.

### HasAcsTransactionId

`func (o *ThreeDSAuthentication) HasAcsTransactionId() bool`

HasAcsTransactionId returns a boolean if a field has been set.

### SetAcsTransactionIdNil

`func (o *ThreeDSAuthentication) SetAcsTransactionIdNil(b bool)`

 SetAcsTransactionIdNil sets the value for AcsTransactionId to be an explicit nil

### UnsetAcsTransactionId
`func (o *ThreeDSAuthentication) UnsetAcsTransactionId()`

UnsetAcsTransactionId ensures that no value is present for AcsTransactionId, not even an explicit nil
### GetDsTransactionId

`func (o *ThreeDSAuthentication) GetDsTransactionId() string`

GetDsTransactionId returns the DsTransactionId field if non-nil, zero value otherwise.

### GetDsTransactionIdOk

`func (o *ThreeDSAuthentication) GetDsTransactionIdOk() (*string, bool)`

GetDsTransactionIdOk returns a tuple with the DsTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsTransactionId

`func (o *ThreeDSAuthentication) SetDsTransactionId(v string)`

SetDsTransactionId sets DsTransactionId field to given value.

### HasDsTransactionId

`func (o *ThreeDSAuthentication) HasDsTransactionId() bool`

HasDsTransactionId returns a boolean if a field has been set.

### SetDsTransactionIdNil

`func (o *ThreeDSAuthentication) SetDsTransactionIdNil(b bool)`

 SetDsTransactionIdNil sets the value for DsTransactionId to be an explicit nil

### UnsetDsTransactionId
`func (o *ThreeDSAuthentication) UnsetDsTransactionId()`

UnsetDsTransactionId ensures that no value is present for DsTransactionId, not even an explicit nil
### GetSdkTransactionId

`func (o *ThreeDSAuthentication) GetSdkTransactionId() string`

GetSdkTransactionId returns the SdkTransactionId field if non-nil, zero value otherwise.

### GetSdkTransactionIdOk

`func (o *ThreeDSAuthentication) GetSdkTransactionIdOk() (*string, bool)`

GetSdkTransactionIdOk returns a tuple with the SdkTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkTransactionId

`func (o *ThreeDSAuthentication) SetSdkTransactionId(v string)`

SetSdkTransactionId sets SdkTransactionId field to given value.

### HasSdkTransactionId

`func (o *ThreeDSAuthentication) HasSdkTransactionId() bool`

HasSdkTransactionId returns a boolean if a field has been set.

### SetSdkTransactionIdNil

`func (o *ThreeDSAuthentication) SetSdkTransactionIdNil(b bool)`

 SetSdkTransactionIdNil sets the value for SdkTransactionId to be an explicit nil

### UnsetSdkTransactionId
`func (o *ThreeDSAuthentication) UnsetSdkTransactionId()`

UnsetSdkTransactionId ensures that no value is present for SdkTransactionId, not even an explicit nil
### GetAcsReferenceNumber

`func (o *ThreeDSAuthentication) GetAcsReferenceNumber() string`

GetAcsReferenceNumber returns the AcsReferenceNumber field if non-nil, zero value otherwise.

### GetAcsReferenceNumberOk

`func (o *ThreeDSAuthentication) GetAcsReferenceNumberOk() (*string, bool)`

GetAcsReferenceNumberOk returns a tuple with the AcsReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsReferenceNumber

`func (o *ThreeDSAuthentication) SetAcsReferenceNumber(v string)`

SetAcsReferenceNumber sets AcsReferenceNumber field to given value.

### HasAcsReferenceNumber

`func (o *ThreeDSAuthentication) HasAcsReferenceNumber() bool`

HasAcsReferenceNumber returns a boolean if a field has been set.

### SetAcsReferenceNumberNil

`func (o *ThreeDSAuthentication) SetAcsReferenceNumberNil(b bool)`

 SetAcsReferenceNumberNil sets the value for AcsReferenceNumber to be an explicit nil

### UnsetAcsReferenceNumber
`func (o *ThreeDSAuthentication) UnsetAcsReferenceNumber()`

UnsetAcsReferenceNumber ensures that no value is present for AcsReferenceNumber, not even an explicit nil
### GetDsReferenceNumber

`func (o *ThreeDSAuthentication) GetDsReferenceNumber() string`

GetDsReferenceNumber returns the DsReferenceNumber field if non-nil, zero value otherwise.

### GetDsReferenceNumberOk

`func (o *ThreeDSAuthentication) GetDsReferenceNumberOk() (*string, bool)`

GetDsReferenceNumberOk returns a tuple with the DsReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsReferenceNumber

`func (o *ThreeDSAuthentication) SetDsReferenceNumber(v string)`

SetDsReferenceNumber sets DsReferenceNumber field to given value.

### HasDsReferenceNumber

`func (o *ThreeDSAuthentication) HasDsReferenceNumber() bool`

HasDsReferenceNumber returns a boolean if a field has been set.

### SetDsReferenceNumberNil

`func (o *ThreeDSAuthentication) SetDsReferenceNumberNil(b bool)`

 SetDsReferenceNumberNil sets the value for DsReferenceNumber to be an explicit nil

### UnsetDsReferenceNumber
`func (o *ThreeDSAuthentication) UnsetDsReferenceNumber()`

UnsetDsReferenceNumber ensures that no value is present for DsReferenceNumber, not even an explicit nil
### GetAuthenticationValue

`func (o *ThreeDSAuthentication) GetAuthenticationValue() string`

GetAuthenticationValue returns the AuthenticationValue field if non-nil, zero value otherwise.

### GetAuthenticationValueOk

`func (o *ThreeDSAuthentication) GetAuthenticationValueOk() (*string, bool)`

GetAuthenticationValueOk returns a tuple with the AuthenticationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationValue

`func (o *ThreeDSAuthentication) SetAuthenticationValue(v string)`

SetAuthenticationValue sets AuthenticationValue field to given value.

### HasAuthenticationValue

`func (o *ThreeDSAuthentication) HasAuthenticationValue() bool`

HasAuthenticationValue returns a boolean if a field has been set.

### SetAuthenticationValueNil

`func (o *ThreeDSAuthentication) SetAuthenticationValueNil(b bool)`

 SetAuthenticationValueNil sets the value for AuthenticationValue to be an explicit nil

### UnsetAuthenticationValue
`func (o *ThreeDSAuthentication) UnsetAuthenticationValue()`

UnsetAuthenticationValue ensures that no value is present for AuthenticationValue, not even an explicit nil
### GetAuthenticationStatus

`func (o *ThreeDSAuthentication) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *ThreeDSAuthentication) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *ThreeDSAuthentication) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *ThreeDSAuthentication) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### SetAuthenticationStatusNil

`func (o *ThreeDSAuthentication) SetAuthenticationStatusNil(b bool)`

 SetAuthenticationStatusNil sets the value for AuthenticationStatus to be an explicit nil

### UnsetAuthenticationStatus
`func (o *ThreeDSAuthentication) UnsetAuthenticationStatus()`

UnsetAuthenticationStatus ensures that no value is present for AuthenticationStatus, not even an explicit nil
### GetAuthenticationStatusReason

`func (o *ThreeDSAuthentication) GetAuthenticationStatusReason() string`

GetAuthenticationStatusReason returns the AuthenticationStatusReason field if non-nil, zero value otherwise.

### GetAuthenticationStatusReasonOk

`func (o *ThreeDSAuthentication) GetAuthenticationStatusReasonOk() (*string, bool)`

GetAuthenticationStatusReasonOk returns a tuple with the AuthenticationStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatusReason

`func (o *ThreeDSAuthentication) SetAuthenticationStatusReason(v string)`

SetAuthenticationStatusReason sets AuthenticationStatusReason field to given value.

### HasAuthenticationStatusReason

`func (o *ThreeDSAuthentication) HasAuthenticationStatusReason() bool`

HasAuthenticationStatusReason returns a boolean if a field has been set.

### SetAuthenticationStatusReasonNil

`func (o *ThreeDSAuthentication) SetAuthenticationStatusReasonNil(b bool)`

 SetAuthenticationStatusReasonNil sets the value for AuthenticationStatusReason to be an explicit nil

### UnsetAuthenticationStatusReason
`func (o *ThreeDSAuthentication) UnsetAuthenticationStatusReason()`

UnsetAuthenticationStatusReason ensures that no value is present for AuthenticationStatusReason, not even an explicit nil
### GetEci

`func (o *ThreeDSAuthentication) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSAuthentication) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSAuthentication) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDSAuthentication) HasEci() bool`

HasEci returns a boolean if a field has been set.

### SetEciNil

`func (o *ThreeDSAuthentication) SetEciNil(b bool)`

 SetEciNil sets the value for Eci to be an explicit nil

### UnsetEci
`func (o *ThreeDSAuthentication) UnsetEci()`

UnsetEci ensures that no value is present for Eci, not even an explicit nil
### GetAcsChallengeMandated

`func (o *ThreeDSAuthentication) GetAcsChallengeMandated() string`

GetAcsChallengeMandated returns the AcsChallengeMandated field if non-nil, zero value otherwise.

### GetAcsChallengeMandatedOk

`func (o *ThreeDSAuthentication) GetAcsChallengeMandatedOk() (*string, bool)`

GetAcsChallengeMandatedOk returns a tuple with the AcsChallengeMandated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsChallengeMandated

`func (o *ThreeDSAuthentication) SetAcsChallengeMandated(v string)`

SetAcsChallengeMandated sets AcsChallengeMandated field to given value.

### HasAcsChallengeMandated

`func (o *ThreeDSAuthentication) HasAcsChallengeMandated() bool`

HasAcsChallengeMandated returns a boolean if a field has been set.

### SetAcsChallengeMandatedNil

`func (o *ThreeDSAuthentication) SetAcsChallengeMandatedNil(b bool)`

 SetAcsChallengeMandatedNil sets the value for AcsChallengeMandated to be an explicit nil

### UnsetAcsChallengeMandated
`func (o *ThreeDSAuthentication) UnsetAcsChallengeMandated()`

UnsetAcsChallengeMandated ensures that no value is present for AcsChallengeMandated, not even an explicit nil
### GetAcsDecoupledAuthentication

`func (o *ThreeDSAuthentication) GetAcsDecoupledAuthentication() string`

GetAcsDecoupledAuthentication returns the AcsDecoupledAuthentication field if non-nil, zero value otherwise.

### GetAcsDecoupledAuthenticationOk

`func (o *ThreeDSAuthentication) GetAcsDecoupledAuthenticationOk() (*string, bool)`

GetAcsDecoupledAuthenticationOk returns a tuple with the AcsDecoupledAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsDecoupledAuthentication

`func (o *ThreeDSAuthentication) SetAcsDecoupledAuthentication(v string)`

SetAcsDecoupledAuthentication sets AcsDecoupledAuthentication field to given value.

### HasAcsDecoupledAuthentication

`func (o *ThreeDSAuthentication) HasAcsDecoupledAuthentication() bool`

HasAcsDecoupledAuthentication returns a boolean if a field has been set.

### SetAcsDecoupledAuthenticationNil

`func (o *ThreeDSAuthentication) SetAcsDecoupledAuthenticationNil(b bool)`

 SetAcsDecoupledAuthenticationNil sets the value for AcsDecoupledAuthentication to be an explicit nil

### UnsetAcsDecoupledAuthentication
`func (o *ThreeDSAuthentication) UnsetAcsDecoupledAuthentication()`

UnsetAcsDecoupledAuthentication ensures that no value is present for AcsDecoupledAuthentication, not even an explicit nil
### GetAuthenticationChallengeType

`func (o *ThreeDSAuthentication) GetAuthenticationChallengeType() string`

GetAuthenticationChallengeType returns the AuthenticationChallengeType field if non-nil, zero value otherwise.

### GetAuthenticationChallengeTypeOk

`func (o *ThreeDSAuthentication) GetAuthenticationChallengeTypeOk() (*string, bool)`

GetAuthenticationChallengeTypeOk returns a tuple with the AuthenticationChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationChallengeType

`func (o *ThreeDSAuthentication) SetAuthenticationChallengeType(v string)`

SetAuthenticationChallengeType sets AuthenticationChallengeType field to given value.

### HasAuthenticationChallengeType

`func (o *ThreeDSAuthentication) HasAuthenticationChallengeType() bool`

HasAuthenticationChallengeType returns a boolean if a field has been set.

### SetAuthenticationChallengeTypeNil

`func (o *ThreeDSAuthentication) SetAuthenticationChallengeTypeNil(b bool)`

 SetAuthenticationChallengeTypeNil sets the value for AuthenticationChallengeType to be an explicit nil

### UnsetAuthenticationChallengeType
`func (o *ThreeDSAuthentication) UnsetAuthenticationChallengeType()`

UnsetAuthenticationChallengeType ensures that no value is present for AuthenticationChallengeType, not even an explicit nil
### GetAcsRenderingType

`func (o *ThreeDSAuthentication) GetAcsRenderingType() ThreeDSAcsRenderingType`

GetAcsRenderingType returns the AcsRenderingType field if non-nil, zero value otherwise.

### GetAcsRenderingTypeOk

`func (o *ThreeDSAuthentication) GetAcsRenderingTypeOk() (*ThreeDSAcsRenderingType, bool)`

GetAcsRenderingTypeOk returns a tuple with the AcsRenderingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsRenderingType

`func (o *ThreeDSAuthentication) SetAcsRenderingType(v ThreeDSAcsRenderingType)`

SetAcsRenderingType sets AcsRenderingType field to given value.

### HasAcsRenderingType

`func (o *ThreeDSAuthentication) HasAcsRenderingType() bool`

HasAcsRenderingType returns a boolean if a field has been set.

### GetAcsSignedContent

`func (o *ThreeDSAuthentication) GetAcsSignedContent() string`

GetAcsSignedContent returns the AcsSignedContent field if non-nil, zero value otherwise.

### GetAcsSignedContentOk

`func (o *ThreeDSAuthentication) GetAcsSignedContentOk() (*string, bool)`

GetAcsSignedContentOk returns a tuple with the AcsSignedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsSignedContent

`func (o *ThreeDSAuthentication) SetAcsSignedContent(v string)`

SetAcsSignedContent sets AcsSignedContent field to given value.

### HasAcsSignedContent

`func (o *ThreeDSAuthentication) HasAcsSignedContent() bool`

HasAcsSignedContent returns a boolean if a field has been set.

### SetAcsSignedContentNil

`func (o *ThreeDSAuthentication) SetAcsSignedContentNil(b bool)`

 SetAcsSignedContentNil sets the value for AcsSignedContent to be an explicit nil

### UnsetAcsSignedContent
`func (o *ThreeDSAuthentication) UnsetAcsSignedContent()`

UnsetAcsSignedContent ensures that no value is present for AcsSignedContent, not even an explicit nil
### GetAcsChallengeUrl

`func (o *ThreeDSAuthentication) GetAcsChallengeUrl() string`

GetAcsChallengeUrl returns the AcsChallengeUrl field if non-nil, zero value otherwise.

### GetAcsChallengeUrlOk

`func (o *ThreeDSAuthentication) GetAcsChallengeUrlOk() (*string, bool)`

GetAcsChallengeUrlOk returns a tuple with the AcsChallengeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsChallengeUrl

`func (o *ThreeDSAuthentication) SetAcsChallengeUrl(v string)`

SetAcsChallengeUrl sets AcsChallengeUrl field to given value.

### HasAcsChallengeUrl

`func (o *ThreeDSAuthentication) HasAcsChallengeUrl() bool`

HasAcsChallengeUrl returns a boolean if a field has been set.

### SetAcsChallengeUrlNil

`func (o *ThreeDSAuthentication) SetAcsChallengeUrlNil(b bool)`

 SetAcsChallengeUrlNil sets the value for AcsChallengeUrl to be an explicit nil

### UnsetAcsChallengeUrl
`func (o *ThreeDSAuthentication) UnsetAcsChallengeUrl()`

UnsetAcsChallengeUrl ensures that no value is present for AcsChallengeUrl, not even an explicit nil
### GetChallengeAttempts

`func (o *ThreeDSAuthentication) GetChallengeAttempts() string`

GetChallengeAttempts returns the ChallengeAttempts field if non-nil, zero value otherwise.

### GetChallengeAttemptsOk

`func (o *ThreeDSAuthentication) GetChallengeAttemptsOk() (*string, bool)`

GetChallengeAttemptsOk returns a tuple with the ChallengeAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeAttempts

`func (o *ThreeDSAuthentication) SetChallengeAttempts(v string)`

SetChallengeAttempts sets ChallengeAttempts field to given value.

### HasChallengeAttempts

`func (o *ThreeDSAuthentication) HasChallengeAttempts() bool`

HasChallengeAttempts returns a boolean if a field has been set.

### SetChallengeAttemptsNil

`func (o *ThreeDSAuthentication) SetChallengeAttemptsNil(b bool)`

 SetChallengeAttemptsNil sets the value for ChallengeAttempts to be an explicit nil

### UnsetChallengeAttempts
`func (o *ThreeDSAuthentication) UnsetChallengeAttempts()`

UnsetChallengeAttempts ensures that no value is present for ChallengeAttempts, not even an explicit nil
### GetChallengeCancelReason

`func (o *ThreeDSAuthentication) GetChallengeCancelReason() string`

GetChallengeCancelReason returns the ChallengeCancelReason field if non-nil, zero value otherwise.

### GetChallengeCancelReasonOk

`func (o *ThreeDSAuthentication) GetChallengeCancelReasonOk() (*string, bool)`

GetChallengeCancelReasonOk returns a tuple with the ChallengeCancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeCancelReason

`func (o *ThreeDSAuthentication) SetChallengeCancelReason(v string)`

SetChallengeCancelReason sets ChallengeCancelReason field to given value.

### HasChallengeCancelReason

`func (o *ThreeDSAuthentication) HasChallengeCancelReason() bool`

HasChallengeCancelReason returns a boolean if a field has been set.

### SetChallengeCancelReasonNil

`func (o *ThreeDSAuthentication) SetChallengeCancelReasonNil(b bool)`

 SetChallengeCancelReasonNil sets the value for ChallengeCancelReason to be an explicit nil

### UnsetChallengeCancelReason
`func (o *ThreeDSAuthentication) UnsetChallengeCancelReason()`

UnsetChallengeCancelReason ensures that no value is present for ChallengeCancelReason, not even an explicit nil
### GetCardholderInfo

`func (o *ThreeDSAuthentication) GetCardholderInfo() string`

GetCardholderInfo returns the CardholderInfo field if non-nil, zero value otherwise.

### GetCardholderInfoOk

`func (o *ThreeDSAuthentication) GetCardholderInfoOk() (*string, bool)`

GetCardholderInfoOk returns a tuple with the CardholderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderInfo

`func (o *ThreeDSAuthentication) SetCardholderInfo(v string)`

SetCardholderInfo sets CardholderInfo field to given value.

### HasCardholderInfo

`func (o *ThreeDSAuthentication) HasCardholderInfo() bool`

HasCardholderInfo returns a boolean if a field has been set.

### SetCardholderInfoNil

`func (o *ThreeDSAuthentication) SetCardholderInfoNil(b bool)`

 SetCardholderInfoNil sets the value for CardholderInfo to be an explicit nil

### UnsetCardholderInfo
`func (o *ThreeDSAuthentication) UnsetCardholderInfo()`

UnsetCardholderInfo ensures that no value is present for CardholderInfo, not even an explicit nil
### GetWhitelistStatus

`func (o *ThreeDSAuthentication) GetWhitelistStatus() string`

GetWhitelistStatus returns the WhitelistStatus field if non-nil, zero value otherwise.

### GetWhitelistStatusOk

`func (o *ThreeDSAuthentication) GetWhitelistStatusOk() (*string, bool)`

GetWhitelistStatusOk returns a tuple with the WhitelistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistStatus

`func (o *ThreeDSAuthentication) SetWhitelistStatus(v string)`

SetWhitelistStatus sets WhitelistStatus field to given value.

### HasWhitelistStatus

`func (o *ThreeDSAuthentication) HasWhitelistStatus() bool`

HasWhitelistStatus returns a boolean if a field has been set.

### SetWhitelistStatusNil

`func (o *ThreeDSAuthentication) SetWhitelistStatusNil(b bool)`

 SetWhitelistStatusNil sets the value for WhitelistStatus to be an explicit nil

### UnsetWhitelistStatus
`func (o *ThreeDSAuthentication) UnsetWhitelistStatus()`

UnsetWhitelistStatus ensures that no value is present for WhitelistStatus, not even an explicit nil
### GetWhitelistStatusSource

`func (o *ThreeDSAuthentication) GetWhitelistStatusSource() string`

GetWhitelistStatusSource returns the WhitelistStatusSource field if non-nil, zero value otherwise.

### GetWhitelistStatusSourceOk

`func (o *ThreeDSAuthentication) GetWhitelistStatusSourceOk() (*string, bool)`

GetWhitelistStatusSourceOk returns a tuple with the WhitelistStatusSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistStatusSource

`func (o *ThreeDSAuthentication) SetWhitelistStatusSource(v string)`

SetWhitelistStatusSource sets WhitelistStatusSource field to given value.

### HasWhitelistStatusSource

`func (o *ThreeDSAuthentication) HasWhitelistStatusSource() bool`

HasWhitelistStatusSource returns a boolean if a field has been set.

### SetWhitelistStatusSourceNil

`func (o *ThreeDSAuthentication) SetWhitelistStatusSourceNil(b bool)`

 SetWhitelistStatusSourceNil sets the value for WhitelistStatusSource to be an explicit nil

### UnsetWhitelistStatusSource
`func (o *ThreeDSAuthentication) UnsetWhitelistStatusSource()`

UnsetWhitelistStatusSource ensures that no value is present for WhitelistStatusSource, not even an explicit nil
### GetMessageExtensions

`func (o *ThreeDSAuthentication) GetMessageExtensions() []ThreeDSMessageExtension`

GetMessageExtensions returns the MessageExtensions field if non-nil, zero value otherwise.

### GetMessageExtensionsOk

`func (o *ThreeDSAuthentication) GetMessageExtensionsOk() (*[]ThreeDSMessageExtension, bool)`

GetMessageExtensionsOk returns a tuple with the MessageExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageExtensions

`func (o *ThreeDSAuthentication) SetMessageExtensions(v []ThreeDSMessageExtension)`

SetMessageExtensions sets MessageExtensions field to given value.

### HasMessageExtensions

`func (o *ThreeDSAuthentication) HasMessageExtensions() bool`

HasMessageExtensions returns a boolean if a field has been set.

### SetMessageExtensionsNil

`func (o *ThreeDSAuthentication) SetMessageExtensionsNil(b bool)`

 SetMessageExtensionsNil sets the value for MessageExtensions to be an explicit nil

### UnsetMessageExtensions
`func (o *ThreeDSAuthentication) UnsetMessageExtensions()`

UnsetMessageExtensions ensures that no value is present for MessageExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


