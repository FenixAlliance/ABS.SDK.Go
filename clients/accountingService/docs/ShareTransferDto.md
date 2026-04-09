# ShareTransferDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**NewShareHolderId** | Pointer to **NullableString** |  | [optional] 
**FormerShareHolderId** | Pointer to **NullableString** |  | [optional] 
**ShareTransferReasonId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShareTransferDto

`func NewShareTransferDto() *ShareTransferDto`

NewShareTransferDto instantiates a new ShareTransferDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareTransferDtoWithDefaults

`func NewShareTransferDtoWithDefaults() *ShareTransferDto`

NewShareTransferDtoWithDefaults instantiates a new ShareTransferDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareTransferDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareTransferDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareTransferDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShareTransferDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ShareTransferDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ShareTransferDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ShareTransferDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShareTransferDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShareTransferDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ShareTransferDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ShareTransferDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ShareTransferDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDescription

`func (o *ShareTransferDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShareTransferDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShareTransferDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShareTransferDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShareTransferDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShareTransferDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetValue

`func (o *ShareTransferDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShareTransferDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShareTransferDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShareTransferDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNewShareHolderId

`func (o *ShareTransferDto) GetNewShareHolderId() string`

GetNewShareHolderId returns the NewShareHolderId field if non-nil, zero value otherwise.

### GetNewShareHolderIdOk

`func (o *ShareTransferDto) GetNewShareHolderIdOk() (*string, bool)`

GetNewShareHolderIdOk returns a tuple with the NewShareHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewShareHolderId

`func (o *ShareTransferDto) SetNewShareHolderId(v string)`

SetNewShareHolderId sets NewShareHolderId field to given value.

### HasNewShareHolderId

`func (o *ShareTransferDto) HasNewShareHolderId() bool`

HasNewShareHolderId returns a boolean if a field has been set.

### SetNewShareHolderIdNil

`func (o *ShareTransferDto) SetNewShareHolderIdNil(b bool)`

 SetNewShareHolderIdNil sets the value for NewShareHolderId to be an explicit nil

### UnsetNewShareHolderId
`func (o *ShareTransferDto) UnsetNewShareHolderId()`

UnsetNewShareHolderId ensures that no value is present for NewShareHolderId, not even an explicit nil
### GetFormerShareHolderId

`func (o *ShareTransferDto) GetFormerShareHolderId() string`

GetFormerShareHolderId returns the FormerShareHolderId field if non-nil, zero value otherwise.

### GetFormerShareHolderIdOk

`func (o *ShareTransferDto) GetFormerShareHolderIdOk() (*string, bool)`

GetFormerShareHolderIdOk returns a tuple with the FormerShareHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormerShareHolderId

`func (o *ShareTransferDto) SetFormerShareHolderId(v string)`

SetFormerShareHolderId sets FormerShareHolderId field to given value.

### HasFormerShareHolderId

`func (o *ShareTransferDto) HasFormerShareHolderId() bool`

HasFormerShareHolderId returns a boolean if a field has been set.

### SetFormerShareHolderIdNil

`func (o *ShareTransferDto) SetFormerShareHolderIdNil(b bool)`

 SetFormerShareHolderIdNil sets the value for FormerShareHolderId to be an explicit nil

### UnsetFormerShareHolderId
`func (o *ShareTransferDto) UnsetFormerShareHolderId()`

UnsetFormerShareHolderId ensures that no value is present for FormerShareHolderId, not even an explicit nil
### GetShareTransferReasonId

`func (o *ShareTransferDto) GetShareTransferReasonId() string`

GetShareTransferReasonId returns the ShareTransferReasonId field if non-nil, zero value otherwise.

### GetShareTransferReasonIdOk

`func (o *ShareTransferDto) GetShareTransferReasonIdOk() (*string, bool)`

GetShareTransferReasonIdOk returns a tuple with the ShareTransferReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareTransferReasonId

`func (o *ShareTransferDto) SetShareTransferReasonId(v string)`

SetShareTransferReasonId sets ShareTransferReasonId field to given value.

### HasShareTransferReasonId

`func (o *ShareTransferDto) HasShareTransferReasonId() bool`

HasShareTransferReasonId returns a boolean if a field has been set.

### SetShareTransferReasonIdNil

`func (o *ShareTransferDto) SetShareTransferReasonIdNil(b bool)`

 SetShareTransferReasonIdNil sets the value for ShareTransferReasonId to be an explicit nil

### UnsetShareTransferReasonId
`func (o *ShareTransferDto) UnsetShareTransferReasonId()`

UnsetShareTransferReasonId ensures that no value is present for ShareTransferReasonId, not even an explicit nil
### GetEnrollmentId

`func (o *ShareTransferDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ShareTransferDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ShareTransferDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ShareTransferDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ShareTransferDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ShareTransferDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTenantId

`func (o *ShareTransferDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShareTransferDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShareTransferDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ShareTransferDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ShareTransferDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ShareTransferDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


