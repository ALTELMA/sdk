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
	"time"
)

// SelfServiceVerificationFlow Used to verify an out-of-band communication channel such as an email address or a phone number.  For more information head over to: https://www.ory.sh/docs/kratos/selfservice/flows/verify-email-account-activation
type SelfServiceVerificationFlow struct {
	// Active, if set, contains the registration method that is being used. It is initially not set.
	Active *string `json:"active,omitempty"`
	// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to verify the address, a new request has to be initiated.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the request occurred.
	IssuedAt *time.Time `json:"issued_at,omitempty"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl *string `json:"request_url,omitempty"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	State SelfServiceVerificationFlowState `json:"state"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
}

// NewSelfServiceVerificationFlow instantiates a new SelfServiceVerificationFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceVerificationFlow(id string, state SelfServiceVerificationFlowState, type_ string, ui UiContainer) *SelfServiceVerificationFlow {
	this := SelfServiceVerificationFlow{}
	this.Id = id
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewSelfServiceVerificationFlowWithDefaults instantiates a new SelfServiceVerificationFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceVerificationFlowWithDefaults() *SelfServiceVerificationFlow {
	this := SelfServiceVerificationFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SelfServiceVerificationFlow) GetActive() string {
	if o == nil || o.Active == nil {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetActiveOk() (*string, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SelfServiceVerificationFlow) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *SelfServiceVerificationFlow) SetActive(v string) {
	o.Active = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *SelfServiceVerificationFlow) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SelfServiceVerificationFlow) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *SelfServiceVerificationFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *SelfServiceVerificationFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SelfServiceVerificationFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise.
func (o *SelfServiceVerificationFlow) GetIssuedAt() time.Time {
	if o == nil || o.IssuedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil || o.IssuedAt == nil {
		return nil, false
	}
	return o.IssuedAt, true
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *SelfServiceVerificationFlow) HasIssuedAt() bool {
	if o != nil && o.IssuedAt != nil {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given time.Time and assigns it to the IssuedAt field.
func (o *SelfServiceVerificationFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = &v
}

// GetRequestUrl returns the RequestUrl field value if set, zero value otherwise.
func (o *SelfServiceVerificationFlow) GetRequestUrl() string {
	if o == nil || o.RequestUrl == nil {
		var ret string
		return ret
	}
	return *o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil || o.RequestUrl == nil {
		return nil, false
	}
	return o.RequestUrl, true
}

// HasRequestUrl returns a boolean if a field has been set.
func (o *SelfServiceVerificationFlow) HasRequestUrl() bool {
	if o != nil && o.RequestUrl != nil {
		return true
	}

	return false
}

// SetRequestUrl gets a reference to the given string and assigns it to the RequestUrl field.
func (o *SelfServiceVerificationFlow) SetRequestUrl(v string) {
	o.RequestUrl = &v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *SelfServiceVerificationFlow) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *SelfServiceVerificationFlow) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *SelfServiceVerificationFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetState returns the State field value
func (o *SelfServiceVerificationFlow) GetState() SelfServiceVerificationFlowState {
	if o == nil {
		var ret SelfServiceVerificationFlowState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetStateOk() (*SelfServiceVerificationFlowState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SelfServiceVerificationFlow) SetState(v SelfServiceVerificationFlowState) {
	o.State = v
}

// GetType returns the Type field value
func (o *SelfServiceVerificationFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SelfServiceVerificationFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *SelfServiceVerificationFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *SelfServiceVerificationFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *SelfServiceVerificationFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o SelfServiceVerificationFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IssuedAt != nil {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if o.RequestUrl != nil {
		toSerialize["request_url"] = o.RequestUrl
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceVerificationFlow struct {
	value *SelfServiceVerificationFlow
	isSet bool
}

func (v NullableSelfServiceVerificationFlow) Get() *SelfServiceVerificationFlow {
	return v.value
}

func (v *NullableSelfServiceVerificationFlow) Set(val *SelfServiceVerificationFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceVerificationFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceVerificationFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceVerificationFlow(val *SelfServiceVerificationFlow) *NullableSelfServiceVerificationFlow {
	return &NullableSelfServiceVerificationFlow{value: val, isSet: true}
}

func (v NullableSelfServiceVerificationFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceVerificationFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


