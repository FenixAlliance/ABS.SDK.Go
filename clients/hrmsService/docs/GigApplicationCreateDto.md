# GigApplicationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Proposal** | Pointer to **NullableString** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**GigId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CurriculumId** | Pointer to **NullableString** |  | [optional] 
**CurriculumCoverId** | Pointer to **NullableString** |  | [optional] 
**JobApplicantProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGigApplicationCreateDto

`func NewGigApplicationCreateDto() *GigApplicationCreateDto`

NewGigApplicationCreateDto instantiates a new GigApplicationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGigApplicationCreateDtoWithDefaults

`func NewGigApplicationCreateDtoWithDefaults() *GigApplicationCreateDto`

NewGigApplicationCreateDtoWithDefaults instantiates a new GigApplicationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GigApplicationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GigApplicationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GigApplicationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GigApplicationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *GigApplicationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GigApplicationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GigApplicationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GigApplicationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetProposal

`func (o *GigApplicationCreateDto) GetProposal() string`

GetProposal returns the Proposal field if non-nil, zero value otherwise.

### GetProposalOk

`func (o *GigApplicationCreateDto) GetProposalOk() (*string, bool)`

GetProposalOk returns a tuple with the Proposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposal

`func (o *GigApplicationCreateDto) SetProposal(v string)`

SetProposal sets Proposal field to given value.

### HasProposal

`func (o *GigApplicationCreateDto) HasProposal() bool`

HasProposal returns a boolean if a field has been set.

### SetProposalNil

`func (o *GigApplicationCreateDto) SetProposalNil(b bool)`

 SetProposalNil sets the value for Proposal to be an explicit nil

### UnsetProposal
`func (o *GigApplicationCreateDto) UnsetProposal()`

UnsetProposal ensures that no value is present for Proposal, not even an explicit nil
### GetCost

`func (o *GigApplicationCreateDto) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GigApplicationCreateDto) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GigApplicationCreateDto) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GigApplicationCreateDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetGigId

`func (o *GigApplicationCreateDto) GetGigId() string`

GetGigId returns the GigId field if non-nil, zero value otherwise.

### GetGigIdOk

`func (o *GigApplicationCreateDto) GetGigIdOk() (*string, bool)`

GetGigIdOk returns a tuple with the GigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGigId

`func (o *GigApplicationCreateDto) SetGigId(v string)`

SetGigId sets GigId field to given value.

### HasGigId

`func (o *GigApplicationCreateDto) HasGigId() bool`

HasGigId returns a boolean if a field has been set.

### SetGigIdNil

`func (o *GigApplicationCreateDto) SetGigIdNil(b bool)`

 SetGigIdNil sets the value for GigId to be an explicit nil

### UnsetGigId
`func (o *GigApplicationCreateDto) UnsetGigId()`

UnsetGigId ensures that no value is present for GigId, not even an explicit nil
### GetCurrencyId

`func (o *GigApplicationCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *GigApplicationCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *GigApplicationCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *GigApplicationCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *GigApplicationCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *GigApplicationCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurriculumId

`func (o *GigApplicationCreateDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *GigApplicationCreateDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *GigApplicationCreateDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *GigApplicationCreateDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *GigApplicationCreateDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *GigApplicationCreateDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetCurriculumCoverId

`func (o *GigApplicationCreateDto) GetCurriculumCoverId() string`

GetCurriculumCoverId returns the CurriculumCoverId field if non-nil, zero value otherwise.

### GetCurriculumCoverIdOk

`func (o *GigApplicationCreateDto) GetCurriculumCoverIdOk() (*string, bool)`

GetCurriculumCoverIdOk returns a tuple with the CurriculumCoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumCoverId

`func (o *GigApplicationCreateDto) SetCurriculumCoverId(v string)`

SetCurriculumCoverId sets CurriculumCoverId field to given value.

### HasCurriculumCoverId

`func (o *GigApplicationCreateDto) HasCurriculumCoverId() bool`

HasCurriculumCoverId returns a boolean if a field has been set.

### SetCurriculumCoverIdNil

`func (o *GigApplicationCreateDto) SetCurriculumCoverIdNil(b bool)`

 SetCurriculumCoverIdNil sets the value for CurriculumCoverId to be an explicit nil

### UnsetCurriculumCoverId
`func (o *GigApplicationCreateDto) UnsetCurriculumCoverId()`

UnsetCurriculumCoverId ensures that no value is present for CurriculumCoverId, not even an explicit nil
### GetJobApplicantProfileId

`func (o *GigApplicationCreateDto) GetJobApplicantProfileId() string`

GetJobApplicantProfileId returns the JobApplicantProfileId field if non-nil, zero value otherwise.

### GetJobApplicantProfileIdOk

`func (o *GigApplicationCreateDto) GetJobApplicantProfileIdOk() (*string, bool)`

GetJobApplicantProfileIdOk returns a tuple with the JobApplicantProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobApplicantProfileId

`func (o *GigApplicationCreateDto) SetJobApplicantProfileId(v string)`

SetJobApplicantProfileId sets JobApplicantProfileId field to given value.

### HasJobApplicantProfileId

`func (o *GigApplicationCreateDto) HasJobApplicantProfileId() bool`

HasJobApplicantProfileId returns a boolean if a field has been set.

### SetJobApplicantProfileIdNil

`func (o *GigApplicationCreateDto) SetJobApplicantProfileIdNil(b bool)`

 SetJobApplicantProfileIdNil sets the value for JobApplicantProfileId to be an explicit nil

### UnsetJobApplicantProfileId
`func (o *GigApplicationCreateDto) UnsetJobApplicantProfileId()`

UnsetJobApplicantProfileId ensures that no value is present for JobApplicantProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


