# JournalCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**ParentJournalId** | Pointer to **NullableString** |  | [optional] 
**JournalTypeId** | Pointer to **NullableString** |  | [optional] 
**LedgerId** | Pointer to **NullableString** |  | [optional] 

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

### GetParentJournalId

`func (o *JournalCreateDto) GetParentJournalId() string`

GetParentJournalId returns the ParentJournalId field if non-nil, zero value otherwise.

### GetParentJournalIdOk

`func (o *JournalCreateDto) GetParentJournalIdOk() (*string, bool)`

GetParentJournalIdOk returns a tuple with the ParentJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJournalId

`func (o *JournalCreateDto) SetParentJournalId(v string)`

SetParentJournalId sets ParentJournalId field to given value.

### HasParentJournalId

`func (o *JournalCreateDto) HasParentJournalId() bool`

HasParentJournalId returns a boolean if a field has been set.

### SetParentJournalIdNil

`func (o *JournalCreateDto) SetParentJournalIdNil(b bool)`

 SetParentJournalIdNil sets the value for ParentJournalId to be an explicit nil

### UnsetParentJournalId
`func (o *JournalCreateDto) UnsetParentJournalId()`

UnsetParentJournalId ensures that no value is present for ParentJournalId, not even an explicit nil
### GetJournalTypeId

`func (o *JournalCreateDto) GetJournalTypeId() string`

GetJournalTypeId returns the JournalTypeId field if non-nil, zero value otherwise.

### GetJournalTypeIdOk

`func (o *JournalCreateDto) GetJournalTypeIdOk() (*string, bool)`

GetJournalTypeIdOk returns a tuple with the JournalTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTypeId

`func (o *JournalCreateDto) SetJournalTypeId(v string)`

SetJournalTypeId sets JournalTypeId field to given value.

### HasJournalTypeId

`func (o *JournalCreateDto) HasJournalTypeId() bool`

HasJournalTypeId returns a boolean if a field has been set.

### SetJournalTypeIdNil

`func (o *JournalCreateDto) SetJournalTypeIdNil(b bool)`

 SetJournalTypeIdNil sets the value for JournalTypeId to be an explicit nil

### UnsetJournalTypeId
`func (o *JournalCreateDto) UnsetJournalTypeId()`

UnsetJournalTypeId ensures that no value is present for JournalTypeId, not even an explicit nil
### GetLedgerId

`func (o *JournalCreateDto) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *JournalCreateDto) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *JournalCreateDto) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *JournalCreateDto) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### SetLedgerIdNil

`func (o *JournalCreateDto) SetLedgerIdNil(b bool)`

 SetLedgerIdNil sets the value for LedgerId to be an explicit nil

### UnsetLedgerId
`func (o *JournalCreateDto) UnsetLedgerId()`

UnsetLedgerId ensures that no value is present for LedgerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


