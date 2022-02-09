/*
Redfish OAPI specification

Partial Redfish OAPI specification for a limited client

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Payload The HTTP and JSON payload details for this Task.
type Payload struct {
	// This represents the HTTP headers used in the operation of this Task.
	HttpHeaders []string `json:"HttpHeaders,omitempty"`
	// The HTTP operation to perform to execute this Task.
	HttpOperation *string `json:"HttpOperation,omitempty"`
	// This property contains the JSON payload to use in the execution of this Task.
	JsonBody *string `json:"JsonBody,omitempty"`
	// The URI of the target for this task.
	TargetUri *string `json:"TargetUri,omitempty"`
}

// NewPayload instantiates a new Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayload() *Payload {
	this := Payload{}
	return &this
}

// NewPayloadWithDefaults instantiates a new Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadWithDefaults() *Payload {
	this := Payload{}
	return &this
}

// GetHttpHeaders returns the HttpHeaders field value if set, zero value otherwise.
func (o *Payload) GetHttpHeaders() []string {
	if o == nil || o.HttpHeaders == nil {
		var ret []string
		return ret
	}
	return o.HttpHeaders
}

// GetHttpHeadersOk returns a tuple with the HttpHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetHttpHeadersOk() ([]string, bool) {
	if o == nil || o.HttpHeaders == nil {
		return nil, false
	}
	return o.HttpHeaders, true
}

// HasHttpHeaders returns a boolean if a field has been set.
func (o *Payload) HasHttpHeaders() bool {
	if o != nil && o.HttpHeaders != nil {
		return true
	}

	return false
}

// SetHttpHeaders gets a reference to the given []string and assigns it to the HttpHeaders field.
func (o *Payload) SetHttpHeaders(v []string) {
	o.HttpHeaders = v
}

// GetHttpOperation returns the HttpOperation field value if set, zero value otherwise.
func (o *Payload) GetHttpOperation() string {
	if o == nil || o.HttpOperation == nil {
		var ret string
		return ret
	}
	return *o.HttpOperation
}

// GetHttpOperationOk returns a tuple with the HttpOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetHttpOperationOk() (*string, bool) {
	if o == nil || o.HttpOperation == nil {
		return nil, false
	}
	return o.HttpOperation, true
}

// HasHttpOperation returns a boolean if a field has been set.
func (o *Payload) HasHttpOperation() bool {
	if o != nil && o.HttpOperation != nil {
		return true
	}

	return false
}

// SetHttpOperation gets a reference to the given string and assigns it to the HttpOperation field.
func (o *Payload) SetHttpOperation(v string) {
	o.HttpOperation = &v
}

// GetJsonBody returns the JsonBody field value if set, zero value otherwise.
func (o *Payload) GetJsonBody() string {
	if o == nil || o.JsonBody == nil {
		var ret string
		return ret
	}
	return *o.JsonBody
}

// GetJsonBodyOk returns a tuple with the JsonBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetJsonBodyOk() (*string, bool) {
	if o == nil || o.JsonBody == nil {
		return nil, false
	}
	return o.JsonBody, true
}

// HasJsonBody returns a boolean if a field has been set.
func (o *Payload) HasJsonBody() bool {
	if o != nil && o.JsonBody != nil {
		return true
	}

	return false
}

// SetJsonBody gets a reference to the given string and assigns it to the JsonBody field.
func (o *Payload) SetJsonBody(v string) {
	o.JsonBody = &v
}

// GetTargetUri returns the TargetUri field value if set, zero value otherwise.
func (o *Payload) GetTargetUri() string {
	if o == nil || o.TargetUri == nil {
		var ret string
		return ret
	}
	return *o.TargetUri
}

// GetTargetUriOk returns a tuple with the TargetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetTargetUriOk() (*string, bool) {
	if o == nil || o.TargetUri == nil {
		return nil, false
	}
	return o.TargetUri, true
}

// HasTargetUri returns a boolean if a field has been set.
func (o *Payload) HasTargetUri() bool {
	if o != nil && o.TargetUri != nil {
		return true
	}

	return false
}

// SetTargetUri gets a reference to the given string and assigns it to the TargetUri field.
func (o *Payload) SetTargetUri(v string) {
	o.TargetUri = &v
}

func (o Payload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpHeaders != nil {
		toSerialize["HttpHeaders"] = o.HttpHeaders
	}
	if o.HttpOperation != nil {
		toSerialize["HttpOperation"] = o.HttpOperation
	}
	if o.JsonBody != nil {
		toSerialize["JsonBody"] = o.JsonBody
	}
	if o.TargetUri != nil {
		toSerialize["TargetUri"] = o.TargetUri
	}
	return json.Marshal(toSerialize)
}

type NullablePayload struct {
	value *Payload
	isSet bool
}

func (v NullablePayload) Get() *Payload {
	return v.value
}

func (v *NullablePayload) Set(val *Payload) {
	v.value = val
	v.isSet = true
}

func (v NullablePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayload(val *Payload) *NullablePayload {
	return &NullablePayload{value: val, isSet: true}
}

func (v NullablePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


