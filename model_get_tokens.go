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

// checks if the GetTokens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTokens{}

// GetTokens struct for GetTokens
type GetTokens struct {
	Id       []string          `json:"id,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Page     NullableInt32     `json:"page,omitempty"`
	Start    NullableString    `json:"start,omitempty"`
	Size     NullableInt32     `json:"size,omitempty"`
}

// NewGetTokens instantiates a new GetTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTokens() *GetTokens {
	this := GetTokens{}
	return &this
}

// NewGetTokensWithDefaults instantiates a new GetTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokensWithDefaults() *GetTokens {
	this := GetTokens{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTokens) GetId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTokens) GetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field is not nil.
func (o *GetTokens) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *GetTokens) SetId(v []string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTokens) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTokens) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field is not nil.
func (o *GetTokens) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *GetTokens) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTokens) GetPage() int32 {
	if o == nil || IsNil(o.Page.Get()) {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTokens) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field is not nil.
func (o *GetTokens) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *GetTokens) SetPage(v int32) {
	o.Page.Set(&v)
}

// SetPageNil sets the value for Page to be an explicit nil
func (o *GetTokens) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *GetTokens) UnsetPage() {
	o.Page.Unset()
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTokens) GetStart() string {
	if o == nil || IsNil(o.Start.Get()) {
		var ret string
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTokens) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field is not nil.
func (o *GetTokens) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableString and assigns it to the Start field.
func (o *GetTokens) SetStart(v string) {
	o.Start.Set(&v)
}

// SetStartNil sets the value for Start to be an explicit nil
func (o *GetTokens) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *GetTokens) UnsetStart() {
	o.Start.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetTokens) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetTokens) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field is not nil.
func (o *GetTokens) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *GetTokens) SetSize(v int32) {
	o.Size.Set(&v)
}

// SetSizeNil sets the value for Size to be an explicit nil
func (o *GetTokens) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *GetTokens) UnsetSize() {
	o.Size.Unset()
}

func (o GetTokens) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTokens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return toSerialize, nil
}

type NullableGetTokens struct {
	value *GetTokens
	isSet bool
}

func (v NullableGetTokens) Get() *GetTokens {
	return v.value
}

func (v *NullableGetTokens) Set(val *GetTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTokens(val *GetTokens) *NullableGetTokens {
	return &NullableGetTokens{value: val, isSet: true}
}

func (v NullableGetTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
