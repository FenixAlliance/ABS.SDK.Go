# JournalUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**ParentJournalId** | Pointer to **NullableString** |  | [optional] 
**JournalTypeId** | Pointer to **NullableString** |  | [optional] 
**LedgerId** | Pointer to **NullableString** |  | [optional] 

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

### GetParentJournalId

`func (o *JournalUpdateDto) GetParentJournalId() string`

GetParentJournalId returns the ParentJournalId field if non-nil, zero value otherwise.

### GetParentJournalIdOk

`func (o *JournalUpdateDto) GetParentJournalIdOk() (*string, bool)`

GetParentJournalIdOk returns a tuple with the ParentJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalId

`func (o *JournalUpdateDto) SetParentJournalId(v string)`

SetParentJournalId sets ParentJournalId field to given value.

### HasParentJournalId

`func (o *JournalUpdateDto) HasParentJournalId() bool`

HasParentJournalId returns a boolean if a field has been set.

### SetParentJournalIdNil

`func (o *JournalUpdateDto) SetParentJournalIdNil(b bool)`

 SetParentJournalIdNil sets the value for ParentJournalId to be an explicit nil

### UnsetParentJournalId
`func (o *JournalUpdateDto) UnsetParentJournalId()`

UnsetParentJournalId ensures that no value is present for ParentJournalId, not even an explicit nil
### GetJournalTypeId

`func (o *JournalUpdateDto) GetJournalTypeId() string`

GetJournalTypeId returns the JournalTypeId field if non-nil, zero value otherwise.

### GetJournalTypeIdOk

`func (o *JournalUpdateDto) GetJournalTypeIdOk() (*string, bool)`

GetJournalTypeIdOk returns a tuple with the JournalTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTypeId

`func (o *JournalUpdateDto) SetJournalTypeId(v string)`

SetJournalTypeId sets JournalTypeId field to given value.

### HasJournalTypeId

`func (o *JournalUpdateDto) HasJournalTypeId() bool`

HasJournalTypeId returns a boolean if a field has been set.

### SetJournalTypeIdNil

`func (o *JournalUpdateDto) SetJournalTypeIdNil(b bool)`

 SetJournalTypeIdNil sets the value for JournalTypeId to be an explicit nil

### UnsetJournalTypeId
`func (o *JournalUpdateDto) UnsetJournalTypeId()`

UnsetJournalTypeId ensures that no value is present for JournalTypeId, not even an explicit nil
### GetLedgerId

`func (o *JournalUpdateDto) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *JournalUpdateDto) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *JournalUpdateDto) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *JournalUpdateDto) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### SetLedgerIdNil

`func (o *JournalUpdateDto) SetLedgerIdNil(b bool)`

 SetLedgerIdNil sets the value for LedgerId to be an explicit nil

### UnsetLedgerId
`func (o *JournalUpdateDto) UnsetLedgerId()`

UnsetLedgerId ensures that no value is present for LedgerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


