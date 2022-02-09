# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpHeaders** | Pointer to **[]string** | This represents the HTTP headers used in the operation of this Task. | [optional] [readonly] 
**HttpOperation** | Pointer to **string** | The HTTP operation to perform to execute this Task. | [optional] [readonly] 
**JsonBody** | Pointer to **string** | This property contains the JSON payload to use in the execution of this Task. | [optional] [readonly] 
**TargetUri** | Pointer to **string** | The URI of the target for this task. | [optional] [readonly] 

## Methods

### NewPayload

`func NewPayload() *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpHeaders

`func (o *Payload) GetHttpHeaders() []string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *Payload) GetHttpHeadersOk() (*[]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *Payload) SetHttpHeaders(v []string)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *Payload) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetHttpOperation

`func (o *Payload) GetHttpOperation() string`

GetHttpOperation returns the HttpOperation field if non-nil, zero value otherwise.

### GetHttpOperationOk

`func (o *Payload) GetHttpOperationOk() (*string, bool)`

GetHttpOperationOk returns a tuple with the HttpOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOperation

`func (o *Payload) SetHttpOperation(v string)`

SetHttpOperation sets HttpOperation field to given value.

### HasHttpOperation

`func (o *Payload) HasHttpOperation() bool`

HasHttpOperation returns a boolean if a field has been set.

### GetJsonBody

`func (o *Payload) GetJsonBody() string`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *Payload) GetJsonBodyOk() (*string, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *Payload) SetJsonBody(v string)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *Payload) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetTargetUri

`func (o *Payload) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *Payload) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *Payload) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *Payload) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


