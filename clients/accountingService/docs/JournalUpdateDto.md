# JournalUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**ParentJournalID** | Pointer to **NullableString** |  | [optional] 
**JournalTypeID** | Pointer to **NullableString** |  | [optional] 
**LedgerID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJournalUpdateDto

`func NewJournalUpdateDto() *JournalUpdateDto`

NewJournalUpdateDto instantiates a new JournalUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalUpdateDtoWithDefaults

`func NewJournalUpdateDtoWithDefaults() *JournalUpdateDto`

NewJournalUpdateDtoWithDefaults instantiates a new JournalUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JournalUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JournalUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JournalUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JournalUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JournalUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JournalUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *JournalUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JournalUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JournalUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateTime

`func (o *JournalUpdateDto) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *JournalUpdateDto) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *JournalUpdateDto) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *JournalUpdateDto) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetParentJournalID

`func (o *JournalUpdateDto) GetParentJournalID() string`

GetParentJournalID returns the ParentJournalID field if non-nil, zero value otherwise.

### GetParentJournalIDOk

`func (o *JournalUpdateDto) GetParentJournalIDOk() (*string, bool)`

GetParentJournalIDOk returns a tuple with the ParentJournalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalID

`func (o *JournalUpdateDto) SetParentJournalID(v string)`

SetParentJournalID sets ParentJournalID field to given value.

### HasParentJournalID

`func (o *JournalUpdateDto) HasParentJournalID() bool`

HasParentJournalID returns a boolean if a field has been set.

### SetParentJournalIDNil

`func (o *JournalUpdateDto) SetParentJournalIDNil(b bool)`

 SetParentJournalIDNil sets the value for ParentJournalID to be an explicit nil

### UnsetParentJournalID
`func (o *JournalUpdateDto) UnsetParentJournalID()`

UnsetParentJournalID ensures that no value is present for ParentJournalID, not even an explicit nil
### GetJournalTypeID

`func (o *JournalUpdateDto) GetJournalTypeID() string`

GetJournalTypeID returns the JournalTypeID field if non-nil, zero value otherwise.

### GetJournalTypeIDOk

`func (o *JournalUpdateDto) GetJournalTypeIDOk() (*string, bool)`

GetJournalTypeIDOk returns a tuple with the JournalTypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTypeID

`func (o *JournalUpdateDto) SetJournalTypeID(v string)`

SetJournalTypeID sets JournalTypeID field to given value.

### HasJournalTypeID

`func (o *JournalUpdateDto) HasJournalTypeID() bool`

HasJournalTypeID returns a boolean if a field has been set.

### SetJournalTypeIDNil

`func (o *JournalUpdateDto) SetJournalTypeIDNil(b bool)`

 SetJournalTypeIDNil sets the value for JournalTypeID to be an explicit nil

### UnsetJournalTypeID
`func (o *JournalUpdateDto) UnsetJournalTypeID()`

UnsetJournalTypeID ensures that no value is present for JournalTypeID, not even an explicit nil
### GetLedgerID

`func (o *JournalUpdateDto) GetLedgerID() string`

GetLedgerID returns the LedgerID field if non-nil, zero value otherwise.

### GetLedgerIDOk

`func (o *JournalUpdateDto) GetLedgerIDOk() (*string, bool)`

GetLedgerIDOk returns a tuple with the LedgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerID

`func (o *JournalUpdateDto) SetLedgerID(v string)`

SetLedgerID sets LedgerID field to given value.

### HasLedgerID

`func (o *JournalUpdateDto) HasLedgerID() bool`

HasLedgerID returns a boolean if a field has been set.

### SetLedgerIDNil

`func (o *JournalUpdateDto) SetLedgerIDNil(b bool)`

 SetLedgerIDNil sets the value for LedgerID to be an explicit nil

### UnsetLedgerID
`func (o *JournalUpdateDto) UnsetLedgerID()`

UnsetLedgerID ensures that no value is present for LedgerID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


