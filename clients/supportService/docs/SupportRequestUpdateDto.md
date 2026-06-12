# SupportRequestUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**ApprovedTimestamp** | Pointer to **time.Time** |  | [optional] 
**SupportEntitlementId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportRequestUpdateDto

`func NewSupportRequestUpdateDto() *SupportRequestUpdateDto`

NewSupportRequestUpdateDto instantiates a new SupportRequestUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestUpdateDtoWithDefaults

`func NewSupportRequestUpdateDtoWithDefaults() *SupportRequestUpdateDto`

NewSupportRequestUpdateDtoWithDefaults instantiates a new SupportRequestUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SupportRequestUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportRequestUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportRequestUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportRequestUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportRequestUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportRequestUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportRequestUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportRequestUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportRequestUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportRequestUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportRequestUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportRequestUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApproved

`func (o *SupportRequestUpdateDto) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *SupportRequestUpdateDto) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *SupportRequestUpdateDto) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *SupportRequestUpdateDto) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetApprovedTimestamp

`func (o *SupportRequestUpdateDto) GetApprovedTimestamp() time.Time`

GetApprovedTimestamp returns the ApprovedTimestamp field if non-nil, zero value otherwise.

### GetApprovedTimestampOk

`func (o *SupportRequestUpdateDto) GetApprovedTimestampOk() (*time.Time, bool)`

GetApprovedTimestampOk returns a tuple with the ApprovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTimestamp

`func (o *SupportRequestUpdateDto) SetApprovedTimestamp(v time.Time)`

SetApprovedTimestamp sets ApprovedTimestamp field to given value.

### HasApprovedTimestamp

`func (o *SupportRequestUpdateDto) HasApprovedTimestamp() bool`

HasApprovedTimestamp returns a boolean if a field has been set.

### GetSupportEntitlementId

`func (o *SupportRequestUpdateDto) GetSupportEntitlementId() string`

GetSupportEntitlementId returns the SupportEntitlementId field if non-nil, zero value otherwise.

### GetSupportEntitlementIdOk

`func (o *SupportRequestUpdateDto) GetSupportEntitlementIdOk() (*string, bool)`

GetSupportEntitlementIdOk returns a tuple with the SupportEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementId

`func (o *SupportRequestUpdateDto) SetSupportEntitlementId(v string)`

SetSupportEntitlementId sets SupportEntitlementId field to given value.

### HasSupportEntitlementId

`func (o *SupportRequestUpdateDto) HasSupportEntitlementId() bool`

HasSupportEntitlementId returns a boolean if a field has been set.

### SetSupportEntitlementIdNil

`func (o *SupportRequestUpdateDto) SetSupportEntitlementIdNil(b bool)`

 SetSupportEntitlementIdNil sets the value for SupportEntitlementId to be an explicit nil

### UnsetSupportEntitlementId
`func (o *SupportRequestUpdateDto) UnsetSupportEntitlementId()`

UnsetSupportEntitlementId ensures that no value is present for SupportEntitlementId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


