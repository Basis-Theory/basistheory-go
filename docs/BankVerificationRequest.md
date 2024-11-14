# BankVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** |  | 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**RoutingNumber** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankVerificationRequest

`func NewBankVerificationRequest(tokenId string, ) *BankVerificationRequest`

NewBankVerificationRequest instantiates a new BankVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankVerificationRequestWithDefaults

`func NewBankVerificationRequestWithDefaults() *BankVerificationRequest`

NewBankVerificationRequestWithDefaults instantiates a new BankVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *BankVerificationRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BankVerificationRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BankVerificationRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCountryCode

`func (o *BankVerificationRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BankVerificationRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BankVerificationRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BankVerificationRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *BankVerificationRequest) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *BankVerificationRequest) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetRoutingNumber

`func (o *BankVerificationRequest) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankVerificationRequest) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankVerificationRequest) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *BankVerificationRequest) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### SetRoutingNumberNil

`func (o *BankVerificationRequest) SetRoutingNumberNil(b bool)`

 SetRoutingNumberNil sets the value for RoutingNumber to be an explicit nil

### UnsetRoutingNumber
`func (o *BankVerificationRequest) UnsetRoutingNumber()`

UnsetRoutingNumber ensures that no value is present for RoutingNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


