# InboundProxy

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

### NewInboundProxy

`func NewInboundProxy() *InboundProxy`

NewInboundProxy instantiates a new InboundProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundProxyWithDefaults

`func NewInboundProxyWithDefaults() *InboundProxy`

NewInboundProxyWithDefaults instantiates a new InboundProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboundProxy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboundProxy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboundProxy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InboundProxy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *InboundProxy) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InboundProxy) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InboundProxy) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InboundProxy) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *InboundProxy) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *InboundProxy) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetTenantId

`func (o *InboundProxy) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InboundProxy) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InboundProxy) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InboundProxy) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *InboundProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InboundProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InboundProxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InboundProxy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InboundProxy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InboundProxy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDestinationUrl

`func (o *InboundProxy) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *InboundProxy) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *InboundProxy) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *InboundProxy) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrlNil

`func (o *InboundProxy) SetDestinationUrlNil(b bool)`

 SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil

### UnsetDestinationUrl
`func (o *InboundProxy) UnsetDestinationUrl()`

UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
### GetRequestReactorId

`func (o *InboundProxy) GetRequestReactorId() string`

GetRequestReactorId returns the RequestReactorId field if non-nil, zero value otherwise.

### GetRequestReactorIdOk

`func (o *InboundProxy) GetRequestReactorIdOk() (*string, bool)`

GetRequestReactorIdOk returns a tuple with the RequestReactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReactorId

`func (o *InboundProxy) SetRequestReactorId(v string)`

SetRequestReactorId sets RequestReactorId field to given value.

### HasRequestReactorId

`func (o *InboundProxy) HasRequestReactorId() bool`

HasRequestReactorId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InboundProxy) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InboundProxy) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InboundProxy) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InboundProxy) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *InboundProxy) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InboundProxy) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *InboundProxy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InboundProxy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InboundProxy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InboundProxy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *InboundProxy) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *InboundProxy) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetModifiedBy

`func (o *InboundProxy) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *InboundProxy) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *InboundProxy) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *InboundProxy) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *InboundProxy) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *InboundProxy) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetModifiedAt

`func (o *InboundProxy) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *InboundProxy) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *InboundProxy) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *InboundProxy) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *InboundProxy) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *InboundProxy) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


