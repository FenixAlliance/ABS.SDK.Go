# JournalCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**TenantID** | Pointer to **NullableString** |  | [optional] 
**EnrollmentID** | Pointer to **NullableString** |  | [optional] 
**ParentJournalID** | Pointer to **NullableString** |  | [optional] 
**JournalTypeID** | Pointer to **NullableString** |  | [optional] 
**LedgerID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJournalCreateDto

`func NewJournalCreateDto(name string, ) *JournalCreateDto`

NewJournalCreateDto instantiates a new JournalCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalCreateDtoWithDefaults

`func NewJournalCreateDtoWithDefaults() *JournalCreateDto`

NewJournalCreateDtoWithDefaults instantiates a new JournalCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JournalCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *JournalCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JournalCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JournalCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JournalCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *JournalCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JournalCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JournalCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JournalCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JournalCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JournalCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateTime

`func (o *JournalCreateDto) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *JournalCreateDto) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *JournalCreateDto) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *JournalCreateDto) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetTenantID

`func (o *JournalCreateDto) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *JournalCreateDto) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *JournalCreateDto) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *JournalCreateDto) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### SetTenantIDNil

`func (o *JournalCreateDto) SetTenantIDNil(b bool)`

 SetTenantIDNil sets the value for TenantID to be an explicit nil

### UnsetTenantID
`func (o *JournalCreateDto) UnsetTenantID()`

UnsetTenantID ensures that no value is present for TenantID, not even an explicit nil
### GetEnrollmentID

`func (o *JournalCreateDto) GetEnrollmentID() string`

GetEnrollmentID returns the EnrollmentID field if non-nil, zero value otherwise.

### GetEnrollmentIDOk

`func (o *JournalCreateDto) GetEnrollmentIDOk() (*string, bool)`

GetEnrollmentIDOk returns a tuple with the EnrollmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentID

`func (o *JournalCreateDto) SetEnrollmentID(v string)`

SetEnrollmentID sets EnrollmentID field to given value.

### HasEnrollmentID

`func (o *JournalCreateDto) HasEnrollmentID() bool`

HasEnrollmentID returns a boolean if a field has been set.

### SetEnrollmentIDNil

`func (o *JournalCreateDto) SetEnrollmentIDNil(b bool)`

 SetEnrollmentIDNil sets the value for EnrollmentID to be an explicit nil

### UnsetEnrollmentID
`func (o *JournalCreateDto) UnsetEnrollmentID()`

UnsetEnrollmentID ensures that no value is present for EnrollmentID, not even an explicit nil
### GetParentJournalID

`func (o *JournalCreateDto) GetParentJournalID() string`

GetParentJournalID returns the ParentJournalID field if non-nil, zero value otherwise.

### GetParentJournalIDOk

`func (o *JournalCreateDto) GetParentJournalIDOk() (*string, bool)`

GetParentJournalIDOk returns a tuple with the ParentJournalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalID

`func (o *JournalCreateDto) SetParentJournalID(v string)`

SetParentJournalID sets ParentJournalID field to given value.

### HasParentJournalID

`func (o *JournalCreateDto) HasParentJournalID() bool`

HasParentJournalID returns a boolean if a field has been set.

### SetParentJournalIDNil

`func (o *JournalCreateDto) SetParentJournalIDNil(b bool)`

 SetParentJournalIDNil sets the value for ParentJournalID to be an explicit nil

### UnsetParentJournalID
`func (o *JournalCreateDto) UnsetParentJournalID()`

UnsetParentJournalID ensures that no value is present for ParentJournalID, not even an explicit nil
### GetJournalTypeID

`func (o *JournalCreateDto) GetJournalTypeID() string`

GetJournalTypeID returns the JournalTypeID field if non-nil, zero value otherwise.

### GetJournalTypeIDOk

`func (o *JournalCreateDto) GetJournalTypeIDOk() (*string, bool)`

GetJournalTypeIDOk returns a tuple with the JournalTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTypeID

`func (o *JournalCreateDto) SetJournalTypeID(v string)`

SetJournalTypeID sets JournalTypeID field to given value.

### HasJournalTypeID

`func (o *JournalCreateDto) HasJournalTypeID() bool`

HasJournalTypeID returns a boolean if a field has been set.

### SetJournalTypeIDNil

`func (o *JournalCreateDto) SetJournalTypeIDNil(b bool)`

 SetJournalTypeIDNil sets the value for JournalTypeID to be an explicit nil

### UnsetJournalTypeID
`func (o *JournalCreateDto) UnsetJournalTypeID()`

UnsetJournalTypeID ensures that no value is present for JournalTypeID, not even an explicit nil
### GetLedgerID

`func (o *JournalCreateDto) GetLedgerID() string`

GetLedgerID returns the LedgerID field if non-nil, zero value otherwise.

### GetLedgerIDOk

`func (o *JournalCreateDto) GetLedgerIDOk() (*string, bool)`

GetLedgerIDOk returns a tuple with the LedgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerID

`func (o *JournalCreateDto) SetLedgerID(v string)`

SetLedgerID sets LedgerID field to given value.

### HasLedgerID

`func (o *JournalCreateDto) HasLedgerID() bool`

HasLedgerID returns a boolean if a field has been set.

### SetLedgerIDNil

`func (o *JournalCreateDto) SetLedgerIDNil(b bool)`

 SetLedgerIDNil sets the value for LedgerID to be an explicit nil

### UnsetLedgerID
`func (o *JournalCreateDto) UnsetLedgerID()`

UnsetLedgerID ensures that no value is present for LedgerID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


