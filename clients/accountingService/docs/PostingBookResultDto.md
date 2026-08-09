# PostingBookResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**FinancialBookId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**JournalEntryId** | Pointer to **NullableString** |  | [optional] 
**FailureCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostingBookResultDto

`func NewPostingBookResultDto() *PostingBookResultDto`

NewPostingBookResultDto instantiates a new PostingBookResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostingBookResultDtoWithDefaults

`func NewPostingBookResultDtoWithDefaults() *PostingBookResultDto`

NewPostingBookResultDtoWithDefaults instantiates a new PostingBookResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostingBookResultDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostingBookResultDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostingBookResultDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostingBookResultDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PostingBookResultDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PostingBookResultDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PostingBookResultDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PostingBookResultDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PostingBookResultDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PostingBookResultDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PostingBookResultDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PostingBookResultDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetFinancialBookId

`func (o *PostingBookResultDto) GetFinancialBookId() string`

GetFinancialBookId returns the FinancialBookId field if non-nil, zero value otherwise.

### GetFinancialBookIdOk

`func (o *PostingBookResultDto) GetFinancialBookIdOk() (*string, bool)`

GetFinancialBookIdOk returns a tuple with the FinancialBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialBookId

`func (o *PostingBookResultDto) SetFinancialBookId(v string)`

SetFinancialBookId sets FinancialBookId field to given value.

### HasFinancialBookId

`func (o *PostingBookResultDto) HasFinancialBookId() bool`

HasFinancialBookId returns a boolean if a field has been set.

### SetFinancialBookIdNil

`func (o *PostingBookResultDto) SetFinancialBookIdNil(b bool)`

 SetFinancialBookIdNil sets the value for FinancialBookId to be an explicit nil

### UnsetFinancialBookId
`func (o *PostingBookResultDto) UnsetFinancialBookId()`

UnsetFinancialBookId ensures that no value is present for FinancialBookId, not even an explicit nil
### GetStatus

`func (o *PostingBookResultDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostingBookResultDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostingBookResultDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostingBookResultDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJournalEntryId

`func (o *PostingBookResultDto) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *PostingBookResultDto) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *PostingBookResultDto) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.

### HasJournalEntryId

`func (o *PostingBookResultDto) HasJournalEntryId() bool`

HasJournalEntryId returns a boolean if a field has been set.

### SetJournalEntryIdNil

`func (o *PostingBookResultDto) SetJournalEntryIdNil(b bool)`

 SetJournalEntryIdNil sets the value for JournalEntryId to be an explicit nil

### UnsetJournalEntryId
`func (o *PostingBookResultDto) UnsetJournalEntryId()`

UnsetJournalEntryId ensures that no value is present for JournalEntryId, not even an explicit nil
### GetFailureCode

`func (o *PostingBookResultDto) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PostingBookResultDto) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PostingBookResultDto) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PostingBookResultDto) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *PostingBookResultDto) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *PostingBookResultDto) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


