/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.81
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdentitySchemaLocation struct for IdentitySchemaLocation
type IdentitySchemaLocation struct {
	Location *string `json:"location,omitempty"`
}

// NewIdentitySchemaLocation instantiates a new IdentitySchemaLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySchemaLocation() *IdentitySchemaLocation {
	this := IdentitySchemaLocation{}
	return &this
}

// NewIdentitySchemaLocationWithDefaults instantiates a new IdentitySchemaLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySchemaLocationWithDefaults() *IdentitySchemaLocation {
	this := IdentitySchemaLocation{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *IdentitySchemaLocation) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySchemaLocation) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *IdentitySchemaLocation) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *IdentitySchemaLocation) SetLocation(v string) {
	o.Location = &v
}

func (o IdentitySchemaLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySchemaLocation struct {
	value *IdentitySchemaLocation
	isSet bool
}

func (v NullableIdentitySchemaLocation) Get() *IdentitySchemaLocation {
	return v.value
}

func (v *NullableIdentitySchemaLocation) Set(val *IdentitySchemaLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySchemaLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySchemaLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySchemaLocation(val *IdentitySchemaLocation) *NullableIdentitySchemaLocation {
	return &NullableIdentitySchemaLocation{value: val, isSet: true}
}

func (v NullableIdentitySchemaLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySchemaLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


