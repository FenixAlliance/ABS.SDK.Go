# ShareTransferCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**NewShareHolderId** | Pointer to **NullableString** |  | [optional] 
**FormerShareHolderId** | Pointer to **NullableString** |  | [optional] 
**ShareTransferReasonId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShareTransferCreateDto

`func NewShareTransferCreateDto() *ShareTransferCreateDto`

NewShareTransferCreateDto instantiates a new ShareTransferCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareTransferCreateDtoWithDefaults

`func NewShareTransferCreateDtoWithDefaults() *ShareTransferCreateDto`

NewShareTransferCreateDtoWithDefaults instantiates a new ShareTransferCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareTransferCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareTransferCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareTransferCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShareTransferCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ShareTransferCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShareTransferCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShareTransferCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShareTransferCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *ShareTransferCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShareTransferCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShareTransferCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShareTransferCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShareTransferCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShareTransferCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetValue

`func (o *ShareTransferCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShareTransferCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShareTransferCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShareTransferCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNewShareHolderId

`func (o *ShareTransferCreateDto) GetNewShareHolderId() string`

GetNewShareHolderId returns the NewShareHolderId field if non-nil, zero value otherwise.

### GetNewShareHolderIdOk

`func (o *ShareTransferCreateDto) GetNewShareHolderIdOk() (*string, bool)`

GetNewShareHolderIdOk returns a tuple with the NewShareHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewShareHolderId

`func (o *ShareTransferCreateDto) SetNewShareHolderId(v string)`

SetNewShareHolderId sets NewShareHolderId field to given value.

### HasNewShareHolderId

`func (o *ShareTransferCreateDto) HasNewShareHolderId() bool`

HasNewShareHolderId returns a boolean if a field has been set.

### SetNewShareHolderIdNil

`func (o *ShareTransferCreateDto) SetNewShareHolderIdNil(b bool)`

 SetNewShareHolderIdNil sets the value for NewShareHolderId to be an explicit nil

### UnsetNewShareHolderId
`func (o *ShareTransferCreateDto) UnsetNewShareHolderId()`

UnsetNewShareHolderId ensures that no value is present for NewShareHolderId, not even an explicit nil
### GetFormerShareHolderId

`func (o *ShareTransferCreateDto) GetFormerShareHolderId() string`

GetFormerShareHolderId returns the FormerShareHolderId field if non-nil, zero value otherwise.

### GetFormerShareHolderIdOk

`func (o *ShareTransferCreateDto) GetFormerShareHolderIdOk() (*string, bool)`

GetFormerShareHolderIdOk returns a tuple with the FormerShareHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormerShareHolderId

`func (o *ShareTransferCreateDto) SetFormerShareHolderId(v string)`

SetFormerShareHolderId sets FormerShareHolderId field to given value.

### HasFormerShareHolderId

`func (o *ShareTransferCreateDto) HasFormerShareHolderId() bool`

HasFormerShareHolderId returns a boolean if a field has been set.

### SetFormerShareHolderIdNil

`func (o *ShareTransferCreateDto) SetFormerShareHolderIdNil(b bool)`

 SetFormerShareHolderIdNil sets the value for FormerShareHolderId to be an explicit nil

### UnsetFormerShareHolderId
`func (o *ShareTransferCreateDto) UnsetFormerShareHolderId()`

UnsetFormerShareHolderId ensures that no value is present for FormerShareHolderId, not even an explicit nil
### GetShareTransferReasonId

`func (o *ShareTransferCreateDto) GetShareTransferReasonId() string`

GetShareTransferReasonId returns the ShareTransferReasonId field if non-nil, zero value otherwise.

### GetShareTransferReasonIdOk

`func (o *ShareTransferCreateDto) GetShareTransferReasonIdOk() (*string, bool)`

GetShareTransferReasonIdOk returns a tuple with the ShareTransferReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareTransferReasonId

`func (o *ShareTransferCreateDto) SetShareTransferReasonId(v string)`

SetShareTransferReasonId sets ShareTransferReasonId field to given value.

### HasShareTransferReasonId

`func (o *ShareTransferCreateDto) HasShareTransferReasonId() bool`

HasShareTransferReasonId returns a boolean if a field has been set.

### SetShareTransferReasonIdNil

`func (o *ShareTransferCreateDto) SetShareTransferReasonIdNil(b bool)`

 SetShareTransferReasonIdNil sets the value for ShareTransferReasonId to be an explicit nil

### UnsetShareTransferReasonId
`func (o *ShareTransferCreateDto) UnsetShareTransferReasonId()`

UnsetShareTransferReasonId ensures that no value is present for ShareTransferReasonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


