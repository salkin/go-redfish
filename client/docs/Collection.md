# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | Pointer to **string** | context | [optional] 
**OdataEtag** | Pointer to **string** | etag | [optional] 
**OdataId** | **string** | id | 
**OdataType** | **string** | type | 
**Description** | Pointer to **NullableString** | description | [optional] [readonly] 
**Members** | [**[]IdRef**](IdRef.md) | Contains the members of this collection. | [readonly] 
**MembersodataCount** | Pointer to **int32** | The number of items in a collection. | [optional] [readonly] 
**MembersodataNextLink** | Pointer to **string** | The URI to the resource containing the next set of partial members. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 

## Methods

### NewCollection

`func NewCollection(odataId string, odataType string, members []IdRef, name string, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOdataContext

`func (o *Collection) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *Collection) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *Collection) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.

### HasOdataContext

`func (o *Collection) HasOdataContext() bool`

HasOdataContext returns a boolean if a field has been set.

### GetOdataEtag

`func (o *Collection) GetOdataEtag() string`

GetOdataEtag returns the OdataEtag field if non-nil, zero value otherwise.

### GetOdataEtagOk

`func (o *Collection) GetOdataEtagOk() (*string, bool)`

GetOdataEtagOk returns a tuple with the OdataEtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataEtag

`func (o *Collection) SetOdataEtag(v string)`

SetOdataEtag sets OdataEtag field to given value.

### HasOdataEtag

`func (o *Collection) HasOdataEtag() bool`

HasOdataEtag returns a boolean if a field has been set.

### GetOdataId

`func (o *Collection) GetOdataId() string`

GetOdataId returns the OdataId field if non-nil, zero value otherwise.

### GetOdataIdOk

`func (o *Collection) GetOdataIdOk() (*string, bool)`

GetOdataIdOk returns a tuple with the OdataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataId

`func (o *Collection) SetOdataId(v string)`

SetOdataId sets OdataId field to given value.


### GetOdataType

`func (o *Collection) GetOdataType() string`

GetOdataType returns the OdataType field if non-nil, zero value otherwise.

### GetOdataTypeOk

`func (o *Collection) GetOdataTypeOk() (*string, bool)`

GetOdataTypeOk returns a tuple with the OdataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataType

`func (o *Collection) SetOdataType(v string)`

SetOdataType sets OdataType field to given value.


### GetDescription

`func (o *Collection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Collection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Collection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Collection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Collection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Collection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMembers

`func (o *Collection) GetMembers() []IdRef`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Collection) GetMembersOk() (*[]IdRef, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Collection) SetMembers(v []IdRef)`

SetMembers sets Members field to given value.


### GetMembersodataCount

`func (o *Collection) GetMembersodataCount() int32`

GetMembersodataCount returns the MembersodataCount field if non-nil, zero value otherwise.

### GetMembersodataCountOk

`func (o *Collection) GetMembersodataCountOk() (*int32, bool)`

GetMembersodataCountOk returns a tuple with the MembersodataCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersodataCount

`func (o *Collection) SetMembersodataCount(v int32)`

SetMembersodataCount sets MembersodataCount field to given value.

### HasMembersodataCount

`func (o *Collection) HasMembersodataCount() bool`

HasMembersodataCount returns a boolean if a field has been set.

### GetMembersodataNextLink

`func (o *Collection) GetMembersodataNextLink() string`

GetMembersodataNextLink returns the MembersodataNextLink field if non-nil, zero value otherwise.

### GetMembersodataNextLinkOk

`func (o *Collection) GetMembersodataNextLinkOk() (*string, bool)`

GetMembersodataNextLinkOk returns a tuple with the MembersodataNextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersodataNextLink

`func (o *Collection) SetMembersodataNextLink(v string)`

SetMembersodataNextLink sets MembersodataNextLink field to given value.

### HasMembersodataNextLink

`func (o *Collection) HasMembersodataNextLink() bool`

HasMembersodataNextLink returns a boolean if a field has been set.

### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


