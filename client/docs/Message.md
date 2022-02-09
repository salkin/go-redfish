# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] [readonly] 
**MessageArgs** | Pointer to **[]string** |  | [optional] [readonly] 
**MessageId** | **string** |  | [readonly] 
**RelatedProperties** | Pointer to **[]string** |  | [optional] [readonly] 
**Resolution** | Pointer to **string** |  | [optional] [readonly] 
**Severity** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMessage

`func NewMessage(messageId string, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Message) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Message) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Message) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Message) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageArgs

`func (o *Message) GetMessageArgs() []string`

GetMessageArgs returns the MessageArgs field if non-nil, zero value otherwise.

### GetMessageArgsOk

`func (o *Message) GetMessageArgsOk() (*[]string, bool)`

GetMessageArgsOk returns a tuple with the MessageArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageArgs

`func (o *Message) SetMessageArgs(v []string)`

SetMessageArgs sets MessageArgs field to given value.

### HasMessageArgs

`func (o *Message) HasMessageArgs() bool`

HasMessageArgs returns a boolean if a field has been set.

### GetMessageId

`func (o *Message) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Message) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Message) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetRelatedProperties

`func (o *Message) GetRelatedProperties() []string`

GetRelatedProperties returns the RelatedProperties field if non-nil, zero value otherwise.

### GetRelatedPropertiesOk

`func (o *Message) GetRelatedPropertiesOk() (*[]string, bool)`

GetRelatedPropertiesOk returns a tuple with the RelatedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProperties

`func (o *Message) SetRelatedProperties(v []string)`

SetRelatedProperties sets RelatedProperties field to given value.

### HasRelatedProperties

`func (o *Message) HasRelatedProperties() bool`

HasRelatedProperties returns a boolean if a field has been set.

### GetResolution

`func (o *Message) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Message) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Message) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Message) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeverity

`func (o *Message) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Message) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Message) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Message) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


