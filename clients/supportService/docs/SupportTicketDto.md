# SupportTicketDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SupportTicketStatus** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SupportTicketTypeId** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementId** | Pointer to **NullableString** |  | [optional] 
**SupportPriorityId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportTicketDto

`func NewSupportTicketDto() *SupportTicketDto`

NewSupportTicketDto instantiates a new SupportTicketDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketDtoWithDefaults

`func NewSupportTicketDtoWithDefaults() *SupportTicketDto`

NewSupportTicketDtoWithDefaults instantiates a new SupportTicketDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicketDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportTicketDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportTicketDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SupportTicketDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportTicketDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportTicketDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportTicketDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SupportTicketDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SupportTicketDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SupportTicketDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportTicketDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportTicketDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportTicketDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportTicketDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportTicketDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportTicketDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicketDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicketDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicketDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportTicketDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportTicketDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupportTicketStatus

`func (o *SupportTicketDto) GetSupportTicketStatus() string`

GetSupportTicketStatus returns the SupportTicketStatus field if non-nil, zero value otherwise.

### GetSupportTicketStatusOk

`func (o *SupportTicketDto) GetSupportTicketStatusOk() (*string, bool)`

GetSupportTicketStatusOk returns a tuple with the SupportTicketStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketStatus

`func (o *SupportTicketDto) SetSupportTicketStatus(v string)`

SetSupportTicketStatus sets SupportTicketStatus field to given value.

### HasSupportTicketStatus

`func (o *SupportTicketDto) HasSupportTicketStatus() bool`

HasSupportTicketStatus returns a boolean if a field has been set.

### GetUserId

`func (o *SupportTicketDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SupportTicketDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SupportTicketDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SupportTicketDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SupportTicketDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SupportTicketDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetContactId

`func (o *SupportTicketDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *SupportTicketDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *SupportTicketDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *SupportTicketDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *SupportTicketDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *SupportTicketDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *SupportTicketDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SupportTicketDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SupportTicketDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SupportTicketDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SupportTicketDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SupportTicketDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *SupportTicketDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SupportTicketDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SupportTicketDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SupportTicketDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SupportTicketDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SupportTicketDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSupportTicketTypeId

`func (o *SupportTicketDto) GetSupportTicketTypeId() string`

GetSupportTicketTypeId returns the SupportTicketTypeId field if non-nil, zero value otherwise.

### GetSupportTicketTypeIdOk

`func (o *SupportTicketDto) GetSupportTicketTypeIdOk() (*string, bool)`

GetSupportTicketTypeIdOk returns a tuple with the SupportTicketTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketTypeId

`func (o *SupportTicketDto) SetSupportTicketTypeId(v string)`

SetSupportTicketTypeId sets SupportTicketTypeId field to given value.

### HasSupportTicketTypeId

`func (o *SupportTicketDto) HasSupportTicketTypeId() bool`

HasSupportTicketTypeId returns a boolean if a field has been set.

### SetSupportTicketTypeIdNil

`func (o *SupportTicketDto) SetSupportTicketTypeIdNil(b bool)`

 SetSupportTicketTypeIdNil sets the value for SupportTicketTypeId to be an explicit nil

### UnsetSupportTicketTypeId
`func (o *SupportTicketDto) UnsetSupportTicketTypeId()`

UnsetSupportTicketTypeId ensures that no value is present for SupportTicketTypeId, not even an explicit nil
### GetSupportEntitlementId

`func (o *SupportTicketDto) GetSupportEntitlementId() string`

GetSupportEntitlementId returns the SupportEntitlementId field if non-nil, zero value otherwise.

### GetSupportEntitlementIdOk

`func (o *SupportTicketDto) GetSupportEntitlementIdOk() (*string, bool)`

GetSupportEntitlementIdOk returns a tuple with the SupportEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementId

`func (o *SupportTicketDto) SetSupportEntitlementId(v string)`

SetSupportEntitlementId sets SupportEntitlementId field to given value.

### HasSupportEntitlementId

`func (o *SupportTicketDto) HasSupportEntitlementId() bool`

HasSupportEntitlementId returns a boolean if a field has been set.

### SetSupportEntitlementIdNil

`func (o *SupportTicketDto) SetSupportEntitlementIdNil(b bool)`

 SetSupportEntitlementIdNil sets the value for SupportEntitlementId to be an explicit nil

### UnsetSupportEntitlementId
`func (o *SupportTicketDto) UnsetSupportEntitlementId()`

UnsetSupportEntitlementId ensures that no value is present for SupportEntitlementId, not even an explicit nil
### GetSupportPriorityId

`func (o *SupportTicketDto) GetSupportPriorityId() string`

GetSupportPriorityId returns the SupportPriorityId field if non-nil, zero value otherwise.

### GetSupportPriorityIdOk

`func (o *SupportTicketDto) GetSupportPriorityIdOk() (*string, bool)`

GetSupportPriorityIdOk returns a tuple with the SupportPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPriorityId

`func (o *SupportTicketDto) SetSupportPriorityId(v string)`

SetSupportPriorityId sets SupportPriorityId field to given value.

### HasSupportPriorityId

`func (o *SupportTicketDto) HasSupportPriorityId() bool`

HasSupportPriorityId returns a boolean if a field has been set.

### SetSupportPriorityIdNil

`func (o *SupportTicketDto) SetSupportPriorityIdNil(b bool)`

 SetSupportPriorityIdNil sets the value for SupportPriorityId to be an explicit nil

### UnsetSupportPriorityId
`func (o *SupportTicketDto) UnsetSupportPriorityId()`

UnsetSupportPriorityId ensures that no value is present for SupportPriorityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


