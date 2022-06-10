# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** |  | [optional] 
**ExpirationMonth** | Pointer to **NullableInt32** |  | [optional] 
**ExpirationYear** | Pointer to **NullableInt32** |  | [optional] 
**Cvc** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCard

`func NewCard() *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Card) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Card) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Card) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Card) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *Card) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *Card) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetExpirationMonth

`func (o *Card) GetExpirationMonth() int32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *Card) GetExpirationMonthOk() (*int32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *Card) SetExpirationMonth(v int32)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *Card) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### SetExpirationMonthNil

`func (o *Card) SetExpirationMonthNil(b bool)`

 SetExpirationMonthNil sets the value for ExpirationMonth to be an explicit nil

### UnsetExpirationMonth
`func (o *Card) UnsetExpirationMonth()`

UnsetExpirationMonth ensures that no value is present for ExpirationMonth, not even an explicit nil
### GetExpirationYear

`func (o *Card) GetExpirationYear() int32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *Card) GetExpirationYearOk() (*int32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *Card) SetExpirationYear(v int32)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *Card) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### SetExpirationYearNil

`func (o *Card) SetExpirationYearNil(b bool)`

 SetExpirationYearNil sets the value for ExpirationYear to be an explicit nil

### UnsetExpirationYear
`func (o *Card) UnsetExpirationYear()`

UnsetExpirationYear ensures that no value is present for ExpirationYear, not even an explicit nil
### GetCvc

`func (o *Card) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *Card) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *Card) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *Card) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### SetCvcNil

`func (o *Card) SetCvcNil(b bool)`

 SetCvcNil sets the value for Cvc to be an explicit nil

### UnsetCvc
`func (o *Card) UnsetCvc()`

UnsetCvc ensures that no value is present for Cvc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


