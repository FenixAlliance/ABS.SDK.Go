# RecordSigningParticipantOutcomeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | **string** |  | 
**OutcomeAtUtc** | Pointer to **NullableTime** |  | [optional] 
**DeclineReason** | Pointer to **NullableString** |  | [optional] 
**ExternalReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRecordSigningParticipantOutcomeDto

`func NewRecordSigningParticipantOutcomeDto(outcome string, ) *RecordSigningParticipantOutcomeDto`

NewRecordSigningParticipantOutcomeDto instantiates a new RecordSigningParticipantOutcomeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSigningParticipantOutcomeDtoWithDefaults

`func NewRecordSigningParticipantOutcomeDtoWithDefaults() *RecordSigningParticipantOutcomeDto`

NewRecordSigningParticipantOutcomeDtoWithDefaults instantiates a new RecordSigningParticipantOutcomeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *RecordSigningParticipantOutcomeDto) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RecordSigningParticipantOutcomeDto) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RecordSigningParticipantOutcomeDto) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetOutcomeAtUtc

`func (o *RecordSigningParticipantOutcomeDto) GetOutcomeAtUtc() time.Time`

GetOutcomeAtUtc returns the OutcomeAtUtc field if non-nil, zero value otherwise.

### GetOutcomeAtUtcOk

`func (o *RecordSigningParticipantOutcomeDto) GetOutcomeAtUtcOk() (*time.Time, bool)`

GetOutcomeAtUtcOk returns a tuple with the OutcomeAtUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeAtUtc

`func (o *RecordSigningParticipantOutcomeDto) SetOutcomeAtUtc(v time.Time)`

SetOutcomeAtUtc sets OutcomeAtUtc field to given value.

### HasOutcomeAtUtc

`func (o *RecordSigningParticipantOutcomeDto) HasOutcomeAtUtc() bool`

HasOutcomeAtUtc returns a boolean if a field has been set.

### SetOutcomeAtUtcNil

`func (o *RecordSigningParticipantOutcomeDto) SetOutcomeAtUtcNil(b bool)`

 SetOutcomeAtUtcNil sets the value for OutcomeAtUtc to be an explicit nil

### UnsetOutcomeAtUtc
`func (o *RecordSigningParticipantOutcomeDto) UnsetOutcomeAtUtc()`

UnsetOutcomeAtUtc ensures that no value is present for OutcomeAtUtc, not even an explicit nil
### GetDeclineReason

`func (o *RecordSigningParticipantOutcomeDto) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *RecordSigningParticipantOutcomeDto) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *RecordSigningParticipantOutcomeDto) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *RecordSigningParticipantOutcomeDto) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### SetDeclineReasonNil

`func (o *RecordSigningParticipantOutcomeDto) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *RecordSigningParticipantOutcomeDto) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil
### GetExternalReference

`func (o *RecordSigningParticipantOutcomeDto) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *RecordSigningParticipantOutcomeDto) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *RecordSigningParticipantOutcomeDto) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *RecordSigningParticipantOutcomeDto) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### SetExternalReferenceNil

`func (o *RecordSigningParticipantOutcomeDto) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *RecordSigningParticipantOutcomeDto) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


