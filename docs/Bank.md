# Bank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingNumber** | **string** |  | 
**AccountNumber** | **string** |  | 

## Methods

### NewBank

`func NewBank(routingNumber string, accountNumber string, ) *Bank`

NewBank instantiates a new Bank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankWithDefaults

`func NewBankWithDefaults() *Bank`

NewBankWithDefaults instantiates a new Bank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingNumber

`func (o *Bank) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *Bank) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *Bank) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetAccountNumber

`func (o *Bank) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Bank) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Bank) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


