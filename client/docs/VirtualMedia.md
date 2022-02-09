# VirtualMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**OdataId** | **string** | The unique identifier for a resource. | [readonly] 
**OdataContext** | Pointer to **string** | The OData description of a payload. | [optional] [readonly] 
**RedfishCopyright** | Pointer to **string** | redfish copyright | [optional] 
**Description** | Pointer to **NullableString** | description | [optional] [readonly] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **NullableString** |  | [optional] [readonly] 
**Inserted** | Pointer to **NullableBool** |  | [optional] 
**ConnectedVia** | Pointer to [**ConnectedVia**](ConnectedVia.md) |  | [optional] 
**MediaTypes** | Pointer to **[]string** |  | [optional] [readonly] 
**WriteProtected** | Pointer to **NullableBool** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**TransferMethod** | Pointer to [**TransferMethod**](TransferMethod.md) |  | [optional] 
**TransferProtocolType** | Pointer to [**TransferProtocolType**](TransferProtocolType.md) |  | [optional] 
**Actions** | Pointer to [**VirtualMediaActions**](VirtualMediaActions.md) |  | [optional] 

## Methods

### NewVirtualMedia

`func NewVirtualMedia(name string, odataType string, odataId string, ) *VirtualMedia`

NewVirtualMedia instantiates a new VirtualMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMediaWithDefaults

`func NewVirtualMediaWithDefaults() *VirtualMedia`

NewVirtualMediaWithDefaults instantiates a new VirtualMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualMedia) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualMedia) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualMedia) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualMedia) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VirtualMedia) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMedia) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMedia) SetName(v string)`

SetName sets Name field to given value.


### GetOdataType

`func (o *VirtualMedia) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *VirtualMedia) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *VirtualMedia) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetOdataId

`func (o *VirtualMedia) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *VirtualMedia) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *VirtualMedia) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataContext

`func (o *VirtualMedia) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *VirtualMedia) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *VirtualMedia) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *VirtualMedia) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.

### GetRedfishCopyright

`func (o *VirtualMedia) GetRedfishCopyright() string`

GetRedfishCopyright returns the RedfishCopyright field if non-nil, zero value otherwise.

### GetRedfishCopyrightOk

`func (o *VirtualMedia) GetRedfishCopyrightOk() (*string, bool)`

GetRedfishCopyrightOk returns a tuple with the RedfishCopyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishCopyright

`func (o *VirtualMedia) SetRedfishCopyright(v string)`

SetRedfishCopyright sets RedfishCopyright field to given value.

### HasRedfishCopyright

`func (o *VirtualMedia) HasRedfishCopyright() bool`

HasRedfishCopyright returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualMedia) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualMedia) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualMedia) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualMedia) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VirtualMedia) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VirtualMedia) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImage

`func (o *VirtualMedia) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VirtualMedia) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VirtualMedia) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VirtualMedia) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *VirtualMedia) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *VirtualMedia) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetImageName

`func (o *VirtualMedia) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *VirtualMedia) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *VirtualMedia) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *VirtualMedia) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *VirtualMedia) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *VirtualMedia) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetInserted

`func (o *VirtualMedia) GetInserted() bool`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *VirtualMedia) GetInsertedOk() (*bool, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *VirtualMedia) SetInserted(v bool)`

SetInserted sets Inserted field to given value.

### HasInserted

`func (o *VirtualMedia) HasInserted() bool`

HasInserted returns a boolean if a field has been set.

### SetInsertedNil

`func (o *VirtualMedia) SetInsertedNil(b bool)`

 SetInsertedNil sets the value for Inserted to be an explicit nil

### UnsetInserted
`func (o *VirtualMedia) UnsetInserted()`

UnsetInserted ensures that no value is present for Inserted, not even an explicit nil
### GetConnectedVia

`func (o *VirtualMedia) GetConnectedVia() ConnectedVia`

GetConnectedVia returns the ConnectedVia field if non-nil, zero value otherwise.

### GetConnectedViaOk

`func (o *VirtualMedia) GetConnectedViaOk() (*ConnectedVia, bool)`

GetConnectedViaOk returns a tuple with the ConnectedVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVia

`func (o *VirtualMedia) SetConnectedVia(v ConnectedVia)`

SetConnectedVia sets ConnectedVia field to given value.

### HasConnectedVia

`func (o *VirtualMedia) HasConnectedVia() bool`

HasConnectedVia returns a boolean if a field has been set.

### GetMediaTypes

`func (o *VirtualMedia) GetMediaTypes() []string`

GetMediaTypes returns the MediaTypes field if non-nil, zero value otherwise.

### GetMediaTypesOk

`func (o *VirtualMedia) GetMediaTypesOk() (*[]string, bool)`

GetMediaTypesOk returns a tuple with the MediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypes

`func (o *VirtualMedia) SetMediaTypes(v []string)`

SetMediaTypes sets MediaTypes field to given value.

### HasMediaTypes

`func (o *VirtualMedia) HasMediaTypes() bool`

HasMediaTypes returns a boolean if a field has been set.

### GetWriteProtected

`func (o *VirtualMedia) GetWriteProtected() bool`

GetWriteProtected returns the WriteProtected field if non-nil, zero value otherwise.

### GetWriteProtectedOk

`func (o *VirtualMedia) GetWriteProtectedOk() (*bool, bool)`

GetWriteProtectedOk returns a tuple with the WriteProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteProtected

`func (o *VirtualMedia) SetWriteProtected(v bool)`

SetWriteProtected sets WriteProtected field to given value.

### HasWriteProtected

`func (o *VirtualMedia) HasWriteProtected() bool`

HasWriteProtected returns a boolean if a field has been set.

### SetWriteProtectedNil

`func (o *VirtualMedia) SetWriteProtectedNil(b bool)`

 SetWriteProtectedNil sets the value for WriteProtected to be an explicit nil

### UnsetWriteProtected
`func (o *VirtualMedia) UnsetWriteProtected()`

UnsetWriteProtected ensures that no value is present for WriteProtected, not even an explicit nil
### GetUserName

`func (o *VirtualMedia) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VirtualMedia) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VirtualMedia) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VirtualMedia) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *VirtualMedia) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *VirtualMedia) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPassword

`func (o *VirtualMedia) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VirtualMedia) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VirtualMedia) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VirtualMedia) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *VirtualMedia) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *VirtualMedia) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTransferMethod

`func (o *VirtualMedia) GetTransferMethod() TransferMethod`

GetTransferMethod returns the TransferMethod field if non-nil, zero value otherwise.

### GetTransferMethodOk

`func (o *VirtualMedia) GetTransferMethodOk() (*TransferMethod, bool)`

GetTransferMethodOk returns a tuple with the TransferMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMethod

`func (o *VirtualMedia) SetTransferMethod(v TransferMethod)`

SetTransferMethod sets TransferMethod field to given value.

### HasTransferMethod

`func (o *VirtualMedia) HasTransferMethod() bool`

HasTransferMethod returns a boolean if a field has been set.

### GetTransferProtocolType

`func (o *VirtualMedia) GetTransferProtocolType() TransferProtocolType`

GetTransferProtocolType returns the TransferProtocolType field if non-nil, zero value otherwise.

### GetTransferProtocolTypeOk

`func (o *VirtualMedia) GetTransferProtocolTypeOk() (*TransferProtocolType, bool)`

GetTransferProtocolTypeOk returns a tuple with the TransferProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferProtocolType

`func (o *VirtualMedia) SetTransferProtocolType(v TransferProtocolType)`

SetTransferProtocolType sets TransferProtocolType field to given value.

### HasTransferProtocolType

`func (o *VirtualMedia) HasTransferProtocolType() bool`

HasTransferProtocolType returns a boolean if a field has been set.

### GetActions

`func (o *VirtualMedia) GetActions() VirtualMediaActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *VirtualMedia) GetActionsOk() (*VirtualMediaActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *VirtualMedia) SetActions(v VirtualMediaActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *VirtualMedia) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


