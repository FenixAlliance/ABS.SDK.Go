# CourseCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**CourseCategoryId** | Pointer to **NullableString** |  | [optional] 
**InstructorProfileId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**RegularPrice** | Pointer to **float64** |  | [optional] 
**MaxCourseEnrollments** | Pointer to **int32** |  | [optional] 
**TotalEffortInWeeks** | Pointer to **int32** |  | [optional] 
**TotalHoursPerWeek** | Pointer to **int32** |  | [optional] 
**TotalEffortInHours** | Pointer to **int32** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**InscriptionsStartDateTime** | Pointer to **NullableTime** |  | [optional] 
**InscriptionsEndDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCourseCreateDto

`func NewCourseCreateDto(title string, description string, ) *CourseCreateDto`

NewCourseCreateDto instantiates a new CourseCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseCreateDtoWithDefaults

`func NewCourseCreateDtoWithDefaults() *CourseCreateDto`

NewCourseCreateDtoWithDefaults instantiates a new CourseCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CourseCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CourseCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSku

`func (o *CourseCreateDto) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CourseCreateDto) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CourseCreateDto) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CourseCreateDto) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CourseCreateDto) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CourseCreateDto) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetSummary

`func (o *CourseCreateDto) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CourseCreateDto) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CourseCreateDto) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CourseCreateDto) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *CourseCreateDto) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *CourseCreateDto) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetCode

`func (o *CourseCreateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CourseCreateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CourseCreateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CourseCreateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CourseCreateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CourseCreateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetVersion

`func (o *CourseCreateDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CourseCreateDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CourseCreateDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CourseCreateDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CourseCreateDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CourseCreateDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCourseCategoryId

`func (o *CourseCreateDto) GetCourseCategoryId() string`

GetCourseCategoryId returns the CourseCategoryId field if non-nil, zero value otherwise.

### GetCourseCategoryIdOk

`func (o *CourseCreateDto) GetCourseCategoryIdOk() (*string, bool)`

GetCourseCategoryIdOk returns a tuple with the CourseCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCategoryId

`func (o *CourseCreateDto) SetCourseCategoryId(v string)`

SetCourseCategoryId sets CourseCategoryId field to given value.

### HasCourseCategoryId

`func (o *CourseCreateDto) HasCourseCategoryId() bool`

HasCourseCategoryId returns a boolean if a field has been set.

### SetCourseCategoryIdNil

`func (o *CourseCreateDto) SetCourseCategoryIdNil(b bool)`

 SetCourseCategoryIdNil sets the value for CourseCategoryId to be an explicit nil

### UnsetCourseCategoryId
`func (o *CourseCreateDto) UnsetCourseCategoryId()`

UnsetCourseCategoryId ensures that no value is present for CourseCategoryId, not even an explicit nil
### GetInstructorProfileId

`func (o *CourseCreateDto) GetInstructorProfileId() string`

GetInstructorProfileId returns the InstructorProfileId field if non-nil, zero value otherwise.

### GetInstructorProfileIdOk

`func (o *CourseCreateDto) GetInstructorProfileIdOk() (*string, bool)`

GetInstructorProfileIdOk returns a tuple with the InstructorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorProfileId

`func (o *CourseCreateDto) SetInstructorProfileId(v string)`

SetInstructorProfileId sets InstructorProfileId field to given value.

### HasInstructorProfileId

`func (o *CourseCreateDto) HasInstructorProfileId() bool`

HasInstructorProfileId returns a boolean if a field has been set.

### SetInstructorProfileIdNil

`func (o *CourseCreateDto) SetInstructorProfileIdNil(b bool)`

 SetInstructorProfileIdNil sets the value for InstructorProfileId to be an explicit nil

### UnsetInstructorProfileId
`func (o *CourseCreateDto) UnsetInstructorProfileId()`

UnsetInstructorProfileId ensures that no value is present for InstructorProfileId, not even an explicit nil
### GetCurrencyId

`func (o *CourseCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CourseCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CourseCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CourseCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *CourseCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *CourseCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetRegularPrice

`func (o *CourseCreateDto) GetRegularPrice() float64`

GetRegularPrice returns the RegularPrice field if non-nil, zero value otherwise.

### GetRegularPriceOk

`func (o *CourseCreateDto) GetRegularPriceOk() (*float64, bool)`

GetRegularPriceOk returns a tuple with the RegularPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPrice

`func (o *CourseCreateDto) SetRegularPrice(v float64)`

SetRegularPrice sets RegularPrice field to given value.

### HasRegularPrice

`func (o *CourseCreateDto) HasRegularPrice() bool`

HasRegularPrice returns a boolean if a field has been set.

### GetMaxCourseEnrollments

`func (o *CourseCreateDto) GetMaxCourseEnrollments() int32`

GetMaxCourseEnrollments returns the MaxCourseEnrollments field if non-nil, zero value otherwise.

### GetMaxCourseEnrollmentsOk

`func (o *CourseCreateDto) GetMaxCourseEnrollmentsOk() (*int32, bool)`

GetMaxCourseEnrollmentsOk returns a tuple with the MaxCourseEnrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCourseEnrollments

`func (o *CourseCreateDto) SetMaxCourseEnrollments(v int32)`

SetMaxCourseEnrollments sets MaxCourseEnrollments field to given value.

### HasMaxCourseEnrollments

`func (o *CourseCreateDto) HasMaxCourseEnrollments() bool`

HasMaxCourseEnrollments returns a boolean if a field has been set.

### GetTotalEffortInWeeks

`func (o *CourseCreateDto) GetTotalEffortInWeeks() int32`

GetTotalEffortInWeeks returns the TotalEffortInWeeks field if non-nil, zero value otherwise.

### GetTotalEffortInWeeksOk

`func (o *CourseCreateDto) GetTotalEffortInWeeksOk() (*int32, bool)`

GetTotalEffortInWeeksOk returns a tuple with the TotalEffortInWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEffortInWeeks

`func (o *CourseCreateDto) SetTotalEffortInWeeks(v int32)`

SetTotalEffortInWeeks sets TotalEffortInWeeks field to given value.

### HasTotalEffortInWeeks

`func (o *CourseCreateDto) HasTotalEffortInWeeks() bool`

HasTotalEffortInWeeks returns a boolean if a field has been set.

### GetTotalHoursPerWeek

`func (o *CourseCreateDto) GetTotalHoursPerWeek() int32`

GetTotalHoursPerWeek returns the TotalHoursPerWeek field if non-nil, zero value otherwise.

### GetTotalHoursPerWeekOk

`func (o *CourseCreateDto) GetTotalHoursPerWeekOk() (*int32, bool)`

GetTotalHoursPerWeekOk returns a tuple with the TotalHoursPerWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHoursPerWeek

`func (o *CourseCreateDto) SetTotalHoursPerWeek(v int32)`

SetTotalHoursPerWeek sets TotalHoursPerWeek field to given value.

### HasTotalHoursPerWeek

`func (o *CourseCreateDto) HasTotalHoursPerWeek() bool`

HasTotalHoursPerWeek returns a boolean if a field has been set.

### GetTotalEffortInHours

`func (o *CourseCreateDto) GetTotalEffortInHours() int32`

GetTotalEffortInHours returns the TotalEffortInHours field if non-nil, zero value otherwise.

### GetTotalEffortInHoursOk

`func (o *CourseCreateDto) GetTotalEffortInHoursOk() (*int32, bool)`

GetTotalEffortInHoursOk returns a tuple with the TotalEffortInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEffortInHours

`func (o *CourseCreateDto) SetTotalEffortInHours(v int32)`

SetTotalEffortInHours sets TotalEffortInHours field to given value.

### HasTotalEffortInHours

`func (o *CourseCreateDto) HasTotalEffortInHours() bool`

HasTotalEffortInHours returns a boolean if a field has been set.

### GetStartDateTime

`func (o *CourseCreateDto) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CourseCreateDto) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CourseCreateDto) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CourseCreateDto) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *CourseCreateDto) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *CourseCreateDto) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetEndDateTime

`func (o *CourseCreateDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CourseCreateDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CourseCreateDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CourseCreateDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *CourseCreateDto) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *CourseCreateDto) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetInscriptionsStartDateTime

`func (o *CourseCreateDto) GetInscriptionsStartDateTime() time.Time`

GetInscriptionsStartDateTime returns the InscriptionsStartDateTime field if non-nil, zero value otherwise.

### GetInscriptionsStartDateTimeOk

`func (o *CourseCreateDto) GetInscriptionsStartDateTimeOk() (*time.Time, bool)`

GetInscriptionsStartDateTimeOk returns a tuple with the InscriptionsStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptionsStartDateTime

`func (o *CourseCreateDto) SetInscriptionsStartDateTime(v time.Time)`

SetInscriptionsStartDateTime sets InscriptionsStartDateTime field to given value.

### HasInscriptionsStartDateTime

`func (o *CourseCreateDto) HasInscriptionsStartDateTime() bool`

HasInscriptionsStartDateTime returns a boolean if a field has been set.

### SetInscriptionsStartDateTimeNil

`func (o *CourseCreateDto) SetInscriptionsStartDateTimeNil(b bool)`

 SetInscriptionsStartDateTimeNil sets the value for InscriptionsStartDateTime to be an explicit nil

### UnsetInscriptionsStartDateTime
`func (o *CourseCreateDto) UnsetInscriptionsStartDateTime()`

UnsetInscriptionsStartDateTime ensures that no value is present for InscriptionsStartDateTime, not even an explicit nil
### GetInscriptionsEndDateTime

`func (o *CourseCreateDto) GetInscriptionsEndDateTime() time.Time`

GetInscriptionsEndDateTime returns the InscriptionsEndDateTime field if non-nil, zero value otherwise.

### GetInscriptionsEndDateTimeOk

`func (o *CourseCreateDto) GetInscriptionsEndDateTimeOk() (*time.Time, bool)`

GetInscriptionsEndDateTimeOk returns a tuple with the InscriptionsEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptionsEndDateTime

`func (o *CourseCreateDto) SetInscriptionsEndDateTime(v time.Time)`

SetInscriptionsEndDateTime sets InscriptionsEndDateTime field to given value.

### HasInscriptionsEndDateTime

`func (o *CourseCreateDto) HasInscriptionsEndDateTime() bool`

HasInscriptionsEndDateTime returns a boolean if a field has been set.

### SetInscriptionsEndDateTimeNil

`func (o *CourseCreateDto) SetInscriptionsEndDateTimeNil(b bool)`

 SetInscriptionsEndDateTimeNil sets the value for InscriptionsEndDateTime to be an explicit nil

### UnsetInscriptionsEndDateTime
`func (o *CourseCreateDto) UnsetInscriptionsEndDateTime()`

UnsetInscriptionsEndDateTime ensures that no value is present for InscriptionsEndDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


