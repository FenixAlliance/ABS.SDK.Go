# ShareTransferUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**NewShareHolderId** | Pointer to **NullableString** |  | [optional] 
**FormerShareHolderId** | Pointer to **NullableString** |  | [optional] 
**ShareTransferReasonId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShareTransferUpdateDto

`func NewShareTransferUpdateDto() *ShareTransferUpdateDto`

NewShareTransferUpdateDto instantiates a new ShareTransferUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareTransferUpdateDtoWithDefaults

`func NewShareTransferUpdateDtoWithDefaults() *ShareTransferUpdateDto`

NewShareTransferUpdateDtoWithDefaults instantiates a new ShareTransferUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ShareTransferUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShareTransferUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShareTransferUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShareTransferUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ShareTransferUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ShareTransferUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetValue

`func (o *ShareTransferUpdateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ShareTransferUpdateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ShareTransferUpdateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ShareTransferUpdateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNewShareHolderId

`func (o *ShareTransferUpdateDto) GetNewShareHolderId() string`

GetNewShareHolderId returns the NewShareHolderId field if non-nil, zero value otherwise.

### GetNewShareHolderIdOk

`func (o *ShareTransferUpdateDto) GetNewShareHolderIdOk() (*string, bool)`

GetNewShareHolderIdOk returns a tuple with the NewShareHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewShareHolderId

`func (o *ShareTransferUpdateDto) SetNewShareHolderId(v string)`

SetNewShareHolderId sets NewShareHolderId field to given value.

### HasNewShareHolderId

`func (o *ShareTransferUpdateDto) HasNewShareHolderId() bool`

HasNewShareHolderId returns a boolean if a field has been set.

### SetNewShareHolderIdNil

`func (o *ShareTransferUpdateDto) SetNewShareHolderIdNil(b bool)`

 SetNewShareHolderIdNil sets the value for NewShareHolderId to be an explicit nil

### UnsetNewShareHolderId
`func (o *ShareTransferUpdateDto) UnsetNewShareHolderId()`

UnsetNewShareHolderId ensures that no value is present for NewShareHolderId, not even an explicit nil
### GetFormerShareHolderId

`func (o *ShareTransferUpdateDto) GetFormerShareHolderId() string`

GetFormerShareHolderId returns the FormerShareHolderId field if non-nil, zero value otherwise.

### GetFormerShareHolderIdOk

`func (o *ShareTransferUpdateDto) GetFormerShareHolderIdOk() (*string, bool)`

GetFormerShareHolderIdOk returns a tuple with the FormerShareHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormerShareHolderId

`func (o *ShareTransferUpdateDto) SetFormerShareHolderId(v string)`

SetFormerShareHolderId sets FormerShareHolderId field to given value.

### HasFormerShareHolderId

`func (o *ShareTransferUpdateDto) HasFormerShareHolderId() bool`

HasFormerShareHolderId returns a boolean if a field has been set.

### SetFormerShareHolderIdNil

`func (o *ShareTransferUpdateDto) SetFormerShareHolderIdNil(b bool)`

 SetFormerShareHolderIdNil sets the value for FormerShareHolderId to be an explicit nil

### UnsetFormerShareHolderId
`func (o *ShareTransferUpdateDto) UnsetFormerShareHolderId()`

UnsetFormerShareHolderId ensures that no value is present for FormerShareHolderId, not even an explicit nil
### GetShareTransferReasonId

`func (o *ShareTransferUpdateDto) GetShareTransferReasonId() string`

GetShareTransferReasonId returns the ShareTransferReasonId field if non-nil, zero value otherwise.

### GetShareTransferReasonIdOk

`func (o *ShareTransferUpdateDto) GetShareTransferReasonIdOk() (*string, bool)`

GetShareTransferReasonIdOk returns a tuple with the ShareTransferReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareTransferReasonId

`func (o *ShareTransferUpdateDto) SetShareTransferReasonId(v string)`

SetShareTransferReasonId sets ShareTransferReasonId field to given value.

### HasShareTransferReasonId

`func (o *ShareTransferUpdateDto) HasShareTransferReasonId() bool`

HasShareTransferReasonId returns a boolean if a field has been set.

### SetShareTransferReasonIdNil

`func (o *ShareTransferUpdateDto) SetShareTransferReasonIdNil(b bool)`

 SetShareTransferReasonIdNil sets the value for ShareTransferReasonId to be an explicit nil

### UnsetShareTransferReasonId
`func (o *ShareTransferUpdateDto) UnsetShareTransferReasonId()`

UnsetShareTransferReasonId ensures that no value is present for ShareTransferReasonId, not even an explicit nil
### GetEnrollmentId

`func (o *ShareTransferUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ShareTransferUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ShareTransferUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ShareTransferUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ShareTransferUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ShareTransferUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTenantId

`func (o *ShareTransferUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ShareTransferUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ShareTransferUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ShareTransferUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ShareTransferUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ShareTransferUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


