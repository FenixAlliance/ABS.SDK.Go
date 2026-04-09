# CoursePageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCoursePageDto

`func NewCoursePageDto() *CoursePageDto`

NewCoursePageDto instantiates a new CoursePageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoursePageDtoWithDefaults

`func NewCoursePageDtoWithDefaults() *CoursePageDto`

NewCoursePageDtoWithDefaults instantiates a new CoursePageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoursePageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoursePageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoursePageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoursePageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CoursePageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CoursePageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CoursePageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CoursePageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CoursePageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CoursePageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CoursePageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CoursePageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CoursePageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CoursePageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CoursePageDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CoursePageDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CoursePageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CoursePageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CoursePageDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoursePageDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoursePageDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoursePageDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CoursePageDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CoursePageDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *CoursePageDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CoursePageDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CoursePageDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CoursePageDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CoursePageDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CoursePageDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetSlug

`func (o *CoursePageDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CoursePageDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CoursePageDto) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CoursePageDto) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *CoursePageDto) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *CoursePageDto) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetCourseID

`func (o *CoursePageDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CoursePageDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CoursePageDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CoursePageDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CoursePageDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CoursePageDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetTenantId

`func (o *CoursePageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CoursePageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CoursePageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CoursePageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CoursePageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CoursePageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


