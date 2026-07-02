# JobOfferApplicationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**SalaryExpectation** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CurriculumId** | Pointer to **NullableString** |  | [optional] 
**CurriculumCoverId** | Pointer to **NullableString** |  | [optional] 
**JobOfferId** | Pointer to **NullableString** |  | [optional] 
**PartnerProfileId** | Pointer to **NullableString** |  | [optional] 
**JobApplicantProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobOfferApplicationDto

`func NewJobOfferApplicationDto() *JobOfferApplicationDto`

NewJobOfferApplicationDto instantiates a new JobOfferApplicationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferApplicationDtoWithDefaults

`func NewJobOfferApplicationDtoWithDefaults() *JobOfferApplicationDto`

NewJobOfferApplicationDtoWithDefaults instantiates a new JobOfferApplicationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobOfferApplicationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobOfferApplicationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobOfferApplicationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobOfferApplicationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JobOfferApplicationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JobOfferApplicationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *JobOfferApplicationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobOfferApplicationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobOfferApplicationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobOfferApplicationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *JobOfferApplicationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *JobOfferApplicationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetStatus

`func (o *JobOfferApplicationDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobOfferApplicationDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobOfferApplicationDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobOfferApplicationDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStart

`func (o *JobOfferApplicationDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *JobOfferApplicationDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *JobOfferApplicationDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *JobOfferApplicationDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *JobOfferApplicationDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *JobOfferApplicationDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *JobOfferApplicationDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *JobOfferApplicationDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetSalaryExpectation

`func (o *JobOfferApplicationDto) GetSalaryExpectation() float64`

GetSalaryExpectation returns the SalaryExpectation field if non-nil, zero value otherwise.

### GetSalaryExpectationOk

`func (o *JobOfferApplicationDto) GetSalaryExpectationOk() (*float64, bool)`

GetSalaryExpectationOk returns a tuple with the SalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryExpectation

`func (o *JobOfferApplicationDto) SetSalaryExpectation(v float64)`

SetSalaryExpectation sets SalaryExpectation field to given value.

### HasSalaryExpectation

`func (o *JobOfferApplicationDto) HasSalaryExpectation() bool`

HasSalaryExpectation returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JobOfferApplicationDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobOfferApplicationDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobOfferApplicationDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobOfferApplicationDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobOfferApplicationDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobOfferApplicationDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurriculumId

`func (o *JobOfferApplicationDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *JobOfferApplicationDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *JobOfferApplicationDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *JobOfferApplicationDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *JobOfferApplicationDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *JobOfferApplicationDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetCurriculumCoverId

`func (o *JobOfferApplicationDto) GetCurriculumCoverId() string`

GetCurriculumCoverId returns the CurriculumCoverId field if non-nil, zero value otherwise.

### GetCurriculumCoverIdOk

`func (o *JobOfferApplicationDto) GetCurriculumCoverIdOk() (*string, bool)`

GetCurriculumCoverIdOk returns a tuple with the CurriculumCoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumCoverId

`func (o *JobOfferApplicationDto) SetCurriculumCoverId(v string)`

SetCurriculumCoverId sets CurriculumCoverId field to given value.

### HasCurriculumCoverId

`func (o *JobOfferApplicationDto) HasCurriculumCoverId() bool`

HasCurriculumCoverId returns a boolean if a field has been set.

### SetCurriculumCoverIdNil

`func (o *JobOfferApplicationDto) SetCurriculumCoverIdNil(b bool)`

 SetCurriculumCoverIdNil sets the value for CurriculumCoverId to be an explicit nil

### UnsetCurriculumCoverId
`func (o *JobOfferApplicationDto) UnsetCurriculumCoverId()`

UnsetCurriculumCoverId ensures that no value is present for CurriculumCoverId, not even an explicit nil
### GetJobOfferId

`func (o *JobOfferApplicationDto) GetJobOfferId() string`

GetJobOfferId returns the JobOfferId field if non-nil, zero value otherwise.

### GetJobOfferIdOk

`func (o *JobOfferApplicationDto) GetJobOfferIdOk() (*string, bool)`

GetJobOfferIdOk returns a tuple with the JobOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOfferId

`func (o *JobOfferApplicationDto) SetJobOfferId(v string)`

SetJobOfferId sets JobOfferId field to given value.

### HasJobOfferId

`func (o *JobOfferApplicationDto) HasJobOfferId() bool`

HasJobOfferId returns a boolean if a field has been set.

### SetJobOfferIdNil

`func (o *JobOfferApplicationDto) SetJobOfferIdNil(b bool)`

 SetJobOfferIdNil sets the value for JobOfferId to be an explicit nil

### UnsetJobOfferId
`func (o *JobOfferApplicationDto) UnsetJobOfferId()`

UnsetJobOfferId ensures that no value is present for JobOfferId, not even an explicit nil
### GetPartnerProfileId

`func (o *JobOfferApplicationDto) GetPartnerProfileId() string`

GetPartnerProfileId returns the PartnerProfileId field if non-nil, zero value otherwise.

### GetPartnerProfileIdOk

`func (o *JobOfferApplicationDto) GetPartnerProfileIdOk() (*string, bool)`

GetPartnerProfileIdOk returns a tuple with the PartnerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProfileId

`func (o *JobOfferApplicationDto) SetPartnerProfileId(v string)`

SetPartnerProfileId sets PartnerProfileId field to given value.

### HasPartnerProfileId

`func (o *JobOfferApplicationDto) HasPartnerProfileId() bool`

HasPartnerProfileId returns a boolean if a field has been set.

### SetPartnerProfileIdNil

`func (o *JobOfferApplicationDto) SetPartnerProfileIdNil(b bool)`

 SetPartnerProfileIdNil sets the value for PartnerProfileId to be an explicit nil

### UnsetPartnerProfileId
`func (o *JobOfferApplicationDto) UnsetPartnerProfileId()`

UnsetPartnerProfileId ensures that no value is present for PartnerProfileId, not even an explicit nil
### GetJobApplicantProfileId

`func (o *JobOfferApplicationDto) GetJobApplicantProfileId() string`

GetJobApplicantProfileId returns the JobApplicantProfileId field if non-nil, zero value otherwise.

### GetJobApplicantProfileIdOk

`func (o *JobOfferApplicationDto) GetJobApplicantProfileIdOk() (*string, bool)`

GetJobApplicantProfileIdOk returns a tuple with the JobApplicantProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobApplicantProfileId

`func (o *JobOfferApplicationDto) SetJobApplicantProfileId(v string)`

SetJobApplicantProfileId sets JobApplicantProfileId field to given value.

### HasJobApplicantProfileId

`func (o *JobOfferApplicationDto) HasJobApplicantProfileId() bool`

HasJobApplicantProfileId returns a boolean if a field has been set.

### SetJobApplicantProfileIdNil

`func (o *JobOfferApplicationDto) SetJobApplicantProfileIdNil(b bool)`

 SetJobApplicantProfileIdNil sets the value for JobApplicantProfileId to be an explicit nil

### UnsetJobApplicantProfileId
`func (o *JobOfferApplicationDto) UnsetJobApplicantProfileId()`

UnsetJobApplicantProfileId ensures that no value is present for JobApplicantProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


