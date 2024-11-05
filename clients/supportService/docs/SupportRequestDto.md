# SupportRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**ApprovedTimestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementID** | Pointer to **NullableString** |  | [optional] 
**ContactID** | Pointer to **NullableString** |  | [optional] 
**AccountHolderID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportRequestDto

`func NewSupportRequestDto() *SupportRequestDto`

NewSupportRequestDto instantiates a new SupportRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestDtoWithDefaults

`func NewSupportRequestDtoWithDefaults() *SupportRequestDto`

NewSupportRequestDtoWithDefaults instantiates a new SupportRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportRequestDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportRequestDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportRequestDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportRequestDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportRequestDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportRequestDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SupportRequestDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportRequestDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportRequestDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportRequestDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SupportRequestDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SupportRequestDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SupportRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportRequestDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SupportRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportRequestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportRequestDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportRequestDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApproved

`func (o *SupportRequestDto) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *SupportRequestDto) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *SupportRequestDto) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *SupportRequestDto) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetApprovedTimestamp

`func (o *SupportRequestDto) GetApprovedTimestamp() time.Time`

GetApprovedTimestamp returns the ApprovedTimestamp field if non-nil, zero value otherwise.

### GetApprovedTimestampOk

`func (o *SupportRequestDto) GetApprovedTimestampOk() (*time.Time, bool)`

GetApprovedTimestampOk returns a tuple with the ApprovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTimestamp

`func (o *SupportRequestDto) SetApprovedTimestamp(v time.Time)`

SetApprovedTimestamp sets ApprovedTimestamp field to given value.

### HasApprovedTimestamp

`func (o *SupportRequestDto) HasApprovedTimestamp() bool`

HasApprovedTimestamp returns a boolean if a field has been set.

### GetBusinessID

`func (o *SupportRequestDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportRequestDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportRequestDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportRequestDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportRequestDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportRequestDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportRequestDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportRequestDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportRequestDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportRequestDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportRequestDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportRequestDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetSupportEntitlementID

`func (o *SupportRequestDto) GetSupportEntitlementID() string`

GetSupportEntitlementID returns the SupportEntitlementID field if non-nil, zero value otherwise.

### GetSupportEntitlementIDOk

`func (o *SupportRequestDto) GetSupportEntitlementIDOk() (*string, bool)`

GetSupportEntitlementIDOk returns a tuple with the SupportEntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementID

`func (o *SupportRequestDto) SetSupportEntitlementID(v string)`

SetSupportEntitlementID sets SupportEntitlementID field to given value.

### HasSupportEntitlementID

`func (o *SupportRequestDto) HasSupportEntitlementID() bool`

HasSupportEntitlementID returns a boolean if a field has been set.

### SetSupportEntitlementIDNil

`func (o *SupportRequestDto) SetSupportEntitlementIDNil(b bool)`

 SetSupportEntitlementIDNil sets the value for SupportEntitlementID to be an explicit nil

### UnsetSupportEntitlementID
`func (o *SupportRequestDto) UnsetSupportEntitlementID()`

UnsetSupportEntitlementID ensures that no value is present for SupportEntitlementID, not even an explicit nil
### GetContactID

`func (o *SupportRequestDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *SupportRequestDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *SupportRequestDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactID

`func (o *SupportRequestDto) HasContactID() bool`

HasContactID returns a boolean if a field has been set.

### SetContactIDNil

`func (o *SupportRequestDto) SetContactIDNil(b bool)`

 SetContactIDNil sets the value for ContactID to be an explicit nil

### UnsetContactID
`func (o *SupportRequestDto) UnsetContactID()`

UnsetContactID ensures that no value is present for ContactID, not even an explicit nil
### GetAccountHolderID

`func (o *SupportRequestDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *SupportRequestDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *SupportRequestDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *SupportRequestDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *SupportRequestDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *SupportRequestDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


