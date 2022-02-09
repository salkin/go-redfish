# InsertMediaRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** |  | 
**Inserted** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**TransferMethod** | Pointer to [**TransferMethod**](TransferMethod.md) |  | [optional] 
**TransferProtocolType** | Pointer to [**TransferProtocolType**](TransferProtocolType.md) |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**WriteProtected** | Pointer to **bool** |  | [optional] 

## Methods

### NewInsertMediaRequestBody

`func NewInsertMediaRequestBody(image string, ) *InsertMediaRequestBody`

NewInsertMediaRequestBody instantiates a new InsertMediaRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertMediaRequestBodyWithDefaults

`func NewInsertMediaRequestBodyWithDefaults() *InsertMediaRequestBody`

NewInsertMediaRequestBodyWithDefaults instantiates a new InsertMediaRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *InsertMediaRequestBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InsertMediaRequestBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InsertMediaRequestBody) SetImage(v string)`

SetImage sets Image field to given value.


### GetInserted

`func (o *InsertMediaRequestBody) GetInserted() bool`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *InsertMediaRequestBody) GetInsertedOk() (*bool, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *InsertMediaRequestBody) SetInserted(v bool)`

SetInserted sets Inserted field to given value.

### HasInserted

`func (o *InsertMediaRequestBody) HasInserted() bool`

HasInserted returns a boolean if a field has been set.

### GetPassword

`func (o *InsertMediaRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InsertMediaRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InsertMediaRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InsertMediaRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTransferMethod

`func (o *InsertMediaRequestBody) GetTransferMethod() TransferMethod`

GetTransferMethod returns the TransferMethod field if non-nil, zero value otherwise.

### GetTransferMethodOk

`func (o *InsertMediaRequestBody) GetTransferMethodOk() (*TransferMethod, bool)`

GetTransferMethodOk returns a tuple with the TransferMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMethod

`func (o *InsertMediaRequestBody) SetTransferMethod(v TransferMethod)`

SetTransferMethod sets TransferMethod field to given value.

### HasTransferMethod

`func (o *InsertMediaRequestBody) HasTransferMethod() bool`

HasTransferMethod returns a boolean if a field has been set.

### GetTransferProtocolType

`func (o *InsertMediaRequestBody) GetTransferProtocolType() TransferProtocolType`

GetTransferProtocolType returns the TransferProtocolType field if non-nil, zero value otherwise.

### GetTransferProtocolTypeOk

`func (o *InsertMediaRequestBody) GetTransferProtocolTypeOk() (*TransferProtocolType, bool)`

GetTransferProtocolTypeOk returns a tuple with the TransferProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferProtocolType

`func (o *InsertMediaRequestBody) SetTransferProtocolType(v TransferProtocolType)`

SetTransferProtocolType sets TransferProtocolType field to given value.

### HasTransferProtocolType

`func (o *InsertMediaRequestBody) HasTransferProtocolType() bool`

HasTransferProtocolType returns a boolean if a field has been set.

### GetUserName

`func (o *InsertMediaRequestBody) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *InsertMediaRequestBody) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *InsertMediaRequestBody) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *InsertMediaRequestBody) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWriteProtected

`func (o *InsertMediaRequestBody) GetWriteProtected() bool`

GetWriteProtected returns the WriteProtected field if non-nil, zero value otherwise.

### GetWriteProtectedOk

`func (o *InsertMediaRequestBody) GetWriteProtectedOk() (*bool, bool)`

GetWriteProtectedOk returns a tuple with the WriteProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteProtected

`func (o *InsertMediaRequestBody) SetWriteProtected(v bool)`

SetWriteProtected sets WriteProtected field to given value.

### HasWriteProtected

`func (o *InsertMediaRequestBody) HasWriteProtected() bool`

HasWriteProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


