# JobOfferApplicationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewJobOfferApplicationCreateDto

`func NewJobOfferApplicationCreateDto() *JobOfferApplicationCreateDto`

NewJobOfferApplicationCreateDto instantiates a new JobOfferApplicationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferApplicationCreateDtoWithDefaults

`func NewJobOfferApplicationCreateDtoWithDefaults() *JobOfferApplicationCreateDto`

NewJobOfferApplicationCreateDtoWithDefaults instantiates a new JobOfferApplicationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobOfferApplicationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobOfferApplicationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobOfferApplicationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobOfferApplicationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *JobOfferApplicationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *JobOfferApplicationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *JobOfferApplicationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *JobOfferApplicationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStart

`func (o *JobOfferApplicationCreateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *JobOfferApplicationCreateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *JobOfferApplicationCreateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *JobOfferApplicationCreateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *JobOfferApplicationCreateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *JobOfferApplicationCreateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *JobOfferApplicationCreateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *JobOfferApplicationCreateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetSalaryExpectation

`func (o *JobOfferApplicationCreateDto) GetSalaryExpectation() float64`

GetSalaryExpectation returns the SalaryExpectation field if non-nil, zero value otherwise.

### GetSalaryExpectationOk

`func (o *JobOfferApplicationCreateDto) GetSalaryExpectationOk() (*float64, bool)`

GetSalaryExpectationOk returns a tuple with the SalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryExpectation

`func (o *JobOfferApplicationCreateDto) SetSalaryExpectation(v float64)`

SetSalaryExpectation sets SalaryExpectation field to given value.

### HasSalaryExpectation

`func (o *JobOfferApplicationCreateDto) HasSalaryExpectation() bool`

HasSalaryExpectation returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JobOfferApplicationCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobOfferApplicationCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobOfferApplicationCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobOfferApplicationCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobOfferApplicationCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobOfferApplicationCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurriculumId

`func (o *JobOfferApplicationCreateDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *JobOfferApplicationCreateDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *JobOfferApplicationCreateDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *JobOfferApplicationCreateDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *JobOfferApplicationCreateDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *JobOfferApplicationCreateDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetCurriculumCoverId

`func (o *JobOfferApplicationCreateDto) GetCurriculumCoverId() string`

GetCurriculumCoverId returns the CurriculumCoverId field if non-nil, zero value otherwise.

### GetCurriculumCoverIdOk

`func (o *JobOfferApplicationCreateDto) GetCurriculumCoverIdOk() (*string, bool)`

GetCurriculumCoverIdOk returns a tuple with the CurriculumCoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumCoverId

`func (o *JobOfferApplicationCreateDto) SetCurriculumCoverId(v string)`

SetCurriculumCoverId sets CurriculumCoverId field to given value.

### HasCurriculumCoverId

`func (o *JobOfferApplicationCreateDto) HasCurriculumCoverId() bool`

HasCurriculumCoverId returns a boolean if a field has been set.

### SetCurriculumCoverIdNil

`func (o *JobOfferApplicationCreateDto) SetCurriculumCoverIdNil(b bool)`

 SetCurriculumCoverIdNil sets the value for CurriculumCoverId to be an explicit nil

### UnsetCurriculumCoverId
`func (o *JobOfferApplicationCreateDto) UnsetCurriculumCoverId()`

UnsetCurriculumCoverId ensures that no value is present for CurriculumCoverId, not even an explicit nil
### GetJobOfferId

`func (o *JobOfferApplicationCreateDto) GetJobOfferId() string`

GetJobOfferId returns the JobOfferId field if non-nil, zero value otherwise.

### GetJobOfferIdOk

`func (o *JobOfferApplicationCreateDto) GetJobOfferIdOk() (*string, bool)`

GetJobOfferIdOk returns a tuple with the JobOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobOfferId

`func (o *JobOfferApplicationCreateDto) SetJobOfferId(v string)`

SetJobOfferId sets JobOfferId field to given value.

### HasJobOfferId

`func (o *JobOfferApplicationCreateDto) HasJobOfferId() bool`

HasJobOfferId returns a boolean if a field has been set.

### SetJobOfferIdNil

`func (o *JobOfferApplicationCreateDto) SetJobOfferIdNil(b bool)`

 SetJobOfferIdNil sets the value for JobOfferId to be an explicit nil

### UnsetJobOfferId
`func (o *JobOfferApplicationCreateDto) UnsetJobOfferId()`

UnsetJobOfferId ensures that no value is present for JobOfferId, not even an explicit nil
### GetPartnerProfileId

`func (o *JobOfferApplicationCreateDto) GetPartnerProfileId() string`

GetPartnerProfileId returns the PartnerProfileId field if non-nil, zero value otherwise.

### GetPartnerProfileIdOk

`func (o *JobOfferApplicationCreateDto) GetPartnerProfileIdOk() (*string, bool)`

GetPartnerProfileIdOk returns a tuple with the PartnerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProfileId

`func (o *JobOfferApplicationCreateDto) SetPartnerProfileId(v string)`

SetPartnerProfileId sets PartnerProfileId field to given value.

### HasPartnerProfileId

`func (o *JobOfferApplicationCreateDto) HasPartnerProfileId() bool`

HasPartnerProfileId returns a boolean if a field has been set.

### SetPartnerProfileIdNil

`func (o *JobOfferApplicationCreateDto) SetPartnerProfileIdNil(b bool)`

 SetPartnerProfileIdNil sets the value for PartnerProfileId to be an explicit nil

### UnsetPartnerProfileId
`func (o *JobOfferApplicationCreateDto) UnsetPartnerProfileId()`

UnsetPartnerProfileId ensures that no value is present for PartnerProfileId, not even an explicit nil
### GetJobApplicantProfileId

`func (o *JobOfferApplicationCreateDto) GetJobApplicantProfileId() string`

GetJobApplicantProfileId returns the JobApplicantProfileId field if non-nil, zero value otherwise.

### GetJobApplicantProfileIdOk

`func (o *JobOfferApplicationCreateDto) GetJobApplicantProfileIdOk() (*string, bool)`

GetJobApplicantProfileIdOk returns a tuple with the JobApplicantProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobApplicantProfileId

`func (o *JobOfferApplicationCreateDto) SetJobApplicantProfileId(v string)`

SetJobApplicantProfileId sets JobApplicantProfileId field to given value.

### HasJobApplicantProfileId

`func (o *JobOfferApplicationCreateDto) HasJobApplicantProfileId() bool`

HasJobApplicantProfileId returns a boolean if a field has been set.

### SetJobApplicantProfileIdNil

`func (o *JobOfferApplicationCreateDto) SetJobApplicantProfileIdNil(b bool)`

 SetJobApplicantProfileIdNil sets the value for JobApplicantProfileId to be an explicit nil

### UnsetJobApplicantProfileId
`func (o *JobOfferApplicationCreateDto) UnsetJobApplicantProfileId()`

UnsetJobApplicantProfileId ensures that no value is present for JobApplicantProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


