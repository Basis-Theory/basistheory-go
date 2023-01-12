# CreateSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionKey** | Pointer to **NullableString** |  | [optional] 
**Nonce** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCreateSessionResponse

`func NewCreateSessionResponse() *CreateSessionResponse`

NewCreateSessionResponse instantiates a new CreateSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionResponseWithDefaults

`func NewCreateSessionResponseWithDefaults() *CreateSessionResponse`

NewCreateSessionResponseWithDefaults instantiates a new CreateSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionKey

`func (o *CreateSessionResponse) GetSessionKey() string`

GetSessionKey returns the SessionKey field if non-nil, zero value otherwise.

### GetSessionKeyOk

`func (o *CreateSessionResponse) GetSessionKeyOk() (*string, bool)`

GetSessionKeyOk returns a tuple with the SessionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionKey

`func (o *CreateSessionResponse) SetSessionKey(v string)`

SetSessionKey sets SessionKey field to given value.

### HasSessionKey

`func (o *CreateSessionResponse) HasSessionKey() bool`

HasSessionKey returns a boolean if a field has been set.

### SetSessionKeyNil

`func (o *CreateSessionResponse) SetSessionKeyNil(b bool)`

 SetSessionKeyNil sets the value for SessionKey to be an explicit nil

### UnsetSessionKey
`func (o *CreateSessionResponse) UnsetSessionKey()`

UnsetSessionKey ensures that no value is present for SessionKey, not even an explicit nil
### GetNonce

`func (o *CreateSessionResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateSessionResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateSessionResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CreateSessionResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *CreateSessionResponse) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *CreateSessionResponse) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetExpiresAt

`func (o *CreateSessionResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateSessionResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateSessionResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateSessionResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateSessionResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateSessionResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


