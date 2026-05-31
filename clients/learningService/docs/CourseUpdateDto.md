# CourseUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**CourseCategoryID** | Pointer to **NullableString** |  | [optional] 
**InstructorProfileID** | Pointer to **NullableString** |  | [optional] 
**CurrencyID** | Pointer to **NullableString** |  | [optional] 
**RegularPrice** | Pointer to **NullableFloat64** |  | [optional] 
**MaxCourseEnrollments** | Pointer to **NullableInt32** |  | [optional] 
**TotalEffortInWeeks** | Pointer to **NullableInt32** |  | [optional] 
**TotalHoursPerWeek** | Pointer to **NullableInt32** |  | [optional] 
**TotalEffortInHours** | Pointer to **NullableInt32** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**InscriptionsStartDateTime** | Pointer to **NullableTime** |  | [optional] 
**InscriptionsEndDateTime** | Pointer to **NullableTime** |  | [optional] 
**Published** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCourseUpdateDto

`func NewCourseUpdateDto() *CourseUpdateDto`

NewCourseUpdateDto instantiates a new CourseUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseUpdateDtoWithDefaults

`func NewCourseUpdateDtoWithDefaults() *CourseUpdateDto`

NewCourseUpdateDtoWithDefaults instantiates a new CourseUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CourseUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSku

`func (o *CourseUpdateDto) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CourseUpdateDto) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CourseUpdateDto) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CourseUpdateDto) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CourseUpdateDto) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CourseUpdateDto) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetSummary

`func (o *CourseUpdateDto) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CourseUpdateDto) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CourseUpdateDto) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CourseUpdateDto) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *CourseUpdateDto) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *CourseUpdateDto) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetCode

`func (o *CourseUpdateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CourseUpdateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CourseUpdateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CourseUpdateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CourseUpdateDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CourseUpdateDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetVersion

`func (o *CourseUpdateDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CourseUpdateDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CourseUpdateDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CourseUpdateDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CourseUpdateDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CourseUpdateDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCourseCategoryID

`func (o *CourseUpdateDto) GetCourseCategoryID() string`

GetCourseCategoryID returns the CourseCategoryID field if non-nil, zero value otherwise.

### GetCourseCategoryIDOk

`func (o *CourseUpdateDto) GetCourseCategoryIDOk() (*string, bool)`

GetCourseCategoryIDOk returns a tuple with the CourseCategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseCategoryID

`func (o *CourseUpdateDto) SetCourseCategoryID(v string)`

SetCourseCategoryID sets CourseCategoryID field to given value.

### HasCourseCategoryID

`func (o *CourseUpdateDto) HasCourseCategoryID() bool`

HasCourseCategoryID returns a boolean if a field has been set.

### SetCourseCategoryIDNil

`func (o *CourseUpdateDto) SetCourseCategoryIDNil(b bool)`

 SetCourseCategoryIDNil sets the value for CourseCategoryID to be an explicit nil

### UnsetCourseCategoryID
`func (o *CourseUpdateDto) UnsetCourseCategoryID()`

UnsetCourseCategoryID ensures that no value is present for CourseCategoryID, not even an explicit nil
### GetInstructorProfileID

`func (o *CourseUpdateDto) GetInstructorProfileID() string`

GetInstructorProfileID returns the InstructorProfileID field if non-nil, zero value otherwise.

### GetInstructorProfileIDOk

`func (o *CourseUpdateDto) GetInstructorProfileIDOk() (*string, bool)`

GetInstructorProfileIDOk returns a tuple with the InstructorProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructorProfileID

`func (o *CourseUpdateDto) SetInstructorProfileID(v string)`

SetInstructorProfileID sets InstructorProfileID field to given value.

### HasInstructorProfileID

`func (o *CourseUpdateDto) HasInstructorProfileID() bool`

HasInstructorProfileID returns a boolean if a field has been set.

### SetInstructorProfileIDNil

`func (o *CourseUpdateDto) SetInstructorProfileIDNil(b bool)`

 SetInstructorProfileIDNil sets the value for InstructorProfileID to be an explicit nil

### UnsetInstructorProfileID
`func (o *CourseUpdateDto) UnsetInstructorProfileID()`

UnsetInstructorProfileID ensures that no value is present for InstructorProfileID, not even an explicit nil
### GetCurrencyID

`func (o *CourseUpdateDto) GetCurrencyID() string`

GetCurrencyID returns the CurrencyID field if non-nil, zero value otherwise.

### GetCurrencyIDOk

`func (o *CourseUpdateDto) GetCurrencyIDOk() (*string, bool)`

GetCurrencyIDOk returns a tuple with the CurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyID

`func (o *CourseUpdateDto) SetCurrencyID(v string)`

SetCurrencyID sets CurrencyID field to given value.

### HasCurrencyID

`func (o *CourseUpdateDto) HasCurrencyID() bool`

HasCurrencyID returns a boolean if a field has been set.

### SetCurrencyIDNil

`func (o *CourseUpdateDto) SetCurrencyIDNil(b bool)`

 SetCurrencyIDNil sets the value for CurrencyID to be an explicit nil

### UnsetCurrencyID
`func (o *CourseUpdateDto) UnsetCurrencyID()`

UnsetCurrencyID ensures that no value is present for CurrencyID, not even an explicit nil
### GetRegularPrice

`func (o *CourseUpdateDto) GetRegularPrice() float64`

GetRegularPrice returns the RegularPrice field if non-nil, zero value otherwise.

### GetRegularPriceOk

`func (o *CourseUpdateDto) GetRegularPriceOk() (*float64, bool)`

GetRegularPriceOk returns a tuple with the RegularPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularPrice

`func (o *CourseUpdateDto) SetRegularPrice(v float64)`

SetRegularPrice sets RegularPrice field to given value.

### HasRegularPrice

`func (o *CourseUpdateDto) HasRegularPrice() bool`

HasRegularPrice returns a boolean if a field has been set.

### SetRegularPriceNil

`func (o *CourseUpdateDto) SetRegularPriceNil(b bool)`

 SetRegularPriceNil sets the value for RegularPrice to be an explicit nil

### UnsetRegularPrice
`func (o *CourseUpdateDto) UnsetRegularPrice()`

UnsetRegularPrice ensures that no value is present for RegularPrice, not even an explicit nil
### GetMaxCourseEnrollments

`func (o *CourseUpdateDto) GetMaxCourseEnrollments() int32`

GetMaxCourseEnrollments returns the MaxCourseEnrollments field if non-nil, zero value otherwise.

### GetMaxCourseEnrollmentsOk

`func (o *CourseUpdateDto) GetMaxCourseEnrollmentsOk() (*int32, bool)`

GetMaxCourseEnrollmentsOk returns a tuple with the MaxCourseEnrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCourseEnrollments

`func (o *CourseUpdateDto) SetMaxCourseEnrollments(v int32)`

SetMaxCourseEnrollments sets MaxCourseEnrollments field to given value.

### HasMaxCourseEnrollments

`func (o *CourseUpdateDto) HasMaxCourseEnrollments() bool`

HasMaxCourseEnrollments returns a boolean if a field has been set.

### SetMaxCourseEnrollmentsNil

`func (o *CourseUpdateDto) SetMaxCourseEnrollmentsNil(b bool)`

 SetMaxCourseEnrollmentsNil sets the value for MaxCourseEnrollments to be an explicit nil

### UnsetMaxCourseEnrollments
`func (o *CourseUpdateDto) UnsetMaxCourseEnrollments()`

UnsetMaxCourseEnrollments ensures that no value is present for MaxCourseEnrollments, not even an explicit nil
### GetTotalEffortInWeeks

`func (o *CourseUpdateDto) GetTotalEffortInWeeks() int32`

GetTotalEffortInWeeks returns the TotalEffortInWeeks field if non-nil, zero value otherwise.

### GetTotalEffortInWeeksOk

`func (o *CourseUpdateDto) GetTotalEffortInWeeksOk() (*int32, bool)`

GetTotalEffortInWeeksOk returns a tuple with the TotalEffortInWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEffortInWeeks

`func (o *CourseUpdateDto) SetTotalEffortInWeeks(v int32)`

SetTotalEffortInWeeks sets TotalEffortInWeeks field to given value.

### HasTotalEffortInWeeks

`func (o *CourseUpdateDto) HasTotalEffortInWeeks() bool`

HasTotalEffortInWeeks returns a boolean if a field has been set.

### SetTotalEffortInWeeksNil

`func (o *CourseUpdateDto) SetTotalEffortInWeeksNil(b bool)`

 SetTotalEffortInWeeksNil sets the value for TotalEffortInWeeks to be an explicit nil

### UnsetTotalEffortInWeeks
`func (o *CourseUpdateDto) UnsetTotalEffortInWeeks()`

UnsetTotalEffortInWeeks ensures that no value is present for TotalEffortInWeeks, not even an explicit nil
### GetTotalHoursPerWeek

`func (o *CourseUpdateDto) GetTotalHoursPerWeek() int32`

GetTotalHoursPerWeek returns the TotalHoursPerWeek field if non-nil, zero value otherwise.

### GetTotalHoursPerWeekOk

`func (o *CourseUpdateDto) GetTotalHoursPerWeekOk() (*int32, bool)`

GetTotalHoursPerWeekOk returns a tuple with the TotalHoursPerWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHoursPerWeek

`func (o *CourseUpdateDto) SetTotalHoursPerWeek(v int32)`

SetTotalHoursPerWeek sets TotalHoursPerWeek field to given value.

### HasTotalHoursPerWeek

`func (o *CourseUpdateDto) HasTotalHoursPerWeek() bool`

HasTotalHoursPerWeek returns a boolean if a field has been set.

### SetTotalHoursPerWeekNil

`func (o *CourseUpdateDto) SetTotalHoursPerWeekNil(b bool)`

 SetTotalHoursPerWeekNil sets the value for TotalHoursPerWeek to be an explicit nil

### UnsetTotalHoursPerWeek
`func (o *CourseUpdateDto) UnsetTotalHoursPerWeek()`

UnsetTotalHoursPerWeek ensures that no value is present for TotalHoursPerWeek, not even an explicit nil
### GetTotalEffortInHours

`func (o *CourseUpdateDto) GetTotalEffortInHours() int32`

GetTotalEffortInHours returns the TotalEffortInHours field if non-nil, zero value otherwise.

### GetTotalEffortInHoursOk

`func (o *CourseUpdateDto) GetTotalEffortInHoursOk() (*int32, bool)`

GetTotalEffortInHoursOk returns a tuple with the TotalEffortInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEffortInHours

`func (o *CourseUpdateDto) SetTotalEffortInHours(v int32)`

SetTotalEffortInHours sets TotalEffortInHours field to given value.

### HasTotalEffortInHours

`func (o *CourseUpdateDto) HasTotalEffortInHours() bool`

HasTotalEffortInHours returns a boolean if a field has been set.

### SetTotalEffortInHoursNil

`func (o *CourseUpdateDto) SetTotalEffortInHoursNil(b bool)`

 SetTotalEffortInHoursNil sets the value for TotalEffortInHours to be an explicit nil

### UnsetTotalEffortInHours
`func (o *CourseUpdateDto) UnsetTotalEffortInHours()`

UnsetTotalEffortInHours ensures that no value is present for TotalEffortInHours, not even an explicit nil
### GetStartDateTime

`func (o *CourseUpdateDto) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CourseUpdateDto) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CourseUpdateDto) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CourseUpdateDto) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *CourseUpdateDto) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *CourseUpdateDto) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetEndDateTime

`func (o *CourseUpdateDto) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CourseUpdateDto) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CourseUpdateDto) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CourseUpdateDto) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *CourseUpdateDto) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *CourseUpdateDto) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetInscriptionsStartDateTime

`func (o *CourseUpdateDto) GetInscriptionsStartDateTime() time.Time`

GetInscriptionsStartDateTime returns the InscriptionsStartDateTime field if non-nil, zero value otherwise.

### GetInscriptionsStartDateTimeOk

`func (o *CourseUpdateDto) GetInscriptionsStartDateTimeOk() (*time.Time, bool)`

GetInscriptionsStartDateTimeOk returns a tuple with the InscriptionsStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptionsStartDateTime

`func (o *CourseUpdateDto) SetInscriptionsStartDateTime(v time.Time)`

SetInscriptionsStartDateTime sets InscriptionsStartDateTime field to given value.

### HasInscriptionsStartDateTime

`func (o *CourseUpdateDto) HasInscriptionsStartDateTime() bool`

HasInscriptionsStartDateTime returns a boolean if a field has been set.

### SetInscriptionsStartDateTimeNil

`func (o *CourseUpdateDto) SetInscriptionsStartDateTimeNil(b bool)`

 SetInscriptionsStartDateTimeNil sets the value for InscriptionsStartDateTime to be an explicit nil

### UnsetInscriptionsStartDateTime
`func (o *CourseUpdateDto) UnsetInscriptionsStartDateTime()`

UnsetInscriptionsStartDateTime ensures that no value is present for InscriptionsStartDateTime, not even an explicit nil
### GetInscriptionsEndDateTime

`func (o *CourseUpdateDto) GetInscriptionsEndDateTime() time.Time`

GetInscriptionsEndDateTime returns the InscriptionsEndDateTime field if non-nil, zero value otherwise.

### GetInscriptionsEndDateTimeOk

`func (o *CourseUpdateDto) GetInscriptionsEndDateTimeOk() (*time.Time, bool)`

GetInscriptionsEndDateTimeOk returns a tuple with the InscriptionsEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptionsEndDateTime

`func (o *CourseUpdateDto) SetInscriptionsEndDateTime(v time.Time)`

SetInscriptionsEndDateTime sets InscriptionsEndDateTime field to given value.

### HasInscriptionsEndDateTime

`func (o *CourseUpdateDto) HasInscriptionsEndDateTime() bool`

HasInscriptionsEndDateTime returns a boolean if a field has been set.

### SetInscriptionsEndDateTimeNil

`func (o *CourseUpdateDto) SetInscriptionsEndDateTimeNil(b bool)`

 SetInscriptionsEndDateTimeNil sets the value for InscriptionsEndDateTime to be an explicit nil

### UnsetInscriptionsEndDateTime
`func (o *CourseUpdateDto) UnsetInscriptionsEndDateTime()`

UnsetInscriptionsEndDateTime ensures that no value is present for InscriptionsEndDateTime, not even an explicit nil
### GetPublished

`func (o *CourseUpdateDto) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CourseUpdateDto) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CourseUpdateDto) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CourseUpdateDto) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *CourseUpdateDto) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *CourseUpdateDto) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


