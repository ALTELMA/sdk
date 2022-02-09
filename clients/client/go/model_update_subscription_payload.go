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

// UpdateSubscriptionPayload struct for UpdateSubscriptionPayload
type UpdateSubscriptionPayload struct {
	PlanOrPrice string `json:"plan_or_price"`
	ReturnTo *string `json:"return_to,omitempty"`
}

// NewUpdateSubscriptionPayload instantiates a new UpdateSubscriptionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionPayload(planOrPrice string) *UpdateSubscriptionPayload {
	this := UpdateSubscriptionPayload{}
	this.PlanOrPrice = planOrPrice
	return &this
}

// NewUpdateSubscriptionPayloadWithDefaults instantiates a new UpdateSubscriptionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionPayloadWithDefaults() *UpdateSubscriptionPayload {
	this := UpdateSubscriptionPayload{}
	return &this
}

// GetPlanOrPrice returns the PlanOrPrice field value
func (o *UpdateSubscriptionPayload) GetPlanOrPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanOrPrice
}

// GetPlanOrPriceOk returns a tuple with the PlanOrPrice field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionPayload) GetPlanOrPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlanOrPrice, true
}

// SetPlanOrPrice sets field value
func (o *UpdateSubscriptionPayload) SetPlanOrPrice(v string) {
	o.PlanOrPrice = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *UpdateSubscriptionPayload) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionPayload) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *UpdateSubscriptionPayload) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *UpdateSubscriptionPayload) SetReturnTo(v string) {
	o.ReturnTo = &v
}

func (o UpdateSubscriptionPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plan_or_price"] = o.PlanOrPrice
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubscriptionPayload struct {
	value *UpdateSubscriptionPayload
	isSet bool
}

func (v NullableUpdateSubscriptionPayload) Get() *UpdateSubscriptionPayload {
	return v.value
}

func (v *NullableUpdateSubscriptionPayload) Set(val *UpdateSubscriptionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionPayload(val *UpdateSubscriptionPayload) *NullableUpdateSubscriptionPayload {
	return &NullableUpdateSubscriptionPayload{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


