# CourseCertificateTemplateCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CourseID** | **string** |  | 
**BusinessID** | **string** |  | 
**WebPortalID** | Pointer to **NullableString** |  | [optional] 
**WebsiteThemeID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 
**ParentWebContentID** | Pointer to **NullableString** |  | [optional] 
**ParentWebContentVersionID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseCertificateTemplateCreateDto

`func NewCourseCertificateTemplateCreateDto(courseID string, businessID string, ) *CourseCertificateTemplateCreateDto`

NewCourseCertificateTemplateCreateDto instantiates a new CourseCertificateTemplateCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseCertificateTemplateCreateDtoWithDefaults

`func NewCourseCertificateTemplateCreateDtoWithDefaults() *CourseCertificateTemplateCreateDto`

NewCourseCertificateTemplateCreateDtoWithDefaults instantiates a new CourseCertificateTemplateCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseCertificateTemplateCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseCertificateTemplateCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseCertificateTemplateCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseCertificateTemplateCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CourseCertificateTemplateCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseCertificateTemplateCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseCertificateTemplateCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseCertificateTemplateCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCourseID

`func (o *CourseCertificateTemplateCreateDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseCertificateTemplateCreateDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseCertificateTemplateCreateDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.


### GetBusinessID

`func (o *CourseCertificateTemplateCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *CourseCertificateTemplateCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *CourseCertificateTemplateCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetWebPortalID

`func (o *CourseCertificateTemplateCreateDto) GetWebPortalID() string`

GetWebPortalID returns the WebPortalID field if non-nil, zero value otherwise.

### GetWebPortalIDOk

`func (o *CourseCertificateTemplateCreateDto) GetWebPortalIDOk() (*string, bool)`

GetWebPortalIDOk returns a tuple with the WebPortalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPortalID

`func (o *CourseCertificateTemplateCreateDto) SetWebPortalID(v string)`

SetWebPortalID sets WebPortalID field to given value.

### HasWebPortalID

`func (o *CourseCertificateTemplateCreateDto) HasWebPortalID() bool`

HasWebPortalID returns a boolean if a field has been set.

### SetWebPortalIDNil

`func (o *CourseCertificateTemplateCreateDto) SetWebPortalIDNil(b bool)`

 SetWebPortalIDNil sets the value for WebPortalID to be an explicit nil

### UnsetWebPortalID
`func (o *CourseCertificateTemplateCreateDto) UnsetWebPortalID()`

UnsetWebPortalID ensures that no value is present for WebPortalID, not even an explicit nil
### GetWebsiteThemeID

`func (o *CourseCertificateTemplateCreateDto) GetWebsiteThemeID() string`

GetWebsiteThemeID returns the WebsiteThemeID field if non-nil, zero value otherwise.

### GetWebsiteThemeIDOk

`func (o *CourseCertificateTemplateCreateDto) GetWebsiteThemeIDOk() (*string, bool)`

GetWebsiteThemeIDOk returns a tuple with the WebsiteThemeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteThemeID

`func (o *CourseCertificateTemplateCreateDto) SetWebsiteThemeID(v string)`

SetWebsiteThemeID sets WebsiteThemeID field to given value.

### HasWebsiteThemeID

`func (o *CourseCertificateTemplateCreateDto) HasWebsiteThemeID() bool`

HasWebsiteThemeID returns a boolean if a field has been set.

### SetWebsiteThemeIDNil

`func (o *CourseCertificateTemplateCreateDto) SetWebsiteThemeIDNil(b bool)`

 SetWebsiteThemeIDNil sets the value for WebsiteThemeID to be an explicit nil

### UnsetWebsiteThemeID
`func (o *CourseCertificateTemplateCreateDto) UnsetWebsiteThemeID()`

UnsetWebsiteThemeID ensures that no value is present for WebsiteThemeID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *CourseCertificateTemplateCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *CourseCertificateTemplateCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *CourseCertificateTemplateCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *CourseCertificateTemplateCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *CourseCertificateTemplateCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *CourseCertificateTemplateCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
### GetSocialProfileID

`func (o *CourseCertificateTemplateCreateDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *CourseCertificateTemplateCreateDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *CourseCertificateTemplateCreateDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *CourseCertificateTemplateCreateDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *CourseCertificateTemplateCreateDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *CourseCertificateTemplateCreateDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
### GetParentWebContentID

`func (o *CourseCertificateTemplateCreateDto) GetParentWebContentID() string`

GetParentWebContentID returns the ParentWebContentID field if non-nil, zero value otherwise.

### GetParentWebContentIDOk

`func (o *CourseCertificateTemplateCreateDto) GetParentWebContentIDOk() (*string, bool)`

GetParentWebContentIDOk returns a tuple with the ParentWebContentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentID

`func (o *CourseCertificateTemplateCreateDto) SetParentWebContentID(v string)`

SetParentWebContentID sets ParentWebContentID field to given value.

### HasParentWebContentID

`func (o *CourseCertificateTemplateCreateDto) HasParentWebContentID() bool`

HasParentWebContentID returns a boolean if a field has been set.

### SetParentWebContentIDNil

`func (o *CourseCertificateTemplateCreateDto) SetParentWebContentIDNil(b bool)`

 SetParentWebContentIDNil sets the value for ParentWebContentID to be an explicit nil

### UnsetParentWebContentID
`func (o *CourseCertificateTemplateCreateDto) UnsetParentWebContentID()`

UnsetParentWebContentID ensures that no value is present for ParentWebContentID, not even an explicit nil
### GetParentWebContentVersionID

`func (o *CourseCertificateTemplateCreateDto) GetParentWebContentVersionID() string`

GetParentWebContentVersionID returns the ParentWebContentVersionID field if non-nil, zero value otherwise.

### GetParentWebContentVersionIDOk

`func (o *CourseCertificateTemplateCreateDto) GetParentWebContentVersionIDOk() (*string, bool)`

GetParentWebContentVersionIDOk returns a tuple with the ParentWebContentVersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWebContentVersionID

`func (o *CourseCertificateTemplateCreateDto) SetParentWebContentVersionID(v string)`

SetParentWebContentVersionID sets ParentWebContentVersionID field to given value.

### HasParentWebContentVersionID

`func (o *CourseCertificateTemplateCreateDto) HasParentWebContentVersionID() bool`

HasParentWebContentVersionID returns a boolean if a field has been set.

### SetParentWebContentVersionIDNil

`func (o *CourseCertificateTemplateCreateDto) SetParentWebContentVersionIDNil(b bool)`

 SetParentWebContentVersionIDNil sets the value for ParentWebContentVersionID to be an explicit nil

### UnsetParentWebContentVersionID
`func (o *CourseCertificateTemplateCreateDto) UnsetParentWebContentVersionID()`

UnsetParentWebContentVersionID ensures that no value is present for ParentWebContentVersionID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


