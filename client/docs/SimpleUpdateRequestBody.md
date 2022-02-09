# SimpleUpdateRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageURI** | **string** |  | 
**Targets** | Pointer to **[]string** |  | [optional] 
**TransferProtocolType** | Pointer to [**TransferProtocolType**](TransferProtocolType.md) |  | [optional] 

## Methods

### NewSimpleUpdateRequestBody

`func NewSimpleUpdateRequestBody(imageURI string, ) *SimpleUpdateRequestBody`

NewSimpleUpdateRequestBody instantiates a new SimpleUpdateRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleUpdateRequestBodyWithDefaults

`func NewSimpleUpdateRequestBodyWithDefaults() *SimpleUpdateRequestBody`

NewSimpleUpdateRequestBodyWithDefaults instantiates a new SimpleUpdateRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageURI

`func (o *SimpleUpdateRequestBody) GetImageURI() string`

GetImageURI returns the ImageURI field if non-nil, zero value otherwise.

### GetImageURIOk

`func (o *SimpleUpdateRequestBody) GetImageURIOk() (*string, bool)`

GetImageURIOk returns a tuple with the ImageURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageURI

`func (o *SimpleUpdateRequestBody) SetImageURI(v string)`

SetImageURI sets ImageURI field to given value.


### GetTargets

`func (o *SimpleUpdateRequestBody) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SimpleUpdateRequestBody) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SimpleUpdateRequestBody) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SimpleUpdateRequestBody) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTransferProtocolType

`func (o *SimpleUpdateRequestBody) GetTransferProtocolType() TransferProtocolType`

GetTransferProtocolType returns the TransferProtocolType field if non-nil, zero value otherwise.

### GetTransferProtocolTypeOk

`func (o *SimpleUpdateRequestBody) GetTransferProtocolTypeOk() (*TransferProtocolType, bool)`

GetTransferProtocolTypeOk returns a tuple with the TransferProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferProtocolType

`func (o *SimpleUpdateRequestBody) SetTransferProtocolType(v TransferProtocolType)`

SetTransferProtocolType sets TransferProtocolType field to given value.

### HasTransferProtocolType

`func (o *SimpleUpdateRequestBody) HasTransferProtocolType() bool`

HasTransferProtocolType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


