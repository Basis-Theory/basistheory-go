# CreateThreeDSSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pan** | **string** |  | 
**Type** | Pointer to **NullableString** |  | [optional] 
**Device** | Pointer to **NullableString** |  | [optional] 
**DeviceInfo** | Pointer to [**ThreeDSDeviceInfo**](ThreeDSDeviceInfo.md) |  | [optional] 

## Methods

### NewCreateThreeDSSessionRequest

`func NewCreateThreeDSSessionRequest(pan string, ) *CreateThreeDSSessionRequest`

NewCreateThreeDSSessionRequest instantiates a new CreateThreeDSSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThreeDSSessionRequestWithDefaults

`func NewCreateThreeDSSessionRequestWithDefaults() *CreateThreeDSSessionRequest`

NewCreateThreeDSSessionRequestWithDefaults instantiates a new CreateThreeDSSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPan

`func (o *CreateThreeDSSessionRequest) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *CreateThreeDSSessionRequest) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *CreateThreeDSSessionRequest) SetPan(v string)`

SetPan sets Pan field to given value.


### GetType

`func (o *CreateThreeDSSessionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateThreeDSSessionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateThreeDSSessionRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateThreeDSSessionRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CreateThreeDSSessionRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateThreeDSSessionRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDevice

`func (o *CreateThreeDSSessionRequest) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CreateThreeDSSessionRequest) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CreateThreeDSSessionRequest) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CreateThreeDSSessionRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *CreateThreeDSSessionRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *CreateThreeDSSessionRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetDeviceInfo

`func (o *CreateThreeDSSessionRequest) GetDeviceInfo() ThreeDSDeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *CreateThreeDSSessionRequest) GetDeviceInfoOk() (*ThreeDSDeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *CreateThreeDSSessionRequest) SetDeviceInfo(v ThreeDSDeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *CreateThreeDSSessionRequest) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


