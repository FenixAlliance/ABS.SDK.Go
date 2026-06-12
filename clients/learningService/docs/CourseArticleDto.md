# CourseArticleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**CourseWikiId** | Pointer to **NullableString** |  | [optional] 
**CourseId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseArticleDto

`func NewCourseArticleDto() *CourseArticleDto`

NewCourseArticleDto instantiates a new CourseArticleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseArticleDtoWithDefaults

`func NewCourseArticleDtoWithDefaults() *CourseArticleDto`

NewCourseArticleDtoWithDefaults instantiates a new CourseArticleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseArticleDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseArticleDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseArticleDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseArticleDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseArticleDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseArticleDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseArticleDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseArticleDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseArticleDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseArticleDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseArticleDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseArticleDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CourseArticleDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseArticleDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseArticleDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseArticleDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseArticleDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseArticleDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CourseArticleDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CourseArticleDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CourseArticleDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CourseArticleDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CourseArticleDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CourseArticleDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CourseArticleDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CourseArticleDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CourseArticleDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CourseArticleDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CourseArticleDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CourseArticleDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCourseWikiId

`func (o *CourseArticleDto) GetCourseWikiId() string`

GetCourseWikiId returns the CourseWikiId field if non-nil, zero value otherwise.

### GetCourseWikiIdOk

`func (o *CourseArticleDto) GetCourseWikiIdOk() (*string, bool)`

GetCourseWikiIdOk returns a tuple with the CourseWikiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseWikiId

`func (o *CourseArticleDto) SetCourseWikiId(v string)`

SetCourseWikiId sets CourseWikiId field to given value.

### HasCourseWikiId

`func (o *CourseArticleDto) HasCourseWikiId() bool`

HasCourseWikiId returns a boolean if a field has been set.

### SetCourseWikiIdNil

`func (o *CourseArticleDto) SetCourseWikiIdNil(b bool)`

 SetCourseWikiIdNil sets the value for CourseWikiId to be an explicit nil

### UnsetCourseWikiId
`func (o *CourseArticleDto) UnsetCourseWikiId()`

UnsetCourseWikiId ensures that no value is present for CourseWikiId, not even an explicit nil
### GetCourseId

`func (o *CourseArticleDto) GetCourseId() string`

GetCourseId returns the CourseId field if non-nil, zero value otherwise.

### GetCourseIdOk

`func (o *CourseArticleDto) GetCourseIdOk() (*string, bool)`

GetCourseIdOk returns a tuple with the CourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseId

`func (o *CourseArticleDto) SetCourseId(v string)`

SetCourseId sets CourseId field to given value.

### HasCourseId

`func (o *CourseArticleDto) HasCourseId() bool`

HasCourseId returns a boolean if a field has been set.

### SetCourseIdNil

`func (o *CourseArticleDto) SetCourseIdNil(b bool)`

 SetCourseIdNil sets the value for CourseId to be an explicit nil

### UnsetCourseId
`func (o *CourseArticleDto) UnsetCourseId()`

UnsetCourseId ensures that no value is present for CourseId, not even an explicit nil
### GetTenantId

`func (o *CourseArticleDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseArticleDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseArticleDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseArticleDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseArticleDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseArticleDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CourseArticleDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CourseArticleDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CourseArticleDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CourseArticleDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CourseArticleDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CourseArticleDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


