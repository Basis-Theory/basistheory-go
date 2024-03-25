# ThreeDSDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserAcceptHeader** | Pointer to **NullableString** |  | [optional] 
**BrowserIp** | Pointer to **NullableString** |  | [optional] 
**BrowserJavascriptEnabled** | Pointer to **NullableBool** |  | [optional] 
**BrowserJavaEnabled** | Pointer to **NullableBool** |  | [optional] 
**BrowserLanguage** | Pointer to **NullableString** |  | [optional] 
**BrowserColorDepth** | Pointer to **NullableString** |  | [optional] 
**BrowserScreenHeight** | Pointer to **NullableString** |  | [optional] 
**BrowserScreenWidth** | Pointer to **NullableString** |  | [optional] 
**BrowserTz** | Pointer to **NullableString** |  | [optional] 
**BrowserUserAgent** | Pointer to **NullableString** |  | [optional] 
**SdkTransactionId** | Pointer to **NullableString** |  | [optional] 
**SdkApplicationId** | Pointer to **NullableString** |  | [optional] 
**SdkEncryptionData** | Pointer to **NullableString** |  | [optional] 
**SdkEphemeralPublicKey** | Pointer to **NullableString** |  | [optional] 
**SdkMaxTimeout** | Pointer to **NullableString** |  | [optional] 
**SdkReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**SdkRenderOptions** | Pointer to [**ThreeDSMobileSdkRenderOptions**](ThreeDSMobileSdkRenderOptions.md) |  | [optional] 

## Methods

### NewThreeDSDeviceInfo

`func NewThreeDSDeviceInfo() *ThreeDSDeviceInfo`

NewThreeDSDeviceInfo instantiates a new ThreeDSDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSDeviceInfoWithDefaults

`func NewThreeDSDeviceInfoWithDefaults() *ThreeDSDeviceInfo`

NewThreeDSDeviceInfoWithDefaults instantiates a new ThreeDSDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserAcceptHeader

`func (o *ThreeDSDeviceInfo) GetBrowserAcceptHeader() string`

GetBrowserAcceptHeader returns the BrowserAcceptHeader field if non-nil, zero value otherwise.

### GetBrowserAcceptHeaderOk

`func (o *ThreeDSDeviceInfo) GetBrowserAcceptHeaderOk() (*string, bool)`

GetBrowserAcceptHeaderOk returns a tuple with the BrowserAcceptHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserAcceptHeader

`func (o *ThreeDSDeviceInfo) SetBrowserAcceptHeader(v string)`

SetBrowserAcceptHeader sets BrowserAcceptHeader field to given value.

### HasBrowserAcceptHeader

`func (o *ThreeDSDeviceInfo) HasBrowserAcceptHeader() bool`

HasBrowserAcceptHeader returns a boolean if a field has been set.

### SetBrowserAcceptHeaderNil

`func (o *ThreeDSDeviceInfo) SetBrowserAcceptHeaderNil(b bool)`

 SetBrowserAcceptHeaderNil sets the value for BrowserAcceptHeader to be an explicit nil

### UnsetBrowserAcceptHeader
`func (o *ThreeDSDeviceInfo) UnsetBrowserAcceptHeader()`

UnsetBrowserAcceptHeader ensures that no value is present for BrowserAcceptHeader, not even an explicit nil
### GetBrowserIp

`func (o *ThreeDSDeviceInfo) GetBrowserIp() string`

GetBrowserIp returns the BrowserIp field if non-nil, zero value otherwise.

### GetBrowserIpOk

`func (o *ThreeDSDeviceInfo) GetBrowserIpOk() (*string, bool)`

GetBrowserIpOk returns a tuple with the BrowserIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserIp

`func (o *ThreeDSDeviceInfo) SetBrowserIp(v string)`

SetBrowserIp sets BrowserIp field to given value.

### HasBrowserIp

`func (o *ThreeDSDeviceInfo) HasBrowserIp() bool`

HasBrowserIp returns a boolean if a field has been set.

### SetBrowserIpNil

`func (o *ThreeDSDeviceInfo) SetBrowserIpNil(b bool)`

 SetBrowserIpNil sets the value for BrowserIp to be an explicit nil

### UnsetBrowserIp
`func (o *ThreeDSDeviceInfo) UnsetBrowserIp()`

UnsetBrowserIp ensures that no value is present for BrowserIp, not even an explicit nil
### GetBrowserJavascriptEnabled

`func (o *ThreeDSDeviceInfo) GetBrowserJavascriptEnabled() bool`

GetBrowserJavascriptEnabled returns the BrowserJavascriptEnabled field if non-nil, zero value otherwise.

### GetBrowserJavascriptEnabledOk

`func (o *ThreeDSDeviceInfo) GetBrowserJavascriptEnabledOk() (*bool, bool)`

GetBrowserJavascriptEnabledOk returns a tuple with the BrowserJavascriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserJavascriptEnabled

`func (o *ThreeDSDeviceInfo) SetBrowserJavascriptEnabled(v bool)`

SetBrowserJavascriptEnabled sets BrowserJavascriptEnabled field to given value.

### HasBrowserJavascriptEnabled

`func (o *ThreeDSDeviceInfo) HasBrowserJavascriptEnabled() bool`

HasBrowserJavascriptEnabled returns a boolean if a field has been set.

### SetBrowserJavascriptEnabledNil

`func (o *ThreeDSDeviceInfo) SetBrowserJavascriptEnabledNil(b bool)`

 SetBrowserJavascriptEnabledNil sets the value for BrowserJavascriptEnabled to be an explicit nil

### UnsetBrowserJavascriptEnabled
`func (o *ThreeDSDeviceInfo) UnsetBrowserJavascriptEnabled()`

UnsetBrowserJavascriptEnabled ensures that no value is present for BrowserJavascriptEnabled, not even an explicit nil
### GetBrowserJavaEnabled

`func (o *ThreeDSDeviceInfo) GetBrowserJavaEnabled() bool`

GetBrowserJavaEnabled returns the BrowserJavaEnabled field if non-nil, zero value otherwise.

### GetBrowserJavaEnabledOk

`func (o *ThreeDSDeviceInfo) GetBrowserJavaEnabledOk() (*bool, bool)`

GetBrowserJavaEnabledOk returns a tuple with the BrowserJavaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserJavaEnabled

`func (o *ThreeDSDeviceInfo) SetBrowserJavaEnabled(v bool)`

SetBrowserJavaEnabled sets BrowserJavaEnabled field to given value.

### HasBrowserJavaEnabled

`func (o *ThreeDSDeviceInfo) HasBrowserJavaEnabled() bool`

HasBrowserJavaEnabled returns a boolean if a field has been set.

### SetBrowserJavaEnabledNil

`func (o *ThreeDSDeviceInfo) SetBrowserJavaEnabledNil(b bool)`

 SetBrowserJavaEnabledNil sets the value for BrowserJavaEnabled to be an explicit nil

### UnsetBrowserJavaEnabled
`func (o *ThreeDSDeviceInfo) UnsetBrowserJavaEnabled()`

UnsetBrowserJavaEnabled ensures that no value is present for BrowserJavaEnabled, not even an explicit nil
### GetBrowserLanguage

`func (o *ThreeDSDeviceInfo) GetBrowserLanguage() string`

GetBrowserLanguage returns the BrowserLanguage field if non-nil, zero value otherwise.

### GetBrowserLanguageOk

`func (o *ThreeDSDeviceInfo) GetBrowserLanguageOk() (*string, bool)`

GetBrowserLanguageOk returns a tuple with the BrowserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserLanguage

`func (o *ThreeDSDeviceInfo) SetBrowserLanguage(v string)`

SetBrowserLanguage sets BrowserLanguage field to given value.

### HasBrowserLanguage

`func (o *ThreeDSDeviceInfo) HasBrowserLanguage() bool`

HasBrowserLanguage returns a boolean if a field has been set.

### SetBrowserLanguageNil

`func (o *ThreeDSDeviceInfo) SetBrowserLanguageNil(b bool)`

 SetBrowserLanguageNil sets the value for BrowserLanguage to be an explicit nil

### UnsetBrowserLanguage
`func (o *ThreeDSDeviceInfo) UnsetBrowserLanguage()`

UnsetBrowserLanguage ensures that no value is present for BrowserLanguage, not even an explicit nil
### GetBrowserColorDepth

`func (o *ThreeDSDeviceInfo) GetBrowserColorDepth() string`

GetBrowserColorDepth returns the BrowserColorDepth field if non-nil, zero value otherwise.

### GetBrowserColorDepthOk

`func (o *ThreeDSDeviceInfo) GetBrowserColorDepthOk() (*string, bool)`

GetBrowserColorDepthOk returns a tuple with the BrowserColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserColorDepth

`func (o *ThreeDSDeviceInfo) SetBrowserColorDepth(v string)`

SetBrowserColorDepth sets BrowserColorDepth field to given value.

### HasBrowserColorDepth

`func (o *ThreeDSDeviceInfo) HasBrowserColorDepth() bool`

HasBrowserColorDepth returns a boolean if a field has been set.

### SetBrowserColorDepthNil

`func (o *ThreeDSDeviceInfo) SetBrowserColorDepthNil(b bool)`

 SetBrowserColorDepthNil sets the value for BrowserColorDepth to be an explicit nil

### UnsetBrowserColorDepth
`func (o *ThreeDSDeviceInfo) UnsetBrowserColorDepth()`

UnsetBrowserColorDepth ensures that no value is present for BrowserColorDepth, not even an explicit nil
### GetBrowserScreenHeight

`func (o *ThreeDSDeviceInfo) GetBrowserScreenHeight() string`

GetBrowserScreenHeight returns the BrowserScreenHeight field if non-nil, zero value otherwise.

### GetBrowserScreenHeightOk

`func (o *ThreeDSDeviceInfo) GetBrowserScreenHeightOk() (*string, bool)`

GetBrowserScreenHeightOk returns a tuple with the BrowserScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserScreenHeight

`func (o *ThreeDSDeviceInfo) SetBrowserScreenHeight(v string)`

SetBrowserScreenHeight sets BrowserScreenHeight field to given value.

### HasBrowserScreenHeight

`func (o *ThreeDSDeviceInfo) HasBrowserScreenHeight() bool`

HasBrowserScreenHeight returns a boolean if a field has been set.

### SetBrowserScreenHeightNil

`func (o *ThreeDSDeviceInfo) SetBrowserScreenHeightNil(b bool)`

 SetBrowserScreenHeightNil sets the value for BrowserScreenHeight to be an explicit nil

### UnsetBrowserScreenHeight
`func (o *ThreeDSDeviceInfo) UnsetBrowserScreenHeight()`

UnsetBrowserScreenHeight ensures that no value is present for BrowserScreenHeight, not even an explicit nil
### GetBrowserScreenWidth

`func (o *ThreeDSDeviceInfo) GetBrowserScreenWidth() string`

GetBrowserScreenWidth returns the BrowserScreenWidth field if non-nil, zero value otherwise.

### GetBrowserScreenWidthOk

`func (o *ThreeDSDeviceInfo) GetBrowserScreenWidthOk() (*string, bool)`

GetBrowserScreenWidthOk returns a tuple with the BrowserScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserScreenWidth

`func (o *ThreeDSDeviceInfo) SetBrowserScreenWidth(v string)`

SetBrowserScreenWidth sets BrowserScreenWidth field to given value.

### HasBrowserScreenWidth

`func (o *ThreeDSDeviceInfo) HasBrowserScreenWidth() bool`

HasBrowserScreenWidth returns a boolean if a field has been set.

### SetBrowserScreenWidthNil

`func (o *ThreeDSDeviceInfo) SetBrowserScreenWidthNil(b bool)`

 SetBrowserScreenWidthNil sets the value for BrowserScreenWidth to be an explicit nil

### UnsetBrowserScreenWidth
`func (o *ThreeDSDeviceInfo) UnsetBrowserScreenWidth()`

UnsetBrowserScreenWidth ensures that no value is present for BrowserScreenWidth, not even an explicit nil
### GetBrowserTz

`func (o *ThreeDSDeviceInfo) GetBrowserTz() string`

GetBrowserTz returns the BrowserTz field if non-nil, zero value otherwise.

### GetBrowserTzOk

`func (o *ThreeDSDeviceInfo) GetBrowserTzOk() (*string, bool)`

GetBrowserTzOk returns a tuple with the BrowserTz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserTz

`func (o *ThreeDSDeviceInfo) SetBrowserTz(v string)`

SetBrowserTz sets BrowserTz field to given value.

### HasBrowserTz

`func (o *ThreeDSDeviceInfo) HasBrowserTz() bool`

HasBrowserTz returns a boolean if a field has been set.

### SetBrowserTzNil

`func (o *ThreeDSDeviceInfo) SetBrowserTzNil(b bool)`

 SetBrowserTzNil sets the value for BrowserTz to be an explicit nil

### UnsetBrowserTz
`func (o *ThreeDSDeviceInfo) UnsetBrowserTz()`

UnsetBrowserTz ensures that no value is present for BrowserTz, not even an explicit nil
### GetBrowserUserAgent

`func (o *ThreeDSDeviceInfo) GetBrowserUserAgent() string`

GetBrowserUserAgent returns the BrowserUserAgent field if non-nil, zero value otherwise.

### GetBrowserUserAgentOk

`func (o *ThreeDSDeviceInfo) GetBrowserUserAgentOk() (*string, bool)`

GetBrowserUserAgentOk returns a tuple with the BrowserUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUserAgent

`func (o *ThreeDSDeviceInfo) SetBrowserUserAgent(v string)`

SetBrowserUserAgent sets BrowserUserAgent field to given value.

### HasBrowserUserAgent

`func (o *ThreeDSDeviceInfo) HasBrowserUserAgent() bool`

HasBrowserUserAgent returns a boolean if a field has been set.

### SetBrowserUserAgentNil

`func (o *ThreeDSDeviceInfo) SetBrowserUserAgentNil(b bool)`

 SetBrowserUserAgentNil sets the value for BrowserUserAgent to be an explicit nil

### UnsetBrowserUserAgent
`func (o *ThreeDSDeviceInfo) UnsetBrowserUserAgent()`

UnsetBrowserUserAgent ensures that no value is present for BrowserUserAgent, not even an explicit nil
### GetSdkTransactionId

`func (o *ThreeDSDeviceInfo) GetSdkTransactionId() string`

GetSdkTransactionId returns the SdkTransactionId field if non-nil, zero value otherwise.

### GetSdkTransactionIdOk

`func (o *ThreeDSDeviceInfo) GetSdkTransactionIdOk() (*string, bool)`

GetSdkTransactionIdOk returns a tuple with the SdkTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkTransactionId

`func (o *ThreeDSDeviceInfo) SetSdkTransactionId(v string)`

SetSdkTransactionId sets SdkTransactionId field to given value.

### HasSdkTransactionId

`func (o *ThreeDSDeviceInfo) HasSdkTransactionId() bool`

HasSdkTransactionId returns a boolean if a field has been set.

### SetSdkTransactionIdNil

`func (o *ThreeDSDeviceInfo) SetSdkTransactionIdNil(b bool)`

 SetSdkTransactionIdNil sets the value for SdkTransactionId to be an explicit nil

### UnsetSdkTransactionId
`func (o *ThreeDSDeviceInfo) UnsetSdkTransactionId()`

UnsetSdkTransactionId ensures that no value is present for SdkTransactionId, not even an explicit nil
### GetSdkApplicationId

`func (o *ThreeDSDeviceInfo) GetSdkApplicationId() string`

GetSdkApplicationId returns the SdkApplicationId field if non-nil, zero value otherwise.

### GetSdkApplicationIdOk

`func (o *ThreeDSDeviceInfo) GetSdkApplicationIdOk() (*string, bool)`

GetSdkApplicationIdOk returns a tuple with the SdkApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkApplicationId

`func (o *ThreeDSDeviceInfo) SetSdkApplicationId(v string)`

SetSdkApplicationId sets SdkApplicationId field to given value.

### HasSdkApplicationId

`func (o *ThreeDSDeviceInfo) HasSdkApplicationId() bool`

HasSdkApplicationId returns a boolean if a field has been set.

### SetSdkApplicationIdNil

`func (o *ThreeDSDeviceInfo) SetSdkApplicationIdNil(b bool)`

 SetSdkApplicationIdNil sets the value for SdkApplicationId to be an explicit nil

### UnsetSdkApplicationId
`func (o *ThreeDSDeviceInfo) UnsetSdkApplicationId()`

UnsetSdkApplicationId ensures that no value is present for SdkApplicationId, not even an explicit nil
### GetSdkEncryptionData

`func (o *ThreeDSDeviceInfo) GetSdkEncryptionData() string`

GetSdkEncryptionData returns the SdkEncryptionData field if non-nil, zero value otherwise.

### GetSdkEncryptionDataOk

`func (o *ThreeDSDeviceInfo) GetSdkEncryptionDataOk() (*string, bool)`

GetSdkEncryptionDataOk returns a tuple with the SdkEncryptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEncryptionData

`func (o *ThreeDSDeviceInfo) SetSdkEncryptionData(v string)`

SetSdkEncryptionData sets SdkEncryptionData field to given value.

### HasSdkEncryptionData

`func (o *ThreeDSDeviceInfo) HasSdkEncryptionData() bool`

HasSdkEncryptionData returns a boolean if a field has been set.

### SetSdkEncryptionDataNil

`func (o *ThreeDSDeviceInfo) SetSdkEncryptionDataNil(b bool)`

 SetSdkEncryptionDataNil sets the value for SdkEncryptionData to be an explicit nil

### UnsetSdkEncryptionData
`func (o *ThreeDSDeviceInfo) UnsetSdkEncryptionData()`

UnsetSdkEncryptionData ensures that no value is present for SdkEncryptionData, not even an explicit nil
### GetSdkEphemeralPublicKey

`func (o *ThreeDSDeviceInfo) GetSdkEphemeralPublicKey() string`

GetSdkEphemeralPublicKey returns the SdkEphemeralPublicKey field if non-nil, zero value otherwise.

### GetSdkEphemeralPublicKeyOk

`func (o *ThreeDSDeviceInfo) GetSdkEphemeralPublicKeyOk() (*string, bool)`

GetSdkEphemeralPublicKeyOk returns a tuple with the SdkEphemeralPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEphemeralPublicKey

`func (o *ThreeDSDeviceInfo) SetSdkEphemeralPublicKey(v string)`

SetSdkEphemeralPublicKey sets SdkEphemeralPublicKey field to given value.

### HasSdkEphemeralPublicKey

`func (o *ThreeDSDeviceInfo) HasSdkEphemeralPublicKey() bool`

HasSdkEphemeralPublicKey returns a boolean if a field has been set.

### SetSdkEphemeralPublicKeyNil

`func (o *ThreeDSDeviceInfo) SetSdkEphemeralPublicKeyNil(b bool)`

 SetSdkEphemeralPublicKeyNil sets the value for SdkEphemeralPublicKey to be an explicit nil

### UnsetSdkEphemeralPublicKey
`func (o *ThreeDSDeviceInfo) UnsetSdkEphemeralPublicKey()`

UnsetSdkEphemeralPublicKey ensures that no value is present for SdkEphemeralPublicKey, not even an explicit nil
### GetSdkMaxTimeout

`func (o *ThreeDSDeviceInfo) GetSdkMaxTimeout() string`

GetSdkMaxTimeout returns the SdkMaxTimeout field if non-nil, zero value otherwise.

### GetSdkMaxTimeoutOk

`func (o *ThreeDSDeviceInfo) GetSdkMaxTimeoutOk() (*string, bool)`

GetSdkMaxTimeoutOk returns a tuple with the SdkMaxTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkMaxTimeout

`func (o *ThreeDSDeviceInfo) SetSdkMaxTimeout(v string)`

SetSdkMaxTimeout sets SdkMaxTimeout field to given value.

### HasSdkMaxTimeout

`func (o *ThreeDSDeviceInfo) HasSdkMaxTimeout() bool`

HasSdkMaxTimeout returns a boolean if a field has been set.

### SetSdkMaxTimeoutNil

`func (o *ThreeDSDeviceInfo) SetSdkMaxTimeoutNil(b bool)`

 SetSdkMaxTimeoutNil sets the value for SdkMaxTimeout to be an explicit nil

### UnsetSdkMaxTimeout
`func (o *ThreeDSDeviceInfo) UnsetSdkMaxTimeout()`

UnsetSdkMaxTimeout ensures that no value is present for SdkMaxTimeout, not even an explicit nil
### GetSdkReferenceNumber

`func (o *ThreeDSDeviceInfo) GetSdkReferenceNumber() string`

GetSdkReferenceNumber returns the SdkReferenceNumber field if non-nil, zero value otherwise.

### GetSdkReferenceNumberOk

`func (o *ThreeDSDeviceInfo) GetSdkReferenceNumberOk() (*string, bool)`

GetSdkReferenceNumberOk returns a tuple with the SdkReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkReferenceNumber

`func (o *ThreeDSDeviceInfo) SetSdkReferenceNumber(v string)`

SetSdkReferenceNumber sets SdkReferenceNumber field to given value.

### HasSdkReferenceNumber

`func (o *ThreeDSDeviceInfo) HasSdkReferenceNumber() bool`

HasSdkReferenceNumber returns a boolean if a field has been set.

### SetSdkReferenceNumberNil

`func (o *ThreeDSDeviceInfo) SetSdkReferenceNumberNil(b bool)`

 SetSdkReferenceNumberNil sets the value for SdkReferenceNumber to be an explicit nil

### UnsetSdkReferenceNumber
`func (o *ThreeDSDeviceInfo) UnsetSdkReferenceNumber()`

UnsetSdkReferenceNumber ensures that no value is present for SdkReferenceNumber, not even an explicit nil
### GetSdkRenderOptions

`func (o *ThreeDSDeviceInfo) GetSdkRenderOptions() ThreeDSMobileSdkRenderOptions`

GetSdkRenderOptions returns the SdkRenderOptions field if non-nil, zero value otherwise.

### GetSdkRenderOptionsOk

`func (o *ThreeDSDeviceInfo) GetSdkRenderOptionsOk() (*ThreeDSMobileSdkRenderOptions, bool)`

GetSdkRenderOptionsOk returns a tuple with the SdkRenderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkRenderOptions

`func (o *ThreeDSDeviceInfo) SetSdkRenderOptions(v ThreeDSMobileSdkRenderOptions)`

SetSdkRenderOptions sets SdkRenderOptions field to given value.

### HasSdkRenderOptions

`func (o *ThreeDSDeviceInfo) HasSdkRenderOptions() bool`

HasSdkRenderOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


