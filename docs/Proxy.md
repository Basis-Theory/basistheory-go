# Proxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] [readonly] 
**TenantId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DestinationUrl** | Pointer to **NullableString** |  | [optional] 
**RequestReactorId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewProxy

`func NewProxy() *Proxy`

NewProxy instantiates a new Proxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyWithDefaults

`func NewProxyWithDefaults() *Proxy`

NewProxyWithDefaults instantiates a new Proxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Proxy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Proxy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Proxy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Proxy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Proxy) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Proxy) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Proxy) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Proxy) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *Proxy) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *Proxy) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetTenantId

`func (o *Proxy) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Proxy) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Proxy) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Proxy) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *Proxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Proxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Proxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Proxy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Proxy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Proxy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDestinationUrl

`func (o *Proxy) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *Proxy) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *Proxy) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *Proxy) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrlNil

`func (o *Proxy) SetDestinationUrlNil(b bool)`

 SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil

### UnsetDestinationUrl
`func (o *Proxy) UnsetDestinationUrl()`

UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
### GetRequestReactorId

`func (o *Proxy) GetRequestReactorId() string`

GetRequestReactorId returns the RequestReactorId field if non-nil, zero value otherwise.

### GetRequestReactorIdOk

`func (o *Proxy) GetRequestReactorIdOk() (*string, bool)`

GetRequestReactorIdOk returns a tuple with the RequestReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReactorId

`func (o *Proxy) SetRequestReactorId(v string)`

SetRequestReactorId sets RequestReactorId field to given value.

### HasRequestReactorId

`func (o *Proxy) HasRequestReactorId() bool`

HasRequestReactorId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Proxy) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Proxy) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Proxy) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Proxy) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Proxy) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Proxy) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *Proxy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Proxy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Proxy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Proxy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Proxy) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Proxy) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetModifiedBy

`func (o *Proxy) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Proxy) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Proxy) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Proxy) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *Proxy) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *Proxy) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetModifiedAt

`func (o *Proxy) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Proxy) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Proxy) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Proxy) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *Proxy) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Proxy) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


