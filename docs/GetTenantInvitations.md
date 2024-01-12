# GetTenantInvitations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**TenantInvitationStatus**](TenantInvitationStatus.md) |  | [optional] 
**Page** | Pointer to **NullableInt32** |  | [optional] 
**Start** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGetTenantInvitations

`func NewGetTenantInvitations() *GetTenantInvitations`

NewGetTenantInvitations instantiates a new GetTenantInvitations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenantInvitationsWithDefaults

`func NewGetTenantInvitationsWithDefaults() *GetTenantInvitations`

NewGetTenantInvitationsWithDefaults instantiates a new GetTenantInvitations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTenantInvitations) GetStatus() TenantInvitationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTenantInvitations) GetStatusOk() (*TenantInvitationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTenantInvitations) SetStatus(v TenantInvitationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTenantInvitations) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPage

`func (o *GetTenantInvitations) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetTenantInvitations) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetTenantInvitations) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetTenantInvitations) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *GetTenantInvitations) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *GetTenantInvitations) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil
### GetStart

`func (o *GetTenantInvitations) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *GetTenantInvitations) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *GetTenantInvitations) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *GetTenantInvitations) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *GetTenantInvitations) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *GetTenantInvitations) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetSize

`func (o *GetTenantInvitations) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTenantInvitations) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTenantInvitations) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetTenantInvitations) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GetTenantInvitations) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GetTenantInvitations) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


