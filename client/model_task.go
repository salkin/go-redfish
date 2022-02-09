/*
Redfish OAPI specification

Partial Redfish OAPI specification for a limited client

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// Task This resource contains information about a specific Task scheduled by or being executed by a Redfish service's Task Service.
type Task struct {
	// The OData description of a payload.
	OdataContext *string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag *string `json:"@odata.etag,omitempty"`
	// The name of the resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// description
	Description NullableString `json:"Description,omitempty"`
	// The date-time stamp that the task was last completed.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Indicates that the contents of the Payload should be hidden from view after the Task has been created.  When set to True, the Payload object will not be returned on GET.
	HidePayload *bool `json:"HidePayload,omitempty"`
	// The name of the resource.
	Id string `json:"Id"`
	// This is an array of messages associated with the task.
	Messages []Message `json:"Messages,omitempty"`
	// The name of the resource.
	Name string `json:"Name"`
	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem *string `json:"Oem,omitempty"`
	Payload *Payload `json:"Payload,omitempty"`
	// The date-time stamp that the task was last started.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The URI of the Task Monitor for this task.
	TaskMonitor *string `json:"TaskMonitor,omitempty"`
	TaskState *TaskState `json:"TaskState,omitempty"`
	TaskStatus *Health `json:"TaskStatus,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask(odataId string, odataType string, id string, name string) *Task {
	this := Task{}
	this.OdataId = odataId
	this.OdataType = odataType
	this.Id = id
	this.Name = name
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetOdataContext returns the OdataContext field value if set, zero value otherwise.
func (o *Task) GetOdataContext() string {
	if o == nil || o.OdataContext == nil {
		var ret string
		return ret
	}
	return *o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOdataContextOk() (*string, bool) {
	if o == nil || o.OdataContext == nil {
		return nil, false
	}
	return o.OdataContext, true
}

// HasOdataContext returns a boolean if a field has been set.
func (o *Task) HasOdataContext() bool {
	if o != nil && o.OdataContext != nil {
		return true
	}

	return false
}

// SetOdataContext gets a reference to the given string and assigns it to the OdataContext field.
func (o *Task) SetOdataContext(v string) {
	o.OdataContext = &v
}

// GetOdataEtag returns the OdataEtag field value if set, zero value otherwise.
func (o *Task) GetOdataEtag() string {
	if o == nil || o.OdataEtag == nil {
		var ret string
		return ret
	}
	return *o.OdataEtag
}

// GetOdataEtagOk returns a tuple with the OdataEtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOdataEtagOk() (*string, bool) {
	if o == nil || o.OdataEtag == nil {
		return nil, false
	}
	return o.OdataEtag, true
}

// HasOdataEtag returns a boolean if a field has been set.
func (o *Task) HasOdataEtag() bool {
	if o != nil && o.OdataEtag != nil {
		return true
	}

	return false
}

// SetOdataEtag gets a reference to the given string and assigns it to the OdataEtag field.
func (o *Task) SetOdataEtag(v string) {
	o.OdataEtag = &v
}

// GetOdataId returns the OdataId field value
func (o *Task) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *Task) GetOdataIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *Task) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *Task) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *Task) GetOdataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *Task) SetOdataType(v string) {
	o.OdataType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Task) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Task) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Task) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Task) UnsetDescription() {
	o.Description.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Task) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Task) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *Task) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetHidePayload returns the HidePayload field value if set, zero value otherwise.
func (o *Task) GetHidePayload() bool {
	if o == nil || o.HidePayload == nil {
		var ret bool
		return ret
	}
	return *o.HidePayload
}

// GetHidePayloadOk returns a tuple with the HidePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetHidePayloadOk() (*bool, bool) {
	if o == nil || o.HidePayload == nil {
		return nil, false
	}
	return o.HidePayload, true
}

// HasHidePayload returns a boolean if a field has been set.
func (o *Task) HasHidePayload() bool {
	if o != nil && o.HidePayload != nil {
		return true
	}

	return false
}

// SetHidePayload gets a reference to the given bool and assigns it to the HidePayload field.
func (o *Task) SetHidePayload(v bool) {
	o.HidePayload = &v
}

// GetId returns the Id field value
func (o *Task) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Task) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Task) SetId(v string) {
	o.Id = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Task) GetMessages() []Message {
	if o == nil || o.Messages == nil {
		var ret []Message
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetMessagesOk() ([]Message, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Task) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []Message and assigns it to the Messages field.
func (o *Task) SetMessages(v []Message) {
	o.Messages = v
}

// GetName returns the Name field value
func (o *Task) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Task) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Task) SetName(v string) {
	o.Name = v
}

// GetOem returns the Oem field value if set, zero value otherwise.
func (o *Task) GetOem() string {
	if o == nil || o.Oem == nil {
		var ret string
		return ret
	}
	return *o.Oem
}

// GetOemOk returns a tuple with the Oem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOemOk() (*string, bool) {
	if o == nil || o.Oem == nil {
		return nil, false
	}
	return o.Oem, true
}

// HasOem returns a boolean if a field has been set.
func (o *Task) HasOem() bool {
	if o != nil && o.Oem != nil {
		return true
	}

	return false
}

// SetOem gets a reference to the given string and assigns it to the Oem field.
func (o *Task) SetOem(v string) {
	o.Oem = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Task) GetPayload() Payload {
	if o == nil || o.Payload == nil {
		var ret Payload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPayloadOk() (*Payload, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Task) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given Payload and assigns it to the Payload field.
func (o *Task) SetPayload(v Payload) {
	o.Payload = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Task) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Task) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Task) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetTaskMonitor returns the TaskMonitor field value if set, zero value otherwise.
func (o *Task) GetTaskMonitor() string {
	if o == nil || o.TaskMonitor == nil {
		var ret string
		return ret
	}
	return *o.TaskMonitor
}

// GetTaskMonitorOk returns a tuple with the TaskMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskMonitorOk() (*string, bool) {
	if o == nil || o.TaskMonitor == nil {
		return nil, false
	}
	return o.TaskMonitor, true
}

// HasTaskMonitor returns a boolean if a field has been set.
func (o *Task) HasTaskMonitor() bool {
	if o != nil && o.TaskMonitor != nil {
		return true
	}

	return false
}

// SetTaskMonitor gets a reference to the given string and assigns it to the TaskMonitor field.
func (o *Task) SetTaskMonitor(v string) {
	o.TaskMonitor = &v
}

// GetTaskState returns the TaskState field value if set, zero value otherwise.
func (o *Task) GetTaskState() TaskState {
	if o == nil || o.TaskState == nil {
		var ret TaskState
		return ret
	}
	return *o.TaskState
}

// GetTaskStateOk returns a tuple with the TaskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskStateOk() (*TaskState, bool) {
	if o == nil || o.TaskState == nil {
		return nil, false
	}
	return o.TaskState, true
}

// HasTaskState returns a boolean if a field has been set.
func (o *Task) HasTaskState() bool {
	if o != nil && o.TaskState != nil {
		return true
	}

	return false
}

// SetTaskState gets a reference to the given TaskState and assigns it to the TaskState field.
func (o *Task) SetTaskState(v TaskState) {
	o.TaskState = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *Task) GetTaskStatus() Health {
	if o == nil || o.TaskStatus == nil {
		var ret Health
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskStatusOk() (*Health, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *Task) HasTaskStatus() bool {
	if o != nil && o.TaskStatus != nil {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given Health and assigns it to the TaskStatus field.
func (o *Task) SetTaskStatus(v Health) {
	o.TaskStatus = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OdataContext != nil {
		toSerialize["@odata.context"] = o.OdataContext
	}
	if o.OdataEtag != nil {
		toSerialize["@odata.etag"] = o.OdataEtag
	}
	if true {
		toSerialize["@odata.id"] = o.OdataId
	}
	if true {
		toSerialize["@odata.type"] = o.OdataType
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.HidePayload != nil {
		toSerialize["HidePayload"] = o.HidePayload
	}
	if true {
		toSerialize["Id"] = o.Id
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if o.Oem != nil {
		toSerialize["Oem"] = o.Oem
	}
	if o.Payload != nil {
		toSerialize["Payload"] = o.Payload
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.TaskMonitor != nil {
		toSerialize["TaskMonitor"] = o.TaskMonitor
	}
	if o.TaskState != nil {
		toSerialize["TaskState"] = o.TaskState
	}
	if o.TaskStatus != nil {
		toSerialize["TaskStatus"] = o.TaskStatus
	}
	return json.Marshal(toSerialize)
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


