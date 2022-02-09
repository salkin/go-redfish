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

// UpdateService Redfish Update Service.
type UpdateService struct {
	FirmwareInventory *FirmwareInventory `json:"FirmwareInventory,omitempty"`
	// The name of the resource.
	Id *string `json:"Id,omitempty"`
	// The name of the resource.
	Name string `json:"Name"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The OData description of a payload.
	OdataContext *string `json:"@odata.context,omitempty"`
	// description
	Description NullableString `json:"Description,omitempty"`
	HttpPushUri *string `json:"HttpPushUri,omitempty"`
	ServiceEnabled NullableBool `json:"ServiceEnabled,omitempty"`
	Actions *UpdateServiceActions `json:"Actions,omitempty"`
}

// NewUpdateService instantiates a new UpdateService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateService(name string, odataType string, odataId string) *UpdateService {
	this := UpdateService{}
	this.Name = name
	this.OdataType = odataType
	this.OdataId = odataId
	return &this
}

// NewUpdateServiceWithDefaults instantiates a new UpdateService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceWithDefaults() *UpdateService {
	this := UpdateService{}
	return &this
}

// GetFirmwareInventory returns the FirmwareInventory field value if set, zero value otherwise.
func (o *UpdateService) GetFirmwareInventory() FirmwareInventory {
	if o == nil || o.FirmwareInventory == nil {
		var ret FirmwareInventory
		return ret
	}
	return *o.FirmwareInventory
}

// GetFirmwareInventoryOk returns a tuple with the FirmwareInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateService) GetFirmwareInventoryOk() (*FirmwareInventory, bool) {
	if o == nil || o.FirmwareInventory == nil {
		return nil, false
	}
	return o.FirmwareInventory, true
}

// HasFirmwareInventory returns a boolean if a field has been set.
func (o *UpdateService) HasFirmwareInventory() bool {
	if o != nil && o.FirmwareInventory != nil {
		return true
	}

	return false
}

// SetFirmwareInventory gets a reference to the given FirmwareInventory and assigns it to the FirmwareInventory field.
func (o *UpdateService) SetFirmwareInventory(v FirmwareInventory) {
	o.FirmwareInventory = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateService) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateService) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateService) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateService) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *UpdateService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateService) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateService) SetName(v string) {
	o.Name = v
}

// GetOdataType returns the OdataType field value
func (o *UpdateService) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *UpdateService) GetOdataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *UpdateService) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataId returns the OdataId field value
func (o *UpdateService) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *UpdateService) GetOdataIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *UpdateService) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value if set, zero value otherwise.
func (o *UpdateService) GetOdataContext() string {
	if o == nil || o.OdataContext == nil {
		var ret string
		return ret
	}
	return *o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateService) GetOdataContextOk() (*string, bool) {
	if o == nil || o.OdataContext == nil {
		return nil, false
	}
	return o.OdataContext, true
}

// HasOdataContext returns a boolean if a field has been set.
func (o *UpdateService) HasOdataContext() bool {
	if o != nil && o.OdataContext != nil {
		return true
	}

	return false
}

// SetOdataContext gets a reference to the given string and assigns it to the OdataContext field.
func (o *UpdateService) SetOdataContext(v string) {
	o.OdataContext = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateService) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateService) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateService) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateService) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateService) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateService) UnsetDescription() {
	o.Description.Unset()
}

// GetHttpPushUri returns the HttpPushUri field value if set, zero value otherwise.
func (o *UpdateService) GetHttpPushUri() string {
	if o == nil || o.HttpPushUri == nil {
		var ret string
		return ret
	}
	return *o.HttpPushUri
}

// GetHttpPushUriOk returns a tuple with the HttpPushUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateService) GetHttpPushUriOk() (*string, bool) {
	if o == nil || o.HttpPushUri == nil {
		return nil, false
	}
	return o.HttpPushUri, true
}

// HasHttpPushUri returns a boolean if a field has been set.
func (o *UpdateService) HasHttpPushUri() bool {
	if o != nil && o.HttpPushUri != nil {
		return true
	}

	return false
}

// SetHttpPushUri gets a reference to the given string and assigns it to the HttpPushUri field.
func (o *UpdateService) SetHttpPushUri(v string) {
	o.HttpPushUri = &v
}

// GetServiceEnabled returns the ServiceEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateService) GetServiceEnabled() bool {
	if o == nil || o.ServiceEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ServiceEnabled.Get()
}

// GetServiceEnabledOk returns a tuple with the ServiceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateService) GetServiceEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceEnabled.Get(), o.ServiceEnabled.IsSet()
}

// HasServiceEnabled returns a boolean if a field has been set.
func (o *UpdateService) HasServiceEnabled() bool {
	if o != nil && o.ServiceEnabled.IsSet() {
		return true
	}

	return false
}

// SetServiceEnabled gets a reference to the given NullableBool and assigns it to the ServiceEnabled field.
func (o *UpdateService) SetServiceEnabled(v bool) {
	o.ServiceEnabled.Set(&v)
}
// SetServiceEnabledNil sets the value for ServiceEnabled to be an explicit nil
func (o *UpdateService) SetServiceEnabledNil() {
	o.ServiceEnabled.Set(nil)
}

// UnsetServiceEnabled ensures that no value is present for ServiceEnabled, not even an explicit nil
func (o *UpdateService) UnsetServiceEnabled() {
	o.ServiceEnabled.Unset()
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *UpdateService) GetActions() UpdateServiceActions {
	if o == nil || o.Actions == nil {
		var ret UpdateServiceActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateService) GetActionsOk() (*UpdateServiceActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *UpdateService) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given UpdateServiceActions and assigns it to the Actions field.
func (o *UpdateService) SetActions(v UpdateServiceActions) {
	o.Actions = &v
}

func (o UpdateService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirmwareInventory != nil {
		toSerialize["FirmwareInventory"] = o.FirmwareInventory
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["@odata.type"] = o.OdataType
	}
	if true {
		toSerialize["@odata.id"] = o.OdataId
	}
	if o.OdataContext != nil {
		toSerialize["@odata.context"] = o.OdataContext
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.HttpPushUri != nil {
		toSerialize["HttpPushUri"] = o.HttpPushUri
	}
	if o.ServiceEnabled.IsSet() {
		toSerialize["ServiceEnabled"] = o.ServiceEnabled.Get()
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateService struct {
	value *UpdateService
	isSet bool
}

func (v NullableUpdateService) Get() *UpdateService {
	return v.value
}

func (v *NullableUpdateService) Set(val *UpdateService) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateService) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateService(val *UpdateService) *NullableUpdateService {
	return &NullableUpdateService{value: val, isSet: true}
}

func (v NullableUpdateService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


