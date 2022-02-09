# RedfishErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageExtendedInfo** | Pointer to [**[]Message**](Message.md) |  | [optional] 
**Code** | **string** |  | [readonly] 
**Message** | **string** |  | [readonly] 

## Methods

### NewRedfishErrorError

`func NewRedfishErrorError(code string, message string, ) *RedfishErrorError`

NewRedfishErrorError instantiates a new RedfishErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedfishErrorErrorWithDefaults

`func NewRedfishErrorErrorWithDefaults() *RedfishErrorError`

NewRedfishErrorErrorWithDefaults instantiates a new RedfishErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageExtendedInfo

`func (o *RedfishErrorError) GetMessageExtendedInfo() []Message`

GetMessageExtendedInfo returns the MessageExtendedInfo field if non-nil, zero value otherwise.

### GetMessageExtendedInfoOk

`func (o *RedfishErrorError) GetMessageExtendedInfoOk() (*[]Message, bool)`

GetMessageExtendedInfoOk returns a tuple with the MessageExtendedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageExtendedInfo

`func (o *RedfishErrorError) SetMessageExtendedInfo(v []Message)`

SetMessageExtendedInfo sets MessageExtendedInfo field to given value.

### HasMessageExtendedInfo

`func (o *RedfishErrorError) HasMessageExtendedInfo() bool`

HasMessageExtendedInfo returns a boolean if a field has been set.

### GetCode

`func (o *RedfishErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RedfishErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RedfishErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RedfishErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RedfishErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RedfishErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


