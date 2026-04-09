# SocialPostAttachmentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**ParentFileUploadId** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 
**SocialPostId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialPostAttachmentCreateDto

`func NewSocialPostAttachmentCreateDto() *SocialPostAttachmentCreateDto`

NewSocialPostAttachmentCreateDto instantiates a new SocialPostAttachmentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostAttachmentCreateDtoWithDefaults

`func NewSocialPostAttachmentCreateDtoWithDefaults() *SocialPostAttachmentCreateDto`

NewSocialPostAttachmentCreateDtoWithDefaults instantiates a new SocialPostAttachmentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostAttachmentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostAttachmentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostAttachmentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostAttachmentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SocialPostAttachmentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostAttachmentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostAttachmentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostAttachmentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNotes

`func (o *SocialPostAttachmentCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SocialPostAttachmentCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SocialPostAttachmentCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SocialPostAttachmentCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *SocialPostAttachmentCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SocialPostAttachmentCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTitle

`func (o *SocialPostAttachmentCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialPostAttachmentCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialPostAttachmentCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialPostAttachmentCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialPostAttachmentCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialPostAttachmentCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *SocialPostAttachmentCreateDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SocialPostAttachmentCreateDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SocialPostAttachmentCreateDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SocialPostAttachmentCreateDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SocialPostAttachmentCreateDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SocialPostAttachmentCreateDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *SocialPostAttachmentCreateDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SocialPostAttachmentCreateDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SocialPostAttachmentCreateDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SocialPostAttachmentCreateDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetFileName

`func (o *SocialPostAttachmentCreateDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SocialPostAttachmentCreateDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SocialPostAttachmentCreateDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SocialPostAttachmentCreateDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SocialPostAttachmentCreateDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SocialPostAttachmentCreateDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *SocialPostAttachmentCreateDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SocialPostAttachmentCreateDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SocialPostAttachmentCreateDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SocialPostAttachmentCreateDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *SocialPostAttachmentCreateDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *SocialPostAttachmentCreateDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *SocialPostAttachmentCreateDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *SocialPostAttachmentCreateDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *SocialPostAttachmentCreateDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *SocialPostAttachmentCreateDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *SocialPostAttachmentCreateDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *SocialPostAttachmentCreateDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetValidResponse

`func (o *SocialPostAttachmentCreateDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *SocialPostAttachmentCreateDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *SocialPostAttachmentCreateDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *SocialPostAttachmentCreateDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetParentFileUploadId

`func (o *SocialPostAttachmentCreateDto) GetParentFileUploadId() string`

GetParentFileUploadId returns the ParentFileUploadId field if non-nil, zero value otherwise.

### GetParentFileUploadIdOk

`func (o *SocialPostAttachmentCreateDto) GetParentFileUploadIdOk() (*string, bool)`

GetParentFileUploadIdOk returns a tuple with the ParentFileUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileUploadId

`func (o *SocialPostAttachmentCreateDto) SetParentFileUploadId(v string)`

SetParentFileUploadId sets ParentFileUploadId field to given value.

### HasParentFileUploadId

`func (o *SocialPostAttachmentCreateDto) HasParentFileUploadId() bool`

HasParentFileUploadId returns a boolean if a field has been set.

### SetParentFileUploadIdNil

`func (o *SocialPostAttachmentCreateDto) SetParentFileUploadIdNil(b bool)`

 SetParentFileUploadIdNil sets the value for ParentFileUploadId to be an explicit nil

### UnsetParentFileUploadId
`func (o *SocialPostAttachmentCreateDto) UnsetParentFileUploadId()`

UnsetParentFileUploadId ensures that no value is present for ParentFileUploadId, not even an explicit nil
### GetFilePath

`func (o *SocialPostAttachmentCreateDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SocialPostAttachmentCreateDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SocialPostAttachmentCreateDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SocialPostAttachmentCreateDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *SocialPostAttachmentCreateDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *SocialPostAttachmentCreateDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetSocialPostId

`func (o *SocialPostAttachmentCreateDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *SocialPostAttachmentCreateDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *SocialPostAttachmentCreateDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *SocialPostAttachmentCreateDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *SocialPostAttachmentCreateDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *SocialPostAttachmentCreateDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


