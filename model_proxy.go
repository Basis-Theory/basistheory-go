/*
Basis Theory API

## Getting Started * Sign-in to [Basis Theory](https://basistheory.com) and go to [Applications](https://portal.basistheory.com/applications) * Create a Basis Theory Private Application * All permissions should be selected * Paste the API Key into the `BT-API-KEY` variable

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package basistheory

import (
	"encoding/json"
	"time"
)

// checks if the Proxy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Proxy{}

// Proxy struct for Proxy
type Proxy struct {
	Id                *string           `json:"id,omitempty"`
	Key               NullableString    `json:"key,omitempty"`
	TenantId          *string           `json:"tenant_id,omitempty"`
	Name              NullableString    `json:"name,omitempty"`
	DestinationUrl    NullableString    `json:"destination_url,omitempty"`
	RequestReactorId  NullableString    `json:"request_reactor_id,omitempty"`
	ResponseReactorId NullableString    `json:"response_reactor_id,omitempty"`
	RequireAuth       *bool             `json:"require_auth,omitempty"`
	RequestTransform  *ProxyTransform   `json:"request_transform,omitempty"`
	ResponseTransform *ProxyTransform   `json:"response_transform,omitempty"`
	ApplicationId     NullableString    `json:"application_id,omitempty"`
	Configuration     map[string]string `json:"configuration,omitempty"`
	ProxyHost         NullableString    `json:"proxy_host,omitempty"`
	CreatedBy         NullableString    `json:"created_by,omitempty"`
	CreatedAt         NullableTime      `json:"created_at,omitempty"`
	ModifiedBy        NullableString    `json:"modified_by,omitempty"`
	ModifiedAt        NullableTime      `json:"modified_at,omitempty"`
}

// NewProxy instantiates a new Proxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxy() *Proxy {
	this := Proxy{}
	return &this
}

// NewProxyWithDefaults instantiates a new Proxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyWithDefaults() *Proxy {
	this := Proxy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Proxy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proxy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field is not nil.
func (o *Proxy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Proxy) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field is not nil.
func (o *Proxy) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *Proxy) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *Proxy) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *Proxy) UnsetKey() {
	o.Key.Unset()
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *Proxy) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proxy) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field is not nil.
func (o *Proxy) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *Proxy) SetTenantId(v string) {
	o.TenantId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field is not nil.
func (o *Proxy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Proxy) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Proxy) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Proxy) UnsetName() {
	o.Name.Unset()
}

// GetDestinationUrl returns the DestinationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetDestinationUrl() string {
	if o == nil || IsNil(o.DestinationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DestinationUrl.Get()
}

// GetDestinationUrlOk returns a tuple with the DestinationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetDestinationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationUrl.Get(), o.DestinationUrl.IsSet()
}

// HasDestinationUrl returns a boolean if a field is not nil.
func (o *Proxy) HasDestinationUrl() bool {
	if o != nil && !IsNil(o.DestinationUrl) {
		return true
	}

	return false
}

// SetDestinationUrl gets a reference to the given NullableString and assigns it to the DestinationUrl field.
func (o *Proxy) SetDestinationUrl(v string) {
	o.DestinationUrl.Set(&v)
}

// SetDestinationUrlNil sets the value for DestinationUrl to be an explicit nil
func (o *Proxy) SetDestinationUrlNil() {
	o.DestinationUrl.Set(nil)
}

// UnsetDestinationUrl ensures that no value is present for DestinationUrl, not even an explicit nil
func (o *Proxy) UnsetDestinationUrl() {
	o.DestinationUrl.Unset()
}

// GetRequestReactorId returns the RequestReactorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetRequestReactorId() string {
	if o == nil || IsNil(o.RequestReactorId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestReactorId.Get()
}

// GetRequestReactorIdOk returns a tuple with the RequestReactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetRequestReactorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestReactorId.Get(), o.RequestReactorId.IsSet()
}

// HasRequestReactorId returns a boolean if a field is not nil.
func (o *Proxy) HasRequestReactorId() bool {
	if o != nil && !IsNil(o.RequestReactorId) {
		return true
	}

	return false
}

// SetRequestReactorId gets a reference to the given NullableString and assigns it to the RequestReactorId field.
func (o *Proxy) SetRequestReactorId(v string) {
	o.RequestReactorId.Set(&v)
}

// SetRequestReactorIdNil sets the value for RequestReactorId to be an explicit nil
func (o *Proxy) SetRequestReactorIdNil() {
	o.RequestReactorId.Set(nil)
}

// UnsetRequestReactorId ensures that no value is present for RequestReactorId, not even an explicit nil
func (o *Proxy) UnsetRequestReactorId() {
	o.RequestReactorId.Unset()
}

// GetResponseReactorId returns the ResponseReactorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetResponseReactorId() string {
	if o == nil || IsNil(o.ResponseReactorId.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseReactorId.Get()
}

// GetResponseReactorIdOk returns a tuple with the ResponseReactorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetResponseReactorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseReactorId.Get(), o.ResponseReactorId.IsSet()
}

// HasResponseReactorId returns a boolean if a field is not nil.
func (o *Proxy) HasResponseReactorId() bool {
	if o != nil && !IsNil(o.ResponseReactorId) {
		return true
	}

	return false
}

// SetResponseReactorId gets a reference to the given NullableString and assigns it to the ResponseReactorId field.
func (o *Proxy) SetResponseReactorId(v string) {
	o.ResponseReactorId.Set(&v)
}

// SetResponseReactorIdNil sets the value for ResponseReactorId to be an explicit nil
func (o *Proxy) SetResponseReactorIdNil() {
	o.ResponseReactorId.Set(nil)
}

// UnsetResponseReactorId ensures that no value is present for ResponseReactorId, not even an explicit nil
func (o *Proxy) UnsetResponseReactorId() {
	o.ResponseReactorId.Unset()
}

// GetRequireAuth returns the RequireAuth field value if set, zero value otherwise.
func (o *Proxy) GetRequireAuth() bool {
	if o == nil || IsNil(o.RequireAuth) {
		var ret bool
		return ret
	}
	return *o.RequireAuth
}

// GetRequireAuthOk returns a tuple with the RequireAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proxy) GetRequireAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireAuth) {
		return nil, false
	}
	return o.RequireAuth, true
}

// HasRequireAuth returns a boolean if a field is not nil.
func (o *Proxy) HasRequireAuth() bool {
	if o != nil && !IsNil(o.RequireAuth) {
		return true
	}

	return false
}

// SetRequireAuth gets a reference to the given bool and assigns it to the RequireAuth field.
func (o *Proxy) SetRequireAuth(v bool) {
	o.RequireAuth = &v
}

// GetRequestTransform returns the RequestTransform field value if set, zero value otherwise.
func (o *Proxy) GetRequestTransform() ProxyTransform {
	if o == nil || IsNil(o.RequestTransform) {
		var ret ProxyTransform
		return ret
	}
	return *o.RequestTransform
}

// GetRequestTransformOk returns a tuple with the RequestTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proxy) GetRequestTransformOk() (*ProxyTransform, bool) {
	if o == nil || IsNil(o.RequestTransform) {
		return nil, false
	}
	return o.RequestTransform, true
}

// HasRequestTransform returns a boolean if a field is not nil.
func (o *Proxy) HasRequestTransform() bool {
	if o != nil && !IsNil(o.RequestTransform) {
		return true
	}

	return false
}

// SetRequestTransform gets a reference to the given ProxyTransform and assigns it to the RequestTransform field.
func (o *Proxy) SetRequestTransform(v ProxyTransform) {
	o.RequestTransform = &v
}

// GetResponseTransform returns the ResponseTransform field value if set, zero value otherwise.
func (o *Proxy) GetResponseTransform() ProxyTransform {
	if o == nil || IsNil(o.ResponseTransform) {
		var ret ProxyTransform
		return ret
	}
	return *o.ResponseTransform
}

// GetResponseTransformOk returns a tuple with the ResponseTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proxy) GetResponseTransformOk() (*ProxyTransform, bool) {
	if o == nil || IsNil(o.ResponseTransform) {
		return nil, false
	}
	return o.ResponseTransform, true
}

// HasResponseTransform returns a boolean if a field is not nil.
func (o *Proxy) HasResponseTransform() bool {
	if o != nil && !IsNil(o.ResponseTransform) {
		return true
	}

	return false
}

// SetResponseTransform gets a reference to the given ProxyTransform and assigns it to the ResponseTransform field.
func (o *Proxy) SetResponseTransform(v ProxyTransform) {
	o.ResponseTransform = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId.Get()) {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field is not nil.
func (o *Proxy) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given NullableString and assigns it to the ApplicationId field.
func (o *Proxy) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}

// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil
func (o *Proxy) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
func (o *Proxy) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field is not nil.
func (o *Proxy) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *Proxy) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

// GetProxyHost returns the ProxyHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetProxyHost() string {
	if o == nil || IsNil(o.ProxyHost.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyHost.Get()
}

// GetProxyHostOk returns a tuple with the ProxyHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetProxyHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyHost.Get(), o.ProxyHost.IsSet()
}

// HasProxyHost returns a boolean if a field is not nil.
func (o *Proxy) HasProxyHost() bool {
	if o != nil && !IsNil(o.ProxyHost) {
		return true
	}

	return false
}

// SetProxyHost gets a reference to the given NullableString and assigns it to the ProxyHost field.
func (o *Proxy) SetProxyHost(v string) {
	o.ProxyHost.Set(&v)
}

// SetProxyHostNil sets the value for ProxyHost to be an explicit nil
func (o *Proxy) SetProxyHostNil() {
	o.ProxyHost.Set(nil)
}

// UnsetProxyHost ensures that no value is present for ProxyHost, not even an explicit nil
func (o *Proxy) UnsetProxyHost() {
	o.ProxyHost.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field is not nil.
func (o *Proxy) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *Proxy) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *Proxy) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *Proxy) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field is not nil.
func (o *Proxy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *Proxy) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Proxy) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Proxy) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field is not nil.
func (o *Proxy) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableString and assigns it to the ModifiedBy field.
func (o *Proxy) SetModifiedBy(v string) {
	o.ModifiedBy.Set(&v)
}

// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *Proxy) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *Proxy) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Proxy) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Proxy) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field is not nil.
func (o *Proxy) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given NullableTime and assigns it to the ModifiedAt field.
func (o *Proxy) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil
func (o *Proxy) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
func (o *Proxy) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

func (o Proxy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Proxy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DestinationUrl.IsSet() {
		toSerialize["destination_url"] = o.DestinationUrl.Get()
	}
	if o.RequestReactorId.IsSet() {
		toSerialize["request_reactor_id"] = o.RequestReactorId.Get()
	}
	if o.ResponseReactorId.IsSet() {
		toSerialize["response_reactor_id"] = o.ResponseReactorId.Get()
	}
	if !IsNil(o.RequireAuth) {
		toSerialize["require_auth"] = o.RequireAuth
	}
	if !IsNil(o.RequestTransform) {
		toSerialize["request_transform"] = o.RequestTransform
	}
	if !IsNil(o.ResponseTransform) {
		toSerialize["response_transform"] = o.ResponseTransform
	}
	if o.ApplicationId.IsSet() {
		toSerialize["application_id"] = o.ApplicationId.Get()
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ProxyHost.IsSet() {
		toSerialize["proxy_host"] = o.ProxyHost.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["created_by"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modified_by"] = o.ModifiedBy.Get()
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return toSerialize, nil
}

type NullableProxy struct {
	value *Proxy
	isSet bool
}

func (v NullableProxy) Get() *Proxy {
	return v.value
}

func (v *NullableProxy) Set(val *Proxy) {
	v.value = val
	v.isSet = true
}

func (v NullableProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxy(val *Proxy) *NullableProxy {
	return &NullableProxy{value: val, isSet: true}
}

func (v NullableProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
