# SupportTicketCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SupportTicketStatus** | Pointer to **string** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**SupportTicketTypeId** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementId** | Pointer to **NullableString** |  | [optional] 
**SupportPriorityId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportTicketCreateDto

`func NewSupportTicketCreateDto() *SupportTicketCreateDto`

NewSupportTicketCreateDto instantiates a new SupportTicketCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketCreateDtoWithDefaults

`func NewSupportTicketCreateDtoWithDefaults() *SupportTicketCreateDto`

NewSupportTicketCreateDtoWithDefaults instantiates a new SupportTicketCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicketCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SupportTicketCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportTicketCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportTicketCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportTicketCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *SupportTicketCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportTicketCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportTicketCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportTicketCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportTicketCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportTicketCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportTicketCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicketCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicketCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicketCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportTicketCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportTicketCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupportTicketStatus

`func (o *SupportTicketCreateDto) GetSupportTicketStatus() string`

GetSupportTicketStatus returns the SupportTicketStatus field if non-nil, zero value otherwise.

### GetSupportTicketStatusOk

`func (o *SupportTicketCreateDto) GetSupportTicketStatusOk() (*string, bool)`

GetSupportTicketStatusOk returns a tuple with the SupportTicketStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketStatus

`func (o *SupportTicketCreateDto) SetSupportTicketStatus(v string)`

SetSupportTicketStatus sets SupportTicketStatus field to given value.

### HasSupportTicketStatus

`func (o *SupportTicketCreateDto) HasSupportTicketStatus() bool`

HasSupportTicketStatus returns a boolean if a field has been set.

### GetContactId

`func (o *SupportTicketCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SupportTicketCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SupportTicketCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SupportTicketCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SupportTicketCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SupportTicketCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetSupportTicketTypeId

`func (o *SupportTicketCreateDto) GetSupportTicketTypeId() string`

GetSupportTicketTypeId returns the SupportTicketTypeId field if non-nil, zero value otherwise.

### GetSupportTicketTypeIdOk

`func (o *SupportTicketCreateDto) GetSupportTicketTypeIdOk() (*string, bool)`

GetSupportTicketTypeIdOk returns a tuple with the SupportTicketTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketTypeId

`func (o *SupportTicketCreateDto) SetSupportTicketTypeId(v string)`

SetSupportTicketTypeId sets SupportTicketTypeId field to given value.

### HasSupportTicketTypeId

`func (o *SupportTicketCreateDto) HasSupportTicketTypeId() bool`

HasSupportTicketTypeId returns a boolean if a field has been set.

### SetSupportTicketTypeIdNil

`func (o *SupportTicketCreateDto) SetSupportTicketTypeIdNil(b bool)`

 SetSupportTicketTypeIdNil sets the value for SupportTicketTypeId to be an explicit nil

### UnsetSupportTicketTypeId
`func (o *SupportTicketCreateDto) UnsetSupportTicketTypeId()`

UnsetSupportTicketTypeId ensures that no value is present for SupportTicketTypeId, not even an explicit nil
### GetSupportEntitlementId

`func (o *SupportTicketCreateDto) GetSupportEntitlementId() string`

GetSupportEntitlementId returns the SupportEntitlementId field if non-nil, zero value otherwise.

### GetSupportEntitlementIdOk

`func (o *SupportTicketCreateDto) GetSupportEntitlementIdOk() (*string, bool)`

GetSupportEntitlementIdOk returns a tuple with the SupportEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementId

`func (o *SupportTicketCreateDto) SetSupportEntitlementId(v string)`

SetSupportEntitlementId sets SupportEntitlementId field to given value.

### HasSupportEntitlementId

`func (o *SupportTicketCreateDto) HasSupportEntitlementId() bool`

HasSupportEntitlementId returns a boolean if a field has been set.

### SetSupportEntitlementIdNil

`func (o *SupportTicketCreateDto) SetSupportEntitlementIdNil(b bool)`

 SetSupportEntitlementIdNil sets the value for SupportEntitlementId to be an explicit nil

### UnsetSupportEntitlementId
`func (o *SupportTicketCreateDto) UnsetSupportEntitlementId()`

UnsetSupportEntitlementId ensures that no value is present for SupportEntitlementId, not even an explicit nil
### GetSupportPriorityId

`func (o *SupportTicketCreateDto) GetSupportPriorityId() string`

GetSupportPriorityId returns the SupportPriorityId field if non-nil, zero value otherwise.

### GetSupportPriorityIdOk

`func (o *SupportTicketCreateDto) GetSupportPriorityIdOk() (*string, bool)`

GetSupportPriorityIdOk returns a tuple with the SupportPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPriorityId

`func (o *SupportTicketCreateDto) SetSupportPriorityId(v string)`

SetSupportPriorityId sets SupportPriorityId field to given value.

### HasSupportPriorityId

`func (o *SupportTicketCreateDto) HasSupportPriorityId() bool`

HasSupportPriorityId returns a boolean if a field has been set.

### SetSupportPriorityIdNil

`func (o *SupportTicketCreateDto) SetSupportPriorityIdNil(b bool)`

 SetSupportPriorityIdNil sets the value for SupportPriorityId to be an explicit nil

### UnsetSupportPriorityId
`func (o *SupportTicketCreateDto) UnsetSupportPriorityId()`

UnsetSupportPriorityId ensures that no value is present for SupportPriorityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


