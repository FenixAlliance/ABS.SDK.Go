# GigApplicationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**AcceptedTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Accepted** | Pointer to **bool** |  | [optional] 
**Proposal** | Pointer to **NullableString** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**GigId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CurriculumId** | Pointer to **NullableString** |  | [optional] 
**CurriculumCoverId** | Pointer to **NullableString** |  | [optional] 
**JobApplicantProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGigApplicationDto

`func NewGigApplicationDto() *GigApplicationDto`

NewGigApplicationDto instantiates a new GigApplicationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGigApplicationDtoWithDefaults

`func NewGigApplicationDtoWithDefaults() *GigApplicationDto`

NewGigApplicationDtoWithDefaults instantiates a new GigApplicationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GigApplicationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GigApplicationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GigApplicationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GigApplicationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GigApplicationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GigApplicationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *GigApplicationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GigApplicationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GigApplicationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GigApplicationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *GigApplicationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *GigApplicationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetAcceptedTimestamp

`func (o *GigApplicationDto) GetAcceptedTimestamp() time.Time`

GetAcceptedTimestamp returns the AcceptedTimestamp field if non-nil, zero value otherwise.

### GetAcceptedTimestampOk

`func (o *GigApplicationDto) GetAcceptedTimestampOk() (*time.Time, bool)`

GetAcceptedTimestampOk returns a tuple with the AcceptedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTimestamp

`func (o *GigApplicationDto) SetAcceptedTimestamp(v time.Time)`

SetAcceptedTimestamp sets AcceptedTimestamp field to given value.

### HasAcceptedTimestamp

`func (o *GigApplicationDto) HasAcceptedTimestamp() bool`

HasAcceptedTimestamp returns a boolean if a field has been set.

### SetAcceptedTimestampNil

`func (o *GigApplicationDto) SetAcceptedTimestampNil(b bool)`

 SetAcceptedTimestampNil sets the value for AcceptedTimestamp to be an explicit nil

### UnsetAcceptedTimestamp
`func (o *GigApplicationDto) UnsetAcceptedTimestamp()`

UnsetAcceptedTimestamp ensures that no value is present for AcceptedTimestamp, not even an explicit nil
### GetAccepted

`func (o *GigApplicationDto) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *GigApplicationDto) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *GigApplicationDto) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *GigApplicationDto) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetProposal

`func (o *GigApplicationDto) GetProposal() string`

GetProposal returns the Proposal field if non-nil, zero value otherwise.

### GetProposalOk

`func (o *GigApplicationDto) GetProposalOk() (*string, bool)`

GetProposalOk returns a tuple with the Proposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposal

`func (o *GigApplicationDto) SetProposal(v string)`

SetProposal sets Proposal field to given value.

### HasProposal

`func (o *GigApplicationDto) HasProposal() bool`

HasProposal returns a boolean if a field has been set.

### SetProposalNil

`func (o *GigApplicationDto) SetProposalNil(b bool)`

 SetProposalNil sets the value for Proposal to be an explicit nil

### UnsetProposal
`func (o *GigApplicationDto) UnsetProposal()`

UnsetProposal ensures that no value is present for Proposal, not even an explicit nil
### GetCost

`func (o *GigApplicationDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GigApplicationDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GigApplicationDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GigApplicationDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetGigId

`func (o *GigApplicationDto) GetGigId() string`

GetGigId returns the GigId field if non-nil, zero value otherwise.

### GetGigIdOk

`func (o *GigApplicationDto) GetGigIdOk() (*string, bool)`

GetGigIdOk returns a tuple with the GigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGigId

`func (o *GigApplicationDto) SetGigId(v string)`

SetGigId sets GigId field to given value.

### HasGigId

`func (o *GigApplicationDto) HasGigId() bool`

HasGigId returns a boolean if a field has been set.

### SetGigIdNil

`func (o *GigApplicationDto) SetGigIdNil(b bool)`

 SetGigIdNil sets the value for GigId to be an explicit nil

### UnsetGigId
`func (o *GigApplicationDto) UnsetGigId()`

UnsetGigId ensures that no value is present for GigId, not even an explicit nil
### GetCurrencyId

`func (o *GigApplicationDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GigApplicationDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GigApplicationDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GigApplicationDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *GigApplicationDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *GigApplicationDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurriculumId

`func (o *GigApplicationDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *GigApplicationDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *GigApplicationDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *GigApplicationDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *GigApplicationDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *GigApplicationDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetCurriculumCoverId

`func (o *GigApplicationDto) GetCurriculumCoverId() string`

GetCurriculumCoverId returns the CurriculumCoverId field if non-nil, zero value otherwise.

### GetCurriculumCoverIdOk

`func (o *GigApplicationDto) GetCurriculumCoverIdOk() (*string, bool)`

GetCurriculumCoverIdOk returns a tuple with the CurriculumCoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumCoverId

`func (o *GigApplicationDto) SetCurriculumCoverId(v string)`

SetCurriculumCoverId sets CurriculumCoverId field to given value.

### HasCurriculumCoverId

`func (o *GigApplicationDto) HasCurriculumCoverId() bool`

HasCurriculumCoverId returns a boolean if a field has been set.

### SetCurriculumCoverIdNil

`func (o *GigApplicationDto) SetCurriculumCoverIdNil(b bool)`

 SetCurriculumCoverIdNil sets the value for CurriculumCoverId to be an explicit nil

### UnsetCurriculumCoverId
`func (o *GigApplicationDto) UnsetCurriculumCoverId()`

UnsetCurriculumCoverId ensures that no value is present for CurriculumCoverId, not even an explicit nil
### GetJobApplicantProfileId

`func (o *GigApplicationDto) GetJobApplicantProfileId() string`

GetJobApplicantProfileId returns the JobApplicantProfileId field if non-nil, zero value otherwise.

### GetJobApplicantProfileIdOk

`func (o *GigApplicationDto) GetJobApplicantProfileIdOk() (*string, bool)`

GetJobApplicantProfileIdOk returns a tuple with the JobApplicantProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobApplicantProfileId

`func (o *GigApplicationDto) SetJobApplicantProfileId(v string)`

SetJobApplicantProfileId sets JobApplicantProfileId field to given value.

### HasJobApplicantProfileId

`func (o *GigApplicationDto) HasJobApplicantProfileId() bool`

HasJobApplicantProfileId returns a boolean if a field has been set.

### SetJobApplicantProfileIdNil

`func (o *GigApplicationDto) SetJobApplicantProfileIdNil(b bool)`

 SetJobApplicantProfileIdNil sets the value for JobApplicantProfileId to be an explicit nil

### UnsetJobApplicantProfileId
`func (o *GigApplicationDto) UnsetJobApplicantProfileId()`

UnsetJobApplicantProfileId ensures that no value is present for JobApplicantProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


