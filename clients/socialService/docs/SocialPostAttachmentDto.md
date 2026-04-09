# SocialPostAttachmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**IsFolder** | Pointer to **bool** |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 
**FileUrl** | Pointer to **NullableString** |  | [optional] 
**FilePath** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Abstract** | Pointer to **NullableString** |  | [optional] 
**KeyWords** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **NullableString** |  | [optional] 
**FileLength** | Pointer to **int64** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ParentFileId** | Pointer to **NullableString** |  | [optional] 
**ValidResponse** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**FolderPath** | Pointer to **NullableString** |  | [optional] 
**SocialPostId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialPostAttachmentDto

`func NewSocialPostAttachmentDto() *SocialPostAttachmentDto`

NewSocialPostAttachmentDto instantiates a new SocialPostAttachmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostAttachmentDtoWithDefaults

`func NewSocialPostAttachmentDtoWithDefaults() *SocialPostAttachmentDto`

NewSocialPostAttachmentDtoWithDefaults instantiates a new SocialPostAttachmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialPostAttachmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPostAttachmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPostAttachmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPostAttachmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialPostAttachmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialPostAttachmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialPostAttachmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialPostAttachmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialPostAttachmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialPostAttachmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialPostAttachmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialPostAttachmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetNotes

`func (o *SocialPostAttachmentDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SocialPostAttachmentDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SocialPostAttachmentDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SocialPostAttachmentDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *SocialPostAttachmentDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SocialPostAttachmentDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTitle

`func (o *SocialPostAttachmentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialPostAttachmentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialPostAttachmentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialPostAttachmentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialPostAttachmentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialPostAttachmentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *SocialPostAttachmentDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SocialPostAttachmentDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SocialPostAttachmentDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SocialPostAttachmentDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SocialPostAttachmentDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SocialPostAttachmentDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *SocialPostAttachmentDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SocialPostAttachmentDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SocialPostAttachmentDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SocialPostAttachmentDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetHash

`func (o *SocialPostAttachmentDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SocialPostAttachmentDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SocialPostAttachmentDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SocialPostAttachmentDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *SocialPostAttachmentDto) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *SocialPostAttachmentDto) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetFileUrl

`func (o *SocialPostAttachmentDto) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *SocialPostAttachmentDto) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *SocialPostAttachmentDto) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *SocialPostAttachmentDto) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *SocialPostAttachmentDto) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *SocialPostAttachmentDto) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetFilePath

`func (o *SocialPostAttachmentDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SocialPostAttachmentDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SocialPostAttachmentDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SocialPostAttachmentDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *SocialPostAttachmentDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *SocialPostAttachmentDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetFileName

`func (o *SocialPostAttachmentDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SocialPostAttachmentDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SocialPostAttachmentDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SocialPostAttachmentDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SocialPostAttachmentDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SocialPostAttachmentDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *SocialPostAttachmentDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SocialPostAttachmentDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SocialPostAttachmentDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SocialPostAttachmentDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *SocialPostAttachmentDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *SocialPostAttachmentDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *SocialPostAttachmentDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *SocialPostAttachmentDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *SocialPostAttachmentDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *SocialPostAttachmentDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *SocialPostAttachmentDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *SocialPostAttachmentDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetMetadata

`func (o *SocialPostAttachmentDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SocialPostAttachmentDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SocialPostAttachmentDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SocialPostAttachmentDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SocialPostAttachmentDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SocialPostAttachmentDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFileLength

`func (o *SocialPostAttachmentDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *SocialPostAttachmentDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *SocialPostAttachmentDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *SocialPostAttachmentDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetContentType

`func (o *SocialPostAttachmentDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SocialPostAttachmentDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SocialPostAttachmentDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SocialPostAttachmentDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SocialPostAttachmentDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SocialPostAttachmentDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetParentFileId

`func (o *SocialPostAttachmentDto) GetParentFileId() string`

GetParentFileId returns the ParentFileId field if non-nil, zero value otherwise.

### GetParentFileIdOk

`func (o *SocialPostAttachmentDto) GetParentFileIdOk() (*string, bool)`

GetParentFileIdOk returns a tuple with the ParentFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileId

`func (o *SocialPostAttachmentDto) SetParentFileId(v string)`

SetParentFileId sets ParentFileId field to given value.

### HasParentFileId

`func (o *SocialPostAttachmentDto) HasParentFileId() bool`

HasParentFileId returns a boolean if a field has been set.

### SetParentFileIdNil

`func (o *SocialPostAttachmentDto) SetParentFileIdNil(b bool)`

 SetParentFileIdNil sets the value for ParentFileId to be an explicit nil

### UnsetParentFileId
`func (o *SocialPostAttachmentDto) UnsetParentFileId()`

UnsetParentFileId ensures that no value is present for ParentFileId, not even an explicit nil
### GetValidResponse

`func (o *SocialPostAttachmentDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *SocialPostAttachmentDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *SocialPostAttachmentDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *SocialPostAttachmentDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetUserId

`func (o *SocialPostAttachmentDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SocialPostAttachmentDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SocialPostAttachmentDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SocialPostAttachmentDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SocialPostAttachmentDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SocialPostAttachmentDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *SocialPostAttachmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SocialPostAttachmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SocialPostAttachmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SocialPostAttachmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SocialPostAttachmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SocialPostAttachmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *SocialPostAttachmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SocialPostAttachmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SocialPostAttachmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SocialPostAttachmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SocialPostAttachmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SocialPostAttachmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *SocialPostAttachmentDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialPostAttachmentDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialPostAttachmentDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialPostAttachmentDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialPostAttachmentDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialPostAttachmentDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetFolderPath

`func (o *SocialPostAttachmentDto) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *SocialPostAttachmentDto) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *SocialPostAttachmentDto) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *SocialPostAttachmentDto) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### SetFolderPathNil

`func (o *SocialPostAttachmentDto) SetFolderPathNil(b bool)`

 SetFolderPathNil sets the value for FolderPath to be an explicit nil

### UnsetFolderPath
`func (o *SocialPostAttachmentDto) UnsetFolderPath()`

UnsetFolderPath ensures that no value is present for FolderPath, not even an explicit nil
### GetSocialPostId

`func (o *SocialPostAttachmentDto) GetSocialPostId() string`

GetSocialPostId returns the SocialPostId field if non-nil, zero value otherwise.

### GetSocialPostIdOk

`func (o *SocialPostAttachmentDto) GetSocialPostIdOk() (*string, bool)`

GetSocialPostIdOk returns a tuple with the SocialPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialPostId

`func (o *SocialPostAttachmentDto) SetSocialPostId(v string)`

SetSocialPostId sets SocialPostId field to given value.

### HasSocialPostId

`func (o *SocialPostAttachmentDto) HasSocialPostId() bool`

HasSocialPostId returns a boolean if a field has been set.

### SetSocialPostIdNil

`func (o *SocialPostAttachmentDto) SetSocialPostIdNil(b bool)`

 SetSocialPostIdNil sets the value for SocialPostId to be an explicit nil

### UnsetSocialPostId
`func (o *SocialPostAttachmentDto) UnsetSocialPostId()`

UnsetSocialPostId ensures that no value is present for SocialPostId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


