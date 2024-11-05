# SupportTicketCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AccountHolderID** | Pointer to **NullableString** |  | [optional] 
**ContactID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**SupportTicketTypeID** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementID** | Pointer to **NullableString** |  | [optional] 
**SupportPriorityID** | Pointer to **NullableString** |  | [optional] 

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
### GetAccountHolderID

`func (o *SupportTicketCreateDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *SupportTicketCreateDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *SupportTicketCreateDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *SupportTicketCreateDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *SupportTicketCreateDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *SupportTicketCreateDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil
### GetContactID

`func (o *SupportTicketCreateDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *SupportTicketCreateDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *SupportTicketCreateDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactID

`func (o *SupportTicketCreateDto) HasContactID() bool`

HasContactID returns a boolean if a field has been set.

### SetContactIDNil

`func (o *SupportTicketCreateDto) SetContactIDNil(b bool)`

 SetContactIDNil sets the value for ContactID to be an explicit nil

### UnsetContactID
`func (o *SupportTicketCreateDto) UnsetContactID()`

UnsetContactID ensures that no value is present for ContactID, not even an explicit nil
### GetBusinessID

`func (o *SupportTicketCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportTicketCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportTicketCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportTicketCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportTicketCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportTicketCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportTicketCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportTicketCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportTicketCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportTicketCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportTicketCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportTicketCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetSupportTicketTypeID

`func (o *SupportTicketCreateDto) GetSupportTicketTypeID() string`

GetSupportTicketTypeID returns the SupportTicketTypeID field if non-nil, zero value otherwise.

### GetSupportTicketTypeIDOk

`func (o *SupportTicketCreateDto) GetSupportTicketTypeIDOk() (*string, bool)`

GetSupportTicketTypeIDOk returns a tuple with the SupportTicketTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketTypeID

`func (o *SupportTicketCreateDto) SetSupportTicketTypeID(v string)`

SetSupportTicketTypeID sets SupportTicketTypeID field to given value.

### HasSupportTicketTypeID

`func (o *SupportTicketCreateDto) HasSupportTicketTypeID() bool`

HasSupportTicketTypeID returns a boolean if a field has been set.

### SetSupportTicketTypeIDNil

`func (o *SupportTicketCreateDto) SetSupportTicketTypeIDNil(b bool)`

 SetSupportTicketTypeIDNil sets the value for SupportTicketTypeID to be an explicit nil

### UnsetSupportTicketTypeID
`func (o *SupportTicketCreateDto) UnsetSupportTicketTypeID()`

UnsetSupportTicketTypeID ensures that no value is present for SupportTicketTypeID, not even an explicit nil
### GetSupportEntitlementID

`func (o *SupportTicketCreateDto) GetSupportEntitlementID() string`

GetSupportEntitlementID returns the SupportEntitlementID field if non-nil, zero value otherwise.

### GetSupportEntitlementIDOk

`func (o *SupportTicketCreateDto) GetSupportEntitlementIDOk() (*string, bool)`

GetSupportEntitlementIDOk returns a tuple with the SupportEntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementID

`func (o *SupportTicketCreateDto) SetSupportEntitlementID(v string)`

SetSupportEntitlementID sets SupportEntitlementID field to given value.

### HasSupportEntitlementID

`func (o *SupportTicketCreateDto) HasSupportEntitlementID() bool`

HasSupportEntitlementID returns a boolean if a field has been set.

### SetSupportEntitlementIDNil

`func (o *SupportTicketCreateDto) SetSupportEntitlementIDNil(b bool)`

 SetSupportEntitlementIDNil sets the value for SupportEntitlementID to be an explicit nil

### UnsetSupportEntitlementID
`func (o *SupportTicketCreateDto) UnsetSupportEntitlementID()`

UnsetSupportEntitlementID ensures that no value is present for SupportEntitlementID, not even an explicit nil
### GetSupportPriorityID

`func (o *SupportTicketCreateDto) GetSupportPriorityID() string`

GetSupportPriorityID returns the SupportPriorityID field if non-nil, zero value otherwise.

### GetSupportPriorityIDOk

`func (o *SupportTicketCreateDto) GetSupportPriorityIDOk() (*string, bool)`

GetSupportPriorityIDOk returns a tuple with the SupportPriorityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPriorityID

`func (o *SupportTicketCreateDto) SetSupportPriorityID(v string)`

SetSupportPriorityID sets SupportPriorityID field to given value.

### HasSupportPriorityID

`func (o *SupportTicketCreateDto) HasSupportPriorityID() bool`

HasSupportPriorityID returns a boolean if a field has been set.

### SetSupportPriorityIDNil

`func (o *SupportTicketCreateDto) SetSupportPriorityIDNil(b bool)`

 SetSupportPriorityIDNil sets the value for SupportPriorityID to be an explicit nil

### UnsetSupportPriorityID
`func (o *SupportTicketCreateDto) UnsetSupportPriorityID()`

UnsetSupportPriorityID ensures that no value is present for SupportPriorityID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


