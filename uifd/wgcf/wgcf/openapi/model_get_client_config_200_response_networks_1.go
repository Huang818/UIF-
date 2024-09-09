/*
 * untitled API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 536
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetClientConfig200ResponseNetworks1 struct for GetClientConfig200ResponseNetworks1
type GetClientConfig200ResponseNetworks1 struct {
	V4 []GetClientConfig200ResponseNetworks1V4 `json:"v4"`
	V6 []GetClientConfig200ResponseNetworks1V6 `json:"v6"`
}

// NewGetClientConfig200ResponseNetworks1 instantiates a new GetClientConfig200ResponseNetworks1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClientConfig200ResponseNetworks1(v4 []GetClientConfig200ResponseNetworks1V4, v6 []GetClientConfig200ResponseNetworks1V6, ) *GetClientConfig200ResponseNetworks1 {
	this := GetClientConfig200ResponseNetworks1{}
	this.V4 = v4
	this.V6 = v6
	return &this
}

// NewGetClientConfig200ResponseNetworks1WithDefaults instantiates a new GetClientConfig200ResponseNetworks1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClientConfig200ResponseNetworks1WithDefaults() *GetClientConfig200ResponseNetworks1 {
	this := GetClientConfig200ResponseNetworks1{}
	return &this
}

// GetV4 returns the V4 field value
func (o *GetClientConfig200ResponseNetworks1) GetV4() []GetClientConfig200ResponseNetworks1V4 {
	if o == nil  {
		var ret []GetClientConfig200ResponseNetworks1V4
		return ret
	}

	return o.V4
}

// GetV4Ok returns a tuple with the V4 field value
// and a boolean to check if the value has been set.
func (o *GetClientConfig200ResponseNetworks1) GetV4Ok() (*[]GetClientConfig200ResponseNetworks1V4, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.V4, true
}

// SetV4 sets field value
func (o *GetClientConfig200ResponseNetworks1) SetV4(v []GetClientConfig200ResponseNetworks1V4) {
	o.V4 = v
}

// GetV6 returns the V6 field value
func (o *GetClientConfig200ResponseNetworks1) GetV6() []GetClientConfig200ResponseNetworks1V6 {
	if o == nil  {
		var ret []GetClientConfig200ResponseNetworks1V6
		return ret
	}

	return o.V6
}

// GetV6Ok returns a tuple with the V6 field value
// and a boolean to check if the value has been set.
func (o *GetClientConfig200ResponseNetworks1) GetV6Ok() (*[]GetClientConfig200ResponseNetworks1V6, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.V6, true
}

// SetV6 sets field value
func (o *GetClientConfig200ResponseNetworks1) SetV6(v []GetClientConfig200ResponseNetworks1V6) {
	o.V6 = v
}

func (o GetClientConfig200ResponseNetworks1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["v4"] = o.V4
	}
	if true {
		toSerialize["v6"] = o.V6
	}
	return json.Marshal(toSerialize)
}

type NullableGetClientConfig200ResponseNetworks1 struct {
	value *GetClientConfig200ResponseNetworks1
	isSet bool
}

func (v NullableGetClientConfig200ResponseNetworks1) Get() *GetClientConfig200ResponseNetworks1 {
	return v.value
}

func (v *NullableGetClientConfig200ResponseNetworks1) Set(val *GetClientConfig200ResponseNetworks1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClientConfig200ResponseNetworks1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClientConfig200ResponseNetworks1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClientConfig200ResponseNetworks1(val *GetClientConfig200ResponseNetworks1) *NullableGetClientConfig200ResponseNetworks1 {
	return &NullableGetClientConfig200ResponseNetworks1{value: val, isSet: true}
}

func (v NullableGetClientConfig200ResponseNetworks1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClientConfig200ResponseNetworks1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


