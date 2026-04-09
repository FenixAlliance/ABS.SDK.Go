# CourseFileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**FileUploadURL** | Pointer to **NullableString** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**CourseID** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCourseFileDto

`func NewCourseFileDto() *CourseFileDto`

NewCourseFileDto instantiates a new CourseFileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourseFileDtoWithDefaults

`func NewCourseFileDtoWithDefaults() *CourseFileDto`

NewCourseFileDtoWithDefaults instantiates a new CourseFileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CourseFileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CourseFileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CourseFileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CourseFileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CourseFileDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CourseFileDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CourseFileDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CourseFileDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CourseFileDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CourseFileDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CourseFileDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CourseFileDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CourseFileDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CourseFileDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CourseFileDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CourseFileDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CourseFileDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CourseFileDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFileName

`func (o *CourseFileDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CourseFileDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CourseFileDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CourseFileDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *CourseFileDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *CourseFileDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFileUploadURL

`func (o *CourseFileDto) GetFileUploadURL() string`

GetFileUploadURL returns the FileUploadURL field if non-nil, zero value otherwise.

### GetFileUploadURLOk

`func (o *CourseFileDto) GetFileUploadURLOk() (*string, bool)`

GetFileUploadURLOk returns a tuple with the FileUploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadURL

`func (o *CourseFileDto) SetFileUploadURL(v string)`

SetFileUploadURL sets FileUploadURL field to given value.

### HasFileUploadURL

`func (o *CourseFileDto) HasFileUploadURL() bool`

HasFileUploadURL returns a boolean if a field has been set.

### SetFileUploadURLNil

`func (o *CourseFileDto) SetFileUploadURLNil(b bool)`

 SetFileUploadURLNil sets the value for FileUploadURL to be an explicit nil

### UnsetFileUploadURL
`func (o *CourseFileDto) UnsetFileUploadURL()`

UnsetFileUploadURL ensures that no value is present for FileUploadURL, not even an explicit nil
### GetContentType

`func (o *CourseFileDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CourseFileDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CourseFileDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CourseFileDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *CourseFileDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *CourseFileDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileLength

`func (o *CourseFileDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *CourseFileDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *CourseFileDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *CourseFileDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetCourseID

`func (o *CourseFileDto) GetCourseID() string`

GetCourseID returns the CourseID field if non-nil, zero value otherwise.

### GetCourseIDOk

`func (o *CourseFileDto) GetCourseIDOk() (*string, bool)`

GetCourseIDOk returns a tuple with the CourseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourseID

`func (o *CourseFileDto) SetCourseID(v string)`

SetCourseID sets CourseID field to given value.

### HasCourseID

`func (o *CourseFileDto) HasCourseID() bool`

HasCourseID returns a boolean if a field has been set.

### SetCourseIDNil

`func (o *CourseFileDto) SetCourseIDNil(b bool)`

 SetCourseIDNil sets the value for CourseID to be an explicit nil

### UnsetCourseID
`func (o *CourseFileDto) UnsetCourseID()`

UnsetCourseID ensures that no value is present for CourseID, not even an explicit nil
### GetTenantId

`func (o *CourseFileDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CourseFileDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CourseFileDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CourseFileDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CourseFileDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CourseFileDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


