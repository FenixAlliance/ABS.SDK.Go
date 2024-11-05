# SupportRequestAttachmentDto

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
**SupportRequestID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportRequestAttachmentDto

`func NewSupportRequestAttachmentDto() *SupportRequestAttachmentDto`

NewSupportRequestAttachmentDto instantiates a new SupportRequestAttachmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportRequestAttachmentDtoWithDefaults

`func NewSupportRequestAttachmentDtoWithDefaults() *SupportRequestAttachmentDto`

NewSupportRequestAttachmentDtoWithDefaults instantiates a new SupportRequestAttachmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportRequestAttachmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportRequestAttachmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportRequestAttachmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportRequestAttachmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportRequestAttachmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportRequestAttachmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SupportRequestAttachmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportRequestAttachmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportRequestAttachmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportRequestAttachmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SupportRequestAttachmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SupportRequestAttachmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetNotes

`func (o *SupportRequestAttachmentDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SupportRequestAttachmentDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SupportRequestAttachmentDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SupportRequestAttachmentDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *SupportRequestAttachmentDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SupportRequestAttachmentDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTitle

`func (o *SupportRequestAttachmentDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SupportRequestAttachmentDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SupportRequestAttachmentDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SupportRequestAttachmentDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SupportRequestAttachmentDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SupportRequestAttachmentDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAuthor

`func (o *SupportRequestAttachmentDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SupportRequestAttachmentDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SupportRequestAttachmentDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SupportRequestAttachmentDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SupportRequestAttachmentDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SupportRequestAttachmentDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetIsFolder

`func (o *SupportRequestAttachmentDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SupportRequestAttachmentDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SupportRequestAttachmentDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SupportRequestAttachmentDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### GetHash

`func (o *SupportRequestAttachmentDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SupportRequestAttachmentDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SupportRequestAttachmentDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SupportRequestAttachmentDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *SupportRequestAttachmentDto) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *SupportRequestAttachmentDto) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetFileUrl

`func (o *SupportRequestAttachmentDto) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *SupportRequestAttachmentDto) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *SupportRequestAttachmentDto) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *SupportRequestAttachmentDto) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *SupportRequestAttachmentDto) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *SupportRequestAttachmentDto) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetFilePath

`func (o *SupportRequestAttachmentDto) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *SupportRequestAttachmentDto) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *SupportRequestAttachmentDto) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *SupportRequestAttachmentDto) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *SupportRequestAttachmentDto) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *SupportRequestAttachmentDto) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetFileName

`func (o *SupportRequestAttachmentDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SupportRequestAttachmentDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SupportRequestAttachmentDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SupportRequestAttachmentDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *SupportRequestAttachmentDto) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *SupportRequestAttachmentDto) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetAbstract

`func (o *SupportRequestAttachmentDto) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *SupportRequestAttachmentDto) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *SupportRequestAttachmentDto) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *SupportRequestAttachmentDto) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### SetAbstractNil

`func (o *SupportRequestAttachmentDto) SetAbstractNil(b bool)`

 SetAbstractNil sets the value for Abstract to be an explicit nil

### UnsetAbstract
`func (o *SupportRequestAttachmentDto) UnsetAbstract()`

UnsetAbstract ensures that no value is present for Abstract, not even an explicit nil
### GetKeyWords

`func (o *SupportRequestAttachmentDto) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *SupportRequestAttachmentDto) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *SupportRequestAttachmentDto) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *SupportRequestAttachmentDto) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### SetKeyWordsNil

`func (o *SupportRequestAttachmentDto) SetKeyWordsNil(b bool)`

 SetKeyWordsNil sets the value for KeyWords to be an explicit nil

### UnsetKeyWords
`func (o *SupportRequestAttachmentDto) UnsetKeyWords()`

UnsetKeyWords ensures that no value is present for KeyWords, not even an explicit nil
### GetMetadata

`func (o *SupportRequestAttachmentDto) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportRequestAttachmentDto) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportRequestAttachmentDto) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SupportRequestAttachmentDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SupportRequestAttachmentDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SupportRequestAttachmentDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFileLength

`func (o *SupportRequestAttachmentDto) GetFileLength() int64`

GetFileLength returns the FileLength field if non-nil, zero value otherwise.

### GetFileLengthOk

`func (o *SupportRequestAttachmentDto) GetFileLengthOk() (*int64, bool)`

GetFileLengthOk returns a tuple with the FileLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLength

`func (o *SupportRequestAttachmentDto) SetFileLength(v int64)`

SetFileLength sets FileLength field to given value.

### HasFileLength

`func (o *SupportRequestAttachmentDto) HasFileLength() bool`

HasFileLength returns a boolean if a field has been set.

### GetContentType

`func (o *SupportRequestAttachmentDto) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SupportRequestAttachmentDto) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SupportRequestAttachmentDto) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SupportRequestAttachmentDto) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *SupportRequestAttachmentDto) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *SupportRequestAttachmentDto) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetParentFileId

`func (o *SupportRequestAttachmentDto) GetParentFileId() string`

GetParentFileId returns the ParentFileId field if non-nil, zero value otherwise.

### GetParentFileIdOk

`func (o *SupportRequestAttachmentDto) GetParentFileIdOk() (*string, bool)`

GetParentFileIdOk returns a tuple with the ParentFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFileId

`func (o *SupportRequestAttachmentDto) SetParentFileId(v string)`

SetParentFileId sets ParentFileId field to given value.

### HasParentFileId

`func (o *SupportRequestAttachmentDto) HasParentFileId() bool`

HasParentFileId returns a boolean if a field has been set.

### SetParentFileIdNil

`func (o *SupportRequestAttachmentDto) SetParentFileIdNil(b bool)`

 SetParentFileIdNil sets the value for ParentFileId to be an explicit nil

### UnsetParentFileId
`func (o *SupportRequestAttachmentDto) UnsetParentFileId()`

UnsetParentFileId ensures that no value is present for ParentFileId, not even an explicit nil
### GetValidResponse

`func (o *SupportRequestAttachmentDto) GetValidResponse() bool`

GetValidResponse returns the ValidResponse field if non-nil, zero value otherwise.

### GetValidResponseOk

`func (o *SupportRequestAttachmentDto) GetValidResponseOk() (*bool, bool)`

GetValidResponseOk returns a tuple with the ValidResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidResponse

`func (o *SupportRequestAttachmentDto) SetValidResponse(v bool)`

SetValidResponse sets ValidResponse field to given value.

### HasValidResponse

`func (o *SupportRequestAttachmentDto) HasValidResponse() bool`

HasValidResponse returns a boolean if a field has been set.

### GetUserId

`func (o *SupportRequestAttachmentDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SupportRequestAttachmentDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SupportRequestAttachmentDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SupportRequestAttachmentDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SupportRequestAttachmentDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SupportRequestAttachmentDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTenantId

`func (o *SupportRequestAttachmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SupportRequestAttachmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SupportRequestAttachmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SupportRequestAttachmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SupportRequestAttachmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SupportRequestAttachmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *SupportRequestAttachmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SupportRequestAttachmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SupportRequestAttachmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SupportRequestAttachmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SupportRequestAttachmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SupportRequestAttachmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSocialProfileId

`func (o *SupportRequestAttachmentDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SupportRequestAttachmentDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SupportRequestAttachmentDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SupportRequestAttachmentDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SupportRequestAttachmentDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SupportRequestAttachmentDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetFolderPath

`func (o *SupportRequestAttachmentDto) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *SupportRequestAttachmentDto) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *SupportRequestAttachmentDto) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *SupportRequestAttachmentDto) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### SetFolderPathNil

`func (o *SupportRequestAttachmentDto) SetFolderPathNil(b bool)`

 SetFolderPathNil sets the value for FolderPath to be an explicit nil

### UnsetFolderPath
`func (o *SupportRequestAttachmentDto) UnsetFolderPath()`

UnsetFolderPath ensures that no value is present for FolderPath, not even an explicit nil
### GetSupportRequestID

`func (o *SupportRequestAttachmentDto) GetSupportRequestID() string`

GetSupportRequestID returns the SupportRequestID field if non-nil, zero value otherwise.

### GetSupportRequestIDOk

`func (o *SupportRequestAttachmentDto) GetSupportRequestIDOk() (*string, bool)`

GetSupportRequestIDOk returns a tuple with the SupportRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequestID

`func (o *SupportRequestAttachmentDto) SetSupportRequestID(v string)`

SetSupportRequestID sets SupportRequestID field to given value.

### HasSupportRequestID

`func (o *SupportRequestAttachmentDto) HasSupportRequestID() bool`

HasSupportRequestID returns a boolean if a field has been set.

### SetSupportRequestIDNil

`func (o *SupportRequestAttachmentDto) SetSupportRequestIDNil(b bool)`

 SetSupportRequestIDNil sets the value for SupportRequestID to be an explicit nil

### UnsetSupportRequestID
`func (o *SupportRequestAttachmentDto) UnsetSupportRequestID()`

UnsetSupportRequestID ensures that no value is present for SupportRequestID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


