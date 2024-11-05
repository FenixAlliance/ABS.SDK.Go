# SupportTicketPriorityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportTicketPriorityDto

`func NewSupportTicketPriorityDto() *SupportTicketPriorityDto`

NewSupportTicketPriorityDto instantiates a new SupportTicketPriorityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketPriorityDtoWithDefaults

`func NewSupportTicketPriorityDtoWithDefaults() *SupportTicketPriorityDto`

NewSupportTicketPriorityDtoWithDefaults instantiates a new SupportTicketPriorityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicketPriorityDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketPriorityDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketPriorityDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketPriorityDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportTicketPriorityDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportTicketPriorityDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SupportTicketPriorityDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportTicketPriorityDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportTicketPriorityDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportTicketPriorityDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SupportTicketPriorityDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SupportTicketPriorityDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SupportTicketPriorityDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportTicketPriorityDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportTicketPriorityDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportTicketPriorityDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportTicketPriorityDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportTicketPriorityDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportTicketPriorityDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicketPriorityDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicketPriorityDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicketPriorityDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportTicketPriorityDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportTicketPriorityDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBusinessID

`func (o *SupportTicketPriorityDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportTicketPriorityDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportTicketPriorityDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportTicketPriorityDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportTicketPriorityDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportTicketPriorityDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetSupportEntitlementID

`func (o *SupportTicketPriorityDto) GetSupportEntitlementID() string`

GetSupportEntitlementID returns the SupportEntitlementID field if non-nil, zero value otherwise.

### GetSupportEntitlementIDOk

`func (o *SupportTicketPriorityDto) GetSupportEntitlementIDOk() (*string, bool)`

GetSupportEntitlementIDOk returns a tuple with the SupportEntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementID

`func (o *SupportTicketPriorityDto) SetSupportEntitlementID(v string)`

SetSupportEntitlementID sets SupportEntitlementID field to given value.

### HasSupportEntitlementID

`func (o *SupportTicketPriorityDto) HasSupportEntitlementID() bool`

HasSupportEntitlementID returns a boolean if a field has been set.

### SetSupportEntitlementIDNil

`func (o *SupportTicketPriorityDto) SetSupportEntitlementIDNil(b bool)`

 SetSupportEntitlementIDNil sets the value for SupportEntitlementID to be an explicit nil

### UnsetSupportEntitlementID
`func (o *SupportTicketPriorityDto) UnsetSupportEntitlementID()`

UnsetSupportEntitlementID ensures that no value is present for SupportEntitlementID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


