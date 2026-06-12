# SupportTicketUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SupportTicketStatus** | Pointer to **string** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**SupportTicketTypeId** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementId** | Pointer to **NullableString** |  | [optional] 
**SupportPriorityId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportTicketUpdateDto

`func NewSupportTicketUpdateDto() *SupportTicketUpdateDto`

NewSupportTicketUpdateDto instantiates a new SupportTicketUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketUpdateDtoWithDefaults

`func NewSupportTicketUpdateDtoWithDefaults() *SupportTicketUpdateDto`

NewSupportTicketUpdateDtoWithDefaults instantiates a new SupportTicketUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SupportTicketUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportTicketUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportTicketUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportTicketUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportTicketUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportTicketUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportTicketUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicketUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicketUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicketUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportTicketUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportTicketUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupportTicketStatus

`func (o *SupportTicketUpdateDto) GetSupportTicketStatus() string`

GetSupportTicketStatus returns the SupportTicketStatus field if non-nil, zero value otherwise.

### GetSupportTicketStatusOk

`func (o *SupportTicketUpdateDto) GetSupportTicketStatusOk() (*string, bool)`

GetSupportTicketStatusOk returns a tuple with the SupportTicketStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketStatus

`func (o *SupportTicketUpdateDto) SetSupportTicketStatus(v string)`

SetSupportTicketStatus sets SupportTicketStatus field to given value.

### HasSupportTicketStatus

`func (o *SupportTicketUpdateDto) HasSupportTicketStatus() bool`

HasSupportTicketStatus returns a boolean if a field has been set.

### GetContactId

`func (o *SupportTicketUpdateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SupportTicketUpdateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SupportTicketUpdateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SupportTicketUpdateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SupportTicketUpdateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SupportTicketUpdateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetSupportTicketTypeId

`func (o *SupportTicketUpdateDto) GetSupportTicketTypeId() string`

GetSupportTicketTypeId returns the SupportTicketTypeId field if non-nil, zero value otherwise.

### GetSupportTicketTypeIdOk

`func (o *SupportTicketUpdateDto) GetSupportTicketTypeIdOk() (*string, bool)`

GetSupportTicketTypeIdOk returns a tuple with the SupportTicketTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketTypeId

`func (o *SupportTicketUpdateDto) SetSupportTicketTypeId(v string)`

SetSupportTicketTypeId sets SupportTicketTypeId field to given value.

### HasSupportTicketTypeId

`func (o *SupportTicketUpdateDto) HasSupportTicketTypeId() bool`

HasSupportTicketTypeId returns a boolean if a field has been set.

### SetSupportTicketTypeIdNil

`func (o *SupportTicketUpdateDto) SetSupportTicketTypeIdNil(b bool)`

 SetSupportTicketTypeIdNil sets the value for SupportTicketTypeId to be an explicit nil

### UnsetSupportTicketTypeId
`func (o *SupportTicketUpdateDto) UnsetSupportTicketTypeId()`

UnsetSupportTicketTypeId ensures that no value is present for SupportTicketTypeId, not even an explicit nil
### GetSupportEntitlementId

`func (o *SupportTicketUpdateDto) GetSupportEntitlementId() string`

GetSupportEntitlementId returns the SupportEntitlementId field if non-nil, zero value otherwise.

### GetSupportEntitlementIdOk

`func (o *SupportTicketUpdateDto) GetSupportEntitlementIdOk() (*string, bool)`

GetSupportEntitlementIdOk returns a tuple with the SupportEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementId

`func (o *SupportTicketUpdateDto) SetSupportEntitlementId(v string)`

SetSupportEntitlementId sets SupportEntitlementId field to given value.

### HasSupportEntitlementId

`func (o *SupportTicketUpdateDto) HasSupportEntitlementId() bool`

HasSupportEntitlementId returns a boolean if a field has been set.

### SetSupportEntitlementIdNil

`func (o *SupportTicketUpdateDto) SetSupportEntitlementIdNil(b bool)`

 SetSupportEntitlementIdNil sets the value for SupportEntitlementId to be an explicit nil

### UnsetSupportEntitlementId
`func (o *SupportTicketUpdateDto) UnsetSupportEntitlementId()`

UnsetSupportEntitlementId ensures that no value is present for SupportEntitlementId, not even an explicit nil
### GetSupportPriorityId

`func (o *SupportTicketUpdateDto) GetSupportPriorityId() string`

GetSupportPriorityId returns the SupportPriorityId field if non-nil, zero value otherwise.

### GetSupportPriorityIdOk

`func (o *SupportTicketUpdateDto) GetSupportPriorityIdOk() (*string, bool)`

GetSupportPriorityIdOk returns a tuple with the SupportPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPriorityId

`func (o *SupportTicketUpdateDto) SetSupportPriorityId(v string)`

SetSupportPriorityId sets SupportPriorityId field to given value.

### HasSupportPriorityId

`func (o *SupportTicketUpdateDto) HasSupportPriorityId() bool`

HasSupportPriorityId returns a boolean if a field has been set.

### SetSupportPriorityIdNil

`func (o *SupportTicketUpdateDto) SetSupportPriorityIdNil(b bool)`

 SetSupportPriorityIdNil sets the value for SupportPriorityId to be an explicit nil

### UnsetSupportPriorityId
`func (o *SupportTicketUpdateDto) UnsetSupportPriorityId()`

UnsetSupportPriorityId ensures that no value is present for SupportPriorityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


