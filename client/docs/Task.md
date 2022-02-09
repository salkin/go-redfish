# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | Pointer to **string** | The OData description of a payload. | [optional] [readonly] 
**OdataEtag** | Pointer to **string** | The current ETag of the resource. | [optional] [readonly] 
**OdataId** | **string** | The name of the resource. | [readonly] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**Description** | Pointer to **NullableString** | description | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | The date-time stamp that the task was last completed. | [optional] [readonly] 
**HidePayload** | Pointer to **bool** | Indicates that the contents of the Payload should be hidden from view after the Task has been created.  When set to True, the Payload object will not be returned on GET. | [optional] [readonly] 
**Id** | **string** | The name of the resource. | [readonly] 
**Messages** | Pointer to [**[]Message**](Message.md) | This is an array of messages associated with the task. | [optional] 
**Name** | **string** | The name of the resource. | [readonly] 
**Oem** | Pointer to **string** | This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections. | [optional] 
**Payload** | Pointer to [**Payload**](Payload.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | The date-time stamp that the task was last started. | [optional] [readonly] 
**TaskMonitor** | Pointer to **string** | The URI of the Task Monitor for this task. | [optional] [readonly] 
**TaskState** | Pointer to [**TaskState**](TaskState.md) |  | [optional] 
**TaskStatus** | Pointer to [**Health**](Health.md) |  | [optional] 

## Methods

### NewTask

`func NewTask(odataId string, odataType string, id string, name string, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOdataContext

`func (o *Task) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *Task) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *Task) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *Task) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.

### GetOdataEtag

`func (o *Task) GetOdataEtag() string`

GetOdataEtag returns the OdataEtag field if non-nil, zero value otherwise.

### GetOdataEtagOk

`func (o *Task) GetOdataEtagOk() (*string, bool)`

GetOdataEtagOk returns a tuple with the OdataEtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataEtag

`func (o *Task) SetOdataEtag(v string)`

SetOdataEtag sets OdataEtag field to given value.

### HasOdataEtag

`func (o *Task) HasOdataEtag() bool`

HasOdataEtag returns a boolean if a field has been set.

### GetOdataId

`func (o *Task) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *Task) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *Task) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *Task) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *Task) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *Task) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Task) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Task) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTime

`func (o *Task) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Task) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Task) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Task) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHidePayload

`func (o *Task) GetHidePayload() bool`

GetHidePayload returns the HidePayload field if non-nil, zero value otherwise.

### GetHidePayloadOk

`func (o *Task) GetHidePayloadOk() (*bool, bool)`

GetHidePayloadOk returns a tuple with the HidePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePayload

`func (o *Task) SetHidePayload(v bool)`

SetHidePayload sets HidePayload field to given value.

### HasHidePayload

`func (o *Task) HasHidePayload() bool`

HasHidePayload returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.


### GetMessages

`func (o *Task) GetMessages() []Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Task) GetMessagesOk() (*[]Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Task) SetMessages(v []Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Task) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.


### GetOem

`func (o *Task) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *Task) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *Task) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *Task) HasOem() bool`

HasOem returns a boolean if a field has been set.

### GetPayload

`func (o *Task) GetPayload() Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Task) GetPayloadOk() (*Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Task) SetPayload(v Payload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Task) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStartTime

`func (o *Task) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Task) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Task) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Task) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTaskMonitor

`func (o *Task) GetTaskMonitor() string`

GetTaskMonitor returns the TaskMonitor field if non-nil, zero value otherwise.

### GetTaskMonitorOk

`func (o *Task) GetTaskMonitorOk() (*string, bool)`

GetTaskMonitorOk returns a tuple with the TaskMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMonitor

`func (o *Task) SetTaskMonitor(v string)`

SetTaskMonitor sets TaskMonitor field to given value.

### HasTaskMonitor

`func (o *Task) HasTaskMonitor() bool`

HasTaskMonitor returns a boolean if a field has been set.

### GetTaskState

`func (o *Task) GetTaskState() TaskState`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *Task) GetTaskStateOk() (*TaskState, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *Task) SetTaskState(v TaskState)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *Task) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### GetTaskStatus

`func (o *Task) GetTaskStatus() Health`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *Task) GetTaskStatusOk() (*Health, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *Task) SetTaskStatus(v Health)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *Task) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


