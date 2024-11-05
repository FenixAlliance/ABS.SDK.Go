# SupportTicketDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AccountHolderID** | Pointer to **NullableString** |  | [optional] 
**ContactID** | Pointer to **NullableString** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**SupportTicketTypeID** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementID** | Pointer to **NullableString** |  | [optional] 
**SupportPriorityID** | Pointer to **NullableString** |  | [optional] 

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
### GetAccountHolderID

`func (o *SupportTicketDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *SupportTicketDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *SupportTicketDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *SupportTicketDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *SupportTicketDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *SupportTicketDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil
### GetContactID

`func (o *SupportTicketDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *SupportTicketDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *SupportTicketDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactID

`func (o *SupportTicketDto) HasContactID() bool`

HasContactID returns a boolean if a field has been set.

### SetContactIDNil

`func (o *SupportTicketDto) SetContactIDNil(b bool)`

 SetContactIDNil sets the value for ContactID to be an explicit nil

### UnsetContactID
`func (o *SupportTicketDto) UnsetContactID()`

UnsetContactID ensures that no value is present for ContactID, not even an explicit nil
### GetBusinessID

`func (o *SupportTicketDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportTicketDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportTicketDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportTicketDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportTicketDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportTicketDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportTicketDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportTicketDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportTicketDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportTicketDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportTicketDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportTicketDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetSupportTicketTypeID

`func (o *SupportTicketDto) GetSupportTicketTypeID() string`

GetSupportTicketTypeID returns the SupportTicketTypeID field if non-nil, zero value otherwise.

### GetSupportTicketTypeIDOk

`func (o *SupportTicketDto) GetSupportTicketTypeIDOk() (*string, bool)`

GetSupportTicketTypeIDOk returns a tuple with the SupportTicketTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketTypeID

`func (o *SupportTicketDto) SetSupportTicketTypeID(v string)`

SetSupportTicketTypeID sets SupportTicketTypeID field to given value.

### HasSupportTicketTypeID

`func (o *SupportTicketDto) HasSupportTicketTypeID() bool`

HasSupportTicketTypeID returns a boolean if a field has been set.

### SetSupportTicketTypeIDNil

`func (o *SupportTicketDto) SetSupportTicketTypeIDNil(b bool)`

 SetSupportTicketTypeIDNil sets the value for SupportTicketTypeID to be an explicit nil

### UnsetSupportTicketTypeID
`func (o *SupportTicketDto) UnsetSupportTicketTypeID()`

UnsetSupportTicketTypeID ensures that no value is present for SupportTicketTypeID, not even an explicit nil
### GetSupportEntitlementID

`func (o *SupportTicketDto) GetSupportEntitlementID() string`

GetSupportEntitlementID returns the SupportEntitlementID field if non-nil, zero value otherwise.

### GetSupportEntitlementIDOk

`func (o *SupportTicketDto) GetSupportEntitlementIDOk() (*string, bool)`

GetSupportEntitlementIDOk returns a tuple with the SupportEntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementID

`func (o *SupportTicketDto) SetSupportEntitlementID(v string)`

SetSupportEntitlementID sets SupportEntitlementID field to given value.

### HasSupportEntitlementID

`func (o *SupportTicketDto) HasSupportEntitlementID() bool`

HasSupportEntitlementID returns a boolean if a field has been set.

### SetSupportEntitlementIDNil

`func (o *SupportTicketDto) SetSupportEntitlementIDNil(b bool)`

 SetSupportEntitlementIDNil sets the value for SupportEntitlementID to be an explicit nil

### UnsetSupportEntitlementID
`func (o *SupportTicketDto) UnsetSupportEntitlementID()`

UnsetSupportEntitlementID ensures that no value is present for SupportEntitlementID, not even an explicit nil
### GetSupportPriorityID

`func (o *SupportTicketDto) GetSupportPriorityID() string`

GetSupportPriorityID returns the SupportPriorityID field if non-nil, zero value otherwise.

### GetSupportPriorityIDOk

`func (o *SupportTicketDto) GetSupportPriorityIDOk() (*string, bool)`

GetSupportPriorityIDOk returns a tuple with the SupportPriorityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPriorityID

`func (o *SupportTicketDto) SetSupportPriorityID(v string)`

SetSupportPriorityID sets SupportPriorityID field to given value.

### HasSupportPriorityID

`func (o *SupportTicketDto) HasSupportPriorityID() bool`

HasSupportPriorityID returns a boolean if a field has been set.

### SetSupportPriorityIDNil

`func (o *SupportTicketDto) SetSupportPriorityIDNil(b bool)`

 SetSupportPriorityIDNil sets the value for SupportPriorityID to be an explicit nil

### UnsetSupportPriorityID
`func (o *SupportTicketDto) UnsetSupportPriorityID()`

UnsetSupportPriorityID ensures that no value is present for SupportPriorityID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


