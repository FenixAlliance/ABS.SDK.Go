# SupportRequestCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**ApprovedTimestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementID** | Pointer to **NullableString** |  | [optional] 
**ContactID** | Pointer to **NullableString** |  | [optional] 
**AccountHolderID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportRequestCreateDto

`func NewSupportRequestCreateDto(title string, ) *SupportRequestCreateDto`

NewSupportRequestCreateDto instantiates a new SupportRequestCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestCreateDtoWithDefaults

`func NewSupportRequestCreateDtoWithDefaults() *SupportRequestCreateDto`

NewSupportRequestCreateDtoWithDefaults instantiates a new SupportRequestCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportRequestCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportRequestCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportRequestCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportRequestCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SupportRequestCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportRequestCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportRequestCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportRequestCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *SupportRequestCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportRequestCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportRequestCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SupportRequestCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportRequestCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportRequestCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportRequestCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SupportRequestCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SupportRequestCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApproved

`func (o *SupportRequestCreateDto) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *SupportRequestCreateDto) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *SupportRequestCreateDto) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *SupportRequestCreateDto) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetApprovedTimestamp

`func (o *SupportRequestCreateDto) GetApprovedTimestamp() time.Time`

GetApprovedTimestamp returns the ApprovedTimestamp field if non-nil, zero value otherwise.

### GetApprovedTimestampOk

`func (o *SupportRequestCreateDto) GetApprovedTimestampOk() (*time.Time, bool)`

GetApprovedTimestampOk returns a tuple with the ApprovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTimestamp

`func (o *SupportRequestCreateDto) SetApprovedTimestamp(v time.Time)`

SetApprovedTimestamp sets ApprovedTimestamp field to given value.

### HasApprovedTimestamp

`func (o *SupportRequestCreateDto) HasApprovedTimestamp() bool`

HasApprovedTimestamp returns a boolean if a field has been set.

### GetBusinessID

`func (o *SupportRequestCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *SupportRequestCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *SupportRequestCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *SupportRequestCreateDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *SupportRequestCreateDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *SupportRequestCreateDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *SupportRequestCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *SupportRequestCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *SupportRequestCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *SupportRequestCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *SupportRequestCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *SupportRequestCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetSupportEntitlementID

`func (o *SupportRequestCreateDto) GetSupportEntitlementID() string`

GetSupportEntitlementID returns the SupportEntitlementID field if non-nil, zero value otherwise.

### GetSupportEntitlementIDOk

`func (o *SupportRequestCreateDto) GetSupportEntitlementIDOk() (*string, bool)`

GetSupportEntitlementIDOk returns a tuple with the SupportEntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementID

`func (o *SupportRequestCreateDto) SetSupportEntitlementID(v string)`

SetSupportEntitlementID sets SupportEntitlementID field to given value.

### HasSupportEntitlementID

`func (o *SupportRequestCreateDto) HasSupportEntitlementID() bool`

HasSupportEntitlementID returns a boolean if a field has been set.

### SetSupportEntitlementIDNil

`func (o *SupportRequestCreateDto) SetSupportEntitlementIDNil(b bool)`

 SetSupportEntitlementIDNil sets the value for SupportEntitlementID to be an explicit nil

### UnsetSupportEntitlementID
`func (o *SupportRequestCreateDto) UnsetSupportEntitlementID()`

UnsetSupportEntitlementID ensures that no value is present for SupportEntitlementID, not even an explicit nil
### GetContactID

`func (o *SupportRequestCreateDto) GetContactID() string`

GetContactID returns the ContactID field if non-nil, zero value otherwise.

### GetContactIDOk

`func (o *SupportRequestCreateDto) GetContactIDOk() (*string, bool)`

GetContactIDOk returns a tuple with the ContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactID

`func (o *SupportRequestCreateDto) SetContactID(v string)`

SetContactID sets ContactID field to given value.

### HasContactID

`func (o *SupportRequestCreateDto) HasContactID() bool`

HasContactID returns a boolean if a field has been set.

### SetContactIDNil

`func (o *SupportRequestCreateDto) SetContactIDNil(b bool)`

 SetContactIDNil sets the value for ContactID to be an explicit nil

### UnsetContactID
`func (o *SupportRequestCreateDto) UnsetContactID()`

UnsetContactID ensures that no value is present for ContactID, not even an explicit nil
### GetAccountHolderID

`func (o *SupportRequestCreateDto) GetAccountHolderID() string`

GetAccountHolderID returns the AccountHolderID field if non-nil, zero value otherwise.

### GetAccountHolderIDOk

`func (o *SupportRequestCreateDto) GetAccountHolderIDOk() (*string, bool)`

GetAccountHolderIDOk returns a tuple with the AccountHolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderID

`func (o *SupportRequestCreateDto) SetAccountHolderID(v string)`

SetAccountHolderID sets AccountHolderID field to given value.

### HasAccountHolderID

`func (o *SupportRequestCreateDto) HasAccountHolderID() bool`

HasAccountHolderID returns a boolean if a field has been set.

### SetAccountHolderIDNil

`func (o *SupportRequestCreateDto) SetAccountHolderIDNil(b bool)`

 SetAccountHolderIDNil sets the value for AccountHolderID to be an explicit nil

### UnsetAccountHolderID
`func (o *SupportRequestCreateDto) UnsetAccountHolderID()`

UnsetAccountHolderID ensures that no value is present for AccountHolderID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


