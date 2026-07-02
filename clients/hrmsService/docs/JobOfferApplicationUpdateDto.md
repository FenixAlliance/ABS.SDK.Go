# JobOfferApplicationUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**SalaryExpectation** | Pointer to **float64** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CurriculumId** | Pointer to **NullableString** |  | [optional] 
**CurriculumCoverId** | Pointer to **NullableString** |  | [optional] 
**PartnerProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJobOfferApplicationUpdateDto

`func NewJobOfferApplicationUpdateDto() *JobOfferApplicationUpdateDto`

NewJobOfferApplicationUpdateDto instantiates a new JobOfferApplicationUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOfferApplicationUpdateDtoWithDefaults

`func NewJobOfferApplicationUpdateDtoWithDefaults() *JobOfferApplicationUpdateDto`

NewJobOfferApplicationUpdateDtoWithDefaults instantiates a new JobOfferApplicationUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *JobOfferApplicationUpdateDto) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *JobOfferApplicationUpdateDto) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *JobOfferApplicationUpdateDto) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *JobOfferApplicationUpdateDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *JobOfferApplicationUpdateDto) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *JobOfferApplicationUpdateDto) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *JobOfferApplicationUpdateDto) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *JobOfferApplicationUpdateDto) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetSalaryExpectation

`func (o *JobOfferApplicationUpdateDto) GetSalaryExpectation() float64`

GetSalaryExpectation returns the SalaryExpectation field if non-nil, zero value otherwise.

### GetSalaryExpectationOk

`func (o *JobOfferApplicationUpdateDto) GetSalaryExpectationOk() (*float64, bool)`

GetSalaryExpectationOk returns a tuple with the SalaryExpectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryExpectation

`func (o *JobOfferApplicationUpdateDto) SetSalaryExpectation(v float64)`

SetSalaryExpectation sets SalaryExpectation field to given value.

### HasSalaryExpectation

`func (o *JobOfferApplicationUpdateDto) HasSalaryExpectation() bool`

HasSalaryExpectation returns a boolean if a field has been set.

### GetCurrencyId

`func (o *JobOfferApplicationUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *JobOfferApplicationUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *JobOfferApplicationUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *JobOfferApplicationUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *JobOfferApplicationUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *JobOfferApplicationUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurriculumId

`func (o *JobOfferApplicationUpdateDto) GetCurriculumId() string`

GetCurriculumId returns the CurriculumId field if non-nil, zero value otherwise.

### GetCurriculumIdOk

`func (o *JobOfferApplicationUpdateDto) GetCurriculumIdOk() (*string, bool)`

GetCurriculumIdOk returns a tuple with the CurriculumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumId

`func (o *JobOfferApplicationUpdateDto) SetCurriculumId(v string)`

SetCurriculumId sets CurriculumId field to given value.

### HasCurriculumId

`func (o *JobOfferApplicationUpdateDto) HasCurriculumId() bool`

HasCurriculumId returns a boolean if a field has been set.

### SetCurriculumIdNil

`func (o *JobOfferApplicationUpdateDto) SetCurriculumIdNil(b bool)`

 SetCurriculumIdNil sets the value for CurriculumId to be an explicit nil

### UnsetCurriculumId
`func (o *JobOfferApplicationUpdateDto) UnsetCurriculumId()`

UnsetCurriculumId ensures that no value is present for CurriculumId, not even an explicit nil
### GetCurriculumCoverId

`func (o *JobOfferApplicationUpdateDto) GetCurriculumCoverId() string`

GetCurriculumCoverId returns the CurriculumCoverId field if non-nil, zero value otherwise.

### GetCurriculumCoverIdOk

`func (o *JobOfferApplicationUpdateDto) GetCurriculumCoverIdOk() (*string, bool)`

GetCurriculumCoverIdOk returns a tuple with the CurriculumCoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurriculumCoverId

`func (o *JobOfferApplicationUpdateDto) SetCurriculumCoverId(v string)`

SetCurriculumCoverId sets CurriculumCoverId field to given value.

### HasCurriculumCoverId

`func (o *JobOfferApplicationUpdateDto) HasCurriculumCoverId() bool`

HasCurriculumCoverId returns a boolean if a field has been set.

### SetCurriculumCoverIdNil

`func (o *JobOfferApplicationUpdateDto) SetCurriculumCoverIdNil(b bool)`

 SetCurriculumCoverIdNil sets the value for CurriculumCoverId to be an explicit nil

### UnsetCurriculumCoverId
`func (o *JobOfferApplicationUpdateDto) UnsetCurriculumCoverId()`

UnsetCurriculumCoverId ensures that no value is present for CurriculumCoverId, not even an explicit nil
### GetPartnerProfileId

`func (o *JobOfferApplicationUpdateDto) GetPartnerProfileId() string`

GetPartnerProfileId returns the PartnerProfileId field if non-nil, zero value otherwise.

### GetPartnerProfileIdOk

`func (o *JobOfferApplicationUpdateDto) GetPartnerProfileIdOk() (*string, bool)`

GetPartnerProfileIdOk returns a tuple with the PartnerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProfileId

`func (o *JobOfferApplicationUpdateDto) SetPartnerProfileId(v string)`

SetPartnerProfileId sets PartnerProfileId field to given value.

### HasPartnerProfileId

`func (o *JobOfferApplicationUpdateDto) HasPartnerProfileId() bool`

HasPartnerProfileId returns a boolean if a field has been set.

### SetPartnerProfileIdNil

`func (o *JobOfferApplicationUpdateDto) SetPartnerProfileIdNil(b bool)`

 SetPartnerProfileIdNil sets the value for PartnerProfileId to be an explicit nil

### UnsetPartnerProfileId
`func (o *JobOfferApplicationUpdateDto) UnsetPartnerProfileId()`

UnsetPartnerProfileId ensures that no value is present for PartnerProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


