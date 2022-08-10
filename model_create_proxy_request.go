/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
)

// CreateProxyRequest struct for CreateProxyRequest
type CreateProxyRequest struct {
	Name              string         `json:"name"`
	DestinationUrl    string         `json:"destination_url"`
	RequestReactorId  NullableString `json:"request_reactor_id,omitempty"`
	ResponseReactorId NullableString `json:"response_reactor_id,omitempty"`
	RequireAuth       NullableBool   `json:"require_auth,omitempty"`
}

// NewCreateProxyRequest instantiates a new CreateProxyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProxyRequest(name string, destinationUrl string) *CreateProxyRequest {
	this := CreateProxyRequest{}
	this.Name = name
	this.DestinationUrl = destinationUrl
	return &this
}

// NewCreateProxyRequestWithDefaults instantiates a new CreateProxyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProxyRequestWithDefaults() *CreateProxyRequest {
	this := CreateProxyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateProxyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProxyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProxyRequest) SetName(v string) {
	o.Name = v
}

// GetDestinationUrl returns the DestinationUrl field value
func (o *CreateProxyRequest) GetDestinationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationUrl
}

// GetDestinationUrlOk returns a tuple with the DestinationUrl field value
// and a boolean to check if the value has been set.
func (o *CreateProxyRequest) GetDestinationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationUrl, true
}

// SetDestinationUrl sets field value
func (o *CreateProxyRequest) SetDestinationUrl(v string) {
	o.DestinationUrl = v
}

// GetRequestReactorId returns the RequestReactorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxyRequest) GetRequestReactorId() string {
	if o == nil || o.RequestReactorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestReactorId.Get()
}

// GetRequestReactorIdOk returns a tuple with the RequestReactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxyRequest) GetRequestReactorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestReactorId.Get(), o.RequestReactorId.IsSet()
}

// HasRequestReactorId returns a boolean if a field has been set.
func (o *CreateProxyRequest) HasRequestReactorId() bool {
	if o != nil && o.RequestReactorId.IsSet() {
		return true
	}

	return false
}

// SetRequestReactorId gets a reference to the given NullableString and assigns it to the RequestReactorId field.
func (o *CreateProxyRequest) SetRequestReactorId(v string) {
	o.RequestReactorId.Set(&v)
}

// SetRequestReactorIdNil sets the value for RequestReactorId to be an explicit nil
func (o *CreateProxyRequest) SetRequestReactorIdNil() {
	o.RequestReactorId.Set(nil)
}

// UnsetRequestReactorId ensures that no value is present for RequestReactorId, not even an explicit nil
func (o *CreateProxyRequest) UnsetRequestReactorId() {
	o.RequestReactorId.Unset()
}

// GetResponseReactorId returns the ResponseReactorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxyRequest) GetResponseReactorId() string {
	if o == nil || o.ResponseReactorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResponseReactorId.Get()
}

// GetResponseReactorIdOk returns a tuple with the ResponseReactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxyRequest) GetResponseReactorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseReactorId.Get(), o.ResponseReactorId.IsSet()
}

// HasResponseReactorId returns a boolean if a field has been set.
func (o *CreateProxyRequest) HasResponseReactorId() bool {
	if o != nil && o.ResponseReactorId.IsSet() {
		return true
	}

	return false
}

// SetResponseReactorId gets a reference to the given NullableString and assigns it to the ResponseReactorId field.
func (o *CreateProxyRequest) SetResponseReactorId(v string) {
	o.ResponseReactorId.Set(&v)
}

// SetResponseReactorIdNil sets the value for ResponseReactorId to be an explicit nil
func (o *CreateProxyRequest) SetResponseReactorIdNil() {
	o.ResponseReactorId.Set(nil)
}

// UnsetResponseReactorId ensures that no value is present for ResponseReactorId, not even an explicit nil
func (o *CreateProxyRequest) UnsetResponseReactorId() {
	o.ResponseReactorId.Unset()
}

// GetRequireAuth returns the RequireAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxyRequest) GetRequireAuth() bool {
	if o == nil || o.RequireAuth.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RequireAuth.Get()
}

// GetRequireAuthOk returns a tuple with the RequireAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxyRequest) GetRequireAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequireAuth.Get(), o.RequireAuth.IsSet()
}

// HasRequireAuth returns a boolean if a field has been set.
func (o *CreateProxyRequest) HasRequireAuth() bool {
	if o != nil && o.RequireAuth.IsSet() {
		return true
	}

	return false
}

// SetRequireAuth gets a reference to the given NullableBool and assigns it to the RequireAuth field.
func (o *CreateProxyRequest) SetRequireAuth(v bool) {
	o.RequireAuth.Set(&v)
}

// SetRequireAuthNil sets the value for RequireAuth to be an explicit nil
func (o *CreateProxyRequest) SetRequireAuthNil() {
	o.RequireAuth.Set(nil)
}

// UnsetRequireAuth ensures that no value is present for RequireAuth, not even an explicit nil
func (o *CreateProxyRequest) UnsetRequireAuth() {
	o.RequireAuth.Unset()
}

func (o CreateProxyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["destination_url"] = o.DestinationUrl
	}
	if o.RequestReactorId.IsSet() {
		toSerialize["request_reactor_id"] = o.RequestReactorId.Get()
	}
	if o.ResponseReactorId.IsSet() {
		toSerialize["response_reactor_id"] = o.ResponseReactorId.Get()
	}
	if o.RequireAuth.IsSet() {
		toSerialize["require_auth"] = o.RequireAuth.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProxyRequest struct {
	value *CreateProxyRequest
	isSet bool
}

func (v NullableCreateProxyRequest) Get() *CreateProxyRequest {
	return v.value
}

func (v *NullableCreateProxyRequest) Set(val *CreateProxyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProxyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProxyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProxyRequest(val *CreateProxyRequest) *NullableCreateProxyRequest {
	return &NullableCreateProxyRequest{value: val, isSet: true}
}

func (v NullableCreateProxyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProxyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
